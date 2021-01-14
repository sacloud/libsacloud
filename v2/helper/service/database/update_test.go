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

package database

import (
	"context"
	"testing"

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDatabaseService_convertUpdateRequest(t *testing.T) {
	if testutil.IsAccTest() {
		t.Skip("This test runs only without TESTACC=1")
	}
	ctx := context.Background()
	zone := testutil.TestZone()
	name := testutil.ResourceName("database-service-update")
	caller := testutil.SingletonAPICaller()

	sw, err := sacloud.NewSwitchOp(caller).Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	builder := &databaseBuilder.Builder{
		PlanID:         types.DatabasePlans.DB10GB,
		SwitchID:       sw.ID,
		IPAddresses:    []string{"192.168.0.101"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Name,
			DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Version,
			DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Revision,
		},
		SourceID: 0,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.ToWebUI(true),
			ServicePort:     5432,
			SourceNetwork:   []string{"192.168.0.0/24"},
			DefaultUser:     "default",
			UserPassword:    "password1",
			ReplicaUser:     "replica",
			ReplicaPassword: "password2",
		},
		BackupSetting: &sacloud.DatabaseSettingBackup{
			Rotate:    0,
			Time:      "10:00",
			DayOfWeek: []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
		},
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
		},
		Name:         name,
		Description:  "description",
		Tags:         types.Tags{"tag1", "tag2"},
		SettingsHash: "",
		NoWait:       false,
		Parameters: map[string]interface{}{
			"max_connections": float64(100),
		},
		Client: databaseBuilder.NewAPIClient(caller),
	}
	db, err := builder.Build(ctx, zone)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		dbOp := sacloud.NewDatabaseOp(caller)
		if err := dbOp.Shutdown(ctx, zone, db.ID, &sacloud.ShutdownOption{Force: true}); err != nil {
			return
		}
		if _, err := wait.UntilDatabaseIsDown(ctx, dbOp, zone, db.ID); err != nil {
			return
		}
		if err := dbOp.Delete(ctx, zone, db.ID); err != nil {
			return
		}
		sacloud.NewSwitchOp(caller).Delete(ctx, zone, sw.ID) // nolint
	}()

	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				Zone:   zone,
				ID:     db.ID,
				Name:   pointer.NewString(db.Name + "-upd"),
				NoWait: true,
			},
			expect: &ApplyRequest{
				Zone:                  zone,
				ID:                    db.ID,
				Name:                  db.Name + "-upd",
				Description:           "description",
				Tags:                  types.Tags{"tag1", "tag2"},
				PlanID:                types.DatabasePlans.DB10GB,
				SwitchID:              sw.ID,
				IPAddresses:           []string{"192.168.0.101"},
				NetworkMaskLen:        24,
				DefaultRoute:          "192.168.0.1",
				Port:                  5432,
				SourceNetwork:         []string{"192.168.0.0/24"},
				DatabaseType:          types.RDBMSTypesPostgreSQL.String(),
				Username:              "default",
				Password:              "password1",
				EnableReplication:     true,
				ReplicaUserPassword:   "password2",
				EnableWebUI:           true,
				EnableBackup:          true,
				BackupWeekdays:        []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
				BackupStartTimeHour:   10,
				BackupStartTimeMinute: 0,
				Parameters: map[string]interface{}{
					"max_connections": float64(100),
				},
				NoWait: true,
			},
		},
		{
			in: &UpdateRequest{
				Zone:              zone,
				ID:                db.ID,
				EnableReplication: pointer.NewBool(false),
				EnableBackup:      pointer.NewBool(false),
				Parameters: &map[string]interface{}{
					"work_mem": float64(4096),
				},
				NoWait: true,
			},
			expect: &ApplyRequest{
				Zone:                  zone,
				ID:                    db.ID,
				Name:                  db.Name,
				Description:           "description",
				Tags:                  types.Tags{"tag1", "tag2"},
				PlanID:                types.DatabasePlans.DB10GB,
				SwitchID:              sw.ID,
				IPAddresses:           []string{"192.168.0.101"},
				NetworkMaskLen:        24,
				DefaultRoute:          "192.168.0.1",
				Port:                  5432,
				SourceNetwork:         []string{"192.168.0.0/24"},
				DatabaseType:          types.RDBMSTypesPostgreSQL.String(),
				Username:              "default",
				Password:              "password1",
				EnableReplication:     false,
				ReplicaUserPassword:   "password2",
				EnableWebUI:           true,
				EnableBackup:          false,
				BackupWeekdays:        []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
				BackupStartTimeHour:   10,
				BackupStartTimeMinute: 0,
				Parameters: map[string]interface{}{
					"max_connections": float64(100),
					"work_mem":        float64(4096),
				},
				NoWait: true,
			},
		},
	}

	for _, tc := range cases {
		got, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, got)
	}
}
