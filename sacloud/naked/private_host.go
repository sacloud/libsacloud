package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// PrivateHost 専有ホスト
type PrivateHost struct {
	ID               types.ID            `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Name             string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description      string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags             []string            `json:"" yaml:"tags"`
	Icon             *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt        *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Plan             *ProductPrivateHost `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Host             *Host               `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	AssignedCPU      int                 `json:",omitempty" yaml:"assigned_cpu,omitempty" structs:",omitempty"`
	AssignedMemoryMB int                 `json:",omitempty" yaml:"assigned_memory_mb,omitempty" structs:",omitempty"`
}

// ProductPrivateHost 専有ホストプラン
type ProductPrivateHost struct {
	ID           types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Class        string   `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	CPU          int      `json:",omitempty" yaml:"cpu,omitempty" structs:",omitempty"`
	MemoryMB     int      `json:",omitempty" yaml:"memory_mb,omitempty" structs:",omitempty"`
	ServiceClass string   `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}
