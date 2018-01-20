package api

import (
	"fmt"
	"github.com/sacloud/libsacloud"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var client *Client

func TestMain(m *testing.M) {
	//環境変数にトークン/シークレットがある場合のみテスト実施
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
		os.Exit(0) // exit normal
	}
	region := os.Getenv("SAKURACLOUD_REGION")
	if region == "" {
		region = "tk1v"
	}
	client = NewClient(accessToken, accessTokenSecret, region)
	client.DefaultTimeoutDuration = 30 * time.Minute
	client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
	client.AcceptLanguage = "en-US,en;q=0.9"

	ret := m.Run()
	os.Exit(ret)
}

func TestRetryableClient(t *testing.T) {

	t.Run("Retryable http client", func(t *testing.T) {
		called := 0
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called++
			if called < 3 {
				w.WriteHeader(503)
				return
			}
			w.Write([]byte(`ok`))
		}))
		defer s.Close()

		c := retryableHTTPClient{
			retryInterval: 3 * time.Second,
			retryMax:      2,
		}

		req, err := newRequest("GET", s.URL, nil)
		assert.NoError(t, err)

		start := time.Now()

		res, err := c.Do(req)
		defer res.Body.Close()

		end := time.Now()
		diff := end.Sub(start)

		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, 200)
		assert.Equal(t, called, 3)
		assert.True(t, diff > (c.retryInterval*time.Duration(c.retryMax)))
		t.Logf("Waited %f sec.\n", diff.Seconds())
	})

	t.Run("Retryable http client should fail", func(t *testing.T) {
		called := 0
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called++
			if called < 2 {
				w.WriteHeader(503)
				return
			}
			w.Write([]byte(`ok`))
		}))
		defer s.Close()

		c := retryableHTTPClient{
			retryInterval: 3 * time.Second,
			retryMax:      1,
		}

		req, err := newRequest("GET", s.URL, nil)
		assert.NoError(t, err)

		res, err := c.Do(req)
		assert.Nil(t, res)
		assert.Error(t, err)
		assert.Equal(t, 2, called)
	})
}
