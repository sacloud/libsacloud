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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestArchiveOpCRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewArchiveOp(caller)
			findResult, err := client.Find(ctx, testZone, nil)
			if err != nil {
				return err
			}
			archives := findResult.Archives
			for _, a := range archives {
				if a.GetSizeGB() == 20 && a.Availability.IsAvailable() {
					ctx.Values["archive"] = a.ID
					createArchiveParam.SourceArchiveID = a.ID
					createArchiveExpected.SourceArchiveID = a.ID
					createArchiveExpected.SourceArchiveAvailability = a.Availability
					updateArchiveExpected.SourceArchiveID = a.ID
					updateArchiveExpected.SourceArchiveAvailability = a.Availability
					updateArchiveToMinExpected.SourceArchiveID = a.ID
					updateArchiveToMinExpected.SourceArchiveAvailability = a.Availability

					return nil
				}
			}
			return errors.New("valid archive is not found")
		},

		Create: &testutil.CRUDTestFunc{
			Func: testArchiveCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createArchiveExpected,
				IgnoreFields: ignoreArchiveFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testArchiveRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createArchiveExpected,
				IgnoreFields: ignoreArchiveFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testArchiveUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateArchiveExpected,
					IgnoreFields: ignoreArchiveFields,
				}),
			},
			{
				Func: testArchiveUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateArchiveToMinExpected,
					IgnoreFields: ignoreArchiveFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testArchiveDelete,
		},
	})
}

var (
	ignoreArchiveFields = []string{
		"ID",
		"DisplayOrder",
		"Availability",
		"DiskPlanName",
		"SizeMB",
		"MigratedMB",
		"DiskPlanStorageClass",
		"SourceDiskID",
		"SourceDiskAvailability",
		"BundleInfo",
		"Storage",
		"CreatedAt",
		"ModifiedAt",
		"OriginalArchiveID",
		"SourceInfo",
	}

	createArchiveParam = &sacloud.ArchiveCreateRequest{
		Name:        testutil.ResourceName("archive"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
	}
	createArchiveExpected = &sacloud.Archive{
		Name:        createArchiveParam.Name,
		Description: createArchiveParam.Description,
		Tags:        createArchiveParam.Tags,
		Scope:       types.Scopes.User,
		DiskPlanID:  types.DiskPlans.HDD,
	}

	updateArchiveParam = &sacloud.ArchiveUpdateRequest{
		Name:        testutil.ResourceName("archive-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		IconID:      testIconID,
	}
	updateArchiveExpected = &sacloud.Archive{
		Name:        updateArchiveParam.Name,
		Description: updateArchiveParam.Description,
		Tags:        updateArchiveParam.Tags,
		Scope:       types.Scopes.User,
		DiskPlanID:  types.DiskPlans.HDD,
		IconID:      updateArchiveParam.IconID,
	}
	updateArchiveToMinParam = &sacloud.ArchiveUpdateRequest{
		Name:        testutil.ResourceName("archive-min"),
		Description: "",
	}
	updateArchiveToMinExpected = &sacloud.Archive{
		Name:        updateArchiveToMinParam.Name,
		Description: updateArchiveToMinParam.Description,
		Scope:       types.Scopes.User,
		DiskPlanID:  types.DiskPlans.HDD,
	}
)

func testArchiveCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	return client.Create(ctx, testZone, createArchiveParam)
}

func testArchiveRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testArchiveUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateArchiveParam)
}

func testArchiveUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateArchiveToMinParam)
}

func testArchiveDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewArchiveOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestArchiveOp_CreateBlank(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewArchiveOp(singletonAPICaller())
				archive, ftpServer, err := client.CreateBlank(ctx, testZone, &sacloud.ArchiveCreateBlankRequest{
					SizeMB: 20 * 1024,
					Name:   testutil.ResourceName("archive-blank"),
				})

				if err != nil {
					return nil, err
				}

				assert.NotNil(t, archive)
				assert.NotNil(t, ftpServer)

				return archive, err
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				client := sacloud.NewArchiveOp(singletonAPICaller())
				return client.Delete(ctx, testZone, ctx.ID)
			},
		},
	})
}
