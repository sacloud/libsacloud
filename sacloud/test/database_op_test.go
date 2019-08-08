package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestDatabaseOpCRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: setupSwitchFunc("db",
			createDatabaseParam,
			createDatabaseExpected,
			patchDatabaseExpected,
			updateDatabaseExpected,
			updateDatabaseToFullExpected,
			updateDatabaseToMinExpected,
		),
		Create: &testutil.CRUDTestFunc{
			Func: testDatabaseCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDatabaseExpected,
				IgnoreFields: ignoreDatabaseFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testDatabaseRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createDatabaseExpected,
				IgnoreFields: ignoreDatabaseFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testDatabasePatch,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  patchDatabaseExpected,
					IgnoreFields: ignoreDatabaseFields,
				}),
			},
			{
				Func: testDatabaseUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDatabaseExpected,
					IgnoreFields: ignoreDatabaseFields,
				}),
			},
			{
				Func: testDatabaseUpdateToFull,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDatabaseToFullExpected,
					IgnoreFields: ignoreDatabaseFields,
				}),
			},
			{
				Func: testDatabaseUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDatabaseToMinExpected,
					IgnoreFields: ignoreDatabaseFields,
				}),
			},
		},
		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewDatabaseOp(caller)
			return client.Shutdown(ctx, testZone, ctx.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testDatabaseDelete,
		},

		Cleanup: cleanupSwitchFunc("db"),
	})
}

var (
	ignoreDatabaseFields = []string{
		"ID",
		"Class",
		"Tags", // Create(POST)時は指定したタグが返ってくる。その後利用可能になったらデータベースの種類に応じて@MariaDBxxxのようなタグが付与される
		"Availability",
		"InstanceStatus",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"ZoneID",
		"CreatedAt",
		"ModifiedAt",
		"SettingsHash",
	}

	createDatabaseParam = &sacloud.DatabaseCreateRequest{
		PlanID:         types.DatabasePlans.DB10GB,
		IPAddresses:    []string{"192.168.0.11"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           testutil.ResourceName("db"),
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},

		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     types.RDBMSVersions[types.RDBMSTypesMariaDB].Name,
			DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesMariaDB].Version,
			DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesMariaDB].Revision,
			DefaultUser:      "exa.mple",
			UserPassword:     "LibsacloudExamplePassword01",
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:  5432,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword01",
		},
	}
	createDatabaseExpected = &sacloud.Database{
		Name:           createDatabaseParam.Name,
		Description:    createDatabaseParam.Description,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting:  createDatabaseParam.CommonSetting,
	}
	patchDatabaseParam = &sacloud.DatabasePatchRequest{
		PatchEmptyToDescription: true,
		CommonSetting: &sacloud.DatabaseSettingCommonUpdate{
			ServicePort: 54321,
		},
	}
	patchDatabaseExpected = &sacloud.Database{
		Name:           createDatabaseParam.Name,
		Description:    "",
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:  54321,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword01",
		},
	}
	updateDatabaseParam = &sacloud.DatabaseUpdateRequest{
		Name:        testutil.ResourceName("db-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		CommonSetting: &sacloud.DatabaseSettingCommonUpdate{
			ServicePort:  5432,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword02",
		},
	}
	updateDatabaseExpected = &sacloud.Database{
		Name:           updateDatabaseParam.Name,
		Description:    updateDatabaseParam.Description,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:  5432,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword02",
		},
	}
	updateDatabaseToFullParam = &sacloud.DatabaseUpdateRequest{
		Name:        testutil.ResourceName("db-to-full"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		BackupSetting: &sacloud.DatabaseSettingBackup{
			Rotate: 3,
			Time:   "00:00",
			DayOfWeek: []types.EBackupSpanWeekday{
				types.BackupSpanWeekdays.Sunday,
				types.BackupSpanWeekdays.Monday,
			},
		},
		CommonSetting: &sacloud.DatabaseSettingCommonUpdate{
			ServicePort:   54321,
			DefaultUser:   "exa.mple",
			UserPassword:  "LibsacloudExamplePassword03",
			SourceNetwork: []string{"192.168.11.0/24", "192.168.12.0/24"},
		},
		IconID: testIconID,
	}
	updateDatabaseToFullExpected = &sacloud.Database{
		Name:           updateDatabaseToFullParam.Name,
		Description:    updateDatabaseToFullParam.Description,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:   54321,
			DefaultUser:   "exa.mple",
			UserPassword:  "LibsacloudExamplePassword03",
			SourceNetwork: []string{"192.168.11.0/24", "192.168.12.0/24"},
		},
		BackupSetting: updateDatabaseToFullParam.BackupSetting,
		IconID:        updateDatabaseToFullParam.IconID,
	}
	updateDatabaseToMinParam = &sacloud.DatabaseUpdateRequest{
		Name: testutil.ResourceName("db-to-min"),
		CommonSetting: &sacloud.DatabaseSettingCommonUpdate{
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword04",
		},
	}
	updateDatabaseToMinExpected = &sacloud.Database{
		Name:           updateDatabaseToMinParam.Name,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword04",
		},
	}
)

func testDatabaseCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Create(ctx, testZone, createDatabaseParam)
}

func testDatabaseRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testDatabasePatch(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Patch(ctx, testZone, ctx.ID, patchDatabaseParam)
}

func testDatabaseUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDatabaseParam)
}

func testDatabaseUpdateToFull(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDatabaseToFullParam)
}

func testDatabaseUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDatabaseToMinParam)
}

func testDatabaseDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDatabaseOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
