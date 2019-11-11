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
