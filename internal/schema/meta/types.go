package meta

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/enums"
)

var (
	// TypeID ID型
	TypeID = Static(int64(0))
	// TypeFlag フラグ型(boolean)
	TypeFlag = Static(false)
	// TypeString 文字列
	TypeString = Static("")
	// TypeStringSlice 文字列スライス
	TypeStringSlice = Static([]string{})
	// TypeInt int型
	TypeInt = Static(int(0))
	// TypeIntSlice intスライス
	TypeIntSlice = Static([]int{})
	// TypeTime Time型
	TypeTime = Static(time.Time{})
	// TypeAvailability 有効状態
	TypeAvailability = Static(enums.EAvailability(""))
	// TypeInstanceStatus インスタンスステータス
	TypeInstanceStatus = Static(enums.EServerInstanceStatus(""))
	// TypeScope スコープ
	TypeScope = Static(enums.EScope(""))
)
