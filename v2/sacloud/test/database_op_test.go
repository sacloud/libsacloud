// Copyright 2016-2021 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/helper/power"
	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestDatabaseOpCRUD(t *testing.T) {
	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup: setupSwitchFunc("db",
			createDatabaseParam,
			createDatabaseExpected,
			updateDatabaseSettingsExpected,
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
				Func: testDatabaseUpdateSettings,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updateDatabaseSettingsExpected,
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
			return power.ShutdownDatabase(ctx, client, testZone, ctx.ID, true)
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
			DatabaseRevision: "10.4.12",
			DefaultUser:      "exa.mple",
			UserPassword:     "LibsacloudExamplePassword01",
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:     5432,
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword01",
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
		},
	}
	createDatabaseExpected = &sacloud.Database{
		Name:               createDatabaseParam.Name,
		Description:        createDatabaseParam.Description,
		Availability:       types.Availabilities.Available,
		PlanID:             createDatabaseParam.PlanID,
		DefaultRoute:       createDatabaseParam.DefaultRoute,
		NetworkMaskLen:     createDatabaseParam.NetworkMaskLen,
		IPAddresses:        createDatabaseParam.IPAddresses,
		Conf:               createDatabaseParam.Conf,
		CommonSetting:      createDatabaseParam.CommonSetting,
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
	}
	updateDatabaseSettingsParam = &sacloud.DatabaseUpdateSettingsRequest{
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:     54322,
			DefaultUser:     "exa.mple.upd",
			UserPassword:    "LibsacloudExamplePassword01up1",
			ReplicaUser:     "replica-upd",
			ReplicaPassword: "replica-user-password-upd",
		},
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
	}
	updateDatabaseSettingsExpected = &sacloud.Database{
		Name:           createDatabaseParam.Name,
		Description:    createDatabaseParam.Description,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:     54322,
			DefaultUser:     "exa.mple.upd",
			UserPassword:    "LibsacloudExamplePassword01up1",
			ReplicaUser:     "replica-upd",
			ReplicaPassword: "replica-user-password-upd",
		},
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
	}
	updateDatabaseParam = &sacloud.DatabaseUpdateRequest{
		Name:        testutil.ResourceName("db-upd"),
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:     5432,
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword02",
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
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
			ServicePort:     5432,
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword02",
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
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
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:     54321,
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword03",
			SourceNetwork:   []string{"192.168.11.0/24", "192.168.12.0/24"},
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
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
			ServicePort:     54321,
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword03",
			SourceNetwork:   []string{"192.168.11.0/24", "192.168.12.0/24"},
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		BackupSetting:      updateDatabaseToFullParam.BackupSetting,
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
		IconID:             updateDatabaseToFullParam.IconID,
	}
	updateDatabaseToMinParam = &sacloud.DatabaseUpdateRequest{
		Name: testutil.ResourceName("db-to-min"),
		CommonSetting: &sacloud.DatabaseSettingCommon{
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword04",
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
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
			DefaultUser:     "exa.mple",
			UserPassword:    "LibsacloudExamplePassword04",
			ReplicaUser:     "replica",
			ReplicaPassword: "replica-user-password",
		},
		ReplicationSetting: createDatabaseParam.ReplicationSetting,
	}
)

func testDatabaseCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	db, err := client.Create(ctx, testZone, createDatabaseParam)
	if err != nil {
		return nil, err
	}
	return wait.UntilDatabaseIsUp(ctx, client, testZone, db.ID)
}

func testDatabaseRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testDatabaseUpdateSettings(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.UpdateSettings(ctx, testZone, ctx.ID, updateDatabaseSettingsParam)
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
