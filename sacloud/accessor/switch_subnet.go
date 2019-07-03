package accessor

import (
	"net"

	"github.com/sacloud/libsacloud/v2/pkg/cidr"
)

// AssignedIPAddress スイッチ+ルータの割り当てられたIPアドレスリスト
type AssignedIPAddress interface {
	GetAssignedIPAddressMax() string
	SetAssignedIPAddressMax(v string)
	GetAssignedIPAddressMin() string
	SetAssignedIPAddressMin(v string)
}

// GetAssignedIPAddresses 最小/最大IPアドレスからIPアドレスリストを算出して返す
func GetAssignedIPAddresses(target AssignedIPAddress) []string {
	base := net.ParseIP(target.GetAssignedIPAddressMin())
	max := net.ParseIP(target.GetAssignedIPAddressMax())
	addresses := []string{base.String()}

	for {
		current := cidr.Inc(base)
		addresses = append(addresses, current.String())

		if current.Equal(max) {
			break
		}
		base = current
	}

	return addresses
}
