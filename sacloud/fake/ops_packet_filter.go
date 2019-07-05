package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *PacketFilterOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.PacketFilterFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.PacketFilter
	for _, res := range results {
		dest := &sacloud.PacketFilter{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.PacketFilterFindResult{
		Total:         len(results),
		Count:         len(results),
		From:          0,
		PacketFilters: values,
	}, nil
}

// Create is fake implementation
func (o *PacketFilterOp) Create(ctx context.Context, zone string, param *sacloud.PacketFilterCreateRequest) (*sacloud.PacketFilterCreateResult, error) {
	result := &sacloud.PacketFilter{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	s.setPacketFilter(zone, result)
	return &sacloud.PacketFilterCreateResult{
		IsOk:         true,
		PacketFilter: result,
	}, nil
}

// Read is fake implementation
func (o *PacketFilterOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.PacketFilterReadResult, error) {
	value := s.getPacketFilterByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.PacketFilter{}
	copySameNameField(value, dest)
	return &sacloud.PacketFilterReadResult{
		IsOk:         true,
		PacketFilter: dest,
	}, nil
}

// Update is fake implementation
func (o *PacketFilterOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.PacketFilterUpdateRequest) (*sacloud.PacketFilterUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.PacketFilter
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return &sacloud.PacketFilterUpdateResult{
		IsOk:         true,
		PacketFilter: value,
	}, nil
}

// Delete is fake implementation
func (o *PacketFilterOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}
