package fake

import (
	"context"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Find is fake implementation
func (o *ZoneOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.Zone, error) {
	results, _ := find(o.key, sacloud.DefaultZone, conditions)
	var values []*sacloud.Zone
	for _, res := range results {
		dest := &sacloud.Zone{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
}

// Read is fake implementation
func (o *ZoneOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Zone, error) {
	value := s.getZoneByID(sacloud.DefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Zone{}
	copySameNameField(value, dest)
	return dest, nil
}
