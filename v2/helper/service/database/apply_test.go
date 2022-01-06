// Copyright 2016-2022 The Libsacloud Authors
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

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestCreateRequest_Validate(t *testing.T) {
	cases := []struct {
		in       *ApplyRequest
		hasError bool
	}{
		{
			in:       &ApplyRequest{},
			hasError: true,
		},
		{
			// mininum
			in: &ApplyRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         types.DatabasePlans.DB10GB,
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.11"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: false,
		},
		{
			// invalid ip
			in: &ApplyRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         types.DatabasePlans.DB10GB,
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.999"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: true,
		},
		{
			// invalid ip (out of length)
			in: &ApplyRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         types.DatabasePlans.DB10GB,
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.11", "192.168.0.12", "192.168.0.13"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: true,
		},
		{
			// invalid source range
			in: &ApplyRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         types.DatabasePlans.DB10GB,
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.11"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				SourceNetwork:  []string{"192.168.0.1"}, // require cidr
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: true,
		},
		{
			// replica user password missing
			in: &ApplyRequest{
				Zone:              "tk1a",
				Name:              "test",
				PlanID:            types.DatabasePlans.DB10GB,
				SwitchID:          1,
				IPAddresses:       []string{"192.168.0.11"},
				NetworkMaskLen:    16,
				DefaultRoute:      "192.168.0.1",
				DatabaseType:      "mariadb",
				Username:          "hoge",
				Password:          "pass",
				EnableReplication: true,
			},
			hasError: true,
		},
		{
			// empty plan
			in: &ApplyRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         0, // plan is required
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.11"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: true,
		},
		{
			in: &ApplyRequest{
				Zone:                  "tk1a",
				Name:                  "test",
				Description:           "desc",
				Tags:                  types.Tags{"tag1"},
				IconID:                1,
				PlanID:                types.DatabasePlans.DB10GB,
				SwitchID:              1,
				IPAddresses:           []string{"192.168.0.11"},
				NetworkMaskLen:        16,
				DefaultRoute:          "192.168.0.1",
				Port:                  5432,
				SourceNetwork:         []string{"192.168.0.0/24", "192.168.1.0/24"},
				DatabaseType:          "mariadb",
				Username:              "hoge",
				Password:              "pass",
				EnableReplication:     true,
				ReplicaUserPassword:   "pass2",
				EnableWebUI:           true,
				EnableBackup:          true,
				BackupWeekdays:        []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
				BackupStartTimeHour:   10,
				BackupStartTimeMinute: 15,
			},
			hasError: false,
		},
	}
	for _, tc := range cases {
		err := tc.in.Validate()
		require.Equal(t, tc.hasError, err != nil, "with: %#v error: %s", tc.in, err)
	}
}

func TestDatabaseService_convertToActualBuilder(t *testing.T) {
	caller := testutil.SingletonAPICaller()
	cases := []struct {
		in     *Builder
		expect *databaseBuilder.Builder
	}{
		{
			in: &Builder{
				Caller:       caller,
				DatabaseType: types.RDBMSTypesPostgreSQL.String(),
			},
			expect: &databaseBuilder.Builder{
				CommonSetting: &sacloud.DatabaseSettingCommon{
					WebUI:           types.ToWebUI(false),
					ServicePort:     0,
					SourceNetwork:   nil,
					DefaultUser:     "",
					UserPassword:    "",
					ReplicaUser:     "",
					ReplicaPassword: "",
				},
				Conf: &sacloud.DatabaseRemarkDBConfCommon{
					DatabaseName:     types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Name,
					DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Version,
					DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Revision,
				},
				Client: databaseBuilder.NewAPIClient(caller),
			},
		},
		{
			in: &Builder{
				ID:                    0,
				Zone:                  "is1a",
				Name:                  "name",
				Description:           "description",
				Tags:                  types.Tags{"tag1", "tag2"},
				IconID:                types.ID(1),
				PlanID:                types.DatabasePlans.DB30GB,
				SwitchID:              types.ID(1),
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
				NoWait:                true,
				Caller:                caller,
			},
			expect: &databaseBuilder.Builder{
				PlanID:         types.DatabasePlans.DB30GB,
				SwitchID:       types.ID(1),
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
				Name:         "name",
				Description:  "description",
				Tags:         types.Tags{"tag1", "tag2"},
				IconID:       types.ID(1),
				SettingsHash: "",
				NoWait:       true,
				SetupOptions: nil,
				Client:       databaseBuilder.NewAPIClient(caller),
			},
		},
		{
			in: &Builder{
				ID:                    0,
				Zone:                  "is1a",
				Name:                  "name",
				Description:           "description",
				Tags:                  types.Tags{"tag1", "tag2"},
				IconID:                types.ID(1),
				PlanID:                types.DatabasePlans.DB30GB,
				SwitchID:              types.ID(1),
				IPAddresses:           []string{"192.168.0.101"},
				NetworkMaskLen:        24,
				DefaultRoute:          "192.168.0.1",
				Port:                  5432,
				SourceNetwork:         []string{"192.168.0.0/24"},
				DatabaseType:          types.RDBMSTypesPostgreSQL.String(),
				Username:              "default",
				Password:              "password1",
				EnableReplication:     false, // UPDATE
				ReplicaUserPassword:   "password2",
				EnableWebUI:           true,
				EnableBackup:          false, // UPDATE
				BackupWeekdays:        []types.EBackupSpanWeekday{types.BackupSpanWeekdays.Monday},
				BackupStartTimeHour:   10,
				BackupStartTimeMinute: 0,
				NoWait:                true,
				Caller:                caller,
			},
			expect: &databaseBuilder.Builder{
				PlanID:         types.DatabasePlans.DB30GB,
				SwitchID:       types.ID(1),
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
					ReplicaUser:     "",
					ReplicaPassword: "",
				},
				BackupSetting:      nil,
				ReplicationSetting: nil,
				Name:               "name",
				Description:        "description",
				Tags:               types.Tags{"tag1", "tag2"},
				IconID:             types.ID(1),
				SettingsHash:       "",
				NoWait:             true,
				SetupOptions:       nil,
				Client:             databaseBuilder.NewAPIClient(caller),
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.actualBuilder())
	}
}
