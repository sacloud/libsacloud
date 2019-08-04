package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *IconOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.IconFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Icon
	for _, res := range results {
		dest := &sacloud.Icon{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.IconFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Icons: values,
	}, nil
}

// Create is fake implementation
func (o *IconOp) Create(ctx context.Context, param *sacloud.IconCreateRequest) (*sacloud.Icon, error) {
	result := &sacloud.Icon{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillModifiedAt)

	result.Availability = types.Availabilities.Available
	result.Scope = types.Scopes.User
	result.URL = fmt.Sprintf("https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/icon/%d.png", result.ID)

	putIcon(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *IconOp) Read(ctx context.Context, id types.ID) (*sacloud.Icon, error) {
	value := getIconByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Icon{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *IconOp) Update(ctx context.Context, id types.ID, param *sacloud.IconUpdateRequest) (*sacloud.Icon, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *IconOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
