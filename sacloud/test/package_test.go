package test

import (
	"context"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestMain(m *testing.M) {
	testZone = testutil.TestZone()

	ret := m.Run()

	skipCleanup := os.Getenv("SKIP_CLEANUP")
	if skipCleanup == "" {
		testutil.CleanupTestResources(context.TODO(), singletonAPICaller(), testutil.TestResourcePrefix)
	}

	os.Exit(ret)
}

var testZone string
var testIconID = types.ID(112901627749) // テスト用のアイコンID(shared icon)

func singletonAPICaller() sacloud.APICaller {
	return testutil.SingletonAPICaller()
}

func isAccTest() bool {
	return testutil.IsAccTest()
}

func setupSwitchFunc(targetResource string, dests ...accessor.SwitchID) func(*testutil.CRUDTestContext, sacloud.APICaller) error {
	return func(testContext *testutil.CRUDTestContext, caller sacloud.APICaller) error {
		swClient := sacloud.NewSwitchOp(caller)
		sw, err := swClient.Create(context.Background(), testZone, &sacloud.SwitchCreateRequest{
			Name: testutil.ResourceName("switch-for-" + targetResource),
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

func cleanupSwitchFunc(targetResource string) func(*testutil.CRUDTestContext, sacloud.APICaller) error {
	return func(testContext *testutil.CRUDTestContext, caller sacloud.APICaller) error {
		switchID, ok := testContext.Values[targetResource+"/switch"]
		if !ok {
			return nil
		}

		swClient := sacloud.NewSwitchOp(caller)
		return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
	}
}
