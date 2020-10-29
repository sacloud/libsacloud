package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"

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
			RetryMax:          sacloud.APIDefaultRetryMax,
			RetryWaitMin:      sacloud.APIDefaultRetryWaitMin,
			RetryWaitMax:      sacloud.APIDefaultRetryWaitMax,
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
		require.EqualValues(t, 5, client.HTTPClient.Transport.(*sacloud.RateLimitRoundTripper).RateLimitPerSec)
	})
}
