package fake

import (
	"context"
	"net"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ProxyLBOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ProxyLBFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.ProxyLB
	for _, res := range results {
		dest := &sacloud.ProxyLB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ProxyLBFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		ProxyLBs: values,
	}, nil
}

// Create is fake implementation
func (o *ProxyLBOp) Create(ctx context.Context, zone string, param *sacloud.ProxyLBCreateRequest) (*sacloud.ProxyLB, error) {
	result := &sacloud.ProxyLB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.Class = "proxylb"

	vip := pool.nextSharedIP()
	vipNet := net.IPNet{IP: vip, Mask: []byte{255, 255, 255, 0}}
	result.ProxyNetworks = []string{vipNet.String()}
	if param.UseVIPFailover {
		result.FQDN = "fake.proxylb.sakura.ne.jp"
	}

	status := &sacloud.ProxyLBHealth{
		ActiveConn: 10,
		CPS:        10,
		CurrentVIP: vip.String(),
	}
	for _, server := range param.Servers {
		status.Servers = append(status.Servers, &sacloud.LoadBalancerServerStatus{
			ActiveConn: 10,
			Status:     types.ServerInstanceStatuses.Up,
			IPAddress:  server.IPAddress,
			Port:       types.StringNumber(server.Port),
			CPS:        10,
		})
	}
	s.setWithID(ResourceProxyLB+"Status", zone, status, result.ID)

	s.setProxyLB(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *ProxyLBOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.ProxyLB, error) {
	value := s.getProxyLBByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.ProxyLB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *ProxyLBOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ProxyLBUpdateRequest) (*sacloud.ProxyLB, error) {
	value, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	status := s.getByID(ResourceProxyLB+"Status", zone, id).(*sacloud.ProxyLBHealth)
	status.Servers = []*sacloud.LoadBalancerServerStatus{}
	for _, server := range param.Servers {
		status.Servers = append(status.Servers, &sacloud.LoadBalancerServerStatus{
			ActiveConn: 10,
			Status:     types.ServerInstanceStatuses.Up,
			IPAddress:  server.IPAddress,
			Port:       types.StringNumber(server.Port),
			CPS:        10,
		})
	}
	s.setWithID(ResourceProxyLB+"Status", zone, status, id)

	return value, nil
}

// Delete is fake implementation
func (o *ProxyLBOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}

	s.delete(ResourceProxyLB+"Status", zone, id)
	s.delete(ResourceProxyLB+"Certs", zone, id)
	s.delete(o.key, sacloud.APIDefaultZone, id)

	return nil
}

// ChangePlan is fake implementation
func (o *ProxyLBOp) ChangePlan(ctx context.Context, zone string, id types.ID, param *sacloud.ProxyLBChangePlanRequest) (*sacloud.ProxyLB, error) {
	value, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}

	value.Plan = param.Plan
	return value, err
}

// GetCertificates is fake implementation
func (o *ProxyLBOp) GetCertificates(ctx context.Context, zone string, id types.ID) (*sacloud.ProxyLBCertificates, error) {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}

	v := s.getByID(ResourceProxyLB+"Certs", zone, id)
	if v != nil {
		return v.(*sacloud.ProxyLBCertificates), nil
	}

	return nil, err
}

// SetCertificates is fake implementation
func (o *ProxyLBOp) SetCertificates(ctx context.Context, zone string, id types.ID, param *sacloud.ProxyLBSetCertificatesRequest) (*sacloud.ProxyLBCertificates, error) {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}

	cert := &sacloud.ProxyLBCertificates{}
	copySameNameField(param, cert)
	cert.CertificateCommonName = "dummy-common-name.org"
	cert.CertificateEndDate = time.Now().Add(365 * 24 * time.Hour)

	s.set(ResourceProxyLB+"Certs", zone, cert)
	return cert, nil
}

// DeleteCertificates is fake implementation
func (o *ProxyLBOp) DeleteCertificates(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}

	v := s.getByID(ResourceProxyLB+"Certs", zone, id)
	if v != nil {
		s.delete(ResourceProxyLB+"Certs", zone, id)
	}
	return nil
}

// RenewLetsEncryptCert is fake implementation
func (o *ProxyLBOp) RenewLetsEncryptCert(ctx context.Context, zone string, id types.ID) error {
	return nil
}

// HealthStatus is fake implementation
func (o *ProxyLBOp) HealthStatus(ctx context.Context, zone string, id types.ID) (*sacloud.ProxyLBHealth, error) {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}

	return s.getByID(ResourceProxyLB+"Status", zone, id).(*sacloud.ProxyLBHealth), nil
}
