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

package query

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestFindNFSPlanID(t *testing.T) {
	cases := []struct {
		msg             string
		finder          NoteFinder
		inputDiskPlanID types.ID
		inputDiskSize   types.ENFSSize
		expectedValue   types.ID
		expectedErr     error
	}{
		{
			msg: "finder returns error",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{},
				err:   errors.New("dummy"),
			},
			expectedErr: errors.New("dummy"),
		},
		{
			msg: "finder returns zero",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{},
			},
			expectedErr: errors.New("note[sys-nfs] not found"),
		},
		{
			msg: "not found",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{
					{
						Name:    "sys-nfs",
						Class:   "json",
						Content: `{"plans":{"HDD":[{"size": 100,"availability":"available","planId":1}]}}`,
					},
				},
				err: nil,
			},
			inputDiskPlanID: types.NFSPlans.SSD,
			inputDiskSize:   types.NFSHDDSizes.Size100GB,
			expectedValue:   0,
			expectedErr:     nil,
		},
		{
			msg: "normal",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{
					{
						Name:    "sys-nfs",
						Class:   "json",
						Content: `{"plans":{"HDD":[{"size": 100,"availability":"available","planId":1}]}}`,
					},
				},
				err: nil,
			},
			inputDiskPlanID: types.NFSPlans.HDD,
			inputDiskSize:   types.NFSHDDSizes.Size100GB,
			expectedValue:   1,
			expectedErr:     nil,
		},
	}

	for _, tc := range cases {
		actual, err := FindNFSPlanID(context.Background(), tc.finder, tc.inputDiskPlanID, tc.inputDiskSize)
		if tc.expectedErr != nil {
			require.Equal(t, tc.expectedErr, err, tc.msg)
		} else {
			require.NoError(t, err, tc.msg)
		}
		require.Equal(t, tc.expectedValue, actual, tc.msg)
	}
}

func TestGetPlanInfo(t *testing.T) {
	cases := []struct {
		msg           string
		finder        NoteFinder
		input         types.ID
		expectedValue *NFSPlanInfo
		expectedErr   error
	}{
		{
			msg: "finder returns error",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{},
				err:   errors.New("dummy"),
			},
			expectedErr: errors.New("dummy"),
		},
		{
			msg: "finder returns zero",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{},
			},
			expectedErr: errors.New("note[sys-nfs] not found"),
		},
		{
			msg: "not found",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{
					{
						Name:    "sys-nfs",
						Class:   "json",
						Content: `{"plans":{"HDD":[{"size": 100,"availability":"available","planId":1}]}}`,
					},
				},
				err: nil,
			},
			input:         2,
			expectedValue: nil,
			expectedErr:   fmt.Errorf("nfs plan [id:%d] not found", 2),
		},
		{
			msg: "normal",
			finder: &dummyNoteFinder{
				notes: []*sacloud.Note{
					{
						Name:    "sys-nfs",
						Class:   "json",
						Content: `{"plans":{"HDD":[{"size": 100,"availability":"available","planId":1}]}}`,
					},
				},
				err: nil,
			},
			input: 1,
			expectedValue: &NFSPlanInfo{
				NFSPlanID:  1,
				DiskPlanID: types.NFSPlans.HDD,
				Size:       types.NFSHDDSizes.Size100GB,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range cases {
		actual, err := GetNFSPlanInfo(context.Background(), tc.finder, tc.input)
		if tc.expectedErr != nil {
			require.Equal(t, tc.expectedErr, err, tc.msg)
		} else {
			require.NoError(t, err, tc.msg)
		}
		require.Equal(t, tc.expectedValue, actual, tc.msg)
	}
}
