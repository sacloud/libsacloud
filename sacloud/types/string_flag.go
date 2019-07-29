package types

import (
	"strings"
)

var trueStrings = []string{"true", "on", "1"}

var (
	// StringTrue true値
	StringTrue = StringFlag(true)
	// StringFalse false値
	StringFalse = StringFlag(false)
)

// StringFlag bool型のラッパー、文字列(true/false/on/off/1/0)などをbool値として扱う
//
// - 大文字/小文字の区別はしない
// - 空文字だった場合はfalse
// - 小文字にした場合に次のいずれかにマッチしない場合はfalse [ true / on / 1 ]
type StringFlag bool

// String StringFlagの文字列表現
func (f *StringFlag) String() string {
	if f.Bool() {
		return "True"
	}
	return "False"
}

// Bool StringFlagのbool表現
func (f *StringFlag) Bool() bool {
	return f != nil && bool(*f)
}

// MarshalJSON 文字列でのJSON出力に対応するためのMarshalJSON実装
func (f *StringFlag) MarshalJSON() ([]byte, error) {
	if f != nil && bool(*f) {
		return []byte(`"True"`), nil
	}
	return []byte(`"False"`), nil
}

// UnmarshalJSON 文字列に対応するためのUnmarshalJSON実装
func (f *StringFlag) UnmarshalJSON(b []byte) error {
	s := strings.Replace(strings.ToLower(string(b)), `"`, ``, -1)
	res := false
	for _, strTrue := range trueStrings {
		if s == strTrue {
			res = true
			break
		}
	}
	*f = StringFlag(res)
	return nil
}
