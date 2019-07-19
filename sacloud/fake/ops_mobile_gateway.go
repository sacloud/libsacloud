package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *MobileGatewayOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.MobileGatewayFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.MobileGateway
	for _, res := range results {
		dest := &sacloud.MobileGateway{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.MobileGatewayFindResult{
		Total:          len(results),
		Count:          len(results),
		From:           0,
		MobileGateways: values,
	}, nil
}

// Create is fake implementation
func (o *MobileGatewayOp) Create(ctx context.Context, zone string, param *sacloud.MobileGatewayCreateRequest) (*sacloud.MobileGateway, error) {
	result := &sacloud.MobileGateway{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.Class = "mobilegateway"
	result.ZoneID = zoneIDs[zone]
	result.SettingsHash = ""

	s.setMobileGateway(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *MobileGatewayOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.MobileGateway, error) {
	value := s.getMobileGatewayByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.MobileGateway{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *MobileGatewayOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.MobileGatewayUpdateRequest) (*sacloud.MobileGateway, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *MobileGatewayOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}

// Config is fake implementation
func (o *MobileGatewayOp) Config(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	return err
}

// Boot is fake implementation
func (o *MobileGatewayOp) Boot(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Boot is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// Shutdown is fake implementation
func (o *MobileGatewayOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Shutdown is failed")
	}

	startPowerOff(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return err
}

// Reset is fake implementation
func (o *MobileGatewayOp) Reset(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.InstanceStatus.IsUp() {
		return newErrorConflict(o.key, id, "Reset is failed")
	}

	startPowerOn(o.key, zone, func() (interface{}, error) {
		return o.Read(context.Background(), zone, id)
	})

	return nil
}

// ConnectToSwitch is fake implementation
func (o *MobileGatewayOp) ConnectToSwitch(ctx context.Context, zone string, id types.ID, switchID types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	for _, nic := range value.Interfaces {
		if nic.Index == 1 {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] already connected to switch", 1))
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
	iface, err := ifOp.Create(ctx, zone, &sacloud.InterfaceCreateRequest{ServerID: id})
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	if err := ifOp.ConnectToSwitch(ctx, zone, iface.ID, switchID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	iface, err = ifOp.Read(ctx, zone, iface.ID)
	if err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	mobileGatewayInterface := &sacloud.MobileGatewayInterface{}
	copySameNameField(iface, mobileGatewayInterface)
	value.Interfaces = append(value.Interfaces, mobileGatewayInterface)

	s.setMobileGateway(zone, value)
	return nil
}

// DisconnectFromSwitch is fake implementation
func (o *MobileGatewayOp) DisconnectFromSwitch(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var exists bool
	var nicID types.ID
	var interfaces []*sacloud.MobileGatewayInterface

	for _, nic := range value.Interfaces {
		if nic.Index == 1 {
			exists = true
			nicID = nic.ID
		} else {
			interfaces = append(interfaces, nic)
		}
	}
	if !exists {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("nic[%d] is not exists", 1))
	}

	ifOp := NewInterfaceOp()
	if err := ifOp.DisconnectFromSwitch(ctx, zone, nicID); err != nil {
		return newErrorConflict(o.key, types.ID(0), err.Error())
	}

	value.Interfaces = interfaces
	s.setMobileGateway(zone, value)
	return nil
}

// GetDNS is fake implementation
func (o *MobileGatewayOp) GetDNS(ctx context.Context, zone string, id types.ID) (*sacloud.MobileGatewayDNSSetting, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	dns := s.getByID(o.dnsStoreKey(), zone, id)
	if dns == nil {
		return &sacloud.MobileGatewayDNSSetting{
			DNS1: "133.242.0.1",
			DNS2: "133.242.0.2",
		}, nil
	}
	return dns.(*sacloud.MobileGatewayDNSSetting), nil
}

// SetDNS is fake implementation
func (o *MobileGatewayOp) SetDNS(ctx context.Context, zone string, id types.ID, param *sacloud.MobileGatewayDNSSetting) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.setWithID(o.dnsStoreKey(), zone, param, id)
	return nil
}

// GetSIMRoutes is fake implementation
func (o *MobileGatewayOp) GetSIMRoutes(ctx context.Context, zone string, id types.ID) ([]*sacloud.MobileGatewaySIMRoute, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	routes := s.getByID(o.simRoutesStoreKey(), zone, id)
	if routes == nil {
		return nil, nil
	}
	return routes.([]*sacloud.MobileGatewaySIMRoute), nil
}

// SetSIMRoutes is fake implementation
func (o *MobileGatewayOp) SetSIMRoutes(ctx context.Context, zone string, id types.ID, param []*sacloud.MobileGatewaySIMRouteParam) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.setWithID(o.simRoutesStoreKey(), zone, param, id)
	return nil
}

// ListSIM is fake implementation
func (o *MobileGatewayOp) ListSIM(ctx context.Context, zone string, id types.ID) ([]*sacloud.MobileGatewaySIMInfo, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	sims := s.getByID(o.simsStoreKey(), zone, id)
	if sims == nil {
		return nil, nil
	}
	return sims.([]*sacloud.MobileGatewaySIMInfo), nil
}

// AddSIM is fake implementation
func (o *MobileGatewayOp) AddSIM(ctx context.Context, zone string, id types.ID, param *sacloud.MobileGatewayAddSIMRequest) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var sims []*sacloud.MobileGatewaySIMInfo
	rawSIMs := s.getByID(o.simsStoreKey(), zone, id)
	if rawSIMs != nil {
		sims = rawSIMs.([]*sacloud.MobileGatewaySIMInfo)
		for _, sim := range sims {
			if sim.ResourceID == param.SIMID {
				return newErrorBadRequest(o.key, id, fmt.Sprintf("SIM %s already exists", param.SIMID))
			}
		}
	}

	simOp := NewSIMOp()
	simInfo, err := simOp.Status(context.Background(), types.StringID(param.SIMID))
	if err != nil {
		return err
	}
	sim := &sacloud.MobileGatewaySIMInfo{}
	copySameNameField(simInfo, sim)

	sims = append(sims, sim)

	s.setWithID(o.simsStoreKey(), zone, sims, id)
	return nil
}

// DeleteSIM is fake implementation
func (o *MobileGatewayOp) DeleteSIM(ctx context.Context, zone string, id types.ID, simID types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	var sims, updSIMs []*sacloud.MobileGatewaySIMInfo
	rawSIMs := s.getByID(o.simsStoreKey(), zone, id)
	if rawSIMs != nil {
		sims = rawSIMs.([]*sacloud.MobileGatewaySIMInfo)
		for _, sim := range sims {
			if sim.ResourceID != simID.String() {
				updSIMs = append(updSIMs, sim)
			}
		}
		if len(sims) != len(updSIMs) {
			s.setWithID(o.simsStoreKey(), zone, updSIMs, id)
			return nil
		}
	}
	return newErrorBadRequest(o.key, id, fmt.Sprintf("SIM %d is not exists", simID))
}

// Logs is fake implementation
func (o *MobileGatewayOp) Logs(ctx context.Context, zone string, id types.ID) ([]*sacloud.MobileGatewaySIMLogs, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return []*sacloud.MobileGatewaySIMLogs{
		{
			Date:          time.Now(),
			SessionStatus: "UP",
			ResourceID:    types.ID(1).String(),
		},
	}, nil
}

// GetTrafficConfig is fake implementation
func (o *MobileGatewayOp) GetTrafficConfig(ctx context.Context, zone string, id types.ID) (*sacloud.MobileGatewayTrafficControl, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	config := s.getByID(o.trafficConfigStoreKey(), zone, id)
	if config == nil {
		return nil, nil
	}
	return config.(*sacloud.MobileGatewayTrafficControl), nil
}

// SetTrafficConfig is fake implementation
func (o *MobileGatewayOp) SetTrafficConfig(ctx context.Context, zone string, id types.ID, param *sacloud.MobileGatewayTrafficControl) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.setWithID(o.trafficConfigStoreKey(), zone, param, id)
	return nil
}

// DeleteTrafficConfig is fake implementation
func (o *MobileGatewayOp) DeleteTrafficConfig(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.trafficConfigStoreKey(), zone, id)
	return nil
}

// TrafficStatus is fake implementation
func (o *MobileGatewayOp) TrafficStatus(ctx context.Context, zone string, id types.ID) (*sacloud.MobileGatewayTrafficStatus, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &sacloud.MobileGatewayTrafficStatus{
		UplinkBytes:    100,
		DownlinkBytes:  100,
		TrafficShaping: true,
	}, nil
}

// MonitorInterface is fake implementation
func (o *MobileGatewayOp) MonitorInterface(ctx context.Context, zone string, id types.ID, index int, condition *sacloud.MonitorCondition) (*sacloud.InterfaceActivity, error) {
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

func (o *MobileGatewayOp) dnsStoreKey() string {
	return o.key + "DNS"
}

func (o *MobileGatewayOp) simRoutesStoreKey() string {
	return o.key + "SIMRoutes"
}

func (o *MobileGatewayOp) simsStoreKey() string {
	return o.key + "SIMRoutes"
}

func (o *MobileGatewayOp) trafficConfigStoreKey() string {
	return o.key + "TrafficConfig"
}
