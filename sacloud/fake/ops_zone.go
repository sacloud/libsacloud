package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ZoneOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ZoneFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Zone
	for _, res := range results {
		dest := &sacloud.Zone{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ZoneFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Zones: values,
	}, nil
}

// Read is fake implementation
func (o *ZoneOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.ZoneReadResult, error) {
	value := s.getZoneByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Zone{}
	copySameNameField(value, dest)
	return &sacloud.ZoneReadResult{
		IsOk: true,
		Zone: dest,
	}, nil
}
