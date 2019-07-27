package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestLoadBalancerOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: setupSwitchFunc("lb",
			createLoadBalancerParam,
			createLoadBalancerExpected,
			updateLoadBalancerExpected,
			updateLoadBalancerToMin1Expected,
			updateLoadBalancerToMin2Expected,
		),
		Create: &testutil.CRUDTestFunc{
			Func: testLoadBalancerCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLoadBalancerExpected,
				IgnoreFields: ignoreLoadBalancerFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testLoadBalancerRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createLoadBalancerExpected,
				IgnoreFields: ignoreLoadBalancerFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testLoadBalancerUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLoadBalancerExpected,
					IgnoreFields: ignoreLoadBalancerFields,
				}),
			},
			{
				Func: testLoadBalancerUpdateToMin1,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLoadBalancerToMin1Expected,
					IgnoreFields: ignoreLoadBalancerFields,
				}),
			},
			{
				Func: testLoadBalancerUpdateToMin2,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateLoadBalancerToMin2Expected,
					IgnoreFields: ignoreLoadBalancerFields,
				}),
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewLoadBalancerOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testLoadBalancerDelete,
		},

		Cleanup: cleanupSwitchFunc("lb"),
	})
}

var (
	ignoreLoadBalancerFields = []string{
		"ID",
		"Class",
		"Availability",
		"InstanceStatus",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"Switch",
		"ZoneID",
		"CreatedAt",
		"ModifiedAt",
		"SettingsHash",
	}

	createLoadBalancerParam = &sacloud.LoadBalancerCreateRequest{
		PlanID:         types.LoadBalancerPlans.Premium,
		VRID:           100,
		IPAddresses:    []string{"192.168.0.11", "192.168.0.12"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           "libsacloud-lb",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		VirtualIPAddresses: []*sacloud.LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.101",
				Port:             types.StringNumber(80),
				DelayLoop:        types.StringNumber(10),
				SorryServer:      "192.168.0.2",
				Description:      "vip1 desc",
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:               "192.168.0.201",
						Port:                    types.StringNumber(80),
						Enabled:                 types.StringTrue,
						HealthCheckProtocol:     "http",
						HealthCheckPath:         "/index.html",
						HealthCheckResponseCode: types.StringNumber(200),
					},
					{
						IPAddress:               "192.168.0.202",
						Port:                    types.StringNumber(80),
						Enabled:                 types.StringTrue,
						HealthCheckProtocol:     "http",
						HealthCheckPath:         "/index.html",
						HealthCheckResponseCode: types.StringNumber(200),
					},
				},
			},
			{
				VirtualIPAddress: "192.168.0.102",
				Port:             types.StringNumber(80),
				DelayLoop:        types.StringNumber(10),
				SorryServer:      "192.168.0.2",
				Description:      "vip2 desc",
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:               "192.168.0.203",
						Port:                    types.StringNumber(80),
						Enabled:                 types.StringTrue,
						HealthCheckProtocol:     "http",
						HealthCheckPath:         "/index.html",
						HealthCheckResponseCode: types.StringNumber(200),
					},
					{
						IPAddress:               "192.168.0.204",
						Port:                    types.StringNumber(80),
						Enabled:                 types.StringTrue,
						HealthCheckProtocol:     "http",
						HealthCheckPath:         "/index.html",
						HealthCheckResponseCode: types.StringNumber(200),
					},
				},
			},
		},
	}
	createLoadBalancerExpected = &sacloud.LoadBalancer{
		Name:               createLoadBalancerParam.Name,
		Description:        createLoadBalancerParam.Description,
		Tags:               createLoadBalancerParam.Tags,
		Availability:       types.Availabilities.Available,
		InstanceStatus:     types.ServerInstanceStatuses.Up,
		PlanID:             createLoadBalancerParam.PlanID,
		DefaultRoute:       createLoadBalancerParam.DefaultRoute,
		NetworkMaskLen:     createLoadBalancerParam.NetworkMaskLen,
		IPAddresses:        createLoadBalancerParam.IPAddresses,
		VRID:               createLoadBalancerParam.VRID,
		VirtualIPAddresses: createLoadBalancerParam.VirtualIPAddresses,
	}
	updateLoadBalancerParam = &sacloud.LoadBalancerUpdateRequest{
		Name:        "libsacloud-lb-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		IconID:      testIconID,
		VirtualIPAddresses: []*sacloud.LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.111",
				Port:             types.StringNumber(81),
				DelayLoop:        types.StringNumber(11),
				SorryServer:      "192.168.0.3",
				Description:      "vip1 desc-upd",
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:               "192.168.0.211",
						Port:                    types.StringNumber(81),
						Enabled:                 types.StringFalse,
						HealthCheckProtocol:     "https",
						HealthCheckPath:         "/index-upd.html",
						HealthCheckResponseCode: types.StringNumber(201),
					},
					{
						IPAddress:               "192.168.0.212",
						Port:                    types.StringNumber(81),
						Enabled:                 types.StringFalse,
						HealthCheckProtocol:     "https",
						HealthCheckPath:         "/index-upd.html",
						HealthCheckResponseCode: types.StringNumber(201),
					},
				},
			},
			{
				VirtualIPAddress: "192.168.0.112",
				Port:             types.StringNumber(81),
				DelayLoop:        types.StringNumber(11),
				SorryServer:      "192.168.0.3",
				Description:      "vip2 desc-upd",
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:               "192.168.0.213",
						Port:                    types.StringNumber(81),
						Enabled:                 types.StringFalse,
						HealthCheckProtocol:     "https",
						HealthCheckPath:         "/index-upd.html",
						HealthCheckResponseCode: types.StringNumber(201),
					},
					{
						IPAddress:               "192.168.0.214",
						Port:                    types.StringNumber(81),
						Enabled:                 types.StringFalse,
						HealthCheckProtocol:     "https",
						HealthCheckPath:         "/index-upd.html",
						HealthCheckResponseCode: types.StringNumber(201),
					},
				},
			},
		},
	}
	updateLoadBalancerExpected = &sacloud.LoadBalancer{
		Name:               updateLoadBalancerParam.Name,
		Description:        updateLoadBalancerParam.Description,
		Tags:               updateLoadBalancerParam.Tags,
		IconID:             testIconID,
		Availability:       types.Availabilities.Available,
		PlanID:             createLoadBalancerParam.PlanID,
		InstanceStatus:     types.ServerInstanceStatuses.Up,
		DefaultRoute:       createLoadBalancerParam.DefaultRoute,
		NetworkMaskLen:     createLoadBalancerParam.NetworkMaskLen,
		IPAddresses:        createLoadBalancerParam.IPAddresses,
		VRID:               createLoadBalancerParam.VRID,
		VirtualIPAddresses: updateLoadBalancerParam.VirtualIPAddresses,
	}
	updateLoadBalancerToMin1Param = &sacloud.LoadBalancerUpdateRequest{
		Name: "libsacloud-lb-to-min1",
		VirtualIPAddresses: []*sacloud.LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.111",
				Port:             80,
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:           "192.168.0.211",
						Enabled:             types.StringTrue,
						Port:                80,
						HealthCheckProtocol: "ping",
					},
				},
			},
		},
	}
	updateLoadBalancerToMin1Expected = &sacloud.LoadBalancer{
		Name:           updateLoadBalancerToMin1Param.Name,
		Availability:   types.Availabilities.Available,
		PlanID:         createLoadBalancerParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createLoadBalancerParam.DefaultRoute,
		NetworkMaskLen: createLoadBalancerParam.NetworkMaskLen,
		IPAddresses:    createLoadBalancerParam.IPAddresses,
		VRID:           createLoadBalancerParam.VRID,
		VirtualIPAddresses: []*sacloud.LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.111",
				Port:             80,
				DelayLoop:        10, // default value
				Servers: []*sacloud.LoadBalancerServer{
					{
						IPAddress:           "192.168.0.211",
						Enabled:             types.StringTrue,
						Port:                80,
						HealthCheckProtocol: "ping",
					},
				},
			},
		},
	}
	updateLoadBalancerToMin2Param = &sacloud.LoadBalancerUpdateRequest{
		Name: "libsacloud-lb-to-min2",
	}
	updateLoadBalancerToMin2Expected = &sacloud.LoadBalancer{
		Name:           updateLoadBalancerToMin2Param.Name,
		Availability:   types.Availabilities.Available,
		PlanID:         createLoadBalancerParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createLoadBalancerParam.DefaultRoute,
		NetworkMaskLen: createLoadBalancerParam.NetworkMaskLen,
		IPAddresses:    createLoadBalancerParam.IPAddresses,
		VRID:           createLoadBalancerParam.VRID,
	}
)

func testLoadBalancerCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Create(ctx, testZone, createLoadBalancerParam)
}

func testLoadBalancerRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testLoadBalancerUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateLoadBalancerParam)
}

func testLoadBalancerUpdateToMin1(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateLoadBalancerToMin1Param)
}

func testLoadBalancerUpdateToMin2(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateLoadBalancerToMin2Param)
}

func testLoadBalancerDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
