package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Archive アーカイブ
type Archive struct {
	ID              types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name            string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description     string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags            []string            `json:"" yaml:"tags"`
	Icon            *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt       *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt      *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability    types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	DisplayOrder    int                 `json:",omitempty" yaml:"display_order,omitempty" structs:",omitempty"`
	ServiceClass    string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	SizeMB          int                 `json:",omitempty" yaml:"size_mb,omitempty" structs:",omitempty"`
	MigratedMB      int                 `json:",omitempty" yaml:"migrated_mb,omitempty" structs:",omitempty"`
	JobStatus       *MigrationJobStatus `json:",omitempty" yaml:"job_status,omitempty" structs:",omitempty"`
	Plan            *DiskPlan           `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	SourceDisk      *Disk               `json:",omitempty" yaml:"source_disk,omitempty" structs:",omitempty"`
	SourceArchive   *Archive            `json:",omitempty" yaml:"source_archive,omitempty" structs:",omitempty"`
	BundleInfo      *BundleInfo         `json:",omitempty" yaml:"bundle_info,omitempty" structs:",omitempty"`
	Storage         *Storage            `json:",omitempty" yaml:"storage,omitempty" structs:",omitempty"`
	Scope           types.EScope        `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
	OriginalArchive *OriginalArchive    `json:",omitempty" yaml:"original_archive,omitempty" structs:",omitempty"`
	SourceInfo      *SourceArchive      `json:",omitempty" yaml:"source_info,omitempty" structs:",omitempty"`
}

// SourceArchive 他ゾーンから転送したアーカイブの情報
type SourceArchive struct {
	ArchiveUnderZone *SourceArchiveInfo `json:",omitempty" yaml:"archive_under_zone,omitempty" structs:",omitempty"`
}

// SourceArchiveInfo 他ゾーンから転送したアーカイブの情報
type SourceArchiveInfo struct {
	ID      types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Account *struct {
		ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	} `json:",omitempty" yaml:"account,omitempty" structs:",omitempty"`
	Zone *struct {
		ID   types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
		Name string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	} `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}

// OriginalArchive オリジナルアーカイブ
type OriginalArchive struct {
	ID types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
}
