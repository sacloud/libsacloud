package naked

import (
	"time"

	"github.com/sacloud/libsacloud/sacloud/types"
)

// Subnet サブネット
type Subnet struct {
	ID             types.ID    `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	ServiceClass   string      `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	CreatedAt      *time.Time  `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	DefaultRoute   string      `json:",omitempty" yaml:"default_route,omitempty" structs:",omitempty"`
	NetworkAddress string      `json:",omitempty" yaml:"network_address,omitempty" structs:",omitempty"`
	NetworkMaskLen int         `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
	ServiceID      int64       `json:",omitempty" yaml:"service_id,omitempty" structs:",omitempty"`
	StaticRoute    string      `json:",omitempty" yaml:"static_route,omitempty" structs:",omitempty"`
	NextHop        string      `json:",omitempty" yaml:"next_hop,omitempty" structs:",omitempty"`
	Switch         *Switch     `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	Internet       *Internet   `json:",omitempty" yaml:"internet,omitempty" structs:",omitempty"`
	IPAddresses    interface{} `json:",omitempty" yaml:"ip_addresses,omitempty" structs:",omitempty"`
}

// SubnetIPAddressRange ルータ+スイッチのスイッチ配下から参照できるSubnetでの割り当てられているIPアドレス範囲
type SubnetIPAddressRange struct {
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}
