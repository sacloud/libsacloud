// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sacloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/stretchr/testify/require"
)

type dummyState struct {
	state interface{}
	err   error
}

func testStateCheckFunc(target interface{}) (bool, error) {
	state, ok := target.(*dummyState)
	if !ok {
		return false, fmt.Errorf("got invalid state type: %+v", target)
	}
	return state.state != nil, state.err
}

func TestStatePollingWaiter_withStateCheckFunc(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				return &dummyState{}, nil
			},
			StateCheckFunc:  testStateCheckFunc,
			Timeout:         5 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)
		require.Error(t, err)
		require.EqualError(t, err, "context deadline exceeded")
	})

	t.Run("parent context was canceled", func(t *testing.T) {
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				return &dummyState{}, nil
			},
			StateCheckFunc:  testStateCheckFunc,
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		_, err := waiter.WaitForState(ctx)
		go func() {
			time.Sleep(5 * time.Millisecond)
			cancel()
		}()

		require.Error(t, err)
		require.EqualError(t, err, "context deadline exceeded")
	})

	t.Run("ReadFunc got 404", func(t *testing.T) {
		retry := 5
		read := 0
		waiter := &StatePollingWaiter{
			NotFoundRetry: 10,
			ReadFunc: func() (interface{}, error) {
				read++
				if read < retry {
					return nil, &apiError{responseCode: http.StatusNotFound}
				}
				return &dummyState{state: "done"}, nil
			},
			StateCheckFunc:  testStateCheckFunc,
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.NoError(t, err)
		require.Equal(t, retry, read)
	})

	t.Run("404 errors exceeded maximum", func(t *testing.T) {
		retry := 5
		read := 0
		waiter := &StatePollingWaiter{
			NotFoundRetry: 2,
			ReadFunc: func() (interface{}, error) {
				read++
				if read < retry {
					return nil, &apiError{responseCode: http.StatusNotFound}
				}
				return &dummyState{state: "done"}, nil
			},
			StateCheckFunc:  testStateCheckFunc,
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.Error(t, err)
		require.True(t, IsNotFoundError(err))
		require.Equal(t, waiter.NotFoundRetry+1, read)
	})

	t.Run("ReadFunc got unexpected error", func(t *testing.T) {
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				return &dummyState{}, errors.New("dummy")
			},
			StateCheckFunc:  testStateCheckFunc,
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.Error(t, err)
		require.EqualError(t, err, "dummy")
	})
}

type dummyInstanceStatus struct {
	status       string
	availability string
}

func (d *dummyInstanceStatus) GetInstanceStatus() types.EServerInstanceStatus {
	return types.EServerInstanceStatus(d.status)
}

func (d *dummyInstanceStatus) SetInstanceStatus(s types.EServerInstanceStatus) {
	d.status = string(s)
}
func (d *dummyInstanceStatus) GetAvailability() types.EAvailability {
	return types.EAvailability(d.availability)
}
func (d *dummyInstanceStatus) SetAvailability(a types.EAvailability) {
	d.availability = string(a)
}

func TestStatePollingWaiter_withTargetStates(t *testing.T) {
	t.Run("ReadFunc got unknown state with RaiseErrorWithUnknownState=false", func(t *testing.T) {
		counter := 0
		max := 3
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				counter++
				status := ""
				if counter > max {
					status = string(types.ServerInstanceStatuses.Up)
				}
				return &dummyInstanceStatus{status: status}, nil
			},
			TargetInstanceStatus: []types.EServerInstanceStatus{types.ServerInstanceStatuses.Up},
			Timeout:              100 * time.Millisecond,
			PollingInterval:      1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.NoError(t, err)
	})
	t.Run("ReadFunc got unknown state with RaiseErrorWithUnknownState=true", func(t *testing.T) {
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				return &dummyInstanceStatus{status: "unknown-instance-status"}, nil
			},
			TargetInstanceStatus:       []types.EServerInstanceStatus{types.ServerInstanceStatuses.Up},
			Timeout:                    100 * time.Millisecond,
			PollingInterval:            1 * time.Millisecond,
			RaiseErrorWithUnknownState: true,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.Error(t, err)
		require.EqualError(t, err, `got unexpected value of InstanceState: got "unknown-instance-status"`)
	})

	t.Run("ReadFunc got unknown availability with RaiseErrorWithUnknownState=false", func(t *testing.T) {
		counter := 0
		max := 3
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				counter++
				availability := ""
				if counter > max {
					availability = string(types.Availabilities.Available)
				}
				return &dummyInstanceStatus{availability: availability}, nil
			},
			TargetAvailability: []types.EAvailability{types.Availabilities.Available},
			Timeout:            100 * time.Millisecond,
			PollingInterval:    1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.NoError(t, err)
	})

	t.Run("ReadFunc got unknown availability with RaiseErrorWithUnknownState=true", func(t *testing.T) {
		waiter := &StatePollingWaiter{
			ReadFunc: func() (interface{}, error) {
				return &dummyInstanceStatus{availability: "unknown-availability"}, nil
			},
			TargetAvailability:         []types.EAvailability{types.Availabilities.Available},
			Timeout:                    100 * time.Millisecond,
			PollingInterval:            1 * time.Millisecond,
			RaiseErrorWithUnknownState: true,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.Error(t, err)
		require.EqualError(t, err, `got unexpected value of Availability: got "unknown-availability"`)
	})
}
