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

package nfs

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestNFSService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("nfs-service")
	zone := testutil.TestZone()

	// setup
	swOp := sacloud.NewSwitchOp(caller)
	sw, err := swOp.Create(ctx, zone, &sacloud.SwitchCreateRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}

	current, err := New(caller).CreateWithContext(ctx, &CreateRequest{
		Zone:           zone,
		Name:           name,
		Description:    "desc",
		Tags:           types.Tags{"tag1", "tag2"},
		SwitchID:       sw.ID,
		Plan:           types.NFSPlans.SSD,
		Size:           100,
		IPAddresses:    []string{"192.168.0.101"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		nfsOp := sacloud.NewNFSOp(caller)
		nfsOp.Shutdown(ctx, zone, current.ID, &sacloud.ShutdownOption{Force: true}) // nolint
		wait.UntilNFSIsDown(ctx, nfsOp, zone, current.ID)                           // nolint
		nfsOp.Delete(ctx, zone, current.ID)                                         // nolint
		swOp.Delete(ctx, zone, sw.ID)                                               // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				Zone: zone,
				ID:   current.ID,
				Name: pointer.NewString(current.Name + "-upd"),
			},
			expect: &ApplyRequest{
				ID:             current.ID,
				Zone:           zone,
				Name:           current.Name + "-upd",
				Description:    current.Description,
				Tags:           current.Tags,
				IconID:         current.IconID,
				SwitchID:       current.SwitchID,
				Plan:           types.NFSPlans.SSD,
				Size:           100,
				IPAddresses:    current.IPAddresses,
				NetworkMaskLen: current.NetworkMaskLen,
				DefaultRoute:   current.DefaultRoute,
			},
		},
	}

	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
