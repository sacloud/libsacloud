package meta

import (
	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

var (
	// TypeID ID型
	TypeID = Static(int64(0))
	// TypeString 文字列
	TypeString = Static("")
	// TypeStringSlice 文字列スライス
	TypeStringSlice = Static([]string{})
	// TypeAvailability 有効状態
	TypeAvailability = Static(enums.EAvailability(""))
	// TypeInstanceStatus インスタンスステータス
	TypeInstanceStatus = Static(enums.EServerInstanceStatus(""))
	// TypeScope スコープ
	TypeScope = Static(enums.EScope(""))
)
