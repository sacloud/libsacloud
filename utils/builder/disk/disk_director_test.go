// Copyright 2016-2019 The Libsacloud Authors
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

func TestDiskDirector_Builder(t *testing.T) {
	cases := []struct {
		name string
		in   *DiskDirector
		out  DiskBuilder
	}{
		{
			name: "blank disk",
			in:   &DiskDirector{},
			out:  &BlankDiskBuilder{},
		},
		{
			name: "connect disk",
			in: &DiskDirector{
				DiskID: types.ID(1),
			},
			out: &ConnectedDiskBuilder{
				DiskID: types.ID(1),
			},
		},
		{
			name: "from archive",
			in: &DiskDirector{
				SourceArchiveID: types.ID(1),
			},
			out: &FromDiskOrArchiveDiskBuilder{
				SourceArchiveID: types.ID(1),
			},
		},
		{
			name: "from disk",
			in: &DiskDirector{
				SourceDiskID: types.ID(1),
			},
			out: &FromDiskOrArchiveDiskBuilder{
				SourceDiskID: types.ID(1),
			},
		},
		{
			name: "unix",
			in: &DiskDirector{
				OSType: ostype.CentOS,
				EditParameter: &DiskEditRequest{
					HostName: "example",
				},
			},
			out: &FromUnixDiskBuilder{
				OSType: ostype.CentOS,
				EditParameter: &UnixDiskEditRequest{
					HostName: "example",
				},
			},
		},
		{
			name: "windows",
			in: &DiskDirector{
				OSType: ostype.Windows2019,
				EditParameter: &DiskEditRequest{
					IPAddress: "192.2.0.1",
				},
			},
			out: &FromWindowsDiskBuilder{
				OSType: ostype.Windows2019,
				EditParameter: &WindowsDiskEditRequest{
					IPAddress: "192.2.0.1",
				},
			},
		},
		{
			name: "other",
			in: &DiskDirector{
				OSType: ostype.OPNsense,
			},
			out: &FromFixedArchiveDiskBuilder{
				OSType: ostype.OPNsense,
			},
		},
	}

	for _, tt := range cases {
		builder := tt.in.Builder()
		require.Equal(t, tt.out, builder, tt.name)
	}
}
