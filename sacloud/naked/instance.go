package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Instance サーバなどの起動情報
type Instance struct {
	Host            *Host                       `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	Status          types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	BeforeStatus    types.EServerInstanceStatus `json:",omitempty" yaml:"before_status,omitempty" structs:",omitempty"`
	StatusChangedAt *time.Time                  `json:",omitempty" yaml:"status_changed_at,omitempty" structs:",omitempty"`
	ModifiedAt      *time.Time                  `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Warnings        string                      `json:",omitempty" yaml:"warnings,omitempty" structs:",omitempty"`
	WarningsValue   int                         `json:",omitempty" yaml:"warnings_value,omitempty" structs:",omitempty"`
}
