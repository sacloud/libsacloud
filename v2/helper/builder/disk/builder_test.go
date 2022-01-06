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

package disk

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDiskFromUnixRequest_Validate(t *testing.T) {
	cases := []struct {
		msg string
		in  *FromUnixBuilder
		err error
	}{
		{
			msg: "invalid ostype",
			in: &FromUnixBuilder{
				OSType: ostype.ArchiveOSType(-1),
			},
			err: fmt.Errorf("invalid OSType: %s", ostype.ArchiveOSType(-1)),
		},
		{
			msg: "size not found",
			in: &FromUnixBuilder{
				OSType: ostype.CentOS,
				PlanID: types.DiskPlans.SSD,
				SizeGB: 1,
				Client: &APIClient{
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
			},
			err: fmt.Errorf("disk plan[SSDプラン:1GB] is not found"),
		},
		{
			msg: "invalid disk edit parameter",
			in: &FromUnixBuilder{
				OSType: ostype.CentOS,
				PlanID: types.DiskPlans.SSD,
				SizeGB: 1,
				EditParameter: &UnixEditRequest{
					Notes: []*sacloud.DiskEditNote{
						{ID: 1},
					},
				},
				Client: &APIClient{
					DiskPlan: &dummyDiskPlanReader{
						diskPlan: &sacloud.DiskPlan{
							ID:   types.DiskPlans.SSD,
							Name: "SSDプラン",
							Size: []*sacloud.DiskPlanSizeInfo{
								{
									Availability: types.Availabilities.Available,
									SizeMB:       1 * size.GiB,
								},
							},
						},
					},
					Note: &dummyNoteHandler{
						err: errors.New("dummy"),
					},
				},
			},
			err: errors.New("dummy"),
		},
	}

	for _, tc := range cases {
		err := tc.in.Validate(context.Background(), "tk1v")
		require.Equal(t, tc.err, err)
	}
}
