package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestLoadBalancerOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup:              setupSwitchFunc("lb", createLoadBalancerParam, createLoadBalancerExpected, updateLoadBalancerExpected),
		Create: &CRUDTestFunc{
			Func: testLoadBalancerCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createLoadBalancerExpected,
				IgnoreFields: ignoreLoadBalancerFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testLoadBalancerRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createLoadBalancerExpected,
				IgnoreFields: ignoreLoadBalancerFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testLoadBalancerUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateLoadBalancerExpected,
				IgnoreFields: ignoreLoadBalancerFields,
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewLoadBalancerOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
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
		"IconID",
		"CreatedAt",
		"ModifiedAt",
		"SettingsHash",
	}

	createLoadBalancerParam = &sacloud.LoadBalancerCreateRequest{
		PlanID:         types.ID(2),
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
		Availability:       types.Availabilities.Available,
		PlanID:             createLoadBalancerParam.PlanID,
		InstanceStatus:     types.ServerInstanceStatuses.Up,
		DefaultRoute:       createLoadBalancerParam.DefaultRoute,
		NetworkMaskLen:     createLoadBalancerParam.NetworkMaskLen,
		IPAddresses:        createLoadBalancerParam.IPAddresses,
		VRID:               createLoadBalancerParam.VRID,
		VirtualIPAddresses: updateLoadBalancerParam.VirtualIPAddresses,
	}
)

func testLoadBalancerCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	res, err := client.Create(context.Background(), testZone, createLoadBalancerParam)
	if err != nil {
		return nil, err
	}
	return res.LoadBalancer, nil
}

func testLoadBalancerRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.LoadBalancer, nil
}

func testLoadBalancerUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateLoadBalancerParam)
	if err != nil {
		return nil, err
	}
	return res.LoadBalancer, nil
}

func testLoadBalancerDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewLoadBalancerOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
