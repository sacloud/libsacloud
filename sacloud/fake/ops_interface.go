package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *InterfaceOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.InterfaceFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Interface
	for _, res := range results {
		dest := &sacloud.Interface{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.InterfaceFindResult{
		Total:      len(results),
		Count:      len(results),
		From:       0,
		Interfaces: values,
	}, nil
}

// Create is fake implementation
func (o *InterfaceOp) Create(ctx context.Context, zone string, param *sacloud.InterfaceCreateRequest) (*sacloud.Interface, error) {
	result := &sacloud.Interface{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.MACAddress = pool().nextMACAddress().String()

	putInterface(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *InterfaceOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Interface, error) {
	value := getInterfaceByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Interface{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *InterfaceOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.InterfaceUpdateRequest) (*sacloud.Interface, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putInterface(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *InterfaceOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.InterfacePatchRequest) (*sacloud.Interface, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, value); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	copySameNameField(param, value)

	if param.PatchEmptyToUserIPAddress {
		param.UserIPAddress = ""
	}
	putInterface(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *InterfaceOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// Monitor is fake implementation
func (o *InterfaceOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.InterfaceActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.InterfaceActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorInterfaceValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			Send:    float64(random(1000)),
			Receive: float64(random(1000)),
		})
	}

	return res, nil
}

// ConnectToSharedSegment is fake implementation
func (o *InterfaceOp) ConnectToSharedSegment(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	if !value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to switch[%d]", value.ID, value.SwitchID))
	}

	value.SwitchID = sharedSegmentSwitch.ID
	putInterface(zone, value)
	return nil
}

// ConnectToSwitch is fake implementation
func (o *InterfaceOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to switch[%d]", value.ID, switchID))
	}

	sw, err := NewSwitchOp().Read(ctx, zone, switchID)
	if err != nil {
		return err
	}
	sw.ServerCount++
	putSwitch(zone, sw)

	value.SwitchID = switchID
	putInterface(zone, value)
	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *InterfaceOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.SwitchID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already disconnected", value.ID))
	}

	value.SwitchID = types.ID(0)
	putInterface(zone, value)
	return nil
}

// ConnectToPacketFilter is fake implementation
func (o *InterfaceOp) ConnectToPacketFilter(ctx context.Context, zone string, id types.ID, packetFilterID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.PacketFilterID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already connected to packetfilter[%s]", value.ID, value.PacketFilterID))
	}

	value.PacketFilterID = packetFilterID
	putInterface(zone, value)
	return nil
}

// DisconnectFromPacketFilter is fake implementation
func (o *InterfaceOp) DisconnectFromPacketFilter(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.PacketFilterID.IsEmpty() {
		return newErrorConflict(o.key, id,
			fmt.Sprintf("Interface[%d] is already disconnected", value.ID))
	}

	value.PacketFilterID = types.ID(0)
	putInterface(zone, value)
	return nil
}
