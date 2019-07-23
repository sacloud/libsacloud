package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestSwitchOp_CRUD(t *testing.T) {
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
			{
				Func: testSwitchUpdateToMin,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateSwitchToMinExpected,
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
		IconID:         testIconID,
	}
	updateSwitchExpected = &sacloud.Switch{
		Name:           updateSwitchParam.Name,
		Description:    updateSwitchParam.Description,
		Tags:           updateSwitchParam.Tags,
		DefaultRoute:   updateSwitchParam.DefaultRoute,
		NetworkMaskLen: updateSwitchParam.NetworkMaskLen,
		Scope:          createSwitchExpected.Scope,
		IconID:         testIconID,
	}
	updateSwitchToMinParam = &sacloud.SwitchUpdateRequest{
		Name: "libsacloud-switch-to-min",
	}
	updateSwitchToMinExpected = &sacloud.Switch{
		Name:  updateSwitchToMinParam.Name,
		Scope: createSwitchExpected.Scope,
	}
)

func testSwitchCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Create(ctx, testZone, createSwitchParam)
}

func testSwitchRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testSwitchUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateSwitchParam)
}

func testSwitchUpdateToMin(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateSwitchToMinParam)
}

func testSwitchDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSwitchOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestSwitchOp_BridgeConnection(t *testing.T) {
	caller := singletonAPICaller()

	swOp := sacloud.NewSwitchOp(caller)
	bridgeOp := sacloud.NewBridgeOp(caller)

	var bridgeID types.ID

	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
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
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
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
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					// connect to bridge
					if err := swOp.ConnectToBridge(ctx, testZone, ctx.ID, bridgeID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, ctx.ID)
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

					return DoAsserts(
						AssertEqualFunc(t, sw.ID, bridge.SwitchInZone.ID, "Bridge.SwitchInZone.ID"),
						AssertLenFunc(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo"),
					)
				},
			},
			// disconnect
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					if err := swOp.DisconnectFromBridge(ctx, testZone, ctx.ID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t TestT, ctx *CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, ctx.ID)
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

					return DoAsserts(
						AssertNilFunc(t, bridge.SwitchInZone, "Bridge.SwitchInZone"),
						AssertLenFunc(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo"),
					)
				},
			},
			// bridge delete
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
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
