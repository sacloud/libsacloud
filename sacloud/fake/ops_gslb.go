package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *GSLBOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.GSLBFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.GSLB
	for _, res := range results {
		dest := &sacloud.GSLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.GSLBFindResult{
		Total:              len(results),
		Count:              len(results),
		From:               0,
		CommonServiceItems: values,
	}, nil
}

// Create is fake implementation
func (o *GSLBOp) Create(ctx context.Context, zone string, param *sacloud.GSLBCreateRequest) (*sacloud.GSLBCreateResult, error) {
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

	s.setGSLB(sacloud.APIDefaultZone, result)
	return &sacloud.GSLBCreateResult{
		IsOk:              true,
		CommonServiceItem: result,
	}, nil
}

// Read is fake implementation
func (o *GSLBOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.GSLBReadResult, error) {
	value := s.getGSLBByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.GSLB{}
	copySameNameField(value, dest)
	return &sacloud.GSLBReadResult{
		IsOk:              true,
		CommonServiceItem: dest,
	}, nil
}

// Update is fake implementation
func (o *GSLBOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.GSLBUpdateRequest) (*sacloud.GSLBUpdateResult, error) {
	readResult, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.CommonServiceItem
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.GSLBUpdateResult{
		IsOk:              true,
		CommonServiceItem: value,
	}, nil
}

// Delete is fake implementation
func (o *GSLBOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
