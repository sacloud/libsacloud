package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var testZone string
var apiCaller sacloud.APICaller

var accTestOnce sync.Once
var accTestMu sync.Mutex

func singletonAPICaller() sacloud.APICaller {
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
		client := sacloud.NewClient(accessToken, accessTokenSecret)
		client.DefaultTimeoutDuration = 30 * time.Minute
		client.UserAgent = fmt.Sprintf("test-libsacloud/%s", libsacloud.Version)
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

	if !isAccTest() {
		sacloud.DefaultStatePollInterval = 100 * time.Millisecond
		fake.SwitchFactoryFuncToFake()
	}

	ret := m.Run()

	skipCleanup := os.Getenv("SKIP_CLEANUP")
	if skipCleanup != "" {
		// TODO クリーンアップ処理
	}

	os.Exit(ret)
}

func compositeAPIFunc(funcs ...func(*CRUDTestContext, sacloud.APICaller) error) func(*CRUDTestContext, sacloud.APICaller) error {
	return func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
		for _, f := range funcs {
			if err := f(testContext, caller); err != nil {
				return err
			}
		}
		return nil
	}
}

func setupSwitchFunc(targetResource string, dests ...accessor.SwitchID) func(*CRUDTestContext, sacloud.APICaller) error {
	return func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
		swClient := sacloud.NewSwitchOp(caller)
		swCreateResult, err := swClient.Create(context.Background(), testZone, &sacloud.SwitchCreateRequest{
			Name: "libsacloud-switch-for-" + targetResource,
		})
		if err != nil {
			return err
		}
		sw := swCreateResult.Switch

		testContext.Values[targetResource+"/switch"] = sw.ID
		for _, dest := range dests {
			dest.SetSwitchID(sw.ID)
		}
		return nil
	}
}

func cleanupSwitchFunc(targetResource string) func(*CRUDTestContext, sacloud.APICaller) error {
	return func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
		switchID, ok := testContext.Values[targetResource+"/switch"]
		if !ok {
			return nil
		}

		swClient := sacloud.NewSwitchOp(caller)
		return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
	}
}
