package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// DiskPlan ディスクプラン
type DiskPlan struct {
	ID           types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	StorageClass string   `json:",omitempty" yaml:"storage_class,omitempty" structs:",omitempty"`
}
