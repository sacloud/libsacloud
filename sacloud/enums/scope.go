package enums

// EScope スコープ
type EScope string

// Scopes スコープ
var Scopes = &struct {
	Shared EScope // 共有
	User   EScope // ユーザー
}{
	Shared: EScope("shared"),
	User:   EScope("user"),
}
