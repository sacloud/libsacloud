package search

import "fmt"

// FilterKey 検索条件(Filter)のキー
type FilterKey struct {
	// フィールド名
	Field string
	// 演算子
	Op ComparisonOperator
}

// String Keyの文字列表現
func (k *FilterKey) String() string {
	return fmt.Sprintf("%s%s", k.Field, k.Op)
}

// Key キーの作成
func Key(field string) FilterKey {
	return FilterKey{Field: field}
}

// KeyWithOp 演算子を指定してキーを作成
func KeyWithOp(field string, op ComparisonOperator) FilterKey {
	return FilterKey{Field: field, Op: op}
}
