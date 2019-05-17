package fake

import (
	"context"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Find is fake implementation
func (o *SwitchOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.Switch, error) {
	results, _ := find(ResourceSwitch, zone, conditions)
	var values []*sacloud.Switch
	for _, res := range results {
		values = append(values, res.(*sacloud.Switch))
	}
	return values, nil
}

// Create is fake implementation
func (o *SwitchOp) Create(ctx context.Context, zone string, param *sacloud.SwitchCreateRequest) (*sacloud.Switch, error) {
	result := &sacloud.Switch{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Scope = types.Scopes.User
	s.setSwitch(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *SwitchOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Switch, error) {
	value := s.getSwitchByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(ResourceSwitch, id)
	}
	return value, nil
}

// Update is fake implementation
func (o *SwitchOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.SwitchUpdateRequest) (*sacloud.Switch, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *SwitchOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(ResourceSwitch, zone, id)
	return nil
}

// ConnectToBridge is fake implementation
func (o *SwitchOp) ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	// TODO ブリッジの実装をしてから実装する

	return nil
}

// DisconnectFromBridge is fake implementation
func (o *SwitchOp) DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	// TODO ブリッジの実装をしてから実装する

	return nil
}
