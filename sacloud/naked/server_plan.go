package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// ServerPlan サーバープラン
type ServerPlan struct {
	ID           types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string            `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	CPU          int               `json:",omitempty" yaml:"cpu,omitempty" structs:",omitempty"`
	MemoryMB     int               `json:",omitempty" yaml:"memory_mb,omitempty" structs:",omitempty"`
	Commitment   types.ECommitment `json:",omitempty" yaml:"commitment,omitempty" structs:",omitempty"`
	Generation   int               `json:",omitempty" yaml:"generation,omitempty" structs:",omitempty"`
	ServiceClass string            `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}
