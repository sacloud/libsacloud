package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *BridgeOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.Bridge, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Bridge
	for _, res := range results {
		dest := &sacloud.Bridge{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
}

// Create is fake implementation
func (o *BridgeOp) Create(ctx context.Context, zone string, param *sacloud.BridgeCreateRequest) (*sacloud.Bridge, error) {
	result := &sacloud.Bridge{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	s.setBridge(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *BridgeOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Bridge, error) {
	value := s.getBridgeByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Bridge{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *BridgeOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.BridgeUpdateRequest) (*sacloud.Bridge, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)

	return value, nil
}

// Delete is fake implementation
func (o *BridgeOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}
