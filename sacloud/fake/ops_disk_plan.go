package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *DiskPlanOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.DiskPlanFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.DiskPlan
	for _, res := range results {
		dest := &sacloud.DiskPlan{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.DiskPlanFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		DiskPlans: values,
	}, nil
}

// Read is fake implementation
func (o *DiskPlanOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.DiskPlan, error) {
	value := getDiskPlanByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.DiskPlan{}
	copySameNameField(value, dest)
	return dest, nil
}
