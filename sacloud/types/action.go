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

// IsAllow Allowであるか判定
func (a Action) IsAllow() bool {
	return a == Actions.Allow
}

// IsDeny Denyであるか判定
func (a Action) IsDeny() bool {
	return a == Actions.Deny
}
