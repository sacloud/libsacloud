package types

// EBackupSpanType 自動バックアップのバックアップ間隔種別
type EBackupSpanType string

// BackupSpanTypes 自動バックアップのバックアップ間隔種別
var BackupSpanTypes = struct {
	// Unknown 不明
	Unknown EBackupSpanType
	// Weekdays 曜日
	Weekdays EBackupSpanType
}{
	Unknown:  EBackupSpanType(""),
	Weekdays: EBackupSpanType("weekdays"),
}
