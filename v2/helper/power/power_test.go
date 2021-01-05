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

package power

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestPowerHandler(t *testing.T) {
	t.Parallel()

	defaultInterval := sacloud.DefaultStatePollingInterval

	sacloud.DefaultStatePollingInterval = 10 * time.Millisecond
	BootRetrySpan = time.Millisecond
	ShutdownRetrySpan = time.Millisecond
	defer func() {
		sacloud.DefaultStatePollingInterval = defaultInterval
		BootRetrySpan = DefaultBootRetrySpan
		ShutdownRetrySpan = DefaultShutdownRetrySpan
	}()

	ctx := context.Background()
	t.Run("boot", func(t *testing.T) {
		handler := &dummyPowerHandler{
			ignoreBootCount: 3,
			instanceStatus:  types.ServerInstanceStatuses.Down,
		}
		err := boot(ctx, handler)
		require.NoError(t, err)
		require.Equal(t, handler.ignoreBootCount+1, handler.bootCount)
	})
	t.Run("shutdown", func(t *testing.T) {
		handler := &dummyPowerHandler{
			ignoreShutdownCount: 3,
			instanceStatus:      types.ServerInstanceStatuses.Up,
		}
		err := shutdown(ctx, handler, true)
		require.NoError(t, err)
		require.Equal(t, handler.ignoreShutdownCount+1, handler.shutdownCount)
	})
}

type dummyPowerHandler struct {
	bootCount           int
	shutdownCount       int
	ignoreBootCount     int
	ignoreShutdownCount int
	instanceStatus      types.EServerInstanceStatus

	mu sync.Mutex
}

func (d *dummyPowerHandler) boot() error {
	d.bootCount++
	if d.bootCount > d.ignoreBootCount {
		go d.toggleInstanceStatus()
		return sacloud.NewAPIError("DUMMY", nil, "dummy", http.StatusConflict, nil)
	}
	return nil
}
func (d *dummyPowerHandler) shutdown(force bool) error {
	d.shutdownCount++
	if d.shutdownCount > d.ignoreShutdownCount {
		go d.toggleInstanceStatus()
		return sacloud.NewAPIError("DUMMY", nil, "dummy", http.StatusConflict, nil)
	}
	return nil
}

func (d *dummyPowerHandler) read() (interface{}, error) {
	return d, nil
}

func (d *dummyPowerHandler) toggleInstanceStatus() {
	time.Sleep(100 * time.Millisecond)

	d.mu.Lock()
	defer d.mu.Unlock()

	switch d.instanceStatus {
	case types.ServerInstanceStatuses.Up:
		d.instanceStatus = types.ServerInstanceStatuses.Down
	case types.ServerInstanceStatuses.Down:
		d.instanceStatus = types.ServerInstanceStatuses.Up
	}
}

// GetInstanceStatus .
func (d *dummyPowerHandler) GetInstanceStatus() types.EServerInstanceStatus {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.instanceStatus
}

// SetInstanceStatus .
func (d *dummyPowerHandler) SetInstanceStatus(v types.EServerInstanceStatus) {
	d.instanceStatus = v
}
