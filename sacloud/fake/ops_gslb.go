package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *GSLBOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.GSLBFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.GSLB
	for _, res := range results {
		dest := &sacloud.GSLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.GSLBFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		GSLBs: values,
	}, nil
}

// Create is fake implementation
func (o *GSLBOp) Create(ctx context.Context, param *sacloud.GSLBCreateRequest) (*sacloud.GSLB, error) {
	result := &sacloud.GSLB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability)

	result.FQDN = fmt.Sprintf("site-%d.gslb7.example.ne.jp", result.ID)
	result.SettingsHash = "settingshash"

	s.setGSLB(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *GSLBOp) Read(ctx context.Context, id types.ID) (*sacloud.GSLB, error) {
	value := s.getGSLBByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.GSLB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *GSLBOp) Update(ctx context.Context, id types.ID, param *sacloud.GSLBUpdateRequest) (*sacloud.GSLB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	if param.DelayLoop == 0 {
		param.DelayLoop = 10 // default value
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *GSLBOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	s.delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
