package server

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/libsacloud/v2/utils/server/ostype"
	"github.com/stretchr/testify/require"
)

func TestDiskFromUnixRequest_Validate(t *testing.T) {
	cases := []struct {
		msg    string
		in     *DiskFromUnixRequest
		client *BuildersAPIClient
		err    error
	}{
		{
			msg: "invalid ostype",
			in: &DiskFromUnixRequest{
				OSType: ostype.UnixPublicArchiveType(-1),
			},
			err: fmt.Errorf("invalid OSType: %s", ostype.UnixPublicArchiveType(-1)),
		},
		{
			msg: "size not found",
			in: &DiskFromUnixRequest{
				OSType: ostype.CentOS,
				PlanID: types.DiskPlans.SSD,
				SizeGB: 1,
			},
			client: &BuildersAPIClient{
				DiskPlan: &dummyDiskPlanReader{
					diskPlan: &sacloud.DiskPlan{
						ID:   types.DiskPlans.SSD,
						Name: "SSDプラン",
						Size: []*sacloud.DiskPlanSizeInfo{
							{
								Availability: types.Availabilities.Available,
								SizeMB:       0,
							},
						},
					},
				},
			},
			err: fmt.Errorf("disk plan[SSDプラン:1GB] is not found"),
		},
		{
			msg: "invalid disk edit parameter",
			in: &DiskFromUnixRequest{
				OSType: ostype.CentOS,
				PlanID: types.DiskPlans.SSD,
				SizeGB: 1,
				EditParameter: &UnixDiskEditRequest{
					NoteIDs: []types.ID{1},
				},
			},
			client: &BuildersAPIClient{
				DiskPlan: &dummyDiskPlanReader{
					diskPlan: &sacloud.DiskPlan{
						ID:   types.DiskPlans.SSD,
						Name: "SSDプラン",
						Size: []*sacloud.DiskPlanSizeInfo{
							{
								Availability: types.Availabilities.Available,
								SizeMB:       1024,
							},
						},
					},
				},
				Note: &dummyNoteHandler{
					err: errors.New("dummy"),
				},
			},
			err: errors.New("dummy"),
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate(context.Background(), tc.client, "tk1v")
		require.Equal(t, tc.err, err)
	}
}
