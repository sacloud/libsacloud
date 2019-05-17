package fake

import (
	"context"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Find is fake implementation
func (o *LoadBalancerOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.LoadBalancer, error) {
	results, _ := find(ResourceLoadBalancer, zone, conditions)
	var values []*sacloud.LoadBalancer
	for _, res := range results {
		values = append(values, res.(*sacloud.LoadBalancer))
	}
	return values, nil
}

// Create is fake implementation
func (o *LoadBalancerOp) Create(ctx context.Context, zone string, param *sacloud.LoadBalancerCreateRequest) (*sacloud.LoadBalancer, error) {
	result := &sacloud.LoadBalancer{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Class = "loadbalancer"
	result.Availability = types.Availabilities.Migrating
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""

	s.setLoadBalancer(zone, result)

	startPowerOn(ResourceLoadBalancer, zone, result.ID)
	return result, nil
}

// Read is fake implementation
func (o *LoadBalancerOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.LoadBalancer, error) {
	value := s.getLoadBalancerByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(ResourceLoadBalancer, id)
	}
	return value, nil
}

// Update is fake implementation
func (o *LoadBalancerOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.LoadBalancerUpdateRequest) (*sacloud.LoadBalancer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *LoadBalancerOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(ResourceLoadBalancer, zone, id)
	return nil
}

// Config is fake implementation
func (o *LoadBalancerOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	return err
}

// Boot is fake implementation
func (o *LoadBalancerOp) Boot(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceLoadBalancer, id, "Boot is failed")
	}

	startPowerOn(ResourceLoadBalancer, zone, id)

	return err
}

// Shutdown is fake implementation
func (o *LoadBalancerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceLoadBalancer, id, "Shutdown is failed")
	}

	startPowerOff(ResourceLoadBalancer, zone, id)

	return err
}

// Reset is fake implementation
func (o *LoadBalancerOp) Reset(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(ResourceLoadBalancer, id, "Reset is failed")
	}

	startPowerOn(ResourceLoadBalancer, zone, id)

	return nil
}

// MonitorInterface is fake implementation
func (o *LoadBalancerOp) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.InterfaceActivity, error) {
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

// Status is fake implementation
func (o *LoadBalancerOp) Status(ctx context.Context, zone string, id types.ID) ([]*sacloud.LoadBalancerStatus, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

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

	return results, nil

}
