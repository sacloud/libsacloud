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

package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCouponOp_Find(t *testing.T) {
	t.Parallel()

	// get account ID
	authStatusOp := sacloud.NewAuthStatusOp(singletonAPICaller())
	authStatus, err := authStatusOp.Read(context.Background())
	if !assert.NoError(t, err) {
		return
	}

	// check current permission
	if !authStatus.ExternalPermission.PermittedBill() {
		t.Skip("current account is not permitted to viewing bills")
	}

	client := sacloud.NewCouponOp(singletonAPICaller())
	searched, err := client.Find(context.Background(), authStatus.AccountID)
	assert.NoError(t, err)

	if searched.Count > 0 {
		err = testutil.DoAsserts(
			testutil.AssertNotEmptyFunc(t, searched.Coupons[0].ID, "Coupon.ID"),
			testutil.AssertNotEmptyFunc(t, searched.Coupons[0].MemberID, "Coupon.MemberID"),
			testutil.AssertNotEmptyFunc(t, searched.Coupons[0].ContractID, "Coupon.ContractID"),
			testutil.AssertNotEmptyFunc(t, searched.Coupons[0].ServiceClassID, "Coupon.ServiceClassID"),
		)
		assert.NoError(t, err)
	}
}
