package archive

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

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

func TestCanEditDisk(t *testing.T) {

	cases := []struct {
		msg            string
		id             types.ID
		readers        *SourceInfoReader
		expectedResult bool
		expectedErr    error
	}{
		{
			msg: "disk reader returns unexpected error",
			id:  1,
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
			readers: &SourceInfoReader{
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
