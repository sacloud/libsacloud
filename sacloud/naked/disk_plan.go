package naked

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// DiskPlan ディスクプラン
type DiskPlan struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	StorageClass string              `json:",omitempty" yaml:"storage_class,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Size         []*DiskPlanSizeInfo `json:",omitempty" yaml:"size,omitempty" structs:",omitempty"`
}

// DiskPlanSizeInfo ディスクプランに含まれる利用可能なサイズ情報
type DiskPlanSizeInfo struct {
	Availability  types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	DisplaySize   int                 `json:",omitempty" yaml:"display_size,omitempty" structs:",omitempty"`
	DisplaySuffix string              `json:",omitempty" yaml:"display_suffix,omitempty" structs:",omitempty"`
	ServiceClass  string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	SizeMB        int                 `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
}
