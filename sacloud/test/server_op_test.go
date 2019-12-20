// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestServerOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testServerCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testServerRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createServerExpected,
				IgnoreFields: ignoreServerFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testServerUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateServerExpected,
					IgnoreFields: ignoreServerFields,
				}),
			},
			{
				Func: testServerUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateServerToMinExpected,
					IgnoreFields: ignoreServerFields,
				}),
			},
			// Insert CDROM
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					cdOp := sacloud.NewCDROMOp(caller)
					serverOp := sacloud.NewServerOp(caller)

					// find cdrom
					searched, err := cdOp.Find(ctx, testZone, &sacloud.FindCondition{
						Filter: search.Filter{
							search.Key(keys.Scope): types.Scopes.Shared,
						},
						Count: 1,
					})
					if err != nil {
						return nil, err
					}
					cdromID := searched.CDROMs[0].ID
					ctx.Values["server/cdrom"] = cdromID

					// insert
					if err := serverOp.InsertCDROM(ctx, testZone, ctx.ID, &sacloud.InsertCDROMRequest{ID: cdromID}); err != nil {
						return nil, err
					}
					return serverOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					server := v.(*sacloud.Server)
					return testutil.AssertFalse(t, server.CDROMID.IsEmpty(), "Server.CDROMID")
				},
				SkipExtractID: true,
			},
			// Eject CDROM
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					serverOp := sacloud.NewServerOp(caller)
					cdromID := ctx.Values["server/cdrom"].(types.ID)

					if err := serverOp.EjectCDROM(ctx, testZone, ctx.ID, &sacloud.EjectCDROMRequest{ID: cdromID}); err != nil {
						return nil, err
					}
					return serverOp.Read(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					server := v.(*sacloud.Server)
					return testutil.AssertTrue(t, server.CDROMID.IsEmpty(), "Server.CDROMID")
				},
				SkipExtractID: true,
			},
			// VNC Info
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					serverOp := sacloud.NewServerOp(caller)
					return serverOp.GetVNCProxy(ctx, testZone, ctx.ID)
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					vnc := v.(*sacloud.VNCProxyInfo)
					return testutil.DoAsserts(
						testutil.AssertNotNilFunc(t, vnc, "VNCProxyInfo"),
						testutil.AssertNotEmptyFunc(t, vnc.Status, "VNCProxyInfo.Status"),
						testutil.AssertNotEmptyFunc(t, vnc.Host, "VNCProxyInfo.Host"),
						testutil.AssertNotEmptyFunc(t, vnc.IOServerHost, "VNCProxyInfo.IOServerHost"),
						testutil.AssertNotEmptyFunc(t, vnc.Port, "VNCProxyInfo.Port"),
						testutil.AssertNotEmptyFunc(t, vnc.Password, "VNCProxyInfo.Password"),
						testutil.AssertNotEmptyFunc(t, vnc.VNCFile, "VNCProxyInfo.VNCFile"),
					)
				},
				SkipExtractID: true,
			},
		},

		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewServerOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &testutil.CRUDTestDeleteFunc{
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
		"HostName",
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
		"CreatedAt",
		"ModifiedAt",
	}
	createServerParam = &sacloud.ServerCreateRequest{
		CPU:      1,
		MemoryMB: 1 * size.GiB,
		ConnectedSwitches: []*sacloud.ConnectedSwitch{
			{
				Scope: types.Scopes.Shared,
			},
		},
		InterfaceDriver:   types.InterfaceDrivers.VirtIO,
		Name:              testutil.ResourceName("server"),
		Description:       "desc",
		Tags:              []string{"tag1", "tag2"},
		WaitDiskMigration: false,
	}
	createServerExpected = &sacloud.Server{
		Name:            createServerParam.Name,
		Description:     createServerParam.Description,
		Tags:            createServerParam.Tags,
		InterfaceDriver: createServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
	updateServerParam = &sacloud.ServerUpdateRequest{
		Name:            testutil.ResourceName("server-upd"),
		Tags:            []string{"tag1-upd", "tag2-upd"},
		Description:     "desc-upd",
		IconID:          testIconID,
		InterfaceDriver: types.InterfaceDrivers.VirtIO,
	}
	updateServerExpected = &sacloud.Server{
		Name:            updateServerParam.Name,
		Description:     updateServerParam.Description,
		Tags:            updateServerParam.Tags,
		InterfaceDriver: updateServerParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
		IconID:          testIconID,
	}
	updateServerToMinParam = &sacloud.ServerUpdateRequest{
		Name:            testutil.ResourceName("server-to-min"),
		InterfaceDriver: types.InterfaceDrivers.VirtIO,
	}
	updateServerToMinExpected = &sacloud.Server{
		Name:            updateServerToMinParam.Name,
		InterfaceDriver: updateServerToMinParam.InterfaceDriver,
		CPU:             createServerParam.CPU,
		MemoryMB:        createServerParam.MemoryMB,
	}
)

func testServerCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	server, err := client.Create(ctx, testZone, createServerParam)
	if err != nil {
		return nil, err
	}
	if err := client.Boot(ctx, testZone, server.ID); err != nil {
		return nil, err
	}
	return server, nil
}

func testServerRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testServerUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateServerParam)
}

func testServerUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewServerOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateServerToMinParam)
}

func testServerDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewServerOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestServerOp_ChangePlan(t *testing.T) {
	client := sacloud.NewServerOp(singletonAPICaller())
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		IgnoreStartupWait:  true,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				return client.Create(ctx, testZone, &sacloud.ServerCreateRequest{
					CPU:      1,
					MemoryMB: 1 * size.GiB,
					ConnectedSwitches: []*sacloud.ConnectedSwitch{
						{
							Scope: types.Scopes.Shared,
						},
					},
					InterfaceDriver:   types.InterfaceDrivers.VirtIO,
					Name:              testutil.ResourceName("server"),
					Description:       "desc",
					Tags:              []string{"tag1", "tag2"},
					WaitDiskMigration: false,
				})
			},
			CheckFunc: func(t testutil.TestT, _ *testutil.CRUDTestContext, v interface{}) error {
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
		Read: &testutil.CRUDTestFunc{
			Func: testServerRead,
		},
		Updates: []*testutil.CRUDTestFunc{
			// change plan
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return client.ChangePlan(ctx, testZone, ctx.ID, &sacloud.ServerChangePlanRequest{
						CPU:      2,
						MemoryMB: 4 * size.GiB,
					})
				},
				CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
					newServer := v.(*sacloud.Server)
					if !assert.Equal(t, newServer.CPU, 2) {
						return errors.New("unexpected state: Server.CPU")
					}
					if !assert.Equal(t, newServer.GetMemoryGB(), 4) {
						return errors.New("unexpected state: Server.GerMemoryGB()")
					}
					if !assert.NotEqual(t, ctx.ID, newServer.ID) {
						return errors.New("unexpected state: Server.ID(renew)")
					}
					return nil
				},
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testServerDelete,
		},
	})
}

func TestServerOp_Interfaces(t *testing.T) {
	var serverID, switchID types.ID

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		IgnoreStartupWait:  true,

		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				// create server with interfaces[ disconnected, disconnected, switch ]
				switchOp := sacloud.NewSwitchOp(caller)
				sw, err := switchOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{Name: "libsacloud-switch-for-server"})
				if err != nil {
					return nil, err
				}

				serverOp := sacloud.NewServerOp(caller)
				server, err := serverOp.Create(ctx, testZone, &sacloud.ServerCreateRequest{
					Name:     testutil.ResourceName("server-disconnected-nics"),
					CPU:      1,
					MemoryMB: 1 * size.GiB,
					ConnectedSwitches: []*sacloud.ConnectedSwitch{
						nil,
						nil,
						{ID: sw.ID},
					},
				})
				if err != nil {
					return nil, err
				}

				serverID = server.ID
				switchID = sw.ID

				return server, err
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, v interface{}) error {
				server := v.(*sacloud.Server)
				return testutil.DoAsserts(
					testutil.AssertLenFunc(t, server.Interfaces, 3, "Server.Interfaces"),
				)
			},
		},

		Read: &testutil.CRUDTestFunc{
			Func: testServerRead,
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				switchOp := sacloud.NewSwitchOp(caller)
				serverOp := sacloud.NewServerOp(caller)

				server, _ := serverOp.Read(ctx, testZone, serverID)
				if server != nil && server.InstanceStatus.IsUp() {
					if err := serverOp.Shutdown(ctx, testZone, server.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
						return err
					}
					_, err := sacloud.WaiterForDown(func() (interface{}, error) {
						return serverOp.Read(ctx, testZone, server.ID)
					}).WaitForState(ctx)
					if err != nil {
						return err
					}
				}
				if err := serverOp.Delete(ctx, testZone, server.ID); err != nil {
					return err
				}
				sw, err := switchOp.Read(ctx, testZone, switchID)
				if sw != nil {
					return switchOp.Delete(ctx, testZone, sw.ID)
				}
				return err
			},
		},
	})
}
