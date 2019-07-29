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

// EBackupSpanWeekday バックアップ取得曜日
type EBackupSpanWeekday string

// BackupSpanWeekdays バックアップ取得曜日
var BackupSpanWeekdays = struct {
	Monday    EBackupSpanWeekday
	Tuesday   EBackupSpanWeekday
	Wednesday EBackupSpanWeekday
	Thursday  EBackupSpanWeekday
	Friday    EBackupSpanWeekday
	Saturday  EBackupSpanWeekday
	Sunday    EBackupSpanWeekday
}{
	Monday:    EBackupSpanWeekday("mon"),
	Tuesday:   EBackupSpanWeekday("tue"),
	Wednesday: EBackupSpanWeekday("wed"),
	Thursday:  EBackupSpanWeekday("thu"),
	Friday:    EBackupSpanWeekday("fri"),
	Saturday:  EBackupSpanWeekday("sat"),
	Sunday:    EBackupSpanWeekday("sun"),
}
