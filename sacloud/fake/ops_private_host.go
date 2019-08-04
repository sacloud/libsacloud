package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *PrivateHostOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.PrivateHostFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.PrivateHost
	for _, res := range results {
		dest := &sacloud.PrivateHost{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.PrivateHostFindResult{
		Total:        len(results),
		Count:        len(results),
		From:         0,
		PrivateHosts: values,
	}, nil
}

// Create is fake implementation
func (o *PrivateHostOp) Create(ctx context.Context, zone string, param *sacloud.PrivateHostCreateRequest) (*sacloud.PrivateHost, error) {
	planOp := NewPrivateHostPlanOp()
	plan, err := planOp.Read(ctx, zone, param.PlanID)
	if err != nil {
		return nil, err
	}

	result := &sacloud.PrivateHost{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.PlanName = plan.Name
	result.PlanClass = plan.Class
	result.CPU = plan.CPU
	result.MemoryMB = plan.MemoryMB
	result.HostName = "sac-zone-svNNN"
	putPrivateHost(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *PrivateHostOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.PrivateHost, error) {
	value := getPrivateHostByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.PrivateHost{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *PrivateHostOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.PrivateHostUpdateRequest) (*sacloud.PrivateHost, error) {
	value, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *PrivateHostOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
