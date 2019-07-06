package fake

import (
	"context"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *LoadBalancerOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.LoadBalancerFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.LoadBalancer
	for _, res := range results {
		dest := &sacloud.LoadBalancer{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.LoadBalancerFindResult{
		Total:         len(results),
		Count:         len(results),
		From:          0,
		LoadBalancers: values,
	}, nil
}

// Create is fake implementation
func (o *LoadBalancerOp) Create(ctx context.Context, zone string, param *sacloud.LoadBalancerCreateRequest) (*sacloud.LoadBalancerCreateResult, error) {
	result := &sacloud.LoadBalancer{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "loadbalancer"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""

	s.setLoadBalancer(zone, result)

	id := result.ID
	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.LoadBalancer, nil
	})
	return &sacloud.LoadBalancerCreateResult{
		IsOk:         true,
		LoadBalancer: result,
	}, nil
}

// Read is fake implementation
func (o *LoadBalancerOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.LoadBalancerReadResult, error) {
	value := s.getLoadBalancerByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.LoadBalancer{}
	copySameNameField(value, dest)
	return &sacloud.LoadBalancerReadResult{
		IsOk:         true,
		LoadBalancer: dest,
	}, nil
}

// Update is fake implementation
func (o *LoadBalancerOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.LoadBalancerUpdateRequest) (*sacloud.LoadBalancerUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.LoadBalancer

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.LoadBalancerUpdateResult{
		IsOk:         true,
		LoadBalancer: value,
	}, nil
}

// Delete is fake implementation
func (o *LoadBalancerOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *LoadBalancerOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	return err
}

// Boot is fake implementation
func (o *LoadBalancerOp) Boot(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.LoadBalancer
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.LoadBalancer, nil
	})

	return err
}

// Shutdown is fake implementation
func (o *LoadBalancerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.LoadBalancer
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.LoadBalancer, nil
	})

	return err
}

// Reset is fake implementation
func (o *LoadBalancerOp) Reset(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.LoadBalancer
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.LoadBalancer, nil
	})

	return nil
}

// MonitorInterface is fake implementation
func (o *LoadBalancerOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.LoadBalancerMonitorInterfaceResult, error) {
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

	return &sacloud.LoadBalancerMonitorInterfaceResult{
		IsOk:              true,
		InterfaceActivity: res,
	}, nil
}

// Status is fake implementation
func (o *LoadBalancerOp) Status(ctx context.Context, zone string, id types.ID) (*sacloud.LoadBalancerStatusResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.LoadBalancer

	var results []*sacloud.LoadBalancerStatus
	for _, vip := range value.VirtualIPAddresses {
		status := &sacloud.LoadBalancerStatus{
			VirtualIPAddress: vip.VirtualIPAddress,
			Port:             vip.Port,
			CPS:              types.StringNumber(random(100)),
		}
		var servers []*sacloud.LoadBalancerServerStatus
		for _, server := range vip.Servers {
			servers = append(servers, &sacloud.LoadBalancerServerStatus{
				ActiveConn: types.StringNumber(random(10)),
				Status:     types.ServerInstanceStatuses.Up,
				IPAddress:  server.IPAddress,
				Port:       server.Port,
				CPS:        types.StringNumber(random(100)),
			})
		}
		status.Servers = servers

		results = append(results, status)
	}

	return &sacloud.LoadBalancerStatusResult{
		Status: results,
	}, nil

}
