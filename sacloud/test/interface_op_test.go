package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestInterface_Operations(t *testing.T) {

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:          true,
		IgnoreStartupWait: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
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

		Create: &testutil.CRUDTestFunc{
			Func: testInterfaceCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testInterfaceRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testInterfaceUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateInterfaceExpected,
					IgnoreFields: ignoreInterfaceFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testInterfaceDelete,
		},

		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
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

func testInterfaceCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Create(ctx, testZone, createInterfaceParam)
}

func testInterfaceRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testInterfaceUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateInterfaceParam)
}

func testInterfaceDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInterfaceOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
