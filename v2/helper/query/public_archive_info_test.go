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

package query

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestCanEditDisk(t *testing.T) {
	cases := []struct {
		msg            string
		id             types.ID
		readers        *ArchiveSourceReader
		expectedResult bool
		expectedErr    error
	}{
		{
			msg: "disk reader returns unexpected error",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{},
				DiskReader: &dummyDiskReader{
					err: errors.New("dummy"),
				},
			},
			expectedErr: errors.New("dummy"),
		},
		{
			msg: "from empty disk",
			id:  2,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{ID: 2},
					},
				},
			},
			expectedResult: false,
		},
		{
			msg: "disk copied from disk",
			id:  2,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:   1,
							Tags: types.Tags{"os-linux"},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{ID: 2, SourceDiskID: 3},
						{ID: 3, SourceArchiveID: 1},
					},
				},
			},
			expectedResult: true,
		},
		{
			msg: "archive reader returns error",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					err: errors.New("dummy"),
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedErr: errors.New("dummy"),
		},
		{
			msg: "archive with bundle info",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID: 1,
							BundleInfo: &sacloud.BundleInfo{
								HostClass: bundleInfoWindowsHostClass,
							},
							Tags: types.Tags{"os-linux"},
						},
					},
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedResult: false,
		},
		{
			msg: "sophos UTM: service class",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID: 1,
							BundleInfo: &sacloud.BundleInfo{
								ServiceClass: "hoge/dummy/sophosutm",
							},
							Tags: types.Tags{"os-linux"},
						},
					},
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedResult: false,
		},
		{
			msg: "sophos UTM: tag",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:   1,
							Tags: types.Tags{"pkg-sophosutm"},
						},
					},
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedResult: false,
		},
		{
			msg: "OPNsense",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:   1,
							Tags: types.Tags{"distro-opnsense"},
						},
					},
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedResult: false,
		},
		{
			msg: "Netwiser VE",
			id:  1,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:   1,
							Tags: types.Tags{"distro-netwiserve"},
						},
					},
				},
				DiskReader: &dummyDiskReader{},
			},
			expectedResult: false,
		},
		{
			msg: "Nested",
			id:  4,
			readers: &ArchiveSourceReader{
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:   1,
							Tags: types.Tags{"os-unix"},
						},
						{ID: 2, SourceDiskID: 5},
						{ID: 3, SourceArchiveID: 1},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{ID: 4, SourceArchiveID: 2},
						{ID: 5, SourceDiskID: 6},
						{ID: 6, SourceArchiveID: 3},
					},
				},
			},
			expectedResult: true,
		},
	}

	for _, tc := range cases {
		res, err := CanEditDisk(context.Background(), "tk1v", tc.readers, tc.id)
		if tc.expectedErr != nil {
			require.Equal(t, tc.expectedErr, err, tc.msg)
		} else {
			require.NoError(t, err, tc.msg)
		}
		require.Equal(t, tc.expectedResult, res, tc.msg)
	}
}
