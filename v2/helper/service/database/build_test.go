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

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestCreateRequest_Validate(t *testing.T) {
	cases := []struct {
		in       *BuildRequest
		hasError bool
	}{
		{
			in:       &BuildRequest{},
			hasError: true,
		},
		{
			// mininum
			in: &BuildRequest{
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
			in: &BuildRequest{
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
			in: &BuildRequest{
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
			in: &BuildRequest{
				Zone:           "tk1a",
				Name:           "test",
				PlanID:         types.DatabasePlans.DB10GB,
				SwitchID:       1,
				IPAddresses:    []string{"192.168.0.11"},
				NetworkMaskLen: 16,
				DefaultRoute:   "192.168.0.1",
				SourceNetworks: []string{"192.168.0.1"}, // require cidr
				DatabaseType:   "mariadb",
				Username:       "hoge",
				Password:       "pass",
			},
			hasError: true,
		},
		{
			// replica user password missing
			in: &BuildRequest{
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
			in: &BuildRequest{
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
			in: &BuildRequest{
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
				SourceNetworks:        []string{"192.168.0.0/24", "192.168.1.0/24"},
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
