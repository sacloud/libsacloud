package search

import "fmt"

// Key 検索条件(Filter)のキー
type Key struct {
	// フィールド名
	Field string
	// 演算子
	Op ComparisonOperator
}

// String Keyの文字列表現
func (k *Key) String() string {
	return fmt.Sprintf("%s%s", k.Field, k.Op)
}

// NewKey キーの作成
func NewKey(field string) Key {
	return Key{Field: field}
}

// NewKeyWithOp 演算子を指定してキーを作成
func NewKeyWithOp(field string, op ComparisonOperator) Key {
	return Key{Field: field, Op: op}
}
