package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

// List is fake implementation
func (o *IPAddressOp) List(ctx context.Context, zone string) (*sacloud.IPAddressListResult, error) {
	return &sacloud.IPAddressListResult{
		Total: 1,
		Count: 1,
		From:  0,
		IPAddress: []*sacloud.IPAddress{
			{
				HostName:  "",
				IPAddress: "192.0.2.1",
			},
		},
	}, nil
}

// Read is fake implementation
func (o *IPAddressOp) Read(ctx context.Context, zone string, ipAddress string) (*sacloud.IPAddress, error) {
	return &sacloud.IPAddress{
		HostName:  "",
		IPAddress: ipAddress,
	}, nil
}

// UpdateHostName is fake implementation
func (o *IPAddressOp) UpdateHostName(ctx context.Context, zone string, ipAddress string, hostName string) (*sacloud.IPAddress, error) {
	return &sacloud.IPAddress{
		HostName:  hostName,
		IPAddress: ipAddress,
	}, nil
}
