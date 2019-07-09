package types

// ESimpleMonitorHealth シンプル監視ステータス
type ESimpleMonitorHealth string

// SimpleMonitorHealth シンプル監視ステータス
var SimpleMonitorHealth = struct {
	// Up アップ
	Up ESimpleMonitorHealth
	// Down ダウン
	Down ESimpleMonitorHealth
}{
	Up:   ESimpleMonitorHealth("UP"),
	Down: ESimpleMonitorHealth("DOWN"),
}

// IsUp アップ
func (e ESimpleMonitorHealth) IsUp() bool {
	return e == SimpleMonitorHealth.Up
}

// IsDown ダウン
func (e ESimpleMonitorHealth) IsDown() bool {
	return e == SimpleMonitorHealth.Down
}
