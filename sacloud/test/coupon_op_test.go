package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
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
		err = DoAsserts(
			AssertNotEmptyFunc(t, searched.Coupons[0].ID, "Coupon.ID"),
			AssertNotEmptyFunc(t, searched.Coupons[0].MemberID, "Coupon.MemberID"),
			AssertNotEmptyFunc(t, searched.Coupons[0].ContractID, "Coupon.ContractID"),
			AssertNotEmptyFunc(t, searched.Coupons[0].ServiceClassID, "Coupon.ServiceClassID"),
		)
		assert.NoError(t, err)
	}
}
