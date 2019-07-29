package test

import (
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
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
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
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
		"SwitchID",
		"IPAddresses",
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
		IPAddresses:    createVPCRouterParam.IPAddresses,
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
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					if isAccTest() {
						// 起動直後だとシャットダウンできない場合があるため10秒ほど待つ
						time.Sleep(10 * time.Second)
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
						InternetConnectionEnabled: true,
						Interfaces: []*sacloud.VPCRouterInterfaceSetting{
							withRouterCreateVPCRouterParam.Settings.Interfaces[0],
							{
								VirtualIPAddress: "192.0.2.1",
								IPAddress:        []string{"192.0.2.11", "192.0.2.12"},
								NetworkMaskLen:   24,
								Index:            2,
							},
						},
						StaticNAT: []*sacloud.VPCRouterStaticNAT{
							{
								GlobalAddress:  withRouterCreateVPCRouterParam.Settings.Interfaces[0].IPAliases[0],
								PrivateAddress: "192.0.2.1",
							},
						},
						DHCPServer: []*sacloud.VPCRouterDHCPServer{
							{
								Interface:  "eth2",
								RangeStart: "192.0.2.51",
								RangeStop:  "192.0.2.60",
							},
						},
						DHCPStaticMapping: []*sacloud.VPCRouterDHCPStaticMapping{
							{
								MACAddress: "aa:bb:cc:dd:ee:ff",
								IPAddress:  "192.0.2.21",
							},
						},
						PPTPServer: &sacloud.VPCRouterPPTPServer{
							RangeStart: "192.0.2.61",
							RangeStop:  "192.0.2.70",
						},
						PPTPServerEnabled: true,
						L2TPIPsecServer: &sacloud.VPCRouterL2TPIPsecServer{
							RangeStart:      "192.0.2.71",
							RangeStop:       "192.0.2.80",
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
								Peer:            "10.0.0.1",
								PreSharedSecret: "presharedsecret",
								RemoteID:        "10.0.0.1",
								Routes:          []string{"192.0.2.248/28"},
								LocalPrefix:     []string{"192.0.2.0/24"},
							},
						},
						StaticRoute: []*sacloud.VPCRouterStaticRoute{
							{
								Prefix:  "172.16.0.0/16",
								NextHop: "192.0.2.11",
							},
						},
					}

					withRouterUpdateVPCRouterExpected.Settings = p.Settings
					return testVPCRouterUpdate(updateVPCRouterParam)(ctx, caller)
				},
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateVPCRouterExpected,
					IgnoreFields: ignoreVPCRouterFields,
				}),
			},
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// setup update param
					p := withRouterUpdateVPCRouterToMinParam
					p.Settings = &sacloud.VPCRouterSetting{
						InternetConnectionEnabled: false,
						Interfaces: []*sacloud.VPCRouterInterfaceSetting{
							withRouterCreateVPCRouterParam.Settings.Interfaces[0],
						},
					}

					withRouterUpdateVPCRouterToMinExpected.Settings = p.Settings
					return testVPCRouterUpdate(updateVPCRouterParam)(ctx, caller)
				},
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateVPCRouterExpected,
					IgnoreFields: ignoreVPCRouterFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewVPCRouterOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
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
		PlanID:      types.VPCRouterPlans.Standard,
		Name:        testutil.ResourceName("vpc-router"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
	}
	withRouterCreateVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           createVPCRouterParam.Name,
		Description:    createVPCRouterParam.Description,
		Tags:           createVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createVPCRouterParam.PlanID,
		IPAddresses:    createVPCRouterParam.IPAddresses,
		Settings:       createVPCRouterParam.Settings,
	}
	withRouterUpdateVPCRouterParam = &sacloud.VPCRouterUpdateRequest{
		Name:        testutil.ResourceName("vpc-router-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		IconID:      testIconID,
	}
	withRouterUpdateVPCRouterExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           updateVPCRouterParam.Name,
		Description:    updateVPCRouterParam.Description,
		Tags:           updateVPCRouterParam.Tags,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createVPCRouterParam.PlanID,
		Settings:       withRouterUpdateVPCRouterParam.Settings,
		IconID:         testIconID,
	}
	withRouterUpdateVPCRouterToMinParam = &sacloud.VPCRouterUpdateRequest{
		Name: testutil.ResourceName("vpc-router-to-min"),
	}
	withRouterUpdateVPCRouterToMinExpected = &sacloud.VPCRouter{
		Class:          "vpcrouter",
		Name:           updateVPCRouterParam.Name,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createVPCRouterParam.PlanID,
		Settings:       withRouterUpdateVPCRouterToMinParam.Settings,
	}
)
