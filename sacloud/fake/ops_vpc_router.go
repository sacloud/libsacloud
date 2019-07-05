package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *VPCRouterOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.VPCRouterFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.VPCRouter
	for _, res := range results {
		dest := &sacloud.VPCRouter{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.VPCRouterFindResult{
		Total:      len(results),
		Count:      len(results),
		From:       0,
		Appliances: values,
	}, nil
}

// Create is fake implementation
func (o *VPCRouterOp) Create(ctx context.Context, zone string, param *sacloud.VPCRouterCreateRequest) (*sacloud.VPCRouterCreateResult, error) {
	result := &sacloud.VPCRouter{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "vpcrouter"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""

	ifOp := NewInterfaceOp()
	swOp := NewSwitchOp()

	ifCreateParam := &sacloud.InterfaceCreateRequest{}
	if param.Switch.Scope == types.Scopes.Shared {
		ifCreateParam.ServerID = result.ID
	} else {
		_, err := swOp.Read(ctx, zone, param.Switch.ID)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	}

	ifCreateResult, err := ifOp.Create(ctx, zone, ifCreateParam)
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}
	iface := ifCreateResult.Interface

	if param.Switch.Scope == types.Scopes.Shared {
		if err := ifOp.ConnectToSharedSegment(ctx, zone, iface.ID); err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	} else {
		if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, param.Switch.ID); err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
	}

	ifReadResult, err := ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return nil, newErrorConflict(o.key, types.ID(0), err.Error())
	}
	iface = ifReadResult.Interface

	vpcRouterInterface := &sacloud.VPCRouterInterface{}
	copySameNameField(iface, vpcRouterInterface)
	result.Interfaces = append(result.Interfaces, vpcRouterInterface)

	s.setVPCRouter(zone, result)

	id := result.ID
	startMigration(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Appliance, nil
	})
	return &sacloud.VPCRouterCreateResult{
		IsOk:      true,
		Appliance: result,
	}, nil
}

// Read is fake implementation
func (o *VPCRouterOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.VPCRouterReadResult, error) {
	value := s.getVPCRouterByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.VPCRouter{}
	copySameNameField(value, dest)
	return &sacloud.VPCRouterReadResult{
		IsOk:      true,
		Appliance: dest,
	}, nil
}

// Update is fake implementation
func (o *VPCRouterOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.VPCRouterUpdateRequest) (*sacloud.VPCRouterUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Appliance
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.VPCRouterUpdateResult{
		IsOk:      true,
		Appliance: value,
	}, nil
}

// Delete is fake implementation
func (o *VPCRouterOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *VPCRouterOp) Config(ctx context.Context, zone string, id types.ID) error {
	return nil
}

// Boot is fake implementation
func (o *VPCRouterOp) Boot(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Appliance
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Appliance, nil
	})

	return err
}

// Shutdown is fake implementation
func (o *VPCRouterOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Appliance
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Appliance, nil
	})

	return err
}

// Reset is fake implementation
func (o *VPCRouterOp) Reset(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Appliance
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Appliance, nil
	})

	return nil
}

// ConnectToSwitch is fake implementation
func (o *VPCRouterOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, nicIndex int, switchID types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Appliance

	for _, nic := range value.Interfaces {
		if nic.Index == nicIndex {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] already connected to switch", nicIndex))
		}
	}

	// find switch
	swOp := NewSwitchOp()
	_, err = swOp.Read(ctx, zone, switchID)
	if err != nil {
		return fmt.Errorf("ConnectToSwitch is failed: %s", err)
	}

	// create interface
	ifOp := NewInterfaceOp()
	ifCreateResult, err := ifOp.Create(ctx, zone, &sacloud.InterfaceCreateRequest{ServerID: id})
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}
	iface := ifCreateResult.Interface

	if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, switchID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	ifReadResult, err := ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}
	iface = ifReadResult.Interface

	vpcRouterInterface := &sacloud.VPCRouterInterface{}
	copySameNameField(iface, vpcRouterInterface)
	value.Interfaces = append(value.Interfaces, vpcRouterInterface)

	s.setVPCRouter(zone, value)
	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *VPCRouterOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID, nicIndex int) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Appliance

	var exists bool
	var nicID types.ID
	var interfaces []*sacloud.VPCRouterInterface

	for _, nic := range value.Interfaces {
		if nic.Index == nicIndex {
			exists = true
			nicID = nic.ID
		} else {
			interfaces = append(interfaces, nic)
		}
	}
	if !exists {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] is not exists", nicIndex))
	}

	ifOp := NewInterfaceOp()
	if err := ifOp.DisconnectFromSwitch(ctx, zone, nicID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	value.Interfaces = interfaces
	s.setVPCRouter(zone, value)
	return nil
}

// MonitorInterface is fake implementation
func (o *VPCRouterOp) MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *sacloud.MonitorCondition) (*sacloud.VPCRouterMonitorInterfaceResult, error) {
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

	return &sacloud.VPCRouterMonitorInterfaceResult{
		IsOk: true,
		Data: res,
	}, nil
}
