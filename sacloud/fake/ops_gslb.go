package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/types"
)

// Find is fake implementation
func (o *GSLBOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.GSLB, error) {
	results, _ := find(o.key, sacloud.DefaultZone, conditions)
	var values []*sacloud.GSLB
	for _, res := range results {
		dest := &sacloud.GSLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
}

// Create is fake implementation
func (o *GSLBOp) Create(ctx context.Context, zone string, param *sacloud.GSLBCreateRequest) (*sacloud.GSLB, error) {
	result := &sacloud.GSLB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability)

	result.FQDN = fmt.Sprintf("site-%d.gslb7.example.ne.jp", result.ID)
	result.SettingsHash = "settingshash"
	// TODO mapconvで設定しているデフォルト値をどう扱うか?
	for _, server := range result.DestinationServers {
		if server.Weight.Int() == 0 {
			server.Weight = types.StringNumber(1)
		}
	}

	s.setGSLB(sacloud.DefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *GSLBOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.GSLB, error) {
	value := s.getGSLBByID(sacloud.DefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.GSLB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *GSLBOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.GSLBUpdateRequest) (*sacloud.GSLB, error) {
	value, err := o.Read(ctx, sacloud.DefaultZone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *GSLBOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.DefaultZone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, sacloud.DefaultZone, id)
	return nil
}
