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

package server

import (
	"testing"

	diskBuilder "github.com/sacloud/libsacloud/v2/helper/builder/disk"
	serverBuilder "github.com/sacloud/libsacloud/v2/helper/builder/server"
	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestServerService_convertApplyRequest(t *testing.T) {
	caller := testutil.SingletonAPICaller()

	cases := []struct {
		in     *ApplyRequest
		expect *serverBuilder.Builder
	}{
		{
			in: &ApplyRequest{
				Zone:            "tk1a",
				ID:              101,
				Name:            "name",
				Description:     "desc",
				Tags:            types.Tags{"tag1", "tag2"},
				IconID:          102,
				CPU:             2,
				MemoryGB:        4,
				Commitment:      types.Commitments.DedicatedCPU,
				Generation:      types.PlanGenerations.Default,
				InterfaceDriver: types.InterfaceDrivers.VirtIO,
				BootAfterCreate: true,
				CDROMID:         103,
				PrivateHostID:   104,
				NetworkInterfaces: []*NetworkInterface{
					{Upstream: "shared", PacketFilterID: 105},
					{Upstream: "106", PacketFilterID: 107, UserIPAddress: "192.168.0.101"},
				},
				Disks: []*diskService.ApplyRequest{
					{
						Zone:        "tk1a",
						ID:          201,
						Name:        "name",
						Description: "desc",
						SizeGB:      20,
					},
				},
				NoWait:        true,
				ForceShutdown: true,
			},
			expect: &serverBuilder.Builder{
				Name:            "name",
				CPU:             2,
				MemoryGB:        4,
				Commitment:      types.Commitments.DedicatedCPU,
				Generation:      types.PlanGenerations.Default,
				InterfaceDriver: types.InterfaceDrivers.VirtIO,
				Description:     "desc",
				IconID:          102,
				Tags:            types.Tags{"tag1", "tag2"},
				BootAfterCreate: true,
				CDROMID:         103,
				PrivateHostID:   104,
				NIC:             &serverBuilder.SharedNICSetting{PacketFilterID: 105},
				AdditionalNICs: []serverBuilder.AdditionalNICSettingHolder{
					&serverBuilder.ConnectedNICSetting{
						SwitchID:         106,
						DisplayIPAddress: "192.168.0.101",
						PacketFilterID:   107,
					},
				},
				DiskBuilders: []diskBuilder.Builder{
					&diskBuilder.ConnectedDiskBuilder{
						ID:          201,
						Name:        "name",
						Description: "desc",
						Client:      diskBuilder.NewBuildersAPIClient(caller),
					},
				},
				Client:        serverBuilder.NewBuildersAPIClient(caller),
				NoWait:        true,
				ServerID:      101,
				ForceShutdown: true,
			},
		},
	}

	for _, tc := range cases {
		builder, err := tc.in.Builder(caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, builder)
	}
}
