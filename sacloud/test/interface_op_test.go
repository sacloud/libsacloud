package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestInterface_Operations(t *testing.T) {

	Run(t, &CRUDTestCase{
		Parallel:          true,
		IgnoreStartupWait: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			serverClient := sacloud.NewServerOp(caller)
			server, err := serverClient.Create(ctx, testZone, &sacloud.ServerCreateRequest{
				CPU:      1,
				MemoryMB: 1 * 1024,
				//ConnectedSwitches: []*ConnectedSwitch{
				//	{Scope: types.Scopes.Shared},
				//},
				ServerPlanCommitment: types.Commitments.Standard,
				Name:                 "libsacloud-server-with-interface",
			})
			if !assert.NoError(t, err) {
				return err
			}

			ctx.Values["interface/server"] = server.ID
			createInterfaceParam.ServerID = server.ID
			createInterfaceExpected.ServerID = server.ID
			updateInterfaceExpected.ServerID = server.ID
			return nil
		},

		Create: &CRUDTestFunc{
			Func: testInterfaceCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testInterfaceRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testInterfaceUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateInterfaceExpected,
					IgnoreFields: ignoreInterfaceFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testInterfaceDelete,
		},

		Cleanup: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			serverID, ok := ctx.Values["interface/server"]
			if !ok {
				return nil
			}

			serverClient := sacloud.NewServerOp(caller)
			return serverClient.Delete(ctx, testZone, serverID.(types.ID))
		},
	})

}

var (
	ignoreInterfaceFields = []string{
		"ID",
		"MACAddress",
		"IPAddress",
		"CreatedAt",
		"ModifiedAt",
	}
	createInterfaceParam = &sacloud.InterfaceCreateRequest{}

	createInterfaceExpected = &sacloud.Interface{
		UserIPAddress:  "",
		HostName:       "",
		SwitchID:       types.ID(0),
		PacketFilterID: types.ID(0),
	}
	updateInterfaceParam = &sacloud.InterfaceUpdateRequest{
		UserIPAddress: "192.2.0.1",
	}
	updateInterfaceExpected = &sacloud.Interface{
		UserIPAddress:  "192.2.0.1",
		HostName:       "",
		SwitchID:       types.ID(0),
		PacketFilterID: types.ID(0),
	}
)

func testInterfaceCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Create(ctx, testZone, createInterfaceParam)
}

func testInterfaceRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testInterfaceUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateInterfaceParam)
}

func testInterfaceDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInterfaceOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
