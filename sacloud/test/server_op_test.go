package test

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestServerOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testServerCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testServerRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testServerUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateServerExpected,
					IgnoreFields: ignoreServerFields,
				}),
			},
		},

		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewServerOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testServerDelete,
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
	client := sacloud.NewServerOp(singletonAPICaller())
	ctx := context.Background()

	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		IgnoreStartupWait:  true,
		Create: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return client.Create(ctx, testZone, &sacloud.ServerCreateRequest{
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
			},
			CheckFunc: func(t TestT, _ *CRUDTestContext, v interface{}) error {
				server := v.(*sacloud.Server)

				if !assert.Equal(t, server.CPU, 1) {
					return errors.New("unexpected state: Server.CPU")
				}
				if !assert.Equal(t, server.GetMemoryGB(), 1) {
					return errors.New("unexpected state: Server.GerMemoryGB()")
				}
				return nil
			},
		},
		Read: &CRUDTestFunc{
			Func: testServerRead,
		},
		Updates: []*CRUDTestFunc{
			// change plan
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.ChangePlan(ctx, testZone, testContext.ID, &sacloud.ServerChangePlanRequest{
						CPU:      2,
						MemoryMB: 4 * 1024,
					})

				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					newServer := v.(*sacloud.Server)
					if !assert.Equal(t, newServer.CPU, 2) {
						return errors.New("unexpected state: Server.CPU")
					}
					if !assert.Equal(t, newServer.GetMemoryGB(), 4) {
						return errors.New("unexpected state: Server.GerMemoryGB()")
					}
					if !assert.NotEqual(t, testContext.ID, newServer.ID) {
						return errors.New("unexpected state: Server.ID(renew)")
					}
					return nil
				},
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testServerDelete,
		},
	})
}
