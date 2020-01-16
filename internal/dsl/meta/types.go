// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	// TypeBackupSpanType 自動バックアップの取得間隔種別
	TypeBackupSpanType = Static(types.EBackupSpanType(""))
	// TypeBackupSpanWeekdays 自動バックアップの取得曜日
	TypeBackupSpanWeekdays = Static([]types.EBackupSpanWeekday{})

	// TypeDNSRecordType DNSレコード種別
	TypeDNSRecordType = Static(types.EDNSRecordType(""))

	// TypeSimpleMonitorHealthCheckProtocol シンプル監視 ヘルスチェックプロトコル
	TypeSimpleMonitorHealthCheckProtocol = Static(types.ESimpleMonitorProtocol(""))
)
