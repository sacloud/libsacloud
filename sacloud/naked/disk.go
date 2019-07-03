package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Disk ディスク
type Disk struct {
	ID              types.ID              `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name            string                `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description     string                `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags            []string              `json:"" yaml:"tags"`
	Icon            *Icon                 `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt       *time.Time            `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt      *time.Time            `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability    types.EAvailability   `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass    string                `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	SizeMB          int                   `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
	MigratedMB      int                   `json:",omitempty" yaml:"migrated_mb,omitempty" structs:",omitempty"`
	Connection      types.EDiskConnection `json:",omitempty" yaml:"connection,omitempty" structs:",omitempty"`
	ConnectionOrder int                   `json:",omitempty" yaml:"connection_order,omitempty" structs:",omitempty"`
	ReinstallCount  int                   `json:",omitempty" yaml:"reinstall_count,omitempty" structs:",omitempty"`
	JobStatus       *MigrationJobStatus   `json:",omitempty" yaml:"job_status,omitempty" structs:",omitempty"`
	Plan            *DiskPlan             `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	SourceDisk      *Disk                 `json:",omitempty" yaml:"source_disk,omitempty" structs:",omitempty"`
	SourceArchive   *Archive              `json:",omitempty" yaml:"source_archive,omitempty" structs:",omitempty"`
	BundleInfo      *BundleInfo           `json:",omitempty" yaml:"bundle_info,omitempty" structs:",omitempty"`
	Storage         *Storage              `json:",omitempty" yaml:"storage,omitempty" structs:",omitempty"`
	Server          *Server               `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`
}

// MigrationJobStatus マイグレーションジョブステータス
type MigrationJobStatus struct {
	Status string    `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"` // ステータス
	Delays *struct { // Delays
		Start *struct { // 開始
			Max int `json:",omitempty" yaml:"max,omitempty" structs:",omitempty"` // 最大
			Min int `json:",omitempty" yaml:"min,omitempty" structs:",omitempty"` // 最小
		} `json:",omitempty" yaml:"start,omitempty" structs:",omitempty"`

		Finish *struct { // 終了
			Max int `json:",omitempty" yaml:"max,omitempty" structs:",omitempty"` // 最大
			Min int `json:",omitempty" yaml:"min,omitempty" structs:",omitempty"` // 最小
		} `json:",omitempty" yaml:"finish,omitempty" structs:",omitempty"`
	}
}
