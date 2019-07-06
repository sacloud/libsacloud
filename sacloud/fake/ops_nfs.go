package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *NFSOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.NFSFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.NFS
	for _, res := range results {
		dest := &sacloud.NFS{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.NFSFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		NFS:   values,
	}, nil
}

// Create is fake implementation
func (o *NFSOp) Create(ctx context.Context, zone string, param *sacloud.NFSCreateRequest) (*sacloud.NFSCreateResult, error) {
	result := &sacloud.NFS{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "nfs"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]

	s.setNFS(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.NFS, nil
	})
	return &sacloud.NFSCreateResult{
		IsOk: true,
		NFS:  result,
	}, nil
}

// Read is fake implementation
func (o *NFSOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.NFSReadResult, error) {
	value := s.getNFSByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.NFS{}
	copySameNameField(value, dest)
	return &sacloud.NFSReadResult{
		IsOk: true,
		NFS:  dest,
	}, nil
}

// Update is fake implementation
func (o *NFSOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.NFSUpdateRequest) (*sacloud.NFSUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.NFS

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.NFSUpdateResult{
		IsOk: true,
		NFS:  value,
	}, nil
}

// Delete is fake implementation
func (o *NFSOp) Delete(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.NFS
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, fmt.Sprintf("NFS[%s] is still running", id))
	}

	s.delete(o.key, zone, id)
	return nil
}

// Boot is fake implementation
func (o *NFSOp) Boot(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.NFS
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.NFS, nil
	})

	return err
}

// Shutdown is fake implementation
func (o *NFSOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.NFS
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.NFS, nil
	})

	return err
}

// Reset is fake implementation
func (o *NFSOp) Reset(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.NFS
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.NFS, nil
	})

	return nil
}

// MonitorFreeDiskSize is fake implementation
func (o *NFSOp) MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.NFSMonitorFreeDiskSizeResult, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.FreeDiskSizeActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorFreeDiskSizeValue{
			Time:         now.Add(time.Duration(i*-5) * time.Minute),
			FreeDiskSize: float64(random(1000)),
		})
	}

	return &sacloud.NFSMonitorFreeDiskSizeResult{
		IsOk:                 true,
		FreeDiskSizeActivity: res,
	}, nil
}

// MonitorInterface is fake implementation
func (o *NFSOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.NFSMonitorInterfaceResult, error) {
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

	return &sacloud.NFSMonitorInterfaceResult{
		IsOk:              true,
		InterfaceActivity: res,
	}, nil
}
