package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ServerPlanOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.ServerPlan
	for _, res := range results {
		dest := &sacloud.ServerPlan{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ServerPlanFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		ServerPlans: values,
	}, nil
}

// Read is fake implementation
func (o *ServerPlanOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.ServerPlan, error) {
	value := s.getServerPlanByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.ServerPlan{}
	copySameNameField(value, dest)
	return dest, nil
}
