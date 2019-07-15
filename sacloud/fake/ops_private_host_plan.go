package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *PrivateHostPlanOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.PrivateHostPlanFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.PrivateHostPlan
	for _, res := range results {
		dest := &sacloud.PrivateHostPlan{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.PrivateHostPlanFindResult{
		Total:            len(results),
		Count:            len(results),
		From:             0,
		PrivateHostPlans: values,
	}, nil
}

// Read is fake implementation
func (o *PrivateHostPlanOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.PrivateHostPlan, error) {
	value := s.getPrivateHostPlanByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.PrivateHostPlan{}
	copySameNameField(value, dest)
	return dest, nil
}
