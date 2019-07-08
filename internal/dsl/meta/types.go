package meta

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var (
	// TypeID ID型
	TypeID = Static(types.ID(0))
	// TypeFlag フラグ型(boolean)
	TypeFlag = Static(false)
	// TypeStringFlag 文字列フラグ型
	TypeStringFlag = Static(types.StringFalse)
	// TypeStringNumber 文字列数値型
	TypeStringNumber = Static(types.StringNumber(0))
	// TypeString 文字列
	TypeString = Static("")
	// TypeStringSlice 文字列スライス
	TypeStringSlice = Static([]string{})
	// TypeInt int型
	TypeInt = Static(int(0))
	// TypeIntSlice intスライス
	TypeIntSlice = Static([]int{})
	// TypeInt64 int64型
	TypeInt64 = Static(int64(0))
	// TypeInt64Slice int64スライス
	TypeInt64Slice = Static([]int64{})
	// TypeFloat64 flat64型
	TypeFloat64 = Static(float64(0))
	// TypeTime Time型
	TypeTime = Static(time.Time{})

	// TypeAvailability 有効状態
	TypeAvailability = Static(types.EAvailability(""))
	// TypeCommitment サーバプランCPUコミットメント
	TypeCommitment = Static(types.ECommitment(""))
	// TypeDiskConnection ディスク接続方法
	TypeDiskConnection = Static(types.EDiskConnection(""))
	// TypeInstanceStatus インスタンスステータス
	TypeInstanceStatus = Static(types.EServerInstanceStatus(""))
	// TypeInterfaceDriver インターフェースドライバ
	TypeInterfaceDriver = Static(types.EInterfaceDriver(""))
	// TypePlanGeneration プラン世代
	TypePlanGeneration = Static(types.EPlanGeneration(0))
	// TypeProtocol プロトコル
	TypeProtocol = Static(types.Protocol(""))
	// TypeScope スコープ
	TypeScope = Static(types.EScope(""))

	// TypePacketFilterNetwork パケットフィルタルールでの送信元アドレス/範囲
	TypePacketFilterNetwork = Static(types.PacketFilterNetwork(""))
	// TypePacketFilterPort パケットフィルタルールでのポート
	TypePacketFilterPort = Static(types.PacketFilterPort(""))
	// TypeVPCFirewallNetwork VPCルータのファイアウォールルールでの送信元アドレス/範囲
	TypeVPCFirewallNetwork = Static(types.VPCFirewallNetwork(""))
	// TypeVPCFirewallPort VPCルータのファイアウォールルールでのポート
	TypeVPCFirewallPort = Static(types.VPCFirewallPort(""))
	// TypeAction パケットフィルタルールでのallow/deny動作
	TypeAction = Static(types.Action(""))
)
