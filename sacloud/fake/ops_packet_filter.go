package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
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
func (o *PacketFilterOp) Create(ctx context.Context, zone string, param *sacloud.PacketFilterCreateRequest) (*sacloud.PacketFilter, error) {
	result := &sacloud.PacketFilter{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	putPacketFilter(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *PacketFilterOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.PacketFilter, error) {
	value := getPacketFilterByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.PacketFilter{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *PacketFilterOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.PacketFilterUpdateRequest) (*sacloud.PacketFilter, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putPacketFilter(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *PacketFilterOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.PacketFilterPatchRequest) (*sacloud.PacketFilter, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, value); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	copySameNameField(param, value)

	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToExpression {
		value.Expression = nil
	}

	putPacketFilter(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *PacketFilterOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}
