/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package deliverservice

import (
	"sync"
	"time"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/internal/pkg/peer/blocksprovider"
	"github.com/pkg/errors"
)

const bftHeaderWrongStatusThreshold = 10

//go:generate mockery --dir=. --name=HeaderStreamClient --case=underscore --output=mocks/

// HeaderStreamClient defines the interface for a deliver client
type HeaderStreamClient interface {
	blocksprovider.DeliverClient
}

type bftHeaderReceiver struct {
	mutex              sync.Mutex
	chainID            string
	minBackoffDelay    time.Duration
	maxBackoffDelay    time.Duration
	stop               bool
	stopChan           chan struct{}
	started            bool
	endpoint           string
	client             HeaderStreamClient
	msgCryptoVerifier  blocksprovider.BlockVerifier
	lastHeader         *common.Block // a block with Header & Metadata, without Data (i.e. lastHeader.Data==nil)
	lastHeaderTime     time.Time
	lastHeaderVerified bool
	lastHeaderOK       bool
}

func newBFTHeaderReceiver(
	chainID string,
	endpoint string,
	client HeaderStreamClient,
	msgVerifier blocksprovider.BlockVerifier,
	minBackOff time.Duration,
	maxBackOff time.Duration,
) *bftHeaderReceiver {
	hRcv := &bftHeaderReceiver{
		chainID:           chainID,
		stopChan:          make(chan struct{}, 1),
		endpoint:          endpoint,
		client:            client,
		msgCryptoVerifier: msgVerifier,
		minBackoffDelay:   minBackOff,
		maxBackoffDelay:   maxBackOff,
	}
	return hRcv
}

// DeliverHeaders starts to deliver headers from the stream client
func (hr *bftHeaderReceiver) DeliverHeaders() {
	defer func() {
		hr.CloseSend()
	}()

	bftLogger.Debugf("[%s] Starting to deliver headers from endpoint: %s", hr.chainID, hr.endpoint)
	hr.setStarted()
	var errorStatusCounter int
	var statusCounter uint

	for !hr.isStopped() {
		msg, err := hr.client.Recv()
		if err != nil {
			bftLogger.Warningf("[%s] Receive error: %s", hr.chainID, err.Error())
			return
		}

		switch t := msg.Type.(type) {
		case *orderer.DeliverResponse_Status:
			if t.Status == common.Status_SUCCESS {
				bftLogger.Warningf("[%s] ERROR! Received success for a seek that should never complete", hr.chainID)
				return
			}
			if t.Status == common.Status_BAD_REQUEST || t.Status == common.Status_FORBIDDEN {
				bftLogger.Errorf("[%s] Got error %v", hr.chainID, t)
				errorStatusCounter++
				if errorStatusCounter > bftHeaderWrongStatusThreshold {
					bftLogger.Criticalf("[%s] Wrong statuses threshold passed, stopping bft header receiver", hr.chainID)
					return
				}
			} else {
				errorStatusCounter = 0
				bftLogger.Warningf("[%s] Got error %v", hr.chainID, t)
			}
			dur := backOffDuration(2.0, statusCounter, hr.minBackoffDelay, hr.maxBackoffDelay)
			bftLogger.Debugf("[%s] going to retry in: %s", hr.chainID, dur)
			backOffSleep(dur, hr.stopChan)
			statusCounter++

			hr.client.CloseSend()
			continue

		case *orderer.DeliverResponse_Block:
			errorStatusCounter = 0
			statusCounter = 0
			blockNum := t.Block.Header.Number

			// do not verify, just save for later, in case the block-receiver is suspected of censorship
			bftLogger.Debugf("[%s][%s] Saving block with header & metadata, blockNum = [%d], block = [%v]", hr.chainID, hr.endpoint, blockNum, t.Block)
			hr.mutex.Lock()
			hr.lastHeader = t.Block
			hr.lastHeaderTime = time.Now()
			hr.lastHeaderVerified = false
			hr.lastHeaderOK = false
			hr.mutex.Unlock()

		default:
			bftLogger.Warningf("[%s] Received unknown: %v", hr.chainID, t)
			return
		}
	}

	bftLogger.Debugf("[%s] Stopped to deliver headers from endpoint: %s", hr.chainID, hr.endpoint)
}

func (hr *bftHeaderReceiver) isStopped() bool {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()
	return hr.stop
}

func (hr *bftHeaderReceiver) isStarted() bool {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()
	return hr.started
}

func (hr *bftHeaderReceiver) setStarted() {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()
	hr.started = true
}

// CloseSend closes the client connection
func (hr *bftHeaderReceiver) CloseSend() error {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	if hr.stop {
		return nil
	}

	hr.stop = true
	hr.client.CloseSend()
	close(hr.stopChan)

	return nil
}

// LastBlockNum returns the last block number
func (hr *bftHeaderReceiver) LastBlockNum() (uint64, time.Time, error) {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()

	if hr.lastHeader == nil {
		return 0, time.Unix(0, 0), errors.New("Not found")
	}

	if !hr.lastHeaderVerified {
		hr.lastHeaderVerified = true
		hr.lastHeaderOK = true

		err := hr.msgCryptoVerifier.VerifyHeader(hr.chainID, hr.lastHeader)
		if err != nil {
			hr.lastHeaderOK = false
			bftLogger.Warningf("[%s][%s] Last block verification failed: %s", hr.chainID, hr.endpoint, err)
			return hr.lastHeader.Header.Number, hr.lastHeaderTime, errors.Wrapf(err, "Last block verification failed")
		}
	}

	if !hr.lastHeaderOK {
		bftLogger.Debugf("[%s][%s] Last block verification failed on previous invocation, cached result", hr.chainID, hr.endpoint)
		return hr.lastHeader.Header.Number, hr.lastHeaderTime, errors.New("Last block verification failed")
	}

	return hr.lastHeader.Header.Number, hr.lastHeaderTime, nil
}
