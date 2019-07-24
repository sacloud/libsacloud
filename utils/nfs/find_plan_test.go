package nfs

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

type dummyNoteFinder struct {
	notes []*sacloud.Note
	err   error
}

func (f *dummyNoteFinder) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.NoteFindResult, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &sacloud.NoteFindResult{
		Total: len(f.notes),
		Count: len(f.notes),
		Notes: f.notes,
	}, nil
}

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
