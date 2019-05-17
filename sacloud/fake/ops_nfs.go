package fake

import (
	"context"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Find is fake implementation
func (o *NFSOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.NFS, error) {
	results, _ := find(ResourceNFS, zone, conditions)
	var values []*sacloud.NFS
	for _, res := range results {
		values = append(values, res.(*sacloud.NFS))
	}
	return values, nil
}

// Create is fake implementation
func (o *NFSOp) Create(ctx context.Context, zone string, param *sacloud.NFSCreateRequest) (*sacloud.NFS, error) {
	result := &sacloud.NFS{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "nfs"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]

	s.setNFS(zone, result)

	startPowerOn(ResourceNFS, zone, result.ID)
	return result, nil
}

// Read is fake implementation
func (o *NFSOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.NFS, error) {
	value := s.getNFSByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(ResourceNFS, id)
	}
	return value, nil
}

// Update is fake implementation
func (o *NFSOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.NFSUpdateRequest) (*sacloud.NFS, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *NFSOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(ResourceNFS, zone, id)
	return nil
}

// Boot is fake implementation
func (o *NFSOp) Boot(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceNFS, id, "Boot is failed")
	}

	startPowerOn(ResourceNFS, zone, id)

	return err
}

// Shutdown is fake implementation
func (o *NFSOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceNFS, id, "Shutdown is failed")
	}

	startPowerOff(ResourceNFS, zone, id)

	return err
}

// Reset is fake implementation
func (o *NFSOp) Reset(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceNFS, id, "Reset is failed")
	}

	startPowerOn(ResourceNFS, zone, id)

	return nil
}

// MonitorFreeDiskSize is fake implementation
func (o *NFSOp) MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.FreeDiskSizeActivity, error) {
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
		t := now.Add(time.Duration(i*-5) * time.Minute)
		res.Values = append(res.Values, naked.MonitorFreeDiskSizeValue{
			Time:         t,
			FreeDiskSize: float64(random(1000)),
		})
	}

	return res, nil
}

// MonitorInterface is fake implementation
func (o *NFSOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.InterfaceActivity, error) {
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
		t := now.Add(time.Duration(i*-5) * time.Minute)
		res.Values = append(res.Values, naked.MonitorInterfaceValue{
			Time:    t,
			Send:    float64(random(1000)),
			Receive: float64(random(1000)),
		})
	}

	return res, nil
}
