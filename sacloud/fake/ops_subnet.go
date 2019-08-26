package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *SubnetOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.SubnetFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Subnet
	for _, res := range results {
		dest := &sacloud.Subnet{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.SubnetFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Subnets: values,
	}, nil
}

// Read is fake implementation
func (o *SubnetOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Subnet, error) {
	value := getSubnetByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Subnet{}
	copySameNameField(value, dest)
	return dest, nil
}
