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

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDatabaseService_CreateParameterToBuilder(t *testing.T) {
	caller := testutil.SingletonAPICaller()

	cases := []struct {
		in     *CreateRequest
		expect *Builder
	}{
		{
			in:     &CreateRequest{},
			expect: &Builder{Caller: caller},
		},
		{
			in: &CreateRequest{
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
			},
			expect: &Builder{
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
		},
	}

	for _, tc := range cases {
		got, err := tc.in.Builder(caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, got)
	}
}
