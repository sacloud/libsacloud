package fake

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *SIMOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.SIMFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.SIM
	for _, res := range results {
		dest := &sacloud.SIM{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.SIMFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		SIMs:  values,
	}, nil
}

// Create is fake implementation
func (o *SIMOp) Create(ctx context.Context, param *sacloud.SIMCreateRequest) (*sacloud.SIM, error) {
	result := &sacloud.SIM{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillModifiedAt)

	result.Class = "sim"
	result.Availability = types.Availabilities.Available
	result.Info = &sacloud.SIMInfo{
		ICCID: param.ICCID,
	}

	s.setSIM(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *SIMOp) Read(ctx context.Context, id types.ID) (*sacloud.SIM, error) {
	value := s.getSIMByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.SIM{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *SIMOp) Update(ctx context.Context, id types.ID, param *sacloud.SIMUpdateRequest) (*sacloud.SIM, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *SIMOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	s.delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}

// Activate is fake implementation
func (o *SIMOp) Activate(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.Activated {
		return errors.New("SIM[%d] is already activated")
	}
	value.Info.Activated = true
	value.Info.ActivatedDate = time.Now()
	value.Info.DeactivatedDate = time.Time{}
	return nil
}

// Deactivate is fake implementation
func (o *SIMOp) Deactivate(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if !value.Info.Activated {
		return errors.New("SIM[%d] is already deactivated")
	}
	value.Info.Activated = false
	value.Info.ActivatedDate = time.Time{}
	value.Info.DeactivatedDate = time.Now()
	return nil
}

// AssignIP is fake implementation
func (o *SIMOp) AssignIP(ctx context.Context, id types.ID, param *sacloud.SIMAssignIPRequest) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IP != "" {
		return errors.New("SIM[%d] already has IPAddress")
	}
	value.Info.IP = param.IP
	return nil
}

// ClearIP is fake implementation
func (o *SIMOp) ClearIP(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IP == "" {
		return errors.New("SIM[%d] doesn't have IPAddress")
	}
	value.Info.IP = ""
	return nil
}

// IMEILock is fake implementation
func (o *SIMOp) IMEILock(ctx context.Context, id types.ID, param *sacloud.SIMIMEILockRequest) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if value.Info.IMEILock {
		return errors.New("SIM[%d] is already locked with IMEI")
	}
	value.Info.IMEILock = true
	value.Info.ConnectedIMEI = param.IMEI
	return nil
}

// IMEIUnlock is fake implementation
func (o *SIMOp) IMEIUnlock(ctx context.Context, id types.ID) error {
	value, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	if !value.Info.IMEILock {
		return errors.New("SIM[%d] is not locked with IMEI")
	}
	value.Info.IMEILock = false
	value.Info.ConnectedIMEI = ""
	return nil
}

// Logs is fake implementation
func (o *SIMOp) Logs(ctx context.Context, id types.ID) (*sacloud.SIMLogsResult, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &sacloud.SIMLogsResult{
		Total: 1,
		From:  0,
		Count: 1,
		Logs: []*sacloud.SIMLog{
			{
				Date:          time.Now(),
				SessionStatus: "up",
				ResourceID:    value.ID.String(),
				IMEI:          value.Info.ConnectedIMEI,
				IMSI:          strings.Join(value.Info.IMSI, " "),
			},
		},
	}, nil
}

// GetNetworkOperator is fake implementation
func (o *SIMOp) GetNetworkOperator(ctx context.Context, id types.ID) ([]*sacloud.SIMNetworkOperatorConfig, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return []*sacloud.SIMNetworkOperatorConfig{
		{
			Allow:       true,
			CountryCode: "XX",
			Name:        "name",
		},
	}, nil
}

// SetNetworkOperator is fake implementation
func (o *SIMOp) SetNetworkOperator(ctx context.Context, id types.ID, configs []*sacloud.SIMNetworkOperatorConfig) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// MonitorSIM is fake implementation
func (o *SIMOp) MonitorSIM(ctx context.Context, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.LinkActivity, error) {
	_, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.LinkActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorLinkValue{
			Time:        now.Add(time.Duration(i*-5) * time.Minute),
			UplinkBPS:   float64(random(1000)),
			DownlinkBPS: float64(random(1000)),
		})
	}

	return res, nil
}

// Status is fake implementation
func (o *SIMOp) Status(ctx context.Context, id types.ID) (*sacloud.SIMInfo, error) {
	v, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return v.Info, nil
}
