package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// IPv6Net InternetリソースでのIPv6アドレス帯を表す
type IPv6Net struct {
	ID                 types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	ServiceID          types.ID   `json:",omitempty" yaml:"service_id,omitempty" structs:",omitempty"`
	IPv6Prefix         string     `json:",omitempty" yaml:"ipv6prefix,omitempty" structs:",omitempty"`
	IPv6PrefixLen      int        `json:",omitempty" yaml:"ipv6prefix_len,omitempty" structs:",omitempty"`
	IPv6PrefixTail     string     `json:",omitempty" yaml:"ipv6prefix_tail,omitempty" structs:",omitempty"`
	ServiceClass       string     `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	IPv6Table          *IPv6Table `json:",omitempty" yaml:"ipv6table,omitempty" structs:",omitempty"`
	NamedIPv6AddrCount int        `json:",omitempty" yaml:"named_ipv6addr_count,omitempty" structs:",omitempty"`
	CreatedAt          *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Switch             *Switch    `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
}

// IPv6Table IPv6テーブル
type IPv6Table struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}
