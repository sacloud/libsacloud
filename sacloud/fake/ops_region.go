package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *RegionOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.RegionFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Region
	for _, res := range results {
		dest := &sacloud.Region{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.RegionFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Regions: values,
	}, nil
}

// Read is fake implementation
func (o *RegionOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Region, error) {
	value := s.getRegionByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Region{}
	copySameNameField(value, dest)
	return dest, nil
}
