package types

import "encoding/json"

// Tags タグ
type Tags []string

// MarshalJSON タグを空にする場合への対応
func (t *Tags) MarshalJSON() ([]byte, error) {
	if *t == nil {
		*t = make([]string, 0)
	}

	type alias Tags
	tmp := alias(*t)
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
	return nil
}
