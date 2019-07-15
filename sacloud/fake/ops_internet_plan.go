package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *InternetPlanOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.InternetPlanFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.InternetPlan
	for _, res := range results {
		dest := &sacloud.InternetPlan{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.InternetPlanFindResult{
		Total:         len(results),
		Count:         len(results),
		From:          0,
		InternetPlans: values,
	}, nil
}

// Read is fake implementation
func (o *InternetPlanOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.InternetPlan, error) {
	value := s.getInternetPlanByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.InternetPlan{}
	copySameNameField(value, dest)
	return dest, nil
}
