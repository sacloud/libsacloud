package types

// PacketFilterAction パケットフィルタでのAllow/Denyアクション
type PacketFilterAction string

// PacketFilterActions パケットフィルタでのAllow/Denyアクション
var PacketFilterActions = &struct {
	Allow PacketFilterAction
	Deny  PacketFilterAction
}{
	Allow: PacketFilterAction("allow"),
	Deny:  PacketFilterAction("deny"),
}
