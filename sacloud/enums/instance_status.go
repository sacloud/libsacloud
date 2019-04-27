package enums

// EServerInstanceStatus サーバーインスタンスステータス
type EServerInstanceStatus string

// ServerInstanceStatuses サーバーインスタンスステータス
var ServerInstanceStatuses = &struct {
	Up   EServerInstanceStatus
	Down EServerInstanceStatus
}{
	Up:   EServerInstanceStatus("up"),
	Down: EServerInstanceStatus("down"),
}

// IsUp インスタンスが起動しているか判定
func (e EServerInstanceStatus) IsUp() bool {
	return e == ServerInstanceStatuses.Up
}

// IsDown インスタンスがダウンしているか確認
func (e EServerInstanceStatus) IsDown() bool {
	return e == ServerInstanceStatuses.Down
}
