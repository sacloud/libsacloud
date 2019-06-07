package fake

import (
	"context"
	"errors"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Find is fake implementation
func (o *SIMOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.SIM, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.SIM
	for _, res := range results {
		dest := &sacloud.SIM{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
}

// Create is fake implementation
func (o *SIMOp) Create(ctx context.Context, zone string, param *sacloud.SIMCreateRequest) (*sacloud.SIM, error) {
	result := &sacloud.SIM{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	// TODO core logic is not implemented

	s.setSIM(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *SIMOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.SIM, error) {
	value := s.getSIMByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.SIM{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SIMOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.SIMUpdateRequest) (*sacloud.SIM, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	// TODO core logic is not implemented

	return value, nil
}

// Delete is fake implementation
func (o *SIMOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	// TODO core logic is not implemented

	s.delete(o.key, zone, id)
	return nil
}

// Activate is fake implementation
func (o *SIMOp) Activate(ctx context.Context, zone string, id types.ID) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// Deactivate is fake implementation
func (o *SIMOp) Deactivate(ctx context.Context, zone string, id types.ID) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// AssignIP is fake implementation
func (o *SIMOp) AssignIP(ctx context.Context, zone string, id types.ID, param *sacloud.SIMAssignIPRequest) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// ClearIP is fake implementation
func (o *SIMOp) ClearIP(ctx context.Context, zone string, id types.ID) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// IMEILock is fake implementation
func (o *SIMOp) IMEILock(ctx context.Context, zone string, id types.ID, param *sacloud.SIMIMEILockRequest) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// IMEIUnlock is fake implementation
func (o *SIMOp) IMEIUnlock(ctx context.Context, zone string, id types.ID) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// Logs is fake implementation
func (o *SIMOp) Logs(ctx context.Context, zone string, id types.ID) ([]*sacloud.SIMLog, error) {
	// TODO not implemented
	err := errors.New("not implements")
	return nil, err
}

// GetNetworkOperator is fake implementation
func (o *SIMOp) GetNetworkOperator(ctx context.Context, zone string, id types.ID) ([]*sacloud.SIMNetworkOperatorConfig, error) {
	// TODO not implemented
	err := errors.New("not implements")
	return nil, err
}

// SetNetworkOperator is fake implementation
func (o *SIMOp) SetNetworkOperator(ctx context.Context, zone string, id types.ID, configs *sacloud.SIMNetworkOperatorConfigs) error {
	// TODO not implemented
	err := errors.New("not implements")
	return err
}

// MonitorSIM is fake implementation
func (o *SIMOp) MonitorSIM(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.LinkActivity, error) {
	// TODO not implemented
	err := errors.New("not implements")
	return nil, err
}
