package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ServerOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Server
	for _, res := range results {
		dest := &sacloud.Server{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ServerFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Servers: values,
	}, nil
}

// Create is fake implementation
func (o *ServerOp) Create(ctx context.Context, zone string, param *sacloud.ServerCreateRequest) (*sacloud.ServerCreateResult, error) {
	result := &sacloud.Server{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Migrating
	if param.ServerPlanGeneration == types.PlanGenerations.Default {
		switch zone {
		case "is1a":
			result.ServerPlanGeneration = types.PlanGenerations.G200
		default:
			result.ServerPlanGeneration = types.PlanGenerations.G100
		}
	}
	// TODO プランAPIを実装したら修正する
	result.ServerPlanID = types.StringID(fmt.Sprintf("%03d%03d%03d", result.ServerPlanGeneration, result.GetMemoryGB(), result.CPU))
	result.ServerPlanName = fmt.Sprintf("世代:%03d メモリ:%03d CPU:%03d", result.ServerPlanGeneration, result.GetMemoryGB(), result.CPU)

	for _, cs := range param.ConnectedSwitches {
		ifOp := NewInterfaceOp()
		swOp := NewSwitchOp()

		ifCreateParam := &sacloud.InterfaceCreateRequest{}
		if cs.Scope == types.Scopes.Shared {
			ifCreateParam.ServerID = result.ID
		} else {
			_, err := swOp.Read(ctx, zone, cs.ID)
			if err != nil {
				return nil, newErrorConflict(o.key, types.ID(0), err.Error())
			}
		}

		ifCreateResult, err := ifOp.Create(ctx, zone, ifCreateParam)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
		iface := ifCreateResult.Interface

		if cs.Scope == types.Scopes.Shared {
			if err := ifOp.ConnectToSharedSegment(ctx, zone, iface.ID); err != nil {
				return nil, newErrorConflict(o.key, types.ID(0), err.Error())
			}
		} else {
			if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, cs.ID); err != nil {
				return nil, newErrorConflict(o.key, types.ID(0), err.Error())
			}
		}

		ifReadResult, err := ifOp.Read(ctx, zone, iface.ID)
		if err != nil {
			return nil, newErrorConflict(o.key, types.ID(0), err.Error())
		}
		iface = ifReadResult.Interface
		result.Interfaces = append(result.Interfaces, iface)
	}

	s.setServer(zone, result)
	return &sacloud.ServerCreateResult{
		IsOk:   true,
		Server: result,
	}, nil
}

// Read is fake implementation
func (o *ServerOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.ServerReadResult, error) {
	value := s.getServerByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Server{}
	copySameNameField(value, dest)
	return &sacloud.ServerReadResult{
		IsOk:   true,
		Server: dest,
	}, nil
}

// Update is fake implementation
func (o *ServerOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ServerUpdateRequest) (*sacloud.ServerUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Server

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.ServerUpdateResult{
		IsOk:   true,
		Server: value,
	}, nil
}

// Delete is fake implementation
func (o *ServerOp) Delete(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server

	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, fmt.Sprintf("Server[%s] is still running", id))
	}

	ifOp := NewInterfaceOp()
	for _, iface := range value.Interfaces {
		if err := ifOp.Delete(ctx, zone, iface.ID); err != nil {
			return err
		}
	}

	diskOp := NewDiskOp()
	for _, disk := range value.Disks {
		if err := diskOp.DisconnectFromServer(ctx, zone, disk.ID); err != nil {
			return err
		}
	}

	s.delete(o.key, zone, id)
	return nil
}

// ChangePlan is fake implementation
func (o *ServerOp) ChangePlan(ctx context.Context, zone string, id types.ID, plan *sacloud.ServerChangePlanRequest) (*sacloud.ServerChangePlanResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Server

	if value.InstanceStatus.IsUp() {
		return nil, newErrorConflict(o.key, id, fmt.Sprintf("Server[%d] is running", value.ID))
	}

	copySameNameField(plan, value)
	value.ServerPlanID = types.StringID(fmt.Sprintf("%03d%03d%03d", value.ServerPlanGeneration, value.GetMemoryGB(), value.CPU))
	value.ServerPlanName = fmt.Sprintf("世代:%03d メモリ:%03d CPU:%03d", value.ServerPlanGeneration, value.GetMemoryGB(), value.CPU)

	// ID変更
	s.delete(o.key, zone, value.ID)
	newServer := &sacloud.Server{}
	copySameNameField(value, newServer)
	newServer.ID = pool.generateID()
	s.setServer(zone, newServer)

	return &sacloud.ServerChangePlanResult{
		IsOk:   true,
		Server: newServer,
	}, nil
}

// InsertCDROM is fake implementation
func (o *ServerOp) InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *sacloud.InsertCDROMRequest) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server

	cdromOp := NewCDROMOp()
	if _, err = cdromOp.Read(ctx, zone, insertParam.ID); err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("CDROM[%d] is not exists", insertParam.ID))
	}

	value.CDROMID = insertParam.ID
	s.setServer(zone, value)
	return nil
}

// EjectCDROM is fake implementation
func (o *ServerOp) EjectCDROM(ctx context.Context, zone string, id types.ID, insertParam *sacloud.EjectCDROMRequest) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server

	cdromOp := NewCDROMOp()
	if _, err = cdromOp.Read(ctx, zone, insertParam.ID); err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("CDROM[%d] is not exists", insertParam.ID))
	}

	value.CDROMID = types.ID(0)
	s.setServer(zone, value)
	return nil
}

// Boot is fake implementation
func (o *ServerOp) Boot(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Server, nil
	})

	return err
}

// Shutdown is fake implementation
func (o *ServerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Server, nil
	})

	return err
}

// Reset is fake implementation
func (o *ServerOp) Reset(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Server
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Server, nil
	})

	return nil
}

// Monitor is fake implementation
func (o *ServerOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.ServerMonitorResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Server

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.CPUTimeActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorCPUTimeValue{
			Time:    now.Add(time.Duration(i*-5) * time.Minute),
			CPUTime: float64(random(value.CPU * 1000)),
		})
	}

	return &sacloud.ServerMonitorResult{
		IsOk:            true,
		CPUTimeActivity: res,
	}, nil
}
