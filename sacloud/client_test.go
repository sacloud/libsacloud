// Copyright 2016-2020 The Libsacloud Authors
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
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_Do_Backoff(t *testing.T) {
	h := &dummyHandler{
		responseCode: http.StatusServiceUnavailable,
	}
	dummyServer := httptest.NewServer(h)
	defer dummyServer.Close()

	client := &Client{
		RetryMax:     7,
		RetryWaitMin: 20 * time.Millisecond,
		RetryWaitMax: 640 * time.Millisecond,
	}
	client.Do(context.Background(), http.MethodGet, dummyServer.URL, nil) // nolint

	require.Len(t, h.called, client.RetryMax+1) // initial call + RetryMax
	var previous time.Time
	for i, ct := range h.called {
		if !previous.IsZero() {
			diff := ct.Sub(previous).Truncate(10 * time.Millisecond)
			t.Logf("backoff: retry-%d -> %0.2fs waited\n", i, diff.Seconds())
			require.True(t, client.RetryWaitMin <= diff && diff <= client.RetryWaitMax)
		}
		previous = ct
	}
}

type dummyHandler struct {
	called       []time.Time
	responseCode int
	mu           sync.Mutex
}

func (s *dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.called = append(s.called, time.Now())
	w.WriteHeader(s.responseCode)
}

func (s *dummyHandler) isRetried() bool {
	return len(s.called) > 1
}

func TestClient_Do_CheckRetryWithContext(t *testing.T) {
	client := &Client{RetryMax: 1, RetryWaitMin: 10 * time.Millisecond, RetryWaitMax: 10 * time.Millisecond}

	t.Run("context.Canceled", func(t *testing.T) {
		h := &dummyHandler{
			responseCode: http.StatusServiceUnavailable,
		}
		dummyServer := httptest.NewServer(h)
		defer dummyServer.Close()

		ctx, cancel := context.WithCancel(context.Background())
		// make ctx to Canceled
		cancel()

		client.Do(ctx, http.MethodGet, dummyServer.URL, nil) // nolint
		require.False(t, h.isRetried(), "don't retry when context was canceled")
	})

	t.Run("context.DeadlineExceeded", func(t *testing.T) {
		h := &dummyHandler{
			responseCode: http.StatusServiceUnavailable,
		}
		dummyServer := httptest.NewServer(h)
		defer dummyServer.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		// make ctx to DeadlineExceeded
		time.Sleep(time.Millisecond)

		client.Do(ctx, http.MethodGet, dummyServer.URL, nil) // nolint
		require.False(t, h.isRetried(), "don't retry when context exceeded deadline")
	})
}

func TestClient_RetryByStatusCode(t *testing.T) {
	cases := []struct {
		responseCode int
		shouldRetry  bool
	}{
		{responseCode: http.StatusOK, shouldRetry: false},
		{responseCode: http.StatusCreated, shouldRetry: false},
		{responseCode: http.StatusAccepted, shouldRetry: false},
		{responseCode: http.StatusNoContent, shouldRetry: false},
		{responseCode: http.StatusMovedPermanently, shouldRetry: false},
		{responseCode: http.StatusFound, shouldRetry: false},
		{responseCode: http.StatusBadRequest, shouldRetry: false},
		{responseCode: http.StatusUnauthorized, shouldRetry: false},
		{responseCode: http.StatusForbidden, shouldRetry: false},
		{responseCode: http.StatusNotFound, shouldRetry: false},
		{responseCode: http.StatusLocked, shouldRetry: true}, // Locked: 423
		{responseCode: http.StatusInternalServerError, shouldRetry: false},
		{responseCode: http.StatusBadGateway, shouldRetry: false},
		{responseCode: http.StatusServiceUnavailable, shouldRetry: true},
		{responseCode: http.StatusGatewayTimeout, shouldRetry: false},
	}

	client := &Client{RetryMax: 1, RetryWaitMin: 10 * time.Millisecond, RetryWaitMax: 10 * time.Millisecond}

	for _, tt := range cases {
		h := &dummyHandler{
			responseCode: tt.responseCode,
		}
		dummyServer := httptest.NewServer(h)
		client.Do(context.Background(), http.MethodGet, dummyServer.URL, nil) // nolint
		dummyServer.Close()

		require.Equal(t, tt.shouldRetry, h.isRetried(),
			"got unexpected retry status with status[%d]: expected:%t got:%t", h.responseCode, tt.shouldRetry, h.isRetried())
	}
}
