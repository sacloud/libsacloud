package types

// EServerInstanceStatus サーバーインスタンスステータス
type EServerInstanceStatus string

// ServerInstanceStatuses サーバーインスタンスステータス
var ServerInstanceStatuses = &struct {
	Unknown  EServerInstanceStatus
	Up       EServerInstanceStatus
	Cleaning EServerInstanceStatus
	Down     EServerInstanceStatus
}{
	Unknown:  EServerInstanceStatus(""),
	Up:       EServerInstanceStatus("up"),
	Cleaning: EServerInstanceStatus("cleaning"),
	Down:     EServerInstanceStatus("down"),
}

// IsUp インスタンスが起動しているか判定
func (e EServerInstanceStatus) IsUp() bool {
	return e == ServerInstanceStatuses.Up
}

// IsDown インスタンスがダウンしているか確認
func (e EServerInstanceStatus) IsDown() bool {
	return e == ServerInstanceStatuses.Down
}
