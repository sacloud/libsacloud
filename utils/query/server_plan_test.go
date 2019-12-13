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
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/stretchr/testify/require"
)

func TestFindServerPlan(t *testing.T) {
	cases := []struct {
		msg    string
		in     *FindServerPlanRequest
		finder ServerPlanFinder
		out    *sacloud.ServerPlan
		err    error
	}{
		{
			msg: "finder returns unexpected error",
			in:  nil,
			finder: &dummyServerPlanFinder{
				err: errors.New("dummy"),
			},
			out: nil,
			err: errors.New("dummy"),
		},
		{
			msg: "finder returns multiple result",
			in: &FindServerPlanRequest{
				CPU: 1,
			},
			finder: &dummyServerPlanFinder{
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
			in: &FindServerPlanRequest{
				CPU: 1,
			},
			finder: &dummyServerPlanFinder{
				plans: []*sacloud.ServerPlan{
					{ID: 1},
				},
			},
			out: &sacloud.ServerPlan{ID: 1},
			err: nil,
		},
		{
			msg: "with nil find parameter",
			in: &FindServerPlanRequest{
				CPU: 1,
			},
			finder: &dummyServerPlanFinder{
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
			finder: &dummyServerPlanFinder{
				plans: []*sacloud.ServerPlan{},
			},
			out: nil,
			err: errors.New("server plan not found"),
		},
	}

	for _, tc := range cases {
		res, err := FindServerPlan(context.Background(), tc.finder, "tk1v", tc.in)
		require.Equal(t, tc.err, err, tc.msg)
		require.Equal(t, tc.out, res, tc.msg)
	}
}
