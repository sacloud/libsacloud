package fake

import (
	"context"
	"net"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *InternetOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.InternetFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Internet
	for _, res := range results {
		dest := &sacloud.Internet{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.InternetFindResult{
		Total:     len(results),
		Count:     len(results),
		From:      0,
		Internets: values,
	}, nil
}

// Create is fake implementation
func (o *InternetOp) Create(ctx context.Context, zone string, param *sacloud.InternetCreateRequest) (*sacloud.InternetCreateResult, error) {
	if param.NetworkMaskLen == 0 {
		param.NetworkMaskLen = 28
	}
	if param.BandWidthMbps == 0 {
		param.BandWidthMbps = 100
	}

	result := &sacloud.Internet{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	// assign global address
	subnet := pool.nextSubnet(result.NetworkMaskLen)

	// create switch
	swOp := NewSwitchOp()
	swCreateResult, err := swOp.Create(ctx, zone, &sacloud.SwitchCreateRequest{
		Name:           result.Name,
		NetworkMaskLen: subnet.networkMaskLen,
		DefaultRoute:   subnet.defaultRoute,
	})
	if err != nil {
		return nil, err
	}
	sw := swCreateResult.Switch

	sSubnet := &sacloud.SwitchSubnet{
		ID:                   pool.generateID(),
		DefaultRoute:         subnet.defaultRoute,
		NetworkAddress:       subnet.networkAddress,
		NetworkMaskLen:       subnet.networkMaskLen,
		Internet:             result,
		AssignedIPAddressMax: subnet.addresses[len(subnet.addresses)-1],
		AssignedIPAddressMin: subnet.addresses[0],
	}
	sw.Subnets = append(sw.Subnets, sSubnet)

	// for Internet.Switch
	switchInfo := &sacloud.SwitchInfo{}
	copySameNameField(sw, switchInfo)

	iSubnet := &sacloud.InternetSubnet{
		ID:             sSubnet.ID,
		DefaultRoute:   sSubnet.DefaultRoute,
		NetworkAddress: sSubnet.NetworkAddress,
		NetworkMaskLen: sSubnet.NetworkMaskLen,
	}
	switchInfo.Subnets = []*sacloud.InternetSubnet{iSubnet}
	result.Switch = switchInfo

	s.setSwitch(zone, sw)
	s.setInternet(zone, result)
	return &sacloud.InternetCreateResult{
		IsOk:     true,
		Internet: result,
	}, nil
}

// Read is fake implementation
func (o *InternetOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.InternetReadResult, error) {
	value := s.getInternetByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Internet{}
	copySameNameField(value, dest)
	return &sacloud.InternetReadResult{
		IsOk:     true,
		Internet: dest,
	}, nil
}

// Update is fake implementation
func (o *InternetOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.InternetUpdateRequest) (*sacloud.InternetUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Internet
	copySameNameField(param, value)

	return &sacloud.InternetUpdateResult{
		IsOk:     true,
		Internet: value,
	}, nil
}

// Delete is fake implementation
func (o *InternetOp) Delete(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Internet

	swOp := NewSwitchOp()
	if err := swOp.Delete(ctx, zone, value.Switch.ID); err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}

// UpdateBandWidth is fake implementation
func (o *InternetOp) UpdateBandWidth(ctx context.Context, zone string, id types.ID, param *sacloud.InternetUpdateBandWidthRequest) (*sacloud.InternetUpdateBandWidthResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Internet

	value.BandWidthMbps = param.BandWidthMbps
	s.setInternet(zone, value)
	return &sacloud.InternetUpdateBandWidthResult{
		IsOk:     true,
		Internet: value,
	}, nil
}

// AddSubnet is fake implementation
func (o *InternetOp) AddSubnet(ctx context.Context, zone string, id types.ID, param *sacloud.InternetAddSubnetRequest) (*sacloud.InternetAddSubnetResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Internet

	// assign global address
	subnet := pool.nextSubnetFull(param.NetworkMaskLen, param.NextHop)

	// create switch
	swOp := NewSwitchOp()
	swReadResult, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return nil, err
	}
	sw := swReadResult.Switch

	sSubnet := &sacloud.SwitchSubnet{
		ID:                   pool.generateID(),
		NetworkAddress:       subnet.networkAddress,
		NetworkMaskLen:       subnet.networkMaskLen,
		NextHop:              param.NextHop,
		StaticRoute:          param.NextHop,
		Internet:             value,
		AssignedIPAddressMax: subnet.addresses[len(subnet.addresses)-1],
		AssignedIPAddressMin: subnet.addresses[0],
	}
	sw.Subnets = append(sw.Subnets, sSubnet)

	// for Internet.Switch
	iSubnet := &sacloud.InternetSubnet{
		ID:             sSubnet.ID,
		DefaultRoute:   sSubnet.DefaultRoute,
		NetworkAddress: sSubnet.NetworkAddress,
		NetworkMaskLen: sSubnet.NetworkMaskLen,
	}
	value.Switch.Subnets = append(value.Switch.Subnets, iSubnet)

	s.setSwitch(zone, sw)
	s.setInternet(zone, value)

	return &sacloud.InternetAddSubnetResult{
		IsOk: true,
		Subnet: &sacloud.InternetSubnetOperationResult{
			ID:             sSubnet.ID,
			NextHop:        param.NextHop,
			StaticRoute:    param.NextHop,
			NetworkMaskLen: sSubnet.NetworkMaskLen,
			NetworkAddress: sSubnet.NetworkAddress,
			IPAddresses:    subnet.addresses,
		},
	}, nil
}

// UpdateSubnet is fake implementation
func (o *InternetOp) UpdateSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID, param *sacloud.InternetUpdateSubnetRequest) (*sacloud.InternetUpdateSubnetResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Internet
	// create switch
	swOp := NewSwitchOp()
	swReadResult, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return nil, err
	}
	sw := swReadResult.Switch

	var nwMaskLen int
	var nwAddress, minAddr, maxAddr string
	var addresses []string

	for _, subnet := range sw.Subnets {
		if subnet.ID == subnetID {
			subnet.NextHop = param.NextHop
			subnet.StaticRoute = param.NextHop

			minAddr = subnet.AssignedIPAddressMin
			maxAddr = subnet.AssignedIPAddressMax
			nwMaskLen = subnet.NetworkMaskLen
			nwAddress = subnet.NetworkAddress
		}
	}

	for _, subnet := range value.Switch.Subnets {
		if subnet.ID == subnetID {
			subnet.NextHop = param.NextHop
			subnet.StaticRoute = param.NextHop
		}
	}

	baseIP := net.ParseIP(minAddr).To4()
	min := baseIP[3]
	max := net.ParseIP(maxAddr).To4()[3]

	var i byte
	for (min + i) <= max { //境界含む
		ip := net.IPv4(baseIP[0], baseIP[1], baseIP[2], baseIP[3]+i)
		addresses = append(addresses, ip.String())
		i++
	}

	s.setSwitch(zone, sw)
	s.setInternet(zone, value)
	return &sacloud.InternetUpdateSubnetResult{
		IsOk: true,
		Subnet: &sacloud.InternetSubnetOperationResult{
			ID:             subnetID,
			NextHop:        param.NextHop,
			StaticRoute:    param.NextHop,
			NetworkMaskLen: nwMaskLen,
			NetworkAddress: nwAddress,
			IPAddresses:    addresses,
		},
	}, nil
}

// DeleteSubnet is fake implementation
func (o *InternetOp) DeleteSubnet(ctx context.Context, zone string, id types.ID, subnetID types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Internet

	// create switch
	swOp := NewSwitchOp()
	swReadResult, err := swOp.Read(ctx, zone, value.Switch.ID)
	if err != nil {
		return err
	}
	sw := swReadResult.Switch

	var sSubnets []*sacloud.SwitchSubnet
	for _, subnet := range sw.Subnets {
		if subnet.ID != subnetID {
			sSubnets = append(sSubnets, subnet)
		}
	}
	sw.Subnets = sSubnets

	var iSubnets []*sacloud.InternetSubnet
	for _, subnet := range value.Switch.Subnets {
		if subnet.ID != subnetID {
			iSubnets = append(iSubnets, subnet)
		}
	}
	value.Switch.Subnets = iSubnets

	s.setSwitch(zone, sw)
	s.setInternet(zone, value)
	return nil
}

// Monitor is fake implementation
func (o *InternetOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.InternetMonitorResult, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.RouterActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorRouterValue{
			Time: now.Add(time.Duration(i*-5) * time.Minute),
			In:   float64(random(1000)),
			Out:  float64(random(1000)),
		})
	}

	return &sacloud.InternetMonitorResult{
		IsOk: true,
		Data: res,
	}, nil
}
