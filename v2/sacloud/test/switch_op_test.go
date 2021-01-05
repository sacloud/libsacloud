// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestSwitchOp_CRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testSwitchCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testSwitchRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testSwitchUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSwitchExpected,
					IgnoreFields: ignoreSwitchFields,
				}),
			},
			{
				Func: testSwitchUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateSwitchToMinExpected,
					IgnoreFields: ignoreSwitchFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
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
		Name:           testutil.ResourceName("switch"),
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
		Name:           testutil.ResourceName("switch-upd"),
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
		Name: testutil.ResourceName("switch-to-min"),
	}
	updateSwitchToMinExpected = &sacloud.Switch{
		Name:  updateSwitchToMinParam.Name,
		Scope: createSwitchExpected.Scope,
	}
)

func testSwitchCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Create(ctx, testZone, createSwitchParam)
}

func testSwitchRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testSwitchUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateSwitchParam)
}

func testSwitchUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateSwitchToMinParam)
}

func testSwitchDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSwitchOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestSwitchOp_BridgeConnection(t *testing.T) {
	caller := singletonAPICaller()

	swOp := sacloud.NewSwitchOp(caller)
	bridgeOp := sacloud.NewBridgeOp(caller)

	var bridgeID types.ID

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
				return swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
					Name: testutil.ResourceName("switch-for-bridge"),
				})
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: testSwitchRead,
		},
		Updates: []*testutil.CRUDTestFunc{
			// bridge create and connect
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					bridge, err := bridgeOp.Create(ctx, testZone, &sacloud.BridgeCreateRequest{
						Name: testutil.ResourceName("bridge"),
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
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					// connect to bridge
					if err := swOp.ConnectToBridge(ctx, testZone, ctx.ID, bridgeID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, ctx.ID)
					if err != nil {
						return err
					}
					if err := testutil.AssertEqual(t, bridgeID, sw.BridgeID, "Switch.BridgeID"); err != nil {
						return err
					}

					bridge, err := bridgeOp.Read(ctx, testZone, bridgeID)
					if err != nil {
						return err
					}

					return testutil.DoAsserts(
						testutil.AssertEqualFunc(t, sw.ID, bridge.SwitchInZone.ID, "Bridge.SwitchInZone.ID"),
						testutil.AssertLenFunc(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo"),
					)
				},
			},
			// disconnect
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (i interface{}, e error) {
					if err := swOp.DisconnectFromBridge(ctx, testZone, ctx.ID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, i interface{}) error {
					sw, err := swOp.Read(ctx, testZone, ctx.ID)
					if err != nil {
						return err
					}
					if err := testutil.AssertTrue(t, sw.BridgeID.IsEmpty(), "Switch.BridgeID"); err != nil {
						return err
					}

					bridge, err := bridgeOp.Read(ctx, testZone, bridgeID)
					if err != nil {
						return err
					}

					return testutil.DoAsserts(
						testutil.AssertNilFunc(t, bridge.SwitchInZone, "Bridge.SwitchInZone"),
						testutil.AssertLenFunc(t, bridge.BridgeInfo, 0, "Bridge.BridgeInfo"),
					)
				},
			},
			// bridge delete
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					if err := bridgeOp.Delete(ctx, testZone, bridgeID); err != nil {
						return nil, err
					}
					return nil, nil
				},
				SkipExtractID: true,
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testSwitchDelete,
		},
	})
}
