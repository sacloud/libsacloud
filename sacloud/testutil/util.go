package testutil

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
)

var testZone string
var apiCaller sacloud.APICaller
var httpTrace bool

var accTestOnce sync.Once
var accTestMu sync.Mutex

// SingletonAPICaller 環境変数からシングルトンAPICallerを作成する
func SingletonAPICaller() sacloud.APICaller {

	accTestMu.Lock()
	defer accTestMu.Unlock()
	accTestOnce.Do(func() {
		if !IsAccTest() {
			sacloud.DefaultStatePollInterval = 100 * time.Millisecond
			fake.SwitchFactoryFuncToFake()
		}

		if IsEnableTrace() || IsEnableAPITrace() {
			trace.AddClientFactoryHooks()
		}

		if IsEnableTrace() || IsEnableHTTPTrace() {
			httpTrace = true
		}

		//環境変数にトークン/シークレットがある場合のみテスト実施
		accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
		accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

		if accessToken == "" || accessTokenSecret == "" {
			log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
			os.Exit(0) // exit normal
		}
		client := sacloud.NewClient(accessToken, accessTokenSecret)
		client.DefaultTimeoutDuration = 30 * time.Minute
		client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
		client.AcceptLanguage = "en-US,en;q=0.9"

		client.RetryMax = 20
		client.RetryInterval = 3 * time.Second
		client.HTTPClient = &http.Client{
			Transport: &sacloud.RateLimitRoundTripper{RateLimitPerSec: 1},
		}
		if httpTrace {
			client.HTTPClient.Transport = &sacloud.TracingRoundTripper{
				Transport: client.HTTPClient.Transport,
			}
		}

		apiCaller = client
	})
	return apiCaller
}

// TestZone SAKURACLOUD_ZONE環境変数からテスト対象のゾーンを取得 デフォルトはtk1v
func TestZone() string {
	testZone := os.Getenv("SAKURACLOUD_ZONE")
	if testZone == "" {
		testZone = "tk1v"
	}
	return testZone
}

// IsAccTest TESTACC環境変数が指定されているか
func IsAccTest() bool {
	return os.Getenv("TESTACC") != ""
}

// IsEnableTrace SAKURACLOUD_TRACE環境変数が指定されているか
func IsEnableTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE") != ""
}

// IsEnableAPITrace SAKURACLOUD_TRACE_API環境変数が指定されているか
func IsEnableAPITrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_API") != ""
}

// IsEnableHTTPTrace SAKURACLOUD_TRACE_HTTP環境変数が指定されているか
func IsEnableHTTPTrace() bool {
	return os.Getenv("SAKURACLOUD_TRACE_HTTP") != ""
}
