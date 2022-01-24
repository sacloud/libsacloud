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

func TestGetDefaultUserName(t *testing.T) {
	cases := []struct {
		msg           string
		id            types.ID
		reader        *ServerSourceReader
		expectedValue string
		expectedErr   error
	}{
		{
			msg: "server reader returns unexpected error",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					err: errors.New("dummy"),
				},
				ArchiveReader: &dummyArchiveReader{},
				DiskReader:    &dummyDiskReader{},
			},
			expectedValue: "",
			expectedErr:   errors.New("dummy"),
		},
		{
			msg: "diskless server",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{ID: 1},
					},
				},
				ArchiveReader: &dummyArchiveReader{},
				DiskReader:    &dummyDiskReader{},
			},
			expectedValue: "",
			expectedErr:   nil,
		},
		{
			msg: "disk reader returns unexpected error",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{},
				DiskReader: &dummyDiskReader{
					err: errors.New("dummy"),
				},
			},
			expectedValue: "",
			expectedErr:   errors.New("dummy"),
		},
		{
			msg: "disk with source disk",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:           2,
							SourceDiskID: 3,
						},
						{ID: 3}, // from empty disk
					},
				},
			},
			expectedValue: "",
			expectedErr:   nil,
		},
		{
			msg: "archive reader returns unexpected error",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:              2,
							SourceArchiveID: 3,
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{
					err: errors.New("dummy"),
				},
			},
			expectedValue: "",
			expectedErr:   errors.New("dummy"),
		},
		{
			msg: "from ubuntu",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:              2,
							SourceArchiveID: 3,
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:    3,
							Scope: types.Scopes.Shared,
							Tags:  types.Tags{"distro-ubuntu"},
						},
					},
				},
			},
			expectedValue: "ubuntu",
			expectedErr:   nil,
		},
		{
			msg: "from rancheros",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:              2,
							SourceArchiveID: 3,
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:    3,
							Scope: types.Scopes.Shared,
							Tags:  types.Tags{"distro-rancheros"},
						},
					},
				},
			},
			expectedValue: "rancher",
			expectedErr:   nil,
		},
		{
			msg: "from k3os",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:              2,
							SourceArchiveID: 3,
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:    3,
							Scope: types.Scopes.Shared,
							Tags:  types.Tags{"distro-k3os"},
						},
					},
				},
			},
			expectedValue: "rancher",
			expectedErr:   nil,
		},
		{
			msg: "nested",
			id:  1,
			reader: &ServerSourceReader{
				ServerReader: &dummyServerReader{
					servers: []*sacloud.Server{
						{
							ID: 1,
							Disks: []*sacloud.ServerConnectedDisk{
								{ID: 2},
							},
						},
					},
				},
				DiskReader: &dummyDiskReader{
					disks: []*sacloud.Disk{
						{
							ID:              2,
							SourceArchiveID: 5,
						},
						{
							ID:           3,
							SourceDiskID: 4,
						},
						{
							ID:              4,
							SourceArchiveID: 6,
						},
					},
				},
				ArchiveReader: &dummyArchiveReader{
					archives: []*sacloud.Archive{
						{
							ID:           5,
							SourceDiskID: 3,
						},
						{
							ID:              6,
							SourceArchiveID: 7,
						},
						{
							ID:    7,
							Scope: types.Scopes.Shared,
							Tags:  types.Tags{"distro-ubuntu"},
						},
					},
				},
			},
			expectedValue: "ubuntu",
			expectedErr:   nil,
		},
	}

	for _, tc := range cases {
		actual, err := ServerDefaultUserName(context.Background(), "tk1v", tc.reader, tc.id)
		if tc.expectedErr != nil {
			require.Equal(t, tc.expectedErr, err, tc.msg)
		} else {
			require.NoError(t, err, tc.msg)
		}
		require.Equal(t, tc.expectedValue, actual, tc.msg)
	}
}
