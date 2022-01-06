// Copyright 2016-2022 The Libsacloud Authors
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

package wait

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSimpleStateWaiter(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		waiter := &SimpleStateWaiter{
			ReadStateFunc: func() (bool, error) {
				return false, nil
			},
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)
		require.Error(t, err)
		require.EqualError(t, err, "context deadline exceeded")
	})

	t.Run("parent context was canceled", func(t *testing.T) {
		waiter := &SimpleStateWaiter{
			ReadStateFunc: func() (bool, error) {
				return false, nil
			},
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

	t.Run("ReadStateFunc returns false", func(t *testing.T) {
		retry := 5
		read := 0
		waiter := &SimpleStateWaiter{
			ReadStateFunc: func() (bool, error) {
				read++
				if read < retry {
					return false, nil
				}
				return true, nil
			},
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.NoError(t, err)
		require.Equal(t, retry, read)
	})

	t.Run("ReadStateFunc got unexpected error", func(t *testing.T) {
		waiter := &SimpleStateWaiter{
			ReadStateFunc: func() (bool, error) {
				return false, errors.New("dummy")
			},
			Timeout:         100 * time.Millisecond,
			PollingInterval: 1 * time.Millisecond,
		}
		ctx := context.Background()
		_, err := waiter.WaitForState(ctx)

		require.Error(t, err)
		require.EqualError(t, err, "dummy")
	})
}
