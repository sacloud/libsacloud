package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestAutoBackupOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			diskOp := sacloud.NewDiskOp(caller)
			disk, err := diskOp.Create(context.Background(), testZone, &sacloud.DiskCreateRequest{
				Name:       "libsacloud-disk-with-autobackup",
				SizeMB:     20 * 1024,
				DiskPlanID: types.ID(4), //SSD
			})
			if !assert.NoError(t, err) {
				return err
			}

			_, err = sacloud.WaiterForReady(func() (interface{}, error) {
				return diskOp.Read(context.Background(), testZone, disk.ID)
			}).WaitForState(context.Background())
			if !assert.NoError(t, err) {
				return err
			}

			testContext.Values["autobackup/disk"] = disk.ID
			createAutoBackupParam.DiskID = disk.ID
			createAutoBackupExpected.DiskID = disk.ID
			updateAutoBackupExpected.DiskID = disk.ID
			return err
		},

		Create: &CRUDTestFunc{
			Func: testAutoBackupCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createAutoBackupExpected,
				IgnoreFields: ignoreAutoBackupFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testAutoBackupRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createAutoBackupExpected,
				IgnoreFields: ignoreAutoBackupFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testAutoBackupUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateAutoBackupExpected,
					IgnoreFields: ignoreAutoBackupFields,
				}),
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testAutoBackupDelete,
		},

		Cleanup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			diskID, ok := testContext.Values["autobackup/disk"]
			if !ok {
				return nil
			}

			diskOp := sacloud.NewDiskOp(caller)
			return diskOp.Delete(context.Background(), testZone, diskID.(types.ID))
		},
	})
}

var (
	ignoreAutoBackupFields = []string{
		"ID",
		"Class",
		"SettingsHash",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
		"DiskID",
		"AccountID",
		"ZoneID",
		"ZoneName",
	}
	createAutoBackupParam = &sacloud.AutoBackupCreateRequest{
		Name:        "libsacloud-auto-backup",
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
		Name:        "libsacloud-auto-backup-upd",
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		BackupSpanWeekdays: []types.EBackupSpanWeekday{
			types.BackupSpanWeekdays.Monday,
			types.BackupSpanWeekdays.Tuesday,
			types.BackupSpanWeekdays.Wednesday,
			types.BackupSpanWeekdays.Thursday,
		},
		MaximumNumberOfArchives: 3,
	}
	updateAutoBackupExpected = &sacloud.AutoBackup{
		Name:                    updateAutoBackupParam.Name,
		Description:             updateAutoBackupParam.Description,
		Tags:                    updateAutoBackupParam.Tags,
		Availability:            types.Availabilities.Available,
		BackupSpanWeekdays:      updateAutoBackupParam.BackupSpanWeekdays,
		MaximumNumberOfArchives: updateAutoBackupParam.MaximumNumberOfArchives,
	}
)

func testAutoBackupCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Create(context.Background(), testZone, createAutoBackupParam)
}

func testAutoBackupRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testAutoBackupUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateAutoBackupParam)
}

func testAutoBackupDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewAutoBackupOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
