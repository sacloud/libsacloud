package server

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/require"
)

type dummyPlanFinder struct {
	plans []*sacloud.ServerPlan
	err   error
}

func (f *dummyPlanFinder) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &sacloud.ServerPlanFindResult{
		Total:       len(f.plans),
		Count:       len(f.plans),
		ServerPlans: f.plans,
	}, nil
}

func TestFindPlan(t *testing.T) {
	cases := []struct {
		msg    string
		in     *FindPlanRequest
		finder PlanFinder
		out    *sacloud.ServerPlan
		err    error
	}{
		{
			msg: "finder returns unexpected error",
			in:  nil,
			finder: &dummyPlanFinder{
				err: errors.New("dummy"),
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "finder returns multiple result",
			in: &FindPlanRequest{
				CPU: 1,
			},
			finder: &dummyPlanFinder{
				plans: []*sacloud.ServerPlan{
					{ID: 1},
					{ID: 2},
				},
			},
			out: &sacloud.ServerPlan{ID: 1},
			err: nil,
		},
		{
			msg: "with nil find parameter",
			in: &FindPlanRequest{
				CPU: 1,
			},
			finder: &dummyPlanFinder{
				plans: []*sacloud.ServerPlan{
					{ID: 1},
				},
			},
			out: &sacloud.ServerPlan{ID: 1},
			err: nil,
		},
		{
			msg: "with nil find parameter",
			in: &FindPlanRequest{
				CPU: 1,
			},
			finder: &dummyPlanFinder{
				plans: []*sacloud.ServerPlan{
					{ID: 1},
				},
			},
			out: &sacloud.ServerPlan{ID: 1},
			err: nil,
		},
		{
			msg: "plan not found",
			in:  nil,
			finder: &dummyPlanFinder{
				plans: []*sacloud.ServerPlan{},
			},
			out: nil,
			err: errors.New("server plan not found"),
		},
	}

	for _, tc := range cases {
		res, err := FindPlan(context.Background(), tc.finder, "tk1v", tc.in)
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.out, res, tc.msg)
	}
}
