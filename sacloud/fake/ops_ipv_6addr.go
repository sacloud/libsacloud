package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *IPv6AddrOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.IPv6AddrFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.IPv6Addr
	for _, res := range results {
		dest := &sacloud.IPv6Addr{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.IPv6AddrFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		IPv6Addrs: values,
	}, nil
}

// Create is fake implementation
func (o *IPv6AddrOp) Create(ctx context.Context, zone string, param *sacloud.IPv6AddrCreateRequest) (*sacloud.IPv6Addr, error) {
	result := &sacloud.IPv6Addr{}
	copySameNameField(param, result)

	s.setWithID(ResourceIPv6Addr, zone, result, pool.generateID())
	return result, nil
}

// Read is fake implementation
func (o *IPv6AddrOp) Read(ctx context.Context, zone string, ipv6addr string) (*sacloud.IPv6Addr, error) {
	var value *sacloud.IPv6Addr

	results, _ := find(o.key, zone, nil)
	for _, res := range results {
		v := res.(*sacloud.IPv6Addr)
		if v.IPv6Addr == ipv6addr {
			value = v
			break
		}
	}

	if value == nil {
		return nil, newErrorNotFound(o.key, ipv6addr)
	}
	return value, nil
}

// Update is fake implementation
func (o *IPv6AddrOp) Update(ctx context.Context, zone string, ipv6addr string, param *sacloud.IPv6AddrUpdateRequest) (*sacloud.IPv6Addr, error) {
	found := false
	results := s.values(o.key, zone)
	var value *sacloud.IPv6Addr
	for key, res := range results {
		v := res.(*sacloud.IPv6Addr)
		if v.IPv6Addr == ipv6addr {
			copySameNameField(param, v)
			found = true
			s.setWithID(o.key, zone, v, types.StringID(key))
			value = v
		}
	}

	if !found {
		return nil, newErrorNotFound(o.key, ipv6addr)
	}

	return value, nil
}

// Delete is fake implementation
func (o *IPv6AddrOp) Delete(ctx context.Context, zone string, ipv6addr string) error {
	found := false
	results := s.values(o.key, zone)
	for key, res := range results {
		v := res.(*sacloud.IPv6Addr)
		if v.IPv6Addr == ipv6addr {
			found = true
			s.delete(o.key, zone, types.StringID(key))
		}
	}

	if !found {
		return newErrorNotFound(o.key, ipv6addr)
	}

	return nil
}
