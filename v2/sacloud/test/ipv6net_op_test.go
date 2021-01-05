// Copyright 2016-2021 The Libsacloud Authors
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

package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIPv6NetOp_List_Read(t *testing.T) {
	t.Parallel()

	internetOp := sacloud.NewInternetOp(singletonAPICaller())
	ipv6NetOp := sacloud.NewIPv6NetOp(singletonAPICaller())
	ctx := context.Background()

	internet, err := internetOp.Create(ctx, testZone, &sacloud.InternetCreateRequest{
		Name:           testutil.ResourceName("internet-ipv6"),
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
	})
	assert.NoError(t, err)

	// wait
	waiter := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
		return internetOp.Read(ctx, testZone, internet.ID)
	}, 100)
	if _, err := waiter.WaitForState(context.TODO()); err != nil {
		t.Error("WaitForUp is failed: ", err)
		return
	}

	internet, err = internetOp.Read(ctx, testZone, internet.ID)
	assert.NoError(t, err)

	// Enable IPv6
	ipv6Net, err := internetOp.EnableIPv6(ctx, testZone, internet.ID)
	assert.NoError(t, err)

	err = testutil.DoAsserts(
		testutil.AssertNotEmptyFunc(t, ipv6Net.ID, "IPv6Net.ID"),
		testutil.AssertNotEmptyFunc(t, ipv6Net.IPv6Prefix, "IPv6Net.IPv6Prefix"),
		testutil.AssertNotEmptyFunc(t, ipv6Net.IPv6PrefixLen, "IPv6Net.IPv6PrefixLen"),
	)
	assert.NoError(t, err)

	// find
	searched, err := ipv6NetOp.List(ctx, testZone)
	assert.NoError(t, err)
	assert.True(t, len(searched.IPv6Nets) > 0)

	// read
	read, err := ipv6NetOp.Read(ctx, testZone, ipv6Net.ID)
	assert.NoError(t, err)
	err = testutil.DoAsserts(
		testutil.AssertEqualFunc(t, ipv6Net.ID, read.ID, "IPv6Net.ID"),
		testutil.AssertEqualFunc(t, ipv6Net.IPv6Prefix, read.IPv6Prefix, "IPv6Net.IPv6Prefix"),
		testutil.AssertEqualFunc(t, ipv6Net.IPv6PrefixLen, read.IPv6PrefixLen, "IPv6Net.IPv6PrefixLen"),
	)
	assert.NoError(t, err)

	// Disable IPv6
	err = internetOp.DisableIPv6(ctx, testZone, internet.ID, ipv6Net.ID)
	assert.NoError(t, err)

	// cleanup
	err = internetOp.Delete(ctx, testZone, internet.ID)
	assert.NoError(t, err)
}
