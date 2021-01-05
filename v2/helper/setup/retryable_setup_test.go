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

package setup

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

type dummyIDAccessor struct {
	id types.ID
}

func (d *dummyIDAccessor) GetID() types.ID {
	return d.id
}

func (d *dummyIDAccessor) SetID(id types.ID) {
	d.id = id
}

type dummyAvailabilityAccessor struct {
	available types.EAvailability
}

func (d *dummyAvailabilityAccessor) GetAvailability() types.EAvailability {
	return d.available
}

func (d *dummyAvailabilityAccessor) SetAvailability(a types.EAvailability) {
	d.available = a
}

func TestRetryableSetup_Setup(t *testing.T) {
	ctx := context.Background()
	zone := "tk1v"

	t.Run("create", func(t *testing.T) {
		t.Run("no error", func(t *testing.T) {
			retryable := &RetryableSetup{
				Create: func(context.Context, string) (id accessor.ID, e error) {
					return &dummyIDAccessor{id: 1}, nil
				},
				Read: func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
					return &dummyIDAccessor{id: 1}, nil
				},
				ProvisioningRetryInterval: time.Millisecond,
				DeleteRetryInterval:       time.Millisecond,
				PollingInterval:           time.Millisecond,
			}
			res, err := retryable.Setup(ctx, zone)

			require.NotNil(t, res)
			require.NoError(t, err)
			_, ok := res.(accessor.ID)
			require.True(t, ok)
		})

		t.Run("error", func(t *testing.T) {
			retryable := &RetryableSetup{
				Create: func(context.Context, string) (accessor.ID, error) {
					return nil, fmt.Errorf("error")
				},
				Read: func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
					return &dummyIDAccessor{id: 1}, nil
				},
				ProvisioningRetryInterval: time.Millisecond,
				DeleteRetryInterval:       time.Millisecond,
				PollingInterval:           time.Millisecond,
			}
			res, err := retryable.Setup(ctx, zone)

			require.Nil(t, res)
			require.Error(t, err)
		})
	})
	//
	t.Run("Retry", func(t *testing.T) {
		t.Run("retry under max count", func(t *testing.T) {
			// リトライ3回、ReadFuncは3回呼ばれるまではFailedが返る(4回目以降はAvailable)
			retryable := &RetryableSetup{
				Create: func(context.Context, string) (id accessor.ID, e error) {
					return &dummyIDAccessor{id: 1}, nil
				},
				IsWaitForCopy: true,
				Delete: func(context.Context, string, types.ID) error {
					return nil
				},
				RetryCount: 3,
				Read: withErrorReadFunc(func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
					return &dummyIDAccessor{id: 1}, nil
				}, 3),
				ProvisioningRetryInterval: time.Millisecond,
				DeleteRetryInterval:       time.Millisecond,
				PollingInterval:           time.Millisecond,
			}

			res, err := retryable.Setup(ctx, zone)

			require.NotNil(t, res)
			require.NoError(t, err)
		})

		t.Run("max retry count exceeded", func(t *testing.T) {
			retryable := &RetryableSetup{
				Create: func(context.Context, string) (id accessor.ID, e error) {
					return &dummyIDAccessor{id: 1}, nil
				},
				IsWaitForCopy: true,
				Delete: func(context.Context, string, types.ID) error {
					return nil
				},
				RetryCount: 3,
				Read: withErrorReadFunc(func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
					return &dummyIDAccessor{id: 1}, nil
				}, 5),
				ProvisioningRetryInterval: time.Millisecond,
				DeleteRetryInterval:       time.Millisecond,
				PollingInterval:           time.Millisecond,
			}

			res, err := retryable.Setup(ctx, zone)

			require.Nil(t, res)
			require.Error(t, err)
			_, ok := err.(MaxRetryCountExceededError)
			require.True(t, ok)
		})
	})
}

func withErrorReadFunc(readFunc ReadFunc, errCount int) ReadFunc {
	maxErr := errCount
	return func(ctx context.Context, zone string, id types.ID) (interface{}, error) {
		maxErr--
		if maxErr <= 0 {
			return &dummyAvailabilityAccessor{available: types.Availabilities.Available}, nil
		}
		return &dummyAvailabilityAccessor{available: types.Availabilities.Failed}, nil
	}
}
