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

package query

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testChecker struct {
	_called int
	f       func(called int) (bool, error)
	mu      sync.Mutex
}

func (c *testChecker) isExists() (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c._called++
	return c.f(c._called)
}

func (c *testChecker) called() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c._called
}

func TestWaitWhileReferenced_called(t *testing.T) {
	checker := &testChecker{
		f: func(called int) (bool, error) {
			if called == 2 {
				return false, nil
			}
			return true, nil
		},
	}
	err := waitWhileReferenced(context.Background(), CheckReferencedOption{Tick: time.Millisecond, Timeout: time.Second}, checker.isExists)
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, 2, checker.called())
}

func TestWaitWhileReferenced_contextCanceled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	checker := &testChecker{
		f: func(c int) (bool, error) {
			return true, nil
		},
	}

	err := waitWhileReferenced(ctx, CheckReferencedOption{Tick: time.Millisecond, Timeout: time.Second}, checker.isExists)
	if err.Error() != "context deadline exceeded" {
		t.Errorf("unexpected error: expected: %s, actual: %s", "context deadline exceeded", err)
	}

	if checker.called() < 2 {
		t.Error("check func was not called when retry")
	}
}

func TestWaitWhileReferenced_checkFuncReturnsError(t *testing.T) {
	checker := &testChecker{
		f: func(called int) (bool, error) {
			if called == 2 {
				return false, errors.New("dummy")
			}
			return true, nil
		},
	}

	err := waitWhileReferenced(context.Background(), CheckReferencedOption{Tick: time.Millisecond, Timeout: time.Second}, checker.isExists)
	require.Equal(t, errors.New("dummy"), err)
	require.Equal(t, 2, checker.called())
}

func TestWaitWhileReferenced_optionTimeout(t *testing.T) {
	checker := &testChecker{
		f: func(called int) (bool, error) {
			return true, nil
		},
	}

	err := waitWhileReferenced(context.Background(), CheckReferencedOption{Tick: time.Millisecond, Timeout: 10 * time.Millisecond}, checker.isExists)
	if err.Error() != "context deadline exceeded" {
		t.Errorf("unexpected error: expected: %s, actual: %s", "context deadline exceeded", err)
	}

	if checker.called() < 2 {
		t.Error("check func was not called when retry")
	}
}
