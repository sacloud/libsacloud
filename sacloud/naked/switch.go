package naked

import (
	"time"

	"github.com/sacloud/libsacloud/sacloud/types"
)

// Switch スイッチ
type Switch struct {
	ID          types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string       `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags        []string     `json:"" yaml:"tags"`
	Icon        *Icon        `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt   *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt  *time.Time   `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Scope       types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Subnet      *Subnet      `json:",omitempty" yaml:"subnet,omitempty" structs:",omitempty"`
	UserSubnet  *UserSubnet  `json:",omitempty" yaml:"user_subnet,omitempty" structs:",omitempty"`
	Zone        *Zone        `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Internet    *Internet    `json:",omitempty" yaml:"internet,omitempty" structs:",omitempty"`
	Subnets     []*Subnet    `json:",omitempty" yaml:"subnets,omitempty" structs:",omitempty"`
	IPv6Nets    []*IPv6Net   `json:",omitempty" yaml:"ipv6nets,omitempty" structs:",omitempty"`
	Bridge      *Bridge      `json:",omitempty" yaml:"bridge,omitempty" structs:",omitempty"`
}
