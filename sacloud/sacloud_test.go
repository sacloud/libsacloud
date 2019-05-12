package sacloud

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud-v2"
)

var testZone string
var apiCaller APICaller

var accTestOnce sync.Once
var accTestMu sync.Mutex

func singletonAPICaller() APICaller {
	accTestMu.Lock()
	defer accTestMu.Unlock()
	accTestOnce.Do(func() {
		//環境変数にトークン/シークレットがある場合のみテスト実施
		accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
		accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

		if accessToken == "" || accessTokenSecret == "" {
			log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
			os.Exit(0) // exit normal
		}
		client := NewClient(accessToken, accessTokenSecret)
		client.DefaultTimeoutDuration = 30 * time.Minute
		client.UserAgent = fmt.Sprintf("test-libsacloud-v2/%s", libsacloud.Version)
		client.AcceptLanguage = "en-US,en;q=0.9"

		apiCaller = client
	})
	return apiCaller
}

func TestMain(m *testing.M) {
	testZone = os.Getenv("SAKURACLOUD_ZONE")
	if testZone == "" {
		testZone = "tk1v"
	}
	ret := m.Run()
	os.Exit(ret)
}
