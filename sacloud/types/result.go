package types

import "encoding/json"

// APIResult APIからの戻り値"Success"の別表現
//
// Successにはbool以外にも"Accepted"などの文字列が返ることがある(例:アプライアンス)
// このためAPIResultでUnmarhslJSONを実装してラップする
type APIResult int

const (
	// ResultUnknown 不明
	ResultUnknown APIResult = iota
	// ResultSuccess 成功
	ResultSuccess
	// ResultAccepted 受付成功
	ResultAccepted
	// ResultFailed 失敗
	ResultFailed
)

// UnmarshalJSON bool/string混在型に対応するためのUnmarshalJSON実装
func (r *APIResult) UnmarshalJSON(data []byte) error {
	*r = ResultUnknown

	// try bool
	var b bool
	if err := json.Unmarshal(data, &b); err != nil {

		// try string
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		switch s {
		case "Accepted":
			*r = ResultAccepted
		}

		return nil
	}

	if b {
		*r = ResultSuccess
	} else {
		*r = ResultFailed
	}
	return nil
}
