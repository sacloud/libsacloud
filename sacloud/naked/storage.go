package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Storage ストレージ
type Storage struct {
	ID          types.ID  `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string    `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string    `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Generation  int       `json:",omimtempty" yaml:"generation,omitempty" structs:",omitempty"`
	Class       string    `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	DiskPlan    *DiskPlan `json:",omitempty" yaml:"disk_plan,omitempty" structs:",omitempty"`
	Zone        *Zone     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}
