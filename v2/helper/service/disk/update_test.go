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

package disk

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/wait"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"

	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
)

func TestDiskUpdateRequest_Update(t *testing.T) {
	v := &UpdateRequest{
		Zone:        "tk1a",
		ID:          1,
		Name:        pointer.NewString(""),
		Description: pointer.NewString(""),
	}
	err := v.Validate()
	if err == nil {
		t.Fatalf("no error with: %#v", v)
	}
}

func TestDiskService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("disk-service-update")
	zone := testutil.TestZone()

	diskOp := sacloud.NewDiskOp(caller)
	disk, err := diskOp.Create(context.Background(), zone, &sacloud.DiskCreateRequest{
		DiskPlanID: types.DiskPlans.SSD,
		Connection: types.DiskConnections.VirtIO,
		SizeMB:     20 * size.GiB,
		Name:       name,
	}, []types.ID{})
	if err != nil {
		t.Fatal(err)
	}
	v, err := wait.UntilDiskIsReady(ctx, diskOp, zone, disk.ID)
	if err != nil {
		t.Fatal(err)
	}
	disk = v

	defer func() {
		diskOp.Delete(context.Background(), zone, disk.ID) // nolint
	}()

	var cases = []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				Zone: zone,
				ID:   disk.ID,
				Name: pointer.NewString(name + "-upd"),
				EditParameter: &EditParameter{
					HostName: "hostname",
					Password: "password",
				},
				NoWait: false,
			},
			expect: &ApplyRequest{
				Zone:            zone,
				ID:              disk.ID,
				Name:            name + "-upd",
				Description:     disk.Description,
				Tags:            disk.Tags,
				IconID:          disk.IconID,
				DiskPlanID:      disk.DiskPlanID,
				Connection:      disk.Connection,
				SourceDiskID:    disk.SourceDiskID,
				SourceArchiveID: disk.SourceArchiveID,
				ServerID:        disk.ServerID,
				SizeGB:          disk.GetSizeGB(),
				EditParameter: &EditParameter{
					HostName: "hostname",
					Password: "password",
				},
				NoWait: false,
			},
		},
	}
	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(context.Background(), caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
