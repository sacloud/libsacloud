package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

// Instance サーバなどの起動情報
type Instance struct {
	Host            *Host                       `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	Status          enums.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	StatusChangedAt *time.Time                  `json:",omitempty" yaml:"status_changed_at,omitempty" structs:",omitempty"`
}
