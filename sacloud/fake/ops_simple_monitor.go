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
func (o *SimpleMonitorOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.SimpleMonitorFindResult, error) {
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
func (o *SimpleMonitorOp) Create(ctx context.Context, param *sacloud.SimpleMonitorCreateRequest) (*sacloud.SimpleMonitor, error) {
	result := &sacloud.SimpleMonitor{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Name = param.Target
	result.Class = "simplemon"
	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	if result.DelayLoop == 0 {
		result.DelayLoop = 60
	}

	putSimpleMonitor(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *SimpleMonitorOp) Read(ctx context.Context, id types.ID) (*sacloud.SimpleMonitor, error) {
	value := getSimpleMonitorByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.SimpleMonitor{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SimpleMonitorOp) Update(ctx context.Context, id types.ID, param *sacloud.SimpleMonitorUpdateRequest) (*sacloud.SimpleMonitor, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	if value.DelayLoop == 0 {
		value.DelayLoop = 60
	}
	putSimpleMonitor(sacloud.APIDefaultZone, value)
	return value, nil
}

// Patch is fake implementation
func (o *SimpleMonitorOp) Patch(ctx context.Context, id types.ID, param *sacloud.SimpleMonitorPatchRequest) (*sacloud.SimpleMonitor, error) {
	value, err := o.Read(ctx, id)
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

	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}
	if param.PatchEmptyToDelayLoop {
		value.DelayLoop = 0
	}
	if param.PatchEmptyToEnabled {
		value.Enabled = types.StringFlag(false)
	}
	if param.PatchEmptyToHealthCheck {
		value.HealthCheck = nil
	}
	if param.PatchEmptyToNotifyEmailEnabled {
		value.NotifyEmailEnabled = types.StringFlag(false)
	}
	if param.PatchEmptyToNotifyEmailHTML {
		value.NotifyEmailHTML = types.StringFlag(false)
	}
	if param.PatchEmptyToNotifySlackEnabled {
		value.NotifySlackEnabled = types.StringFlag(false)
	}
	if param.PatchEmptyToSlackWebhooksURL {
		value.SlackWebhooksURL = ""
	}

	putSimpleMonitor(sacloud.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *SimpleMonitorOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}

// MonitorResponseTime is fake implementation
func (o *SimpleMonitorOp) MonitorResponseTime(ctx context.Context, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.ResponseTimeSecActivity, error) {
	_, err := o.Read(ctx, id)
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
func (o *SimpleMonitorOp) HealthStatus(ctx context.Context, id types.ID) (*sacloud.SimpleMonitorHealthStatus, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &sacloud.SimpleMonitorHealthStatus{
		LastCheckedAt:       time.Now(),
		LastHealthChangedAt: time.Now(),
		Health:              types.SimpleMonitorHealth.Up,
	}, nil

}
