package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestServerOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testServerCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testServerRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			},
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testServerUpdate,
				Expect: &CRUDTestExpect{
					ExpectValue:  updateServerExpected,
					IgnoreFields: ignoreServerFields,
				},
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewServerOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testServerDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {

			switchID, ok := testContext.Values["nfs/switch"]
			if !ok {
				return nil
			}

			swClient := sacloud.NewSwitchOp(caller)
			return swClient.Delete(context.Background(), testZone, switchID.(types.ID))
		},
	})
}

var (
	ignoreServerFields = []string{
		"ID",
		"Availability",
		"ServerPlanID",
		"ServerPlanName",
		"ServerPlanGeneration",
		"ServerPlanCommitment",
		"Zone",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatus",
		"InstanceBeforeStatus",
		"InstanceStatusChangedAt",
		"InstanceWarnings",
		"InstanceWarningsValue",
		"Disks",
		"Interfaces",
		"PrivateHostID",
		"PrivateHostName",
		"BundleInfo",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
	}
	createServerParam = &sacloud.ServerCreateRequest{
		CPU:      1,
		MemoryMB: 1 * 1024,
		ConnectedSwitches: []*sacloud.ConnectedSwitch{
			{
				Scope: types.Scopes.Shared,
			},
		},
		InterfaceDriver:   types.InterfaceDrivers.VirtIO,
		HostName:          "libsacloud-server",
		Name:              "libsacloud-server",
		Description:       "desc",
		Tags:              []string{"tag1", "tag2"},
		WaitDiskMigration: false,
	}
	createServerExpected = &sacloud.Server{
		Name:            createServerParam.Name,
		Description:     createServerParam.Description,
		Tags:            createServerParam.Tags,
		HostName:        createServerParam.HostName,
		InterfaceDriver: createServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
	updateServerParam = &sacloud.ServerUpdateRequest{
		Name:        "libsacloud-nfs-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateServerExpected = &sacloud.Server{
		Name:            updateServerParam.Name,
		Description:     updateServerParam.Description,
		Tags:            updateServerParam.Tags,
		HostName:        createServerParam.HostName,
		InterfaceDriver: createServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
)

func testServerCreate(_ *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	server, err := client.Create(context.Background(), testZone, createServerParam)
	if err != nil {
		return nil, err
	}
	if err := client.Boot(context.Background(), testZone, server.ID); err != nil {
		return nil, err
	}
	return server, nil
}

func testServerRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testServerUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateServerParam)
}

func testServerDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewServerOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestServerOp_ChangePlan(t *testing.T) {
	t.Parallel()

	client := sacloud.NewServerOp(singletonAPICaller())
	ctx := context.Background()
	server, err := client.Create(ctx, testZone, &sacloud.ServerCreateRequest{
		CPU:      1,
		MemoryMB: 1 * 1024,
		ConnectedSwitches: []*sacloud.ConnectedSwitch{
			{
				Scope: types.Scopes.Shared,
			},
		},
		InterfaceDriver:   types.InterfaceDrivers.VirtIO,
		HostName:          "libsacloud-server",
		Name:              "libsacloud-server",
		Description:       "desc",
		Tags:              []string{"tag1", "tag2"},
		WaitDiskMigration: false,
	})
	require.NoError(t, err)

	require.Equal(t, server.CPU, 1)
	require.Equal(t, server.GetMemoryGB(), 1)

	// change
	newServer, err := client.ChangePlan(ctx, testZone, server.ID, &sacloud.ServerChangePlanRequest{
		CPU:      2,
		MemoryMB: 4 * 1024,
	})
	require.NoError(t, err)

	require.Equal(t, 2, newServer.CPU)
	require.Equal(t, 4, newServer.GetMemoryGB())
	require.NotEqual(t, server.ID, newServer.ID)

	// cleanup
	err = client.Delete(ctx, testZone, newServer.ID)
	require.NoError(t, err)

}
