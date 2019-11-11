// Copyright 2016-2019 The Libsacloud Authors
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

package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPolling(t *testing.T) {

	t.Run("Normal: Should be no error", func(t *testing.T) {
		f := func() (bool, interface{}, error) {
			return true, "", nil
		}

		c, p, err := poll(f, 10*time.Second)
		for {
			select {
			case <-c:
				return
			case <-p:
				// noop
			case <-err:
				t.Fatal("Invalid chan state: err chan received")
			}
		}

	})

	t.Run("Use progress: Should be no error", func(t *testing.T) {
		counter := 0
		expect := 2

		f := func() (bool, interface{}, error) {
			if counter == expect {
				return true, "", nil
			}
			return false, "", nil
		}

		c, p, err := poll(f, 5*3*time.Second)
		for {
			select {
			case <-c:
				return
			case <-p:
				counter++
			case <-err:
				t.Fatal("Invalid chan state: err chan received")
			}
		}
	})

	t.Run("Timeout: Should return error", func(t *testing.T) {
		counter := 0
		f := func() (bool, interface{}, error) {
			return false, "", nil
		}

		c, p, err := poll(f, 5*time.Second)
		for {
			select {
			case <-c:
				t.Fatal("Invalid chan state: complete chan received")
			case <-p:
				counter++
			case <-err:
				assert.Equal(t, 1, counter)
				return
			}
		}
	})

	t.Run("Handler raise error: Should return error", func(t *testing.T) {
		f := func() (bool, interface{}, error) {
			return false, nil, fmt.Errorf("test")
		}

		c, p, err := poll(f, 5*time.Second)
		for {
			select {
			case <-c:
				t.Fatal("Invalid chan state: complete chan received")
			case <-p:
				t.Fatal("Invalid chan state: complete chan received")
			case <-err:
				return
			}
		}
	})
}

func TestBlockingPoll(t *testing.T) {

	t.Run("Normal: should be no error", func(t *testing.T) {
		f := func() (bool, interface{}, error) {
			return true, "", nil
		}
		done := false
		go func() {
			time.AfterFunc(10*time.Second, func() {
				if done {
					return
				}
				t.Fatal("Invalid timeout")
			})
		}()
		err := blockingPoll(f, 5*time.Second)
		assert.NoError(t, err)
		done = true
	})

	t.Run("Timeout: should return error", func(t *testing.T) {
		f := func() (bool, interface{}, error) {
			time.Sleep(1 * time.Minute)
			return true, "", nil
		}
		done := false
		go func() {
			time.AfterFunc(10*time.Second, func() {
				if done {
					return
				}
				t.Fatal("Invalid timeout")
			})
		}()
		err := blockingPoll(f, 1*time.Second)
		assert.Error(t, err)
		done = true
	})
}

type mockHasAvailableAndFailed struct {
	available bool
	failed    bool
}

func (m *mockHasAvailableAndFailed) IsAvailable() bool {
	return m.available
}
func (m *mockHasAvailableAndFailed) IsFailed() bool {
	return m.failed
}

func TestWaitingForAvailableFunc(t *testing.T) {

	t.Run("No retry: should no error", func(t *testing.T) {
		readFunc := func() (hasAvailable, error) {
			v := &mockHasAvailableAndFailed{
				available: true,
				failed:    false,
			}
			return v, nil
		}
		maxRetry := 1

		f := waitingForAvailableFunc(readFunc, maxRetry)
		err := blockingPoll(f, 5*time.Second)

		assert.NoError(t, err)

	})

	t.Run("Ignore error while maxRetry", func(t *testing.T) {
		counter := 0
		maxRetry := 2
		readFunc := func() (hasAvailable, error) {
			counter++
			if counter < maxRetry {
				return nil, fmt.Errorf("dummy readFunc error")
			}
			return &mockHasAvailableAndFailed{available: true}, nil
		}

		f := waitingForAvailableFunc(readFunc, maxRetry)
		err := blockingPoll(f, 1*time.Minute)

		assert.NoError(t, err)
	})

	t.Run("Raise error when counter become over maxRetry", func(t *testing.T) {
		counter := 0
		maxRetry := 2
		readFunc := func() (hasAvailable, error) {
			counter++
			if counter == 1 {
				return &mockHasAvailableAndFailed{available: false}, nil
			}
			return nil, fmt.Errorf("dummy readFunc error")
		}

		f := waitingForAvailableFunc(readFunc, maxRetry)
		err := blockingPoll(f, 1*time.Minute)

		assert.Error(t, err)
		assert.Equal(t, maxRetry, counter-1) // 一回正常値を返した分を引く
	})

	t.Run("Raise error when instance is failed", func(t *testing.T) {
		counter := 0
		maxRetry := 2
		readFunc := func() (hasAvailable, error) {
			counter++
			return &mockHasAvailableAndFailed{failed: true}, nil
		}

		f := waitingForAvailableFunc(readFunc, maxRetry)
		err := blockingPoll(f, 1*time.Minute)

		assert.Error(t, err)
		assert.True(t, counter < maxRetry)
	})
}
