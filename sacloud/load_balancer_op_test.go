package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func TestLoadBalancerOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,
		Setup:          setupSwitchFunc("lb", createLoadBalancerParam, createLoadBalancerExpected, updateLoadBalancerExpected),
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

		Shutdown: func(testContext *CRUDTestContext, caller APICaller) error {
			client := NewLoadBalancerOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &ShutdownOption{Force: true})
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
		"Icon",
		"CreatedAt",
		"ModifiedAt",
		"SettingsHash",
	}

	createLoadBalancerParam = &LoadBalancerCreateRequest{
		PlanID:         types.ID(2),
		VRID:           100,
		IPAddresses:    []string{"192.168.0.11", "192.168.0.12"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           "libsacloud-v2-lb",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		VirtualIPAddresses: []*LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.101",
				Port:             types.StringNumber(80),
				DelayLoop:        types.StringNumber(10),
				SorryServer:      "192.168.0.2",
				Description:      "vip1 desc",
				Servers: []*LoadBalancerServer{
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
				Servers: []*LoadBalancerServer{
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
	createLoadBalancerExpected = &LoadBalancer{
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
	updateLoadBalancerParam = &LoadBalancerUpdateRequest{
		Name:        "libsacloud-v2-lb-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		VirtualIPAddresses: []*LoadBalancerVirtualIPAddress{
			{
				VirtualIPAddress: "192.168.0.111",
				Port:             types.StringNumber(81),
				DelayLoop:        types.StringNumber(11),
				SorryServer:      "192.168.0.3",
				Description:      "vip1 desc-upd",
				Servers: []*LoadBalancerServer{
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
				Servers: []*LoadBalancerServer{
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
	updateLoadBalancerExpected = &LoadBalancer{
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

func testLoadBalancerCreate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewLoadBalancerOp(caller)
	return client.Create(context.Background(), testZone, createLoadBalancerParam)
}

func testLoadBalancerRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewLoadBalancerOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testLoadBalancerUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewLoadBalancerOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateLoadBalancerParam)
}

func testLoadBalancerDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewLoadBalancerOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
