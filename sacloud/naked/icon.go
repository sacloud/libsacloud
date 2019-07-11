package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Icon アイコン
type Icon struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Tags         []string            `yaml:"tags"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        types.EScope        `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	URL          string              `json:",omitempty" yaml:"url,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`

	Image string `json:",omitempty" yaml:"image,omitempty" structs:",omitempty"` // 画像データBase64文字列(画像アップロード時に利用)
}
