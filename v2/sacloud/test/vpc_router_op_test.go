// Copyright 2016-2020 The Libsacloud Authors
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
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/power"
)

func TestVPCRouterOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testVPCRouterCreate(createVPCRouterParam),
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testVPCRouterRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testVPCRouterUpdate(updateVPCRouterParam),
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateVPCRouterExpected,
					IgnoreFields: ignoreVPCRouterFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewVPCRouterOp(caller)
			return power.ShutdownVPCRouter(ctx, client, testZone, ctx.ID, true)
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testVPCRouterDelete,
		},
	})
}

var (
	ignoreVPCRouterFields = []string{
		"ID",
		"Availability",
		"Class",
		"CreatedAt",
		"SettingsHash",
		"Settings",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatus",
		"InstanceStatusChangedAt",
		"Interfaces",
		"ZoneID",
	}

	createVPCRouterParam = &sacloud.VPCRouterCreateRequest{
		PlanID: types.VPCRouterPlans.Standard,
		Switch: &sacloud.ApplianceConnectedSwitch{
			Scope: types.Scopes.Shared,
		},
		Name:        testutil.ResourceName("vpc-router"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		Settings: &sacloud.VPCRouterSetting{
			InternetConnectionEnabled: true,
			Firewall: []*sacloud.VPCRouterFirewall{
				{
					Receive: []*sacloud.VPCRouterFirewallRule{
						{
							Protocol: types.Protocols.IP,
							Action:   types.Actions.Deny,
						},
					},
				},
			},
		},
	}
	createVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           createVPCRouterParam.Name,
		Description:    createVPCRouterParam.Description,
		Tags:           createVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createVPCRouterParam.PlanID,
		Settings:       createVPCRouterParam.Settings,
	}
	updateVPCRouterParam = &sacloud.VPCRouterUpdateRequest{
		Name:        testutil.ResourceName("vpc-router-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           updateVPCRouterParam.Name,
		Description:    updateVPCRouterParam.Description,
		Tags:           updateVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createVPCRouterParam.PlanID,
	}
)

func testVPCRouterCreate(createParam *sacloud.VPCRouterCreateRequest) func(*testutil.CRUDTestContext, sacloud.APICaller) (interface{}, error) {
	return func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
		client := sacloud.NewVPCRouterOp(caller)
		vpcRouter, err := client.Create(ctx, testZone, createParam)
		if err != nil {
			return nil, err
		}

		n, err := sacloud.WaiterForReady(func() (interface{}, error) {
			return client.Read(ctx, testZone, vpcRouter.ID)
		}).WaitForState(ctx)
		if err != nil {
			return nil, err
		}

		if err := client.Boot(ctx, testZone, vpcRouter.ID); err != nil {
			return nil, err
		}

		return n.(*sacloud.VPCRouter), nil
	}
}

func testVPCRouterRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewVPCRouterOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testVPCRouterUpdate(updateParam *sacloud.VPCRouterUpdateRequest) func(*testutil.CRUDTestContext, sacloud.APICaller) (interface{}, error) {
	return func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
		client := sacloud.NewVPCRouterOp(caller)
		return client.Update(ctx, testZone, ctx.ID, updateParam)
	}
}

func testVPCRouterDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewVPCRouterOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestVPCRouterOp_WithRouterCRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			routerOp := sacloud.NewInternetOp(caller)
			created, err := routerOp.Create(ctx, testZone, &sacloud.InternetCreateRequest{
				Name:           testutil.ResourceName("internet-for-vpc-router"),
				BandWidthMbps:  100,
				NetworkMaskLen: 28,
			})
			if err != nil {
				return err
			}

			ctx.Values["vpcrouter/internet"] = created.ID
			max := 30
			for {
				if max == 0 {
					break
				}
				_, err := routerOp.Read(ctx, testZone, created.ID)
				if err != nil || sacloud.IsNotFoundError(err) {
					max--
					time.Sleep(3 * time.Second)
					continue
				}
				break
			}

			swOp := sacloud.NewSwitchOp(caller)
			sw, err := swOp.Read(ctx, testZone, created.Switch.ID)
			if err != nil {
				return err
			}

			ipaddresses := sw.Subnets[0].GetAssignedIPAddresses()
			p := withRouterCreateVPCRouterParam
			p.Switch = &sacloud.ApplianceConnectedSwitch{
				ID: sw.ID,
			}
			p.IPAddresses = []string{ipaddresses[1], ipaddresses[2]}
			p.Settings = &sacloud.VPCRouterSetting{
				VRID:                      100,
				InternetConnectionEnabled: true,
				Interfaces: []*sacloud.VPCRouterInterfaceSetting{
					{
						VirtualIPAddress: ipaddresses[0],
						IPAddress:        []string{ipaddresses[1], ipaddresses[2]},
						IPAliases:        []string{ipaddresses[3]},
						NetworkMaskLen:   sw.Subnets[0].NetworkMaskLen,
					},
				},
			}

			withRouterCreateVPCRouterExpected.Settings = p.Settings
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: testVPCRouterCreate(withRouterCreateVPCRouterParam),
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  withRouterCreateVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testVPCRouterRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  withRouterCreateVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					if isAccTest() {
						// 起動直後だとシャットダウンできない場合があるため20秒ほど待つ
						time.Sleep(20 * time.Second)
					}

					vpcOp := sacloud.NewVPCRouterOp(caller)
					// shutdown
					if err := vpcOp.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return nil, err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return vpcOp.Read(ctx, testZone, ctx.ID)
					}).WaitForState(ctx)
					if err != nil {
						return nil, err
					}

					swOp := sacloud.NewSwitchOp(caller)
					sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
						Name: testutil.ResourceName("switch-for-vpc-router"),
					})
					if err != nil {
						return nil, err
					}
					ctx.Values["vpcrouter/switch"] = sw.ID

					// connect to switch
					if err := vpcOp.ConnectToSwitch(ctx, testZone, ctx.ID, 2, sw.ID); err != nil {
						return nil, err
					}

					// setup update param
					p := withRouterUpdateVPCRouterParam
					p.Settings = &sacloud.VPCRouterSetting{
						VRID:                      10,
						SyslogHost:                "192.168.2.199",
						InternetConnectionEnabled: true,
						Interfaces: []*sacloud.VPCRouterInterfaceSetting{
							withRouterCreateVPCRouterParam.Settings.Interfaces[0],
							{
								VirtualIPAddress: "192.168.2.1",
								IPAddress:        []string{"192.168.2.11", "192.168.2.12"},
								NetworkMaskLen:   24,
								Index:            2,
							},
						},
						StaticNAT: []*sacloud.VPCRouterStaticNAT{
							{
								GlobalAddress:  withRouterCreateVPCRouterParam.Settings.Interfaces[0].IPAliases[0],
								PrivateAddress: "192.168.2.1",
							},
						},
						PortForwarding: []*sacloud.VPCRouterPortForwarding{
							{
								Protocol:       types.VPCRouterPortForwardingProtocols.TCP,
								GlobalPort:     22,
								PrivateAddress: "192.168.2.2",
								PrivatePort:    10022,
								Description:    "port forwarding",
							},
						},
						DHCPServer: []*sacloud.VPCRouterDHCPServer{
							{
								Interface:  "eth2",
								RangeStart: "192.168.2.51",
								RangeStop:  "192.168.2.60",
							},
						},
						DHCPStaticMapping: []*sacloud.VPCRouterDHCPStaticMapping{
							{
								MACAddress: "aa:bb:cc:dd:ee:ff",
								IPAddress:  "192.168.2.21",
							},
						},
						PPTPServer: &sacloud.VPCRouterPPTPServer{
							RangeStart: "192.168.2.61",
							RangeStop:  "192.168.2.70",
						},
						PPTPServerEnabled: true,
						L2TPIPsecServer: &sacloud.VPCRouterL2TPIPsecServer{
							RangeStart:      "192.168.2.71",
							RangeStop:       "192.168.2.80",
							PreSharedSecret: "presharedsecret",
						},
						L2TPIPsecServerEnabled: true,
						RemoteAccessUsers: []*sacloud.VPCRouterRemoteAccessUser{
							{
								UserName: "user1",
								Password: "password1",
							},
						},
						SiteToSiteIPsecVPN: []*sacloud.VPCRouterSiteToSiteIPsecVPN{
							{
								Peer:            "1.2.3.4",
								PreSharedSecret: "presharedsecret",
								RemoteID:        "1.2.3.4",
								Routes:          []string{"10.0.0.0/24"},
								LocalPrefix:     []string{"192.168.2.0/24"},
							},
						},
						StaticRoute: []*sacloud.VPCRouterStaticRoute{
							{
								Prefix:  "172.16.0.0/16",
								NextHop: "192.168.2.11",
							},
						},
					}

					withRouterUpdateVPCRouterExpected.Settings = p.Settings
					return testVPCRouterUpdate(withRouterUpdateVPCRouterParam)(ctx, caller)
				},
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  withRouterUpdateVPCRouterExpected,
					IgnoreFields: ignoreVPCRouterFields,
				}),
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// setup update param
					p := withRouterUpdateVPCRouterToMinParam
					p.Settings = &sacloud.VPCRouterSetting{
						VRID:                      10,
						InternetConnectionEnabled: false,
						Interfaces: []*sacloud.VPCRouterInterfaceSetting{
							withRouterCreateVPCRouterParam.Settings.Interfaces[0],
						},
					}

					withRouterUpdateVPCRouterToMinExpected.Settings = p.Settings
					return testVPCRouterUpdate(withRouterUpdateVPCRouterToMinParam)(ctx, caller)
				},
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  withRouterUpdateVPCRouterToMinExpected,
					IgnoreFields: ignoreVPCRouterFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewVPCRouterOp(caller)
			return power.ShutdownVPCRouter(ctx, client, testZone, ctx.ID, true)
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testVPCRouterDelete,
		},

		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			routerOp := sacloud.NewInternetOp(caller)
			routerID, ok := ctx.Values["vpcrouter/internet"]
			if ok {
				if err := routerOp.Delete(ctx, testZone, routerID.(types.ID)); err != nil {
					return err
				}
			}

			swOp := sacloud.NewSwitchOp(caller)
			switchID, ok := ctx.Values["vpcrouter/switch"]
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
	withRouterCreateVPCRouterParam = &sacloud.VPCRouterCreateRequest{
		PlanID:      types.VPCRouterPlans.Premium,
		Name:        testutil.ResourceName("vpc-router"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
	}
	withRouterCreateVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           withRouterCreateVPCRouterParam.Name,
		Description:    withRouterCreateVPCRouterParam.Description,
		Tags:           withRouterCreateVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         withRouterCreateVPCRouterParam.PlanID,
		Settings:       withRouterCreateVPCRouterParam.Settings,
	}
	withRouterUpdateVPCRouterParam = &sacloud.VPCRouterUpdateRequest{
		Name:        testutil.ResourceName("vpc-router-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		IconID:      testIconID,
	}
	withRouterUpdateVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           withRouterUpdateVPCRouterParam.Name,
		Description:    withRouterUpdateVPCRouterParam.Description,
		Tags:           withRouterUpdateVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         withRouterCreateVPCRouterParam.PlanID,
		Settings:       withRouterUpdateVPCRouterParam.Settings,
		IconID:         testIconID,
	}
	withRouterUpdateVPCRouterToMinParam = &sacloud.VPCRouterUpdateRequest{
		Name: testutil.ResourceName("vpc-router-to-min"),
	}
	withRouterUpdateVPCRouterToMinExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           withRouterUpdateVPCRouterToMinParam.Name,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         withRouterCreateVPCRouterParam.PlanID,
		Settings:       withRouterUpdateVPCRouterToMinParam.Settings,
	}
)
