package types

import (
	"encoding/json"
	"sort"
)

// Tags タグ
type Tags []string

// Sort 昇順でソートする
func (t *Tags) Sort() {
	sort.Strings([]string(*t))
}

// MarshalJSON タグを空にする場合への対応
func (t Tags) MarshalJSON() ([]byte, error) {
	t.Sort()
	type alias Tags
	tmp := alias(t)
	return json.Marshal(tmp)
}

// UnmarshalJSON タグを空にする場合への対応
func (t *Tags) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		*t = make([]string, 0)
		return nil
	}
	type alias Tags
	tmp := alias(*t)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = Tags(tmp)
	t.Sort()
	return nil
}
