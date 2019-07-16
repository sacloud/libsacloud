package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

// Find is fake implementation
func (o *ServiceClassOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServiceClassFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.ServiceClass
	for _, res := range results {
		dest := &sacloud.ServiceClass{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ServiceClassFindResult{
		Total:          len(results),
		Count:          len(results),
		From:           0,
		ServiceClasses: values,
	}, nil
}
