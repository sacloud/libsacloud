package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *SwitchOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.SwitchFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Switch
	for _, res := range results {
		dest := &sacloud.Switch{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.SwitchFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Switches: values,
	}, nil
}

// Create is fake implementation
func (o *SwitchOp) Create(ctx context.Context, zone string, param *sacloud.SwitchCreateRequest) (*sacloud.SwitchCreateResult, error) {
	result := &sacloud.Switch{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Scope = types.Scopes.User
	s.setSwitch(zone, result)
	return &sacloud.SwitchCreateResult{
		IsOk:   true,
		Switch: result,
	}, nil
}

// Read is fake implementation
func (o *SwitchOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.SwitchReadResult, error) {
	value := s.getSwitchByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Switch{}
	copySameNameField(value, dest)
	return &sacloud.SwitchReadResult{
		IsOk:   true,
		Switch: dest,
	}, nil
}

// Update is fake implementation
func (o *SwitchOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.SwitchUpdateRequest) (*sacloud.SwitchUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Switch

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.SwitchUpdateResult{
		IsOk:   true,
		Switch: value,
	}, nil
}

// Delete is fake implementation
func (o *SwitchOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// ConnectToBridge is fake implementation
func (o *SwitchOp) ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Switch

	bridgeOp := NewBridgeOp()
	bridgeReadResult, err := bridgeOp.Read(ctx, zone, bridgeID)
	if err != nil {
		return fmt.Errorf("ConnectToBridge is failed: %s", err)
	}
	bridge := bridgeReadResult.Bridge

	if bridge.SwitchInZone != nil {
		return newErrorConflict(o.key, id, fmt.Sprintf("Bridge[%d] already connected to switch", bridgeID))
	}

	value.BridgeID = bridgeID

	switchInZone := &sacloud.BridgeSwitchInfo{}
	copySameNameField(value, switchInZone)
	bridge.SwitchInZone = switchInZone

	//bridge.BridgeInfo = append(bridge.BridgeInfo, &sacloud.BridgeInfo{
	//	ID:     value.ID,
	//	Name:   value.Name,
	//	ZoneID: zoneIDs[zone],
	//})

	s.setBridge(zone, bridge)
	s.setSwitch(zone, value)
	return nil
}

// DisconnectFromBridge is fake implementation
func (o *SwitchOp) DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Switch

	if value.BridgeID.IsEmpty() {
		return newErrorConflict(o.key, id, fmt.Sprintf("Switch[%d] already disconnected from switch", id))
	}

	bridgeOp := NewBridgeOp()
	bridgeReadResult, err := bridgeOp.Read(ctx, zone, value.BridgeID)
	if err != nil {
		return fmt.Errorf("DisconnectFromBridge is failed: %s", err)
	}
	bridge := bridgeReadResult.Bridge

	//var bridgeInfo []*sacloud.BridgeInfo
	//for _, i := range bridge.BridgeInfo {
	//	if i.ID != value.ID {
	//		bridgeInfo = append(bridgeInfo, i)
	//	}
	//}

	value.BridgeID = types.ID(0)
	bridge.SwitchInZone = nil
	// fakeドライバーではBridgeInfoに非対応
	//bridge.BridgeInfo = bridgeInfo

	s.setBridge(zone, bridge)
	s.setSwitch(zone, value)
	return nil
}
