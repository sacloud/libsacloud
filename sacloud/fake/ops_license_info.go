package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *LicenseInfoOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.LicenseInfoFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.LicenseInfo
	for _, res := range results {
		dest := &sacloud.LicenseInfo{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.LicenseInfoFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		LicenseInfo: values,
	}, nil
}

// Read is fake implementation
func (o *LicenseInfoOp) Read(ctx context.Context, id types.ID) (*sacloud.LicenseInfo, error) {
	value := getLicenseInfoByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.LicenseInfo{}
	copySameNameField(value, dest)
	return dest, nil
}
