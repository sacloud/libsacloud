// Copyright 2016-2020 The Libsacloud Authors
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
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestAutoBackupOpCRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			diskOp := sacloud.NewDiskOp(caller)
			disk, err := diskOp.Create(ctx, testZone, &sacloud.DiskCreateRequest{
				Name:       testutil.ResourceName("-disk-for-autobackup"),
				SizeMB:     20 * size.GiB,
				DiskPlanID: types.DiskPlans.SSD,
			}, nil)
			if !assert.NoError(t, err) {
				return err
			}

			_, err = sacloud.WaiterForReady(func() (interface{}, error) {
				return diskOp.Read(ctx, testZone, disk.ID)
			}).WaitForState(ctx)
			if !assert.NoError(t, err) {
				return err
			}

			ctx.Values["autobackup/disk"] = disk.ID
			createAutoBackupParam.DiskID = disk.ID
			createAutoBackupExpected.DiskID = disk.ID
			updateAutoBackupExpected.DiskID = disk.ID
			updateAutoBackupToMinExpected.DiskID = disk.ID
			return err
		},

		Create: &testutil.CRUDTestFunc{
			Func: testAutoBackupCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createAutoBackupExpected,
				IgnoreFields: ignoreAutoBackupFields,
			}),
		},

		Read: &testutil.CRUDTestFunc{
			Func: testAutoBackupRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createAutoBackupExpected,
				IgnoreFields: ignoreAutoBackupFields,
			}),
		},

		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testAutoBackupUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateAutoBackupExpected,
					IgnoreFields: ignoreAutoBackupFields,
				}),
			},
			{
				Func: testAutoBackupUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateAutoBackupSettingsExpected,
					IgnoreFields: ignoreAutoBackupFields,
				}),
			},
			{
				Func: testAutoBackupUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateAutoBackupToMinExpected,
					IgnoreFields: ignoreAutoBackupFields,
				}),
			},
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testAutoBackupDelete,
		},

		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			diskID, ok := ctx.Values["autobackup/disk"]
			if !ok {
				return nil
			}

			diskOp := sacloud.NewDiskOp(caller)
			return diskOp.Delete(ctx, testZone, diskID.(types.ID))
		},
	})
}

var (
	ignoreAutoBackupFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"CreatedAt",
		"ModifiedAt",
		"DiskID",
		"AccountID",
		"ZoneID",
		"ZoneName",
	}
	createAutoBackupParam = &sacloud.AutoBackupCreateRequest{
		Name:        testutil.ResourceName("auto-backup"),
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		BackupSpanWeekdays: []types.EBackupSpanWeekday{
			types.BackupSpanWeekdays.Monday,
			types.BackupSpanWeekdays.Tuesday,
		},
		MaximumNumberOfArchives: 2,
	}
	createAutoBackupExpected = &sacloud.AutoBackup{
		Name:                    createAutoBackupParam.Name,
		Description:             createAutoBackupParam.Description,
		Tags:                    createAutoBackupParam.Tags,
		Availability:            types.Availabilities.Available,
		BackupSpanWeekdays:      createAutoBackupParam.BackupSpanWeekdays,
		MaximumNumberOfArchives: createAutoBackupParam.MaximumNumberOfArchives,
	}
	updateAutoBackupParam = &sacloud.AutoBackupUpdateRequest{
		Name:        testutil.ResourceName("auto-backup-upd"),
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		BackupSpanWeekdays: []types.EBackupSpanWeekday{
			types.BackupSpanWeekdays.Monday,
			types.BackupSpanWeekdays.Tuesday,
			types.BackupSpanWeekdays.Wednesday,
			types.BackupSpanWeekdays.Thursday,
		},
		MaximumNumberOfArchives: 3,
		IconID:                  testIconID,
	}
	updateAutoBackupExpected = &sacloud.AutoBackup{
		Name:                    updateAutoBackupParam.Name,
		Description:             updateAutoBackupParam.Description,
		Tags:                    updateAutoBackupParam.Tags,
		Availability:            types.Availabilities.Available,
		BackupSpanWeekdays:      updateAutoBackupParam.BackupSpanWeekdays,
		MaximumNumberOfArchives: updateAutoBackupParam.MaximumNumberOfArchives,
		IconID:                  testIconID,
	}

	updateAutoBackupSettingsParam = &sacloud.AutoBackupUpdateSettingsRequest{
		BackupSpanWeekdays: []types.EBackupSpanWeekday{
			types.BackupSpanWeekdays.Monday,
			types.BackupSpanWeekdays.Tuesday,
			types.BackupSpanWeekdays.Wednesday,
		},
		MaximumNumberOfArchives: 4,
	}
	updateAutoBackupSettingsExpected = &sacloud.AutoBackup{
		Name:                    updateAutoBackupParam.Name,
		Description:             updateAutoBackupParam.Description,
		Tags:                    updateAutoBackupParam.Tags,
		Availability:            types.Availabilities.Available,
		BackupSpanWeekdays:      updateAutoBackupSettingsParam.BackupSpanWeekdays,
		MaximumNumberOfArchives: updateAutoBackupSettingsParam.MaximumNumberOfArchives,
		IconID:                  testIconID,
	}

	updateAutoBackupToMinParam = &sacloud.AutoBackupUpdateRequest{
		Name: testutil.ResourceName("auto-backup-to-min"),
		BackupSpanWeekdays: []types.EBackupSpanWeekday{
			types.BackupSpanWeekdays.Sunday,
		},
		MaximumNumberOfArchives: 1,
	}
	updateAutoBackupToMinExpected = &sacloud.AutoBackup{
		Name:                    updateAutoBackupToMinParam.Name,
		Availability:            types.Availabilities.Available,
		BackupSpanWeekdays:      updateAutoBackupToMinParam.BackupSpanWeekdays,
		MaximumNumberOfArchives: updateAutoBackupToMinParam.MaximumNumberOfArchives,
	}
)

func testAutoBackupCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Create(ctx, testZone, createAutoBackupParam)
}

func testAutoBackupRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testAutoBackupUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateAutoBackupParam)
}

func testAutoBackupUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.UpdateSettings(ctx, testZone, ctx.ID, updateAutoBackupSettingsParam)
}

func testAutoBackupUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateAutoBackupToMinParam)
}

func testAutoBackupDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
