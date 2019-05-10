package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

// Switch スイッチ
type Switch struct {
	ID          int64        `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string       `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string       `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags        []string     `json:"" yaml:"tags"`
	Icon        *Icon        `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt   *time.Time   `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt  *time.Time   `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Scope       enums.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	UserSubnet  *UserSubnet  `json:",omitempty" yaml:"user_subnet,omitempty" structs:",omitempty"`
	Zone        *Zone        `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}
