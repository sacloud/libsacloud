package types

// Action パケットフィルタでのAllow/Denyアクション
type Action string

// Actions パケットフィルタでのAllow/Denyアクション
var Actions = &struct {
	Allow Action
	Deny  Action
}{
	Allow: Action("allow"),
	Deny:  Action("deny"),
}
