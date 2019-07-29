package search

import "encoding/json"

// SortOrder ソート順
type SortOrder int

const (
	// SortAsc 昇順(デフォルト)
	SortAsc SortOrder = iota
	// SortDesc 降順
	SortDesc
)

// SortKeys ソート順指定
type SortKeys []SortKey

// SortKey ソート順指定対象のフィールド名
type SortKey struct {
	Key   string
	Order SortOrder
}

// MarshalJSON キーの文字列表現
func (k SortKey) MarshalJSON() ([]byte, error) {
	s := k.Key
	if k.Order == SortDesc {
		s = "-" + k.Key
	}
	return json.Marshal(s)
}
