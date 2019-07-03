package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestInterface_Operations(t *testing.T) {

	Run(t, &CRUDTestCase{
		Parallel:          true,
		IgnoreStartupWait: true,

		SetupAPICaller: singletonAPICaller,

		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			serverClient := sacloud.NewServerOp(caller)
			server, err := serverClient.Create(context.Background(), testZone, &sacloud.ServerCreateRequest{
				CPU:      1,
				MemoryMB: 1 * 1024,
				//ConnectedSwitches: []*ConnectedSwitch{
				//	{Scope: types.Scopes.Shared},
				//},
				ServerPlanCommitment: types.Commitments.Standard,
				Name:                 "libsacloud-server-with-interface",
			})
			require.NoError(t, err)
			testContext.Values["interface/server"] = server.ID
			createInterfaceParam.ServerID = server.ID
			createInterfaceExpected.ServerID = server.ID
			updateInterfaceExpected.ServerID = server.ID
			return nil
		},

		Create: &CRUDTestFunc{
			Func: testInterfaceCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			},
		},
		Read: &CRUDTestFunc{
			Func: testInterfaceRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			},
		},
		Update: &CRUDTestFunc{
			Func: testInterfaceUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateInterfaceExpected,
				IgnoreFields: ignoreInterfaceFields,
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testInterfaceDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			serverID, ok := testContext.Values["interface/server"]
			if !ok {
				return nil
			}

			serverClient := sacloud.NewServerOp(caller)
			return serverClient.Delete(context.Background(), testZone, serverID.(types.ID))
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

func testInterfaceCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Create(context.Background(), testZone, createInterfaceParam)
}

func testInterfaceRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testInterfaceUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInterfaceOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateInterfaceParam)
}

func testInterfaceDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInterfaceOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
