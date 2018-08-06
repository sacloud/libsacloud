package setup

import (
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestRetryableSetup_Setup(t *testing.T) {

	t.Run("create", func(t *testing.T) {

		t.Run("no error", func(t *testing.T) {
			retryable := &RetryableSetup{
				Create: func() (sacloud.ResourceIDHolder, error) {
					return sacloud.NewResource(1), nil
				},
			}
			res, err := retryable.Setup()

			assert.NotNil(t, res)
			assert.NoError(t, err)
			_, ok := res.(*sacloud.Resource)
			assert.True(t, ok)
		})

		t.Run("error", func(t *testing.T) {
			retryable := &RetryableSetup{
				Create: func() (sacloud.ResourceIDHolder, error) {
					return nil, fmt.Errorf("error")
				},
			}
			res, err := retryable.Setup()

			assert.Nil(t, res)
			assert.Error(t, err)
		})

	})

	t.Run("Retry", func(t *testing.T) {

		t.Run("retry under max count", func(t *testing.T) {
			waiter := &dummyAsyncWaiter{
				reportErrorCount: 2,
			}
			retryable := &RetryableSetup{
				Create: func() (sacloud.ResourceIDHolder, error) {
					return sacloud.NewResource(1), nil
				},
				AsyncWaitForCopy: waiter.asyncWaitForCopy,
				Delete: func(id int64) error {
					return nil
				},
				RetryCount: 3,
			}

			res, err := retryable.Setup()

			assert.NotNil(t, res)
			assert.NoError(t, err)

		})

		t.Run("max retry count exceeded", func(t *testing.T) {
			waiter := &dummyAsyncWaiter{
				reportErrorCount: 3,
			}
			retryable := &RetryableSetup{
				Create: func() (sacloud.ResourceIDHolder, error) {
					return sacloud.NewResource(1), nil
				},
				AsyncWaitForCopy: waiter.asyncWaitForCopy,
				Delete: func(id int64) error {
					return nil
				},
				RetryCount: 3,
			}

			res, err := retryable.Setup()

			assert.Nil(t, res)
			assert.Error(t, err)
			_, ok := err.(MaxRetryCountExceededError)
			assert.True(t, ok)
		})

	})
}

type dummyAsyncWaiter struct {
	reportErrorCount int
	calledCount      int
}

func (d *dummyAsyncWaiter) asyncWaitForCopy(id int64) (chan interface{}, chan interface{}, chan error) {
	d.calledCount++

	compChan := make(chan interface{})
	progChan := make(chan interface{})
	errChan := make(chan error)

	// test scenario
	go func() {
		if d.calledCount <= d.reportErrorCount {
			progChan <- &dummyRetryCreated{isFailed: true}
			compChan <- &dummyRetryCreated{isFailed: true}
			return
		}

		progChan <- &dummyRetryCreated{isFailed: false}
		compChan <- &dummyRetryCreated{isFailed: false}
		return
	}()

	return compChan, progChan, errChan
}

type dummyRetryCreated struct {
	isFailed bool
}

func (d *dummyRetryCreated) IsFailed() bool {
	return d.isFailed
}
