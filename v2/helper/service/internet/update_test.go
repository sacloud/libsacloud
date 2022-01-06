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

package internet

import (
	"context"
	"testing"

	internetBuilder "github.com/sacloud/libsacloud/v2/helper/builder/internet"
	"github.com/sacloud/libsacloud/v2/helper/cleanup"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestInternetService_convertUpdateRequest(t *testing.T) {
	ctx := context.Background()
	zone := testutil.TestZone()
	caller := testutil.SingletonAPICaller()
	name := testutil.ResourceName("internet-service-update")

	// setup
	builder := &internetBuilder.Builder{
		Name:           name,
		Description:    "description",
		Tags:           types.Tags{"tag1", "tag2"},
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
		EnableIPv6:     true,
		NotFoundRetry:  500,
		NoWait:         false,
		Client:         internetBuilder.NewAPIClient(caller),
	}
	internet, err := builder.Build(ctx, zone)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		cleanup.DeleteInternet(ctx, sacloud.NewInternetOp(caller), zone, internet.ID) // nolint
	}()

	// test
	cases := []struct {
		in     *UpdateRequest
		expect *internetBuilder.Builder
	}{
		{
			in: &UpdateRequest{
				Zone:          zone,
				ID:            internet.ID,
				Name:          pointer.NewString(name + "-upd"),
				Description:   pointer.NewString("description-upd"),
				BandWidthMbps: pointer.NewInt(250),
				EnableIPv6:    pointer.NewBool(false),
			},
			expect: &internetBuilder.Builder{
				Name:           name + "-upd",
				Description:    "description-upd",
				Tags:           internet.Tags,
				IconID:         internet.IconID,
				NetworkMaskLen: internet.NetworkMaskLen,
				BandWidthMbps:  250,
				EnableIPv6:     false,
				Client:         internetBuilder.NewAPIClient(caller),
			},
		},
	}

	for _, tc := range cases {
		builder, err := tc.in.Builder(ctx, caller)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, builder)
	}
}
