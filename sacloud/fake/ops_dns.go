package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *DNSOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.DNSFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.DNS
	for _, res := range results {
		dest := &sacloud.DNS{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.DNSFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		DNS:   values,
	}, nil
}

// Create is fake implementation
func (o *DNSOp) Create(ctx context.Context, zone string, param *sacloud.DNSCreateRequest) (*sacloud.DNS, error) {
	result := &sacloud.DNS{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "dns"
	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	result.DNSZone = param.Name

	s.setDNS(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *DNSOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.DNS, error) {
	value := s.getDNSByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.DNS{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *DNSOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.DNSUpdateRequest) (*sacloud.DNS, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *DNSOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}
