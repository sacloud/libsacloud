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
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestDiskOp_BlankDiskCRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &testutil.CRUDTestFunc{
			Func: testDiskCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testDiskRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testDiskUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDiskExpected,
					IgnoreFields: ignoreDiskFields,
				}),
			},
			{
				Func: testDiskUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDiskToMinExpected,
					IgnoreFields: ignoreDiskFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testDiskDelete,
		},
	})
}

var (
	ignoreDiskFields = []string{
		"ID",
		"DisplayOrder",
		"Availability",
		"DiskPlanName",
		"DiskPlanStorageClass",
		"SizeMB",
		"MigratedMB",
		"SourceDiskID",
		"SourceDiskAvailability",
		"SourceArchiveID",
		"SourceArchiveAvailability",
		"BundleInfo",
		"Server",
		"Storage",
		"CreatedAt",
		"ModifiedAt",
	}

	createDiskParam = &sacloud.DiskCreateRequest{
		DiskPlanID:  types.DiskPlans.SSD,
		Connection:  types.DiskConnections.VirtIO,
		Name:        testutil.ResourceName("disk"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		SizeMB:      20 * size.GiB,
	}
	createDiskExpected = &sacloud.Disk{
		Name:        createDiskParam.Name,
		Description: createDiskParam.Description,
		Tags:        createDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
	}
	updateDiskParam = &sacloud.DiskUpdateRequest{
		Name:        testutil.ResourceName("disk-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
	}
	updateDiskExpected = &sacloud.Disk{
		Name:        updateDiskParam.Name,
		Description: updateDiskParam.Description,
		Tags:        updateDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
		IconID:      updateDiskParam.IconID,
	}
	updateDiskToMinParam = &sacloud.DiskUpdateRequest{
		Name: testutil.ResourceName("disk-to-min"),
	}
	updateDiskToMinExpected = &sacloud.Disk{
		Name:       updateDiskToMinParam.Name,
		DiskPlanID: createDiskParam.DiskPlanID,
		Connection: createDiskParam.Connection,
	}
)

func testDiskCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Create(ctx, testZone, createDiskParam, nil)
}

func testDiskRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testDiskUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDiskParam)
}

func testDiskUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDiskToMinParam)
}

func testDiskDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDiskOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestDiskOp_Config(t *testing.T) {
	// source archive
	var archiveID types.ID

	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewArchiveOp(singletonAPICaller())
			searched, err := client.Find(ctx, testZone, &sacloud.FindCondition{
				Filter: search.Filter{
					search.Key("Tags.Name"): search.TagsAndEqual("current-stable", "distro-centos"),
				},
			})
			if !assert.NoError(t, err) {
				return err
			}
			if searched.Count == 0 {
				return errors.New("archive is not found")
			}
			archiveID = searched.Archives[0].ID
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewDiskOp(singletonAPICaller())
				disk, err := client.Create(ctx, testZone, &sacloud.DiskCreateRequest{
					Name:            testutil.ResourceName("disk-edit"),
					DiskPlanID:      types.DiskPlans.SSD,
					SizeMB:          20 * size.GiB,
					SourceArchiveID: archiveID,
				}, nil)
				if err != nil {
					return nil, err
				}
				if _, err = sacloud.WaiterForReady(func() (interface{}, error) {
					return client.Read(ctx, testZone, disk.ID)
				}).WaitForState(ctx); err != nil {
					return disk, err
				}

				return disk, nil
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: testDiskRead,
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// edit disk
					client := sacloud.NewDiskOp(singletonAPICaller())
					err := client.Config(ctx, testZone, ctx.ID, &sacloud.DiskEditRequest{
						Background: true,
						Password:   "password",
						SSHKeys: []*sacloud.DiskEditSSHKey{
							{
								PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4LDQuDiKecOJDPY9InS7EswZ2fPnoRZXc48T1EqyRLyJhgEYGSDWaBiMDs2R/lWgA81Hp37qhrNqZPjFHUkBr93FOXxt9W0m1TNlkNepK0Uyi+14B2n0pdoeqsKEkb3sTevWF0ztxxWrwUd7Mems2hf+wFODITHYye9RlDAKLKPCFRvlQ9xQj4bBWOogQwoaXMSK1znMPjudcm1tRry4KIifLdXmwVKU4qDPGxoXfqs44Dgsikk43UVBStQ7IFoqPgAqcJFSGHLoMS7tPKdTvY9+GME5QidWK84gl69piAkgIdwd+JTMUOc/J+9DXAt220HqZ6l3yhWG5nIgi0x8n",
							},
						},
						DisablePWAuth: true,
						EnableDHCP:    true,
						HostName:      "hostname",
						UserIPAddress: "192.2.0.11",
						UserSubnet: &sacloud.DiskEditUserSubnet{
							DefaultRoute:   "192.2.0.1",
							NetworkMaskLen: 24,
						},
					})
					if err != nil {
						return nil, err
					}
					// wait
					_, err = sacloud.WaiterForReady(func() (interface{}, error) {
						return client.Read(ctx, testZone, ctx.ID)
					}).WaitForState(ctx)
					return nil, err
				},
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testDiskDelete,
		},
	})
}
