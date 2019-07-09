package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// AutoBackup 自動バックアップ
type AutoBackup struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *AutoBackupSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"setting_hash,omitempty" structs:",omitempty"`
	Status       *AutoBackupStatus   `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// AutoBackupSettings 自動バックアップ設定
type AutoBackupSettings struct {
	Autobackup *AutoBackupSetting `json:",omitempty" yaml:"autobackup,omitempty" structs:",omitempty"` // HACK: 注: API側がキャメルケースになっていない
}

// AutoBackupSetting 自動バックアップ設定
type AutoBackupSetting struct {
	BackupSpanType          types.EBackupSpanType      `json:",omitempty" yaml:"backup_span_type,omitempty" structs:",omitempty"`
	BackupSpanWeekdays      []types.EBackupSpanWeekday `json:",omitempty" yaml:"backup_span_weekdays,omitempty" structs:",omitempty"`
	MaximumNumberOfArchives int                        `json:",omitempty" yaml:"maximum_number_of_archives,omitempty" structs:",omitempty"`
}

// AutoBackupStatus 自動バックアップステータス
type AutoBackupStatus struct {
	DiskID    types.ID `json:"DiskId,omitempty" yaml:"disk_id,omitempty" structs:",omitempty"`
	AccountID types.ID `json:"AccountId,omitempty" yaml:"account_id,omitempty" structs:",omitempty"`
	ZoneID    types.ID `json:"ZoneId,omitempty" yaml:"zone_id,omitempty" structs:",omitempty"`
	ZoneName  string   `json:",omitempty" yaml:"zone_name,omitempty" structs:",omitempty"`
}
