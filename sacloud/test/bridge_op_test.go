package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestBridgeOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testBridgeCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createBridgeExpected,
				IgnoreFields: ignoreBridgeFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testBridgeRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createBridgeExpected,
				IgnoreFields: ignoreBridgeFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testBridgeUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateBridgeExpected,
				IgnoreFields: ignoreBridgeFields,
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testBridgeDelete,
		},
	})
}

var (
	ignoreBridgeFields = []string{
		"ID",
		"CreatedAt",
		"Region",
		"SwitchInZone",
		"BridgeInfo",
	}

	createBridgeParam = &sacloud.BridgeCreateRequest{
		Name:        "libsacloud-bridge",
		Description: "desc",
	}
	createBridgeExpected = &sacloud.Bridge{
		Name:        createBridgeParam.Name,
		Description: createBridgeParam.Description,
	}
	updateBridgeParam = &sacloud.BridgeUpdateRequest{
		Name:        "libsacloud-bridge-upd",
		Description: "desc-upd",
	}
	updateBridgeExpected = &sacloud.Bridge{
		Name:        updateBridgeParam.Name,
		Description: updateBridgeParam.Description,
	}
)

func testBridgeCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	res, err := client.Create(context.Background(), testZone, createBridgeParam)
	if err != nil {
		return nil, err
	}
	return res.Bridge, nil
}

func testBridgeRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.Bridge, nil
}

func testBridgeUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewBridgeOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateBridgeParam)
	if err != nil {
		return nil, err
	}
	return res.Bridge, nil
}

func testBridgeDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewBridgeOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
