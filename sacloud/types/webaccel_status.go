package types

// EWebAccelStatus ウェブアクセラレータ ステータス
type EWebAccelStatus string

// WebAccelStatus ウェブアクセラレータ ステータス
var WebAccelStatus = struct {
	// Enabled 有効
	Enabled EWebAccelStatus
	// Disabled 無効
	Disabled EWebAccelStatus
}{
	Enabled:  EWebAccelStatus("enabled"),
	Disabled: EWebAccelStatus("disabled"),
}
