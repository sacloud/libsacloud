package server

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

type dummyServerReader struct {
	servers []*sacloud.Server
	err     error
}

func (r *dummyServerReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Server, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, s := range r.servers {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}

type dummyArchiveReader struct {
	archives []*sacloud.Archive
	err      error
}

func (r *dummyArchiveReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Archive, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, a := range r.archives {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}

type dummyDiskReader struct {
	disks []*sacloud.Disk
	err   error
}

func (r *dummyDiskReader) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Disk, error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, d := range r.disks {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, sacloud.NewAPIError(http.MethodGet, nil, "", http.StatusNotFound, nil)
}

func TestGetDefaultUserName(t *testing.T) {
	cases := []struct {
		msg           string
		id            types.ID
		reader        *SourceInfoReader
		expectedValue string
		expectedErr   error
	}{
		{
			msg: "server reader returns unexpected error",
			id:  1,
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			msg: "from coreos(container linux)",
			id:  1,
			reader: &SourceInfoReader{
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
							Tags:  types.Tags{"distro-coreos"},
						},
					},
				},
			},
			expectedValue: "core",
			expectedErr:   nil,
		},
		{
			msg: "from rancheros",
			id:  1,
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
			reader: &SourceInfoReader{
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
		actual, err := GetDefaultUserName(context.Background(), "tk1v", tc.reader, tc.id)
		if tc.expectedErr != nil {
			require.Equal(t, tc.expectedErr, err, tc.msg)
		} else {
			require.NoError(t, err, tc.msg)
		}
		require.Equal(t, tc.expectedValue, actual, tc.msg)
	}
}
