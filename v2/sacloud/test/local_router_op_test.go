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

func TestLocalRouterOp_CRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("switch-for-local-router"),
			})
			if err != nil {
				return err
			}

			ctx.Values["localrouter/switch"] = sw.ID

			updateLocalRouterParam.Switch.Code = sw.ID.String()
			updateLocalRouterParam.Switch.ZoneID = testZone
			updateLocalRouterExpected.Switch.Code = sw.ID.String()
			updateLocalRouterExpected.Switch.ZoneID = testZone

			updateLocalRouterSettingsParam.Switch.Code = sw.ID.String()
			updateLocalRouterSettingsParam.Switch.ZoneID = testZone
			updateLocalRouterSettingsExpected.Switch.Code = sw.ID.String()
			updateLocalRouterSettingsExpected.Switch.ZoneID = testZone
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: testLocalRouterCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLocalRouterExpected,
				IgnoreFields: ignoreLocalRouterFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testLocalRouterRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLocalRouterExpected,
				IgnoreFields: ignoreLocalRouterFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testLocalRouterUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLocalRouterExpected,
					IgnoreFields: ignoreLocalRouterFields,
				}),
			},
			{
				Func: testLocalRouterUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLocalRouterSettingsExpected,
					IgnoreFields: ignoreLocalRouterFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testLocalRouterDelete,
		},
		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			switchID, ok := ctx.Values["localrouter/switch"]
			if ok {
				if err := swOp.Delete(ctx, testZone, switchID.(types.ID)); err != nil {
					return err
				}
			}
			return nil
		},
	})
}

var (
	ignoreLocalRouterFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"CreatedAt",
		"ModifiedAt",
		"SecretKeys",
	}
	createLocalRouterParam = &sacloud.LocalRouterCreateRequest{
		Name:        testutil.ResourceName("container-registry"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
	}
	createLocalRouterExpected = &sacloud.LocalRouter{
		Name:         createLocalRouterParam.Name,
		Description:  createLocalRouterParam.Description,
		Tags:         createLocalRouterParam.Tags,
		Availability: types.Availabilities.Available,
	}
	updateLocalRouterParam = &sacloud.LocalRouterUpdateRequest{
		Name:        testutil.ResourceName("container-registry-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
		Switch: &sacloud.LocalRouterSwitch{
			Category: "cloud",
		},
		Interface: &sacloud.LocalRouterInterface{
			VirtualIPAddress: "192.168.0.2",
			IPAddress:        []string{"192.168.0.21", "192.168.0.22"},
			NetworkMaskLen:   24,
			VRID:             100,
		},
		StaticRoutes: []*sacloud.LocalRouterStaticRoute{
			{
				Prefix:  "192.168.1.0/24",
				NextHop: "192.168.0.201",
			},
			{
				Prefix:  "192.168.2.0/24",
				NextHop: "192.168.0.202",
			},
		},
	}
	updateLocalRouterExpected = &sacloud.LocalRouter{
		Name:         updateLocalRouterParam.Name,
		Description:  updateLocalRouterParam.Description,
		Tags:         updateLocalRouterParam.Tags,
		Availability: types.Availabilities.Available,
		IconID:       testIconID,
		Switch:       updateLocalRouterParam.Switch,
		Interface:    updateLocalRouterParam.Interface,
		StaticRoutes: updateLocalRouterParam.StaticRoutes,
	}

	updateLocalRouterSettingsParam = &sacloud.LocalRouterUpdateSettingsRequest{
		Switch: &sacloud.LocalRouterSwitch{
			Category: "cloud",
		},
		Interface: &sacloud.LocalRouterInterface{
			VirtualIPAddress: "192.168.0.3",
			IPAddress:        []string{"192.168.0.31", "192.168.0.32"},
			NetworkMaskLen:   24,
			VRID:             100,
		},
		StaticRoutes: []*sacloud.LocalRouterStaticRoute{
			{
				Prefix:  "192.168.1.0/24",
				NextHop: "192.168.0.231",
			},
			{
				Prefix:  "192.168.2.0/24",
				NextHop: "192.168.0.232",
			},
		},
	}
	updateLocalRouterSettingsExpected = &sacloud.LocalRouter{
		Name:         updateLocalRouterParam.Name,
		Description:  updateLocalRouterParam.Description,
		Tags:         updateLocalRouterParam.Tags,
		Availability: types.Availabilities.Available,
		IconID:       testIconID,
		Switch:       updateLocalRouterSettingsParam.Switch,
		Interface:    updateLocalRouterSettingsParam.Interface,
		StaticRoutes: updateLocalRouterSettingsParam.StaticRoutes,
	}
)

func testLocalRouterCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLocalRouterOp(caller)
	return client.Create(ctx, createLocalRouterParam)
}

func testLocalRouterRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLocalRouterOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testLocalRouterUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLocalRouterOp(caller)
	return client.Update(ctx, ctx.ID, updateLocalRouterParam)
}

func testLocalRouterUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLocalRouterOp(caller)
	return client.UpdateSettings(ctx, ctx.ID, updateLocalRouterSettingsParam)
}

func testLocalRouterDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewLocalRouterOp(caller)
	return client.Delete(ctx, ctx.ID)
}

func TestLocalRouter_peering(t *testing.T) {
	var sw1ID, sw2ID types.ID
	var peerLocalRouter1, peerLocalRouter2 *sacloud.LocalRouter

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("switch-for-local-router"),
			})
			if err != nil {
				return err
			}
			sw1ID = sw.ID

			sw2, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("switch-for-local-router"),
			})
			if err != nil {
				return err
			}
			sw2ID = sw2.ID

			lr, err := sacloud.NewLocalRouterOp(caller).Create(ctx, &sacloud.LocalRouterCreateRequest{
				Name: testutil.ResourceName("local-router"),
			})
			if err != nil {
				return err
			}
			peerLocalRouter1 = lr
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				lrOp := sacloud.NewLocalRouterOp(caller)
				lr, err := lrOp.Create(ctx, &sacloud.LocalRouterCreateRequest{
					Name: testutil.ResourceName("local-router"),
				})
				if err != nil {
					return nil, err
				}
				peerLocalRouter2 = lr
				return lr, nil
			},
		},

		Read: &testutil.CRUDTestFunc{
			Func: testLocalRouterRead,
		},

		Updates: []*testutil.CRUDTestFunc{
			// connect to switches
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					lrOp := sacloud.NewLocalRouterOp(caller)
					lr1, err := lrOp.UpdateSettings(ctx, peerLocalRouter1.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch: &sacloud.LocalRouterSwitch{
							Code:     sw1ID.String(),
							Category: "cloud",
							ZoneID:   testZone,
						},
						Interface: &sacloud.LocalRouterInterface{
							VirtualIPAddress: "192.168.0.1",
							IPAddress:        []string{"192.168.0.11", "192.168.0.12"},
							NetworkMaskLen:   24,
							VRID:             100,
						},
						SettingsHash: peerLocalRouter1.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter1 = lr1

					lr2, err := lrOp.UpdateSettings(ctx, peerLocalRouter2.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch: &sacloud.LocalRouterSwitch{
							Code:     sw2ID.String(),
							Category: "cloud",
							ZoneID:   testZone,
						},
						Interface: &sacloud.LocalRouterInterface{
							VirtualIPAddress: "192.168.1.1",
							IPAddress:        []string{"192.168.1.11", "192.168.1.12"},
							NetworkMaskLen:   24,
							VRID:             100,
						},
						SettingsHash: peerLocalRouter2.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter2 = lr2
					return lr2, nil
				},
			},
			// set peer
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					lrOp := sacloud.NewLocalRouterOp(caller)
					lr1, err := lrOp.UpdateSettings(ctx, peerLocalRouter1.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch:    peerLocalRouter1.Switch,
						Interface: peerLocalRouter1.Interface,
						Peers: []*sacloud.LocalRouterPeer{
							{
								ID:          peerLocalRouter2.ID,
								SecretKey:   peerLocalRouter2.SecretKeys[0],
								Enabled:     true,
								Description: "desc",
							},
						},
						SettingsHash: peerLocalRouter1.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter1 = lr1

					lr2, err := lrOp.UpdateSettings(ctx, peerLocalRouter2.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch:    peerLocalRouter2.Switch,
						Interface: peerLocalRouter2.Interface,
						Peers: []*sacloud.LocalRouterPeer{
							{
								ID:          peerLocalRouter1.ID,
								SecretKey:   peerLocalRouter1.SecretKeys[0],
								Enabled:     true,
								Description: "desc",
							},
						},
						SettingsHash: peerLocalRouter2.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter2 = lr2
					return lr2, nil
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, _ interface{}) error {
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, peerLocalRouter1.Peers, "LocalRouter1.Peers"),
						testutil.AssertNotNilFunc(t, peerLocalRouter2.Peers, "LocalRouter2.Peers"),
						testutil.AssertEqualFunc(t, peerLocalRouter1.Peers[0].ID, peerLocalRouter2.ID, "LocalRouter2.Peers[0].ID"),
						testutil.AssertEqualFunc(t, peerLocalRouter2.Peers[0].ID, peerLocalRouter1.ID, "LocalRouter2.Peers[0].ID"),
					)
				},
			},
			// clear peer
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					lrOp := sacloud.NewLocalRouterOp(caller)
					lr1, err := lrOp.UpdateSettings(ctx, peerLocalRouter1.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch:       peerLocalRouter1.Switch,
						Interface:    peerLocalRouter1.Interface,
						SettingsHash: peerLocalRouter1.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter1 = lr1

					lr2, err := lrOp.UpdateSettings(ctx, peerLocalRouter2.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						Switch:       peerLocalRouter2.Switch,
						Interface:    peerLocalRouter2.Interface,
						SettingsHash: peerLocalRouter2.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter2 = lr2
					return lr2, nil
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, _ interface{}) error {
					return testutil.DoAsserts(
						testutil.AssertNilFunc(t, peerLocalRouter1.Peers, "LocalRouter1.Peers"),
						testutil.AssertNilFunc(t, peerLocalRouter2.Peers, "LocalRouter2.Peers"),
					)
				},
			},
			// disconnect from switches
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					lrOp := sacloud.NewLocalRouterOp(caller)
					lr1, err := lrOp.UpdateSettings(ctx, peerLocalRouter1.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						SettingsHash: peerLocalRouter1.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter1 = lr1

					lr2, err := lrOp.UpdateSettings(ctx, peerLocalRouter2.ID, &sacloud.LocalRouterUpdateSettingsRequest{
						SettingsHash: peerLocalRouter2.SettingsHash,
					})
					if err != nil {
						return nil, err
					}
					peerLocalRouter2 = lr2
					return lr2, nil
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, _ interface{}) error {
					return testutil.DoAsserts(
						testutil.AssertNilFunc(t, peerLocalRouter1.Switch, "LocalRouter1.Switch"),
						testutil.AssertNilFunc(t, peerLocalRouter2.Switch, "LocalRouter2.Switch"),
					)
				},
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				lrOp := sacloud.NewLocalRouterOp(caller)
				ids := []types.ID{peerLocalRouter1.ID, peerLocalRouter2.ID}
				for _, id := range ids {
					lrOp.Delete(ctx, id) // nolint
				}

				swOp := sacloud.NewSwitchOp(caller)
				ids = []types.ID{sw1ID, sw2ID}
				for _, id := range ids {
					swOp.Delete(ctx, testZone, id) // nolint
				}
				return nil
			},
		},
	})
}
