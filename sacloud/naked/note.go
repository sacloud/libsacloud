package naked

import (
	"time"

	"github.com/sacloud/libsacloud/sacloud/types"
)

// Note スタートアップスクリプト
type Note struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        types.EScope        `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Content      string              `json:",omitempty" yaml:"content,omitempty" structs:",omitempty"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
}
