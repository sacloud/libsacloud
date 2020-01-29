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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type dummyHandler struct {
	called       []time.Time
	responseCode int
}

func (s *dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.responseCode == http.StatusMovedPermanently {
		w.Header().Set("Location", "/index.html")
	}
	w.WriteHeader(s.responseCode)
	switch s.responseCode {
	case http.StatusMultipleChoices,
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusNotModified,
		http.StatusUseProxy,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect:
		s.responseCode = http.StatusOK
	default:
		s.called = append(s.called, time.Now())
	}
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
			"got unexpected retry status with status[%d]: expected:%t got:%t", tt.responseCode, tt.shouldRetry, h.isRetried())
	}
}
