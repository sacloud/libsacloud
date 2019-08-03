package types

import "sort"

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

// BackupSpanWeekdaysOrder バックアップ取得曜日の並び順(日曜開始)
var BackupSpanWeekdaysOrder = map[EBackupSpanWeekday]int{
	BackupSpanWeekdays.Sunday:    0,
	BackupSpanWeekdays.Monday:    1,
	BackupSpanWeekdays.Tuesday:   2,
	BackupSpanWeekdays.Wednesday: 3,
	BackupSpanWeekdays.Thursday:  4,
	BackupSpanWeekdays.Friday:    5,
	BackupSpanWeekdays.Saturday:  6,
}

// SortBackupSpanWeekdays バックアップ取得曜日のソート(日曜開始)
func SortBackupSpanWeekdays(weekdays []EBackupSpanWeekday) {
	sort.Slice(weekdays, func(i, j int) bool {
		return BackupSpanWeekdaysOrder[weekdays[i]] < BackupSpanWeekdaysOrder[weekdays[j]]
	})
}
