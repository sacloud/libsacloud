package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func TestServerOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,

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

		Update: &CRUDTestFunc{
			Func: testServerUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateServerExpected,
				IgnoreFields: ignoreServerFields,
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller APICaller) error {
			client := NewServerOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testServerDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller APICaller) error {

			switchID, ok := testContext.Values["nfs/switch"]
			if !ok {
				return nil
			}

			swClient := NewSwitchOp(caller)
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
		"Storage",
		"Icon",
		"CreatedAt",
		"ModifiedAt",
	}
	createServerParam = &ServerCreateRequest{
		CPU:      1,
		MemoryMB: 1 * 1024,
		ConnectedSwitches: []*ConnectedSwitch{
			{
				Scope: types.Scopes.Shared,
			},
		},
		InterfaceDriver:   types.InterfaceDrivers.VirtIO,
		HostName:          "libsacloud-v2-server",
		Name:              "libsacloud-v2-server",
		Description:       "desc",
		Tags:              []string{"tag1", "tag2"},
		WaitDiskMigration: false,
	}
	createServerExpected = &Server{
		Name:            createServerParam.Name,
		Description:     createServerParam.Description,
		Tags:            createServerParam.Tags,
		HostName:        createServerParam.HostName,
		InterfaceDriver: createServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
	updateServerParam = &ServerUpdateRequest{
		Name:        "libsacloud-v2-nfs-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateServerExpected = &Server{
		Name:            updateServerParam.Name,
		Description:     updateServerParam.Description,
		Tags:            updateServerParam.Tags,
		HostName:        createServerParam.HostName,
		InterfaceDriver: createServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
)

func testServerCreate(_ *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewServerOp(caller)
	server, err := client.Create(context.Background(), testZone, createServerParam)
	if err != nil {
		return nil, err
	}
	if err := client.Boot(context.Background(), testZone, server.ID); err != nil {
		return nil, err
	}
	return server, nil
}

func testServerRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewServerOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testServerUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewServerOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateServerParam)
}

func testServerDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewServerOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
