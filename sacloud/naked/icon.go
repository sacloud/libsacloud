package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

// Icon アイコン
type Icon struct {
	ID           int64               `json:"ID,omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:"Name,omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Tags         []string            `json:"Tags" yaml:"tags"`
	Availability enums.EAvailability `json:"Availability,omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        string              `json:"Scope,omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	URL          string              `json:"URL,omitempty" yaml:"url,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:"CreatedAt,omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:"ModifiedAt,omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
}
