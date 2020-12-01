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

package server

import (
	"context"
	"testing"

	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestServerService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	zone := testutil.TestZone()
	name := testutil.ResourceName("service-update-server")
	caller := testutil.SingletonAPICaller()

	// setup
	diskOp := sacloud.NewDiskOp(caller)
	disk, err := diskOp.Create(ctx, zone, &sacloud.DiskCreateRequest{
		DiskPlanID:  types.DiskPlans.SSD,
		Connection:  types.DiskConnections.VirtIO,
		SizeMB:      20 * size.GiB,
		Name:        name,
		Description: "desc",
		Tags:        types.Tags{"tag1", "tag2"},
	}, []types.ID{})
	if err != nil {
		t.Fatal(err)
	}

	svc := New(caller)
	server, err := svc.CreateWithContext(ctx, &CreateRequest{
		Zone:            zone,
		Name:            name,
		Description:     "desc",
		Tags:            types.Tags{"tag1", "tag2"},
		CPU:             1,
		MemoryGB:        1,
		Commitment:      types.Commitments.Standard,
		Generation:      types.PlanGenerations.G100,
		InterfaceDriver: types.InterfaceDrivers.VirtIO,
		NetworkInterfaces: []*NetworkInterface{
			{Upstream: "shared"},
		},
		Disks: []*diskService.ApplyRequest{
			{
				Zone:        zone,
				ID:          disk.ID,
				Name:        name,
				Description: "desc",
				DiskPlanID:  types.DiskPlans.SSD,
				Connection:  types.DiskConnections.VirtIO,
				SizeGB:      20,
				NoWait:      true,
			},
		},
		NoWait: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		sacloud.NewServerOp(caller).Delete(ctx, zone, server.ID) // nolint
		diskOp.Delete(ctx, zone, disk.ID)                        // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				Zone:            zone,
				ID:              server.ID,
				Name:            &name,
				Description:     pointer.NewString("desc-upd"),
				Tags:            pointer.NewTags(types.Tags{"tag1-upd", "tag2-upd"}),
				CPU:             pointer.NewInt(2),
				MemoryGB:        pointer.NewInt(4),
				Commitment:      &types.Commitments.DedicatedCPU,
				Generation:      &types.PlanGenerations.G200,
				InterfaceDriver: &types.InterfaceDrivers.E1000,
				CDROMID:         pointer.NewID(102),
				PrivateHostID:   pointer.NewID(103),
				NetworkInterfaces: &[]*NetworkInterface{
					{Upstream: "shared"},
					{Upstream: "105"},
				},
				Disks: &[]*diskService.ApplyRequest{
					{
						Zone:        zone,
						ID:          disk.ID,
						Name:        name,
						Description: "desc",
						DiskPlanID:  types.DiskPlans.SSD,
						Connection:  types.DiskConnections.VirtIO,
						ServerID:    server.ID,
						SizeGB:      20,
						NoWait:      true,
					},
				},
				NoWait:        true,
				ForceShutdown: true,
			},
			expect: &ApplyRequest{
				Zone:            zone,
				ID:              server.ID,
				Name:            name,
				Description:     "desc-upd",
				Tags:            types.Tags{"tag1-upd", "tag2-upd"},
				CPU:             2,
				MemoryGB:        4,
				Commitment:      types.Commitments.DedicatedCPU,
				Generation:      types.PlanGenerations.G200,
				InterfaceDriver: types.InterfaceDrivers.E1000,
				BootAfterCreate: false,
				CDROMID:         102,
				PrivateHostID:   103,
				NetworkInterfaces: []*NetworkInterface{
					{Upstream: "shared"},
					{Upstream: "105"},
				},
				Disks: []*diskService.ApplyRequest{
					{
						Zone:        zone,
						ID:          disk.ID,
						Name:        name,
						Description: "desc",
						DiskPlanID:  types.DiskPlans.SSD,
						Connection:  types.DiskConnections.VirtIO,
						ServerID:    server.ID,
						SizeGB:      20,
						NoWait:      true,
					},
				},
				NoWait:        true,
				ForceShutdown: true,
			},
		},
	}

	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
