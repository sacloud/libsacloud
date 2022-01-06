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

package nfs

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestNFSService_convertApplyRequest(t *testing.T) {
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("nfs-service")
	zone := testutil.TestZone()

	cases := []struct {
		in     *ApplyRequest
		expect *Builder
	}{
		{
			in: &ApplyRequest{
				ID:             101,
				Zone:           zone,
				Name:           name,
				Description:    "desc",
				Tags:           types.Tags{"tag1", "tag2"},
				SwitchID:       102,
				Plan:           types.NFSPlans.SSD,
				Size:           100,
				IPAddresses:    []string{"192.168.0.101"},
				NetworkMaskLen: 24,
				DefaultRoute:   "192.168.0.1",
				NoWait:         true,
			},
			expect: &Builder{
				ID:             101,
				Zone:           zone,
				Name:           name,
				Description:    "desc",
				Tags:           types.Tags{"tag1", "tag2"},
				SwitchID:       102,
				Plan:           types.NFSPlans.SSD,
				Size:           100,
				IPAddresses:    []string{"192.168.0.101"},
				NetworkMaskLen: 24,
				DefaultRoute:   "192.168.0.1",
				NoWait:         true,
				Caller:         caller,
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.Builder(caller))
	}
}
