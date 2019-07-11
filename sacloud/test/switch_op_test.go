package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestSwitchOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testSwitchCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testSwitchRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testSwitchUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateSwitchExpected,
					IgnoreFields: ignoreSwitchFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testSwitchDelete,
		},
	})
}

var (
	ignoreSwitchFields = []string{
		"ID",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
	}
	createSwitchParam = &sacloud.SwitchCreateRequest{
		Name:           "libsacloud-switch",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		DefaultRoute:   "192.168.0.1",
		NetworkMaskLen: 24,
	}
	createSwitchExpected = &sacloud.Switch{
		Name:           createSwitchParam.Name,
		Description:    createSwitchParam.Description,
		Tags:           createSwitchParam.Tags,
		DefaultRoute:   createSwitchParam.DefaultRoute,
		NetworkMaskLen: createSwitchParam.NetworkMaskLen,
		Scope:          types.Scopes.User,
	}
	updateSwitchParam = &sacloud.SwitchUpdateRequest{
		Name:           "libsacloud-switch-upd",
		Tags:           []string{"tag1-upd", "tag2-upd"},
		Description:    "desc-upd",
		DefaultRoute:   "192.168.0.2",
		NetworkMaskLen: 28,
	}
	updateSwitchExpected = &sacloud.Switch{
		Name:           updateSwitchParam.Name,
		Description:    updateSwitchParam.Description,
		Tags:           updateSwitchParam.Tags,
		DefaultRoute:   updateSwitchParam.DefaultRoute,
		NetworkMaskLen: updateSwitchParam.NetworkMaskLen,
		Scope:          createSwitchExpected.Scope,
	}
)

func testSwitchCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Create(context.Background(), testZone, createSwitchParam)
}

func testSwitchRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testSwitchUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateSwitchParam)
}

func testSwitchDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSwitchOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestSwitchOp_BridgeConnection(t *testing.T) {
	ctx := context.Background()
	caller := singletonAPICaller()

	swOp := sacloud.NewSwitchOp(caller)
	bridgeOp := sacloud.NewBridgeOp(caller)

	var bridgeID types.ID

	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
				return swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
					Name: "libsacloud-switch-for-bridge",
				})
			},
		},
		Read: &CRUDTestFunc{
			Func: testSwitchRead,
		},
		Updates: []*CRUDTestFunc{
			// bridge create and connect
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					bridge, err := bridgeOp.Create(ctx, testZone, &sacloud.BridgeCreateRequest{
						Name: "libsacloud-bridge",
					})
					if err != nil {
						return nil, err
					}
					bridgeID = bridge.ID
					return bridge, nil
				},
				SkipExtractID: true,
			},
			// connect
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					// connect to bridge
					if err := swOp.ConnectToBridge(ctx, testZone, testContext.ID, bridgeID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, testContext.ID)
					if err != nil {
						return err
					}
					if err := AssertEqual(t, bridgeID, sw.BridgeID, "Switch.BridgeID"); err != nil {
						return err
					}

					bridge, err := bridgeOp.Read(ctx, testZone, bridgeID)
					if err != nil {
						return err
					}

					if err := DoAsserts(
						func() error { return AssertEqual(t, sw.ID, bridge.SwitchInZone.ID, "Bridge.SwitchInZone.ID") },
						func() error { return AssertLen(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo") },
					); err != nil {
						return err
					}
					return nil
				},
			},
			// disconnect
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					if err := swOp.DisconnectFromBridge(ctx, testZone, testContext.ID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t TestT, testContext *CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, testContext.ID)
					if err != nil {
						return err
					}
					if err := AssertTrue(t, sw.BridgeID.IsEmpty(), "Switch.BridgeID"); err != nil {
						return err
					}

					bridge, err := bridgeOp.Read(ctx, testZone, bridgeID)
					if err != nil {
						return err
					}

					if err := DoAsserts(
						func() error { return AssertNil(t, bridge.SwitchInZone, "Bridge.SwitchInZone") },
						func() error { return AssertLen(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo") },
					); err != nil {
						return err
					}
					return nil
				},
			},
			// bridge delete
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					if err := bridgeOp.Delete(ctx, testZone, bridgeID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testSwitchDelete,
		},
	})
}
