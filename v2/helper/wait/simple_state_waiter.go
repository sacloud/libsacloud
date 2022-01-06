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
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ByFunc デフォルトのパラメータでSimpleStateWaiterを作成して返す
func ByFunc(readStateFunc ReadStateFunc) sacloud.StateWaiter {
	return &SimpleStateWaiter{ReadStateFunc: readStateFunc}
}

type ReadStateFunc func() (bool, error)

// SimpleStateWaiter シンプルな待ち処理のためのsacloud.StateWaiterの実装
//
// sacloud.StatePollingWaiterをラップし、シンプルなfuncのみで待つべきかを判定する
type SimpleStateWaiter struct {
	// ReadStateFunc 待つべきかの判定func
	// trueかつerrorが空の場合は待ち処理を完了させる
	ReadStateFunc ReadStateFunc

	// Timeout タイムアウト
	Timeout time.Duration

	// PollingInterval ポーリング間隔
	PollingInterval time.Duration
}

func (s *SimpleStateWaiter) waiter() sacloud.StateWaiter {
	return &sacloud.StatePollingWaiter{
		ReadFunc: func() (interface{}, error) {
			result, err := s.ReadStateFunc()
			if err != nil {
				return false, err
			}
			return &fakeState{available: result}, nil
		},
		TargetAvailability: []types.EAvailability{
			types.Availabilities.Available,
		},
		PendingAvailability: []types.EAvailability{
			types.Availabilities.Unknown,
		},

		PollingInterval: s.PollingInterval,
		Timeout:         s.Timeout,
	}
}

// WaitForState sacloud.StateWaiterの実装
func (s *SimpleStateWaiter) WaitForState(ctx context.Context) (interface{}, error) {
	return s.waiter().WaitForState(ctx)
}

// AsyncWaitForState sacloud.StateWaiterの実装
func (s *SimpleStateWaiter) AsyncWaitForState(ctx context.Context) (compCh <-chan interface{}, progressCh <-chan interface{}, errorCh <-chan error) {
	return s.waiter().AsyncWaitForState(ctx)
}

// SetPollingTimeout sacloud.StateWaiterの実装
func (s *SimpleStateWaiter) SetPollingTimeout(d time.Duration) {
	s.waiter().SetPollingTimeout(d)
}

// SetPollingInterval sacloud.StateWaiterの実装
func (s *SimpleStateWaiter) SetPollingInterval(d time.Duration) {
	s.waiter().SetPollingInterval(d)
}

type fakeState struct {
	available bool
}

// GetAvailability accessor.Availabilityの実装
func (f *fakeState) GetAvailability() types.EAvailability {
	if f.available {
		return types.Availabilities.Available
	}
	return types.Availabilities.Unknown
}

// SetAvailability accessor.Availabilityの実装
func (f *fakeState) SetAvailability(v types.EAvailability) {
	f.available = v == types.Availabilities.Available
}
