package fake

import (
	"context"
	"fmt"
	"time"

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

	result.MACAddress = pool.nextMACAddress().String()

	s.setInterface(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *InterfaceOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Interface, error) {
	value := s.getInterfaceByID(zone, id)
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
	return value, nil
}

// Delete is fake implementation
func (o *InterfaceOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
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
	s.setInterface(zone, value)
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
			fmt.Sprintf("Interface[%d] is already connected to switch[%d]", value.ID, value.SwitchID))
	}

	value.SwitchID = switchID
	s.setInterface(zone, value)
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
	s.setInterface(zone, value)
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
	s.setInterface(zone, value)
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
	s.setInterface(zone, value)
	return nil
}
