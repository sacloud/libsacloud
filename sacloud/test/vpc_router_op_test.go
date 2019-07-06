package test

import (
	"context"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestVPCRouterOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testVPCRouterCreate(createVPCRouterParam),
			Expect: &CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testVPCRouterRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testVPCRouterUpdate(updateVPCRouterParam),
			Expect: &CRUDTestExpect{
				ExpectValue:  updateVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewVPCRouterOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testVPCRouterDelete,
		},
	})
}

var (
	ignoreVPCRouterFields = []string{
		"ID",
		"Availability",
		"Class",
		"IconID",
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
		PlanID: types.ID(1), // standard  TODO プランIDをどこかで定義する
		Switch: &sacloud.ApplianceConnectedSwitch{
			Scope: types.Scopes.Shared,
		},
		Name:        "libsacloud-vpc-router",
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
		Name:        "libsacloud-vpc-router-upd",
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

func testVPCRouterCreate(createParam *sacloud.VPCRouterCreateRequest) func(*CRUDTestContext, sacloud.APICaller) (interface{}, error) {
	return func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
		client := sacloud.NewVPCRouterOp(caller)
		createResult, err := client.Create(context.Background(), testZone, createParam)
		if err != nil {
			return nil, err
		}

		n, err := sacloud.WaiterForReady(func() (interface{}, error) {
			res, err := client.Read(context.Background(), testZone, createResult.VPCRouter.ID)
			if err != nil {
				return nil, err
			}
			return res.VPCRouter, nil
		}).WaitForState(context.Background())
		if err != nil {
			return nil, err
		}

		if err := client.Boot(context.Background(), testZone, createResult.VPCRouter.ID); err != nil {
			return nil, err
		}

		return n.(*sacloud.VPCRouter), nil
	}
}

func testVPCRouterRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewVPCRouterOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.VPCRouter, nil
}

func testVPCRouterUpdate(updateParam *sacloud.VPCRouterUpdateRequest) func(*CRUDTestContext, sacloud.APICaller) (interface{}, error) {
	return func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
		client := sacloud.NewVPCRouterOp(caller)
		res, err := client.Update(context.Background(), testZone, testContext.ID, updateParam)
		if err != nil {
			return nil, err
		}
		return res.VPCRouter, nil
	}
}

func testVPCRouterDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewVPCRouterOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestVPCRouterOpWithRouterCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			ctx := context.Background()
			routerOp := sacloud.NewInternetOp(caller)
			createResult, err := routerOp.Create(ctx, testZone, &sacloud.InternetCreateRequest{
				Name:           "libsacloud-internet-for-vpc-router",
				BandWidthMbps:  100,
				NetworkMaskLen: 28,
			})
			if err != nil {
				return err
			}
			created := createResult.Internet

			testContext.Values["vpcrouter/internet"] = created.ID
			max := 30
			for {
				if max == 0 {
					break
				}
				_, err := routerOp.Read(context.Background(), testZone, created.ID)
				if err != nil || sacloud.IsNotFoundError(err) {
					max--
					time.Sleep(3 * time.Second)
					continue
				}
				break
			}

			swOp := sacloud.NewSwitchOp(caller)
			swReadResult, err := swOp.Read(ctx, testZone, created.Switch.ID)
			if err != nil {
				return err
			}
			sw := swReadResult.Switch

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
		Create: &CRUDTestFunc{
			Func: testVPCRouterCreate(createVPCRouterParam),
			Expect: &CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testVPCRouterRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				vpcOp := sacloud.NewVPCRouterOp(caller)
				ctx := context.Background()

				// shutdown
				if err := vpcOp.Shutdown(ctx, testZone, testContext.ID, nil); err != nil {
					return nil, err
				}
				_, err := sacloud.WaiterForDown(func() (interface{}, error) {
					res, err := vpcOp.Read(context.Background(), testZone, testContext.ID)
					if err != nil {
						return nil, err
					}
					return res.VPCRouter, nil
				}).WaitForState(context.Background())
				if err != nil {
					return nil, err
				}

				swOp := sacloud.NewSwitchOp(caller)
				swCreateResult, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
					Name: "libsacloud-switch-for-vpc-router",
				})
				if err != nil {
					return nil, err
				}
				sw := swCreateResult.Switch
				testContext.Values["vpcrouter/switch"] = sw.ID

				// connect to switch
				if err := vpcOp.ConnectToSwitch(ctx, testZone, testContext.ID, 2, sw.ID); err != nil {
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
				return testVPCRouterUpdate(updateVPCRouterParam)(testContext, caller)
			},
			Expect: &CRUDTestExpect{
				ExpectValue:  updateVPCRouterExpected,
				IgnoreFields: ignoreVPCRouterFields,
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewVPCRouterOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testVPCRouterDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			ctx := context.Background()
			routerOp := sacloud.NewInternetOp(caller)
			routerID, ok := testContext.Values["vpcrouter/internet"]
			if ok {
				if err := routerOp.Delete(ctx, testZone, routerID.(types.ID)); err != nil {
					return err
				}
			}

			swOp := sacloud.NewSwitchOp(caller)
			switchID, ok := testContext.Values["vpcrouter/switch"]
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
		PlanID:      types.ID(1), // standard  TODO プランIDをどこかで定義する
		Name:        "libsacloud-vpc-router",
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
		Name:        "libsacloud-vpc-router-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
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
	}
)
