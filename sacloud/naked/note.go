package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

// Note スタートアップスクリプト
type Note struct {
	ID           int64               `json:"ID,omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:"Name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:"Description,omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"Tags" yaml:"tags"`
	Availability enums.EAvailability `json:"Availability,omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        enums.EScope        `json:"Scope,omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	Class        string              `json:"Class,omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Content      string              `json:"Content,omitempty" yaml:"content,omitempty" structs:",omitempty"`
	Icon         *Icon               `json:"Icon,omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:"CreatedAt,omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:"ModifiedAt,omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
}
