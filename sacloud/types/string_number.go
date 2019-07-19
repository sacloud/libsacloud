package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringNumber 数値型を文字列で表す型
type StringNumber int64

// MarshalJSON implements json.Marshaler
func (n *StringNumber) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, n.String())), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (n *StringNumber) UnmarshalJSON(b []byte) error {
	if string(b) == `""` {
		*n = StringNumber(0)
		return nil
	}

	var num json.Number
	if err := json.Unmarshal(b, &num); err != nil {
		return err
	}
	number, err := num.Int64()
	if err != nil {
		return err
	}
	*n = StringNumber(number)
	return nil
}

// String returns the literal text of the number.
func (n StringNumber) String() string {
	if n.Int64() == 0 {
		return ""
	}
	return fmt.Sprintf("%d", n)
}

// Int returns the number as an int.
func (n StringNumber) Int() int {
	return int(n)
}

// Int64 returns the number as an int64.
func (n StringNumber) Int64() int64 {
	return int64(n)
}

// ParseStringNumber 文字列からStringNumberへの変換
func ParseStringNumber(s string) (StringNumber, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return StringNumber(0), err
	}
	return StringNumber(n), nil

}
