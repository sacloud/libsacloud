package api

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/sacloud/libsacloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			if called < 3 {
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

func TestCustomHTTPClient(t *testing.T) {
	timeout := 10 * time.Millisecond
	response := `{"data":"ok"}`
	testServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timeout)
		fmt.Fprintf(w, response)
	}))
	defer testServer.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	http.DefaultClient.Transport = tr

	t.Run("Use http.DefaultClient: should success", func(t *testing.T) {
		http.DefaultClient.Timeout = 15 * time.Millisecond

		client := NewClient("token", "secret", "is1a")
		data, err := client.newRequest(http.MethodGet, testServer.URL, nil)
		require.Equal(t, response, string(data))
		require.NoError(t, err)
	})

	t.Run("Use http.DefaultClient: should timeout", func(t *testing.T) {
		http.DefaultClient.Timeout = 5 * time.Millisecond
		client := NewClient("token", "secret", "is1a")
		_, err := client.newRequest(http.MethodGet, testServer.URL, nil)
		require.Error(t, err)
		require.EqualError(t, err, "net/http: request canceled (Client.Timeout exceeded while awaiting headers)")
	})

	t.Run("Use custom http.Client", func(t *testing.T) {
		http.DefaultClient.Timeout = 5 * time.Millisecond
		customClient := &http.Client{
			Timeout:   15 * time.Millisecond,
			Transport: tr,
		}

		client := NewClient("token", "secret", "is1a")
		client.HTTPClient = customClient

		data, err := client.newRequest(http.MethodGet, testServer.URL, nil)
		require.Equal(t, response, string(data))
		require.NoError(t, err)
	})
}
