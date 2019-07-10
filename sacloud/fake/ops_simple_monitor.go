package fake

import (
	"context"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *SimpleMonitorOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.SimpleMonitorFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.SimpleMonitor
	for _, res := range results {
		dest := &sacloud.SimpleMonitor{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.SimpleMonitorFindResult{
		Total:          len(results),
		Count:          len(results),
		From:           0,
		SimpleMonitors: values,
	}, nil
}

// Create is fake implementation
func (o *SimpleMonitorOp) Create(ctx context.Context, zone string, param *sacloud.SimpleMonitorCreateRequest) (*sacloud.SimpleMonitor, error) {
	result := &sacloud.SimpleMonitor{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Name = param.Target
	result.Class = "simplemon"
	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"

	s.setSimpleMonitor(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *SimpleMonitorOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.SimpleMonitor, error) {
	value := s.getSimpleMonitorByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.SimpleMonitor{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SimpleMonitorOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.SimpleMonitorUpdateRequest) (*sacloud.SimpleMonitor, error) {
	value, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *SimpleMonitorOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}

// MonitorResponseTime is fake implementation
func (o *SimpleMonitorOp) MonitorResponseTime(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.ResponseTimeSecActivity, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.ResponseTimeSecActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorResponseTimeSecValue{
			Time:            now.Add(time.Duration(i*-5) * time.Minute),
			ResponseTimeSec: float64(random(1000)),
		})
	}

	return res, nil
}

// HealthStatus is fake implementation
func (o *SimpleMonitorOp) HealthStatus(ctx context.Context, zone string, id types.ID) (*sacloud.SimpleMonitorHealthStatus, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &sacloud.SimpleMonitorHealthStatus{
		LastCheckedAt:       time.Now(),
		LastHealthChangedAt: time.Now(),
		Health:              types.SimpleMonitorHealth.Up,
	}, nil

}
