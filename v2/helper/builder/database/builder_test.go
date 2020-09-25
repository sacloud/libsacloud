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

package database

import (
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/helper/power"

	"github.com/sacloud/libsacloud/v2/helper/builder"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func getSetupOption() *builder.RetryableSetupParameter {
	if testutil.IsAccTest() {
		return nil
	}
	return &builder.RetryableSetupParameter{
		DeleteRetryInterval:       10 * time.Millisecond,
		ProvisioningRetryInterval: 10 * time.Millisecond,
		PollingInterval:           10 * time.Millisecond,
		NICUpdateWaitDuration:     10 * time.Millisecond,
	}
}

func TestBuilder_Build(t *testing.T) {
	var switchID types.ID
	var testZone = testutil.TestZone()

	testutil.RunCRUD(t, &testutil.CRUDTestCase{
		SetupAPICallerFunc: func() sacloud.APICaller {
			return testutil.SingletonAPICaller()
		},
		Parallel:          true,
		IgnoreStartupWait: true,
		Setup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)

			sw, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
				Name: testutil.ResourceName("database-builder"),
			})
			if err != nil {
				return err
			}
			switchID = sw.ID
			return nil
		},
		Create: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				builder := &Builder{
					PlanID:         types.DatabasePlans.DB10GB,
					SwitchID:       switchID,
					IPAddresses:    []string{"192.168.0.11"},
					NetworkMaskLen: 24,
					DefaultRoute:   "192.168.0.1",
					Conf: &sacloud.DatabaseRemarkDBConfCommon{
						DatabaseName:     types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Name,
						DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Version,
						DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Revision,
						DefaultUser:      "builder",
						UserPassword:     "builder-password-dummy",
					},
					CommonSetting: &sacloud.DatabaseSettingCommon{
						DefaultUser:     "builder",
						UserPassword:    "builder-password-dummy",
						ReplicaUser:     "",
						ReplicaPassword: "",
					},
					BackupSetting: &sacloud.DatabaseSettingBackup{
						Rotate:    7,
						Time:      "00:00",
						DayOfWeek: []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
					},
					ReplicationSetting: &sacloud.DatabaseReplicationSetting{},
					Name:               testutil.ResourceName("database-builder"),
					Description:        "description",
					Tags:               types.Tags{"tag1", "tag2"},
					SetupOptions:       getSetupOption(),
					Client:             NewAPIClient(caller),
				}
				return builder.Build(ctx, testZone)
			},
		},
		Read: &testutil.CRUDTestFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				dbOp := sacloud.NewDatabaseOp(caller)
				return dbOp.Read(ctx, testZone, ctx.ID)
			},
			CheckFunc: func(t testutil.TestT, ctx *testutil.CRUDTestContext, value interface{}) error {
				db := value.(*sacloud.Database)
				return testutil.DoAsserts(
					testutil.AssertNotNilFunc(t, db, "Database"),
					testutil.AssertNotNilFunc(t, db.Conf, "Database.Conf"),
				)
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
				dbOp := sacloud.NewDatabaseOp(caller)
				return dbOp.Delete(ctx, testZone, ctx.ID)
			},
		},
		Shutdown: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			dbOp := sacloud.NewDatabaseOp(caller)
			return power.ShutdownDatabase(ctx, dbOp, testZone, ctx.ID, true)
		},
		Cleanup: func(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
			swOp := sacloud.NewSwitchOp(caller)
			return swOp.Delete(ctx, testZone, switchID)
		},
	})
}
