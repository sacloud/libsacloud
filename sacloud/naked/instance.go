package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Instance サーバなどの起動情報
type Instance struct {
	Host            *Host                       `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	Status          types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
	StatusChangedAt *time.Time                  `json:",omitempty" yaml:"status_changed_at,omitempty" structs:",omitempty"`
}
