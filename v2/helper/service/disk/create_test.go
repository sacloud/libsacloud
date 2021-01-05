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

package disk

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDiskService_convertCreateRequest(t *testing.T) {
	cases := []struct {
		in     *CreateRequest
		expect *ApplyRequest
	}{
		{
			in: &CreateRequest{
				Zone:       "is1a",
				Name:       "test",
				DiskPlanID: types.DiskPlans.SSD,
				Connection: types.DiskConnections.VirtIO,
				SizeGB:     20,
				OSType:     ostype.Ubuntu,
				EditParameter: &EditParameter{
					HostName: "hostname",
					Password: "password",
				},
				NoWait: true,
			},
			expect: &ApplyRequest{
				Zone:       "is1a",
				Name:       "test",
				DiskPlanID: types.DiskPlans.SSD,
				Connection: types.DiskConnections.VirtIO,
				SizeGB:     20,
				OSType:     ostype.Ubuntu,
				EditParameter: &EditParameter{
					HostName: "hostname",
					Password: "password",
				},
				NoWait: true,
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.expect, tc.in.ApplyRequest())
	}
}
