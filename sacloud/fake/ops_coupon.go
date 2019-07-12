package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *CouponOp) Find(ctx context.Context, zone string, accountID types.ID) (*sacloud.CouponFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, nil)
	var values []*sacloud.Coupon
	for _, res := range results {
		dest := &sacloud.Coupon{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.CouponFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Coupons: values,
	}, nil
}
