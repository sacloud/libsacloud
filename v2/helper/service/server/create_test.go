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

package server

import (
	"testing"

	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestServerService_convertCreateRequest(t *testing.T) {
	zone := testutil.TestZone()
	name := testutil.ResourceName("service-server-create")

	cases := []struct {
		in     *CreateRequest
		expect *ApplyRequest
	}{
		{
			in: &CreateRequest{
				Zone:            zone,
				Name:            name,
				Description:     "desc",
				Tags:            types.Tags{"tag1", "tag2"},
				IconID:          101,
				CPU:             22,
				MemoryGB:        4,
				Commitment:      types.Commitments.DedicatedCPU,
				Generation:      types.PlanGenerations.Default,
				InterfaceDriver: types.InterfaceDrivers.VirtIO,
				BootAfterCreate: true,
				CDROMID:         102,
				PrivateHostID:   103,
				NetworkInterfaces: []*NetworkInterface{
					{Upstream: "shared", PacketFilterID: 104},
					{Upstream: "105", UserIPAddress: "192.168.0.1"},
				},
				Disks: []*diskService.ApplyRequest{
					{
						Zone:        zone,
						ID:          201,
						Name:        name,
						Description: "desc",
						DiskPlanID:  types.DiskPlans.SSD,
						Connection:  types.DiskConnections.VirtIO,
						SizeGB:      20,
						NoWait:      true,
					},
				},
				NoWait: true,
			},
			expect: &ApplyRequest{
				Zone:            zone,
				Name:            name,
				Description:     "desc",
				Tags:            types.Tags{"tag1", "tag2"},
				IconID:          101,
				CPU:             22,
				MemoryGB:        4,
				GPU:             0,
				Commitment:      types.Commitments.DedicatedCPU,
				Generation:      types.PlanGenerations.Default,
				InterfaceDriver: types.InterfaceDrivers.VirtIO,
				BootAfterCreate: true,
				CDROMID:         102,
				PrivateHostID:   103,
				NetworkInterfaces: []*NetworkInterface{
					{Upstream: "shared", PacketFilterID: 104},
					{Upstream: "105", UserIPAddress: "192.168.0.1"},
				},
				Disks: []*diskService.ApplyRequest{
					{
						Zone:        zone,
						ID:          201,
						Name:        name,
						Description: "desc",
						DiskPlanID:  types.DiskPlans.SSD,
						Connection:  types.DiskConnections.VirtIO,
						SizeGB:      20,
						NoWait:      true,
					},
				},
				NoWait: true,
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.ApplyRequest())
	}
}
