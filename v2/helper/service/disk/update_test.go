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

	diskBuilder "github.com/sacloud/libsacloud/v2/helper/builder/disk"
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
	defer func() {
		diskOp.Delete(context.Background(), zone, disk.ID) // nolint
	}()

	cases := []struct {
		in     *UpdateRequest
		expect diskBuilder.Builder
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
				NoWait: true,
			},
			expect: &diskBuilder.ConnectedDiskBuilder{
				Name:        name + "-upd",
				Connection:  disk.Connection,
				Description: disk.Description,
				Tags:        disk.Tags,
				IconID:      disk.IconID,
				EditParameter: &diskBuilder.UnixEditRequest{
					HostName: "hostname",
					Password: "password",
				},
				Client: diskBuilder.NewBuildersAPIClient(caller),
				NoWait: true,
				ID:     disk.ID,
			},
		},
	}
	for _, tc := range cases {
		builder, err := tc.in.Builder(context.Background(), caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, builder)
	}
}
