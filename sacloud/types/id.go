package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ID さくらのクラウド上のリソースのIDを示す
//
// APIリクエスト/レスポンスに文字列/数値が混在するためここで吸収する
type ID int64

// UnmarshalJSON implememts unmarshal from both of JSON number and JSON string
func (i *ID) UnmarshalJSON(b []byte) error {
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	id, err := n.Int64()
	if err != nil {
		return err
	}
	*i = ID(id)
	return nil
}

// String returns the literal text of the number.
func (i ID) String() string {
	return fmt.Sprintf("%d", i)
}

// Int64 returns the number as an int64.
func (i ID) Int64() int64 {
	return int64(i)
}

// Int64ID creates new ID from int64
func Int64ID(id int64) ID {
	return ID(id)
}

// StringID creates new ID from string
func StringID(id string) ID {
	intID, _ := strconv.ParseInt(id, 10, 64)
	return ID(intID)
}
