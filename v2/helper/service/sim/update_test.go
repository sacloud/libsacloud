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

package sim

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/helper/cleanup"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestSIMService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("sim-service")

	// setup
	simOp := sacloud.NewSIMOp(caller)
	sim, err := New(caller).CreateWithContext(ctx, &CreateRequest{
		Name:        name,
		Description: "desc",
		Tags:        types.Tags{"tag1", "tag2"},
		ICCID:       "aaaaaaaa",
		PassCode:    "bbbbbbbb",
		Activate:    true,
		IMEI:        "cccccccc",
		Carriers: []*sacloud.SIMNetworkOperatorConfig{
			{
				Allow: true,
				Name:  types.SIMOperators.Docomo.String(),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		cleanup.DeleteSIM(ctx, simOp, sim.ID) // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *ApplyRequest
	}{
		{
			in: &UpdateRequest{
				ID:       sim.ID,
				Name:     pointer.NewString(name + "-upd"),
				Activate: pointer.NewBool(false),
				IMEI:     pointer.NewString(""),
				Carriers: &[]*sacloud.SIMNetworkOperatorConfig{
					{Allow: true, Name: types.SIMOperators.SoftBank.String()},
				},
			},
			expect: &ApplyRequest{
				ID:          sim.ID,
				Name:        name + "-upd",
				Description: sim.Description,
				Tags:        sim.Tags,
				IconID:      sim.IconID,
				ICCID:       sim.ICCID,
				PassCode:    "",
				Activate:    false,
				IMEI:        "",
				Carriers: []*sacloud.SIMNetworkOperatorConfig{
					{
						Allow: true,
						Name:  types.SIMOperators.SoftBank.String(),
					},
				},
			},
		},
	}

	for _, tc := range cases {
		req, err := tc.in.ApplyRequest(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, req)
	}
}
