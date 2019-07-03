package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// CDROM ISOイメージ
type CDROM struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	DisplayOrder int                 `json:",omitempty" yaml:"display_order,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Scope        types.EScope        `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	StorageClass string              `json:",omitempty" yaml:"storage_class,omitempty" structs:",omitempty"`
	SizeMB       int                 `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
	Storage      *Storage            `json:",omitempty" yaml:"storage,omitempty" structs:",omitempty"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
}
