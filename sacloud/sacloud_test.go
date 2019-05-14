package sacloud

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud-v2"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
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

		client.RetryMax = 20
		client.RetryInterval = 3 * time.Second

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

func compositeAPIFunc(funcs ...func(*CRUDTestContext, APICaller) error) func(*CRUDTestContext, APICaller) error {
	return func(testContext *CRUDTestContext, caller APICaller) error {
		for _, f := range funcs {
			if err := f(testContext, caller); err != nil {
				return err
			}
		}
		return nil
	}
}

func setupSwitchFunc(targetResource string, dests ...switchIDAccessor) func(*CRUDTestContext, APICaller) error {
	return func(testContext *CRUDTestContext, caller APICaller) error {
		swClient := NewSwitchOp(caller)
		sw, err := swClient.Create(context.Background(), testZone, &SwitchCreateRequest{
			Name: "libsacloud-v2-switch-for-" + targetResource,
		})
		if err != nil {
			return err
		}

		testContext.Values[targetResource+"/switch"] = sw.ID
		for _, dest := range dests {
			dest.SetSwitchID(sw.ID)
		}
		return nil
	}
}

func cleanupSwitchFunc(targetResource string) func(*CRUDTestContext, APICaller) error {
	return func(testContext *CRUDTestContext, caller APICaller) error {
		switchID, ok := testContext.Values[targetResource+"/switch"]
		if !ok {
			return nil
		}

		swClient := NewSwitchOp(caller)
		return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
	}
}
