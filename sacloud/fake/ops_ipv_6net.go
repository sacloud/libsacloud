package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// List is fake implementation
func (o *IPv6NetOp) List(ctx context.Context, zone string) (*sacloud.IPv6NetListResult, error) {
	results, _ := find(o.key, zone, nil)
	var values []*sacloud.IPv6Net
	for _, res := range results {
		dest := &sacloud.IPv6Net{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.IPv6NetListResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		IPv6Nets: values,
	}, nil
}

// Read is fake implementation
func (o *IPv6NetOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.IPv6Net, error) {
	value := s.getIPv6NetByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.IPv6Net{}
	copySameNameField(value, dest)
	return dest, nil
}
