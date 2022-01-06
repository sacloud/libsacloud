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

package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/require"
)

func TestNewCaller(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		caller := newCaller(&CallerOptions{
			AccessToken:       "token",
			AccessTokenSecret: "secret",
		})
		require.NotNil(t, caller)

		client, ok := caller.(*sacloud.Client)
		require.True(t, ok)

		expected := &sacloud.Client{
			AccessToken:       "token",
			AccessTokenSecret: "secret",
			UserAgent:         fmt.Sprintf("libsacloud/%s", libsacloud.Version),
			AcceptLanguage:    "",
			HTTPClient:        http.DefaultClient,
		}
		require.EqualValues(t, expected, client)
	})

	t.Run("custom", func(t *testing.T) {
		caller := newCaller(&CallerOptions{
			AccessToken:          "token",
			AccessTokenSecret:    "secret",
			APIRootURL:           "https://example.com",
			AcceptLanguage:       "ja-JP",
			RetryMax:             1,
			RetryWaitMax:         2,
			RetryWaitMin:         3,
			HTTPRequestTimeout:   4,
			HTTPRequestRateLimit: 5,
			UserAgent:            "dummy",
		})
		require.NotNil(t, caller)

		client, ok := caller.(*sacloud.Client)
		require.True(t, ok)

		expected := &sacloud.Client{
			AccessToken:       "token",
			AccessTokenSecret: "secret",
			UserAgent:         "dummy",
			AcceptLanguage:    "ja-JP",
			RetryMax:          1,
			RetryWaitMax:      2 * time.Second,
			RetryWaitMin:      3 * time.Second,
			HTTPClient:        http.DefaultClient,
		}
		require.EqualValues(t, expected, client)
		require.Equal(t, time.Second*4, client.HTTPClient.Timeout)
		require.EqualValues(t, 5, client.HTTPClient.Transport.(*sacloudhttp.RateLimitRoundTripper).RateLimitPerSec)
	})

	t.Run("multiple-call", func(t *testing.T) {
		caller1 := newCaller(&CallerOptions{
			AccessToken:       "1",
			AccessTokenSecret: "1",
		})
		require.NotNil(t, caller1)

		caller2 := newCaller(&CallerOptions{
			AccessToken:       "2",
			AccessTokenSecret: "2",
		})
		require.NotNil(t, caller2)

		require.NotEqual(t, caller1, caller2)
	})
}
