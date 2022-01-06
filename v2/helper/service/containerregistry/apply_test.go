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

package containerregistry

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestContainerRegistryService_convertApplyRequest(t *testing.T) {
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("container-registry-service")

	cases := []struct {
		in     *ApplyRequest
		expect *Builder
	}{
		{
			in: &ApplyRequest{
				Name:           name,
				Description:    "desc",
				Tags:           types.Tags{"tag1", "tag2"},
				AccessLevel:    types.ContainerRegistryAccessLevels.ReadWrite,
				VirtualDomain:  "container-registry.test.libsacloud.com",
				SubDomainLabel: name,
				Users: []*User{
					{
						UserName:   "user1",
						Password:   "password1",
						Permission: types.ContainerRegistryPermissions.ReadWrite,
					},
				},
				SettingsHash: "aaaaaaaa",
			},
			expect: &Builder{
				ID:             0,
				Name:           name,
				Description:    "desc",
				Tags:           types.Tags{"tag1", "tag2"},
				AccessLevel:    types.ContainerRegistryAccessLevels.ReadWrite,
				VirtualDomain:  "container-registry.test.libsacloud.com",
				SubDomainLabel: name,
				Users: []*User{
					{
						UserName:   "user1",
						Password:   "password1",
						Permission: types.ContainerRegistryPermissions.ReadWrite,
					},
				},
				SettingsHash: "aaaaaaaa",
				Client:       sacloud.NewContainerRegistryOp(caller),
			},
		},
	}

	for _, tc := range cases {
		builder, err := tc.in.Builder(caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, builder)
	}
}
