package naked

import (
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

// Subnet サブネット
//
// TODO 後で
type Subnet struct {
	ID             types.ID   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	ServiceClass   string     `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	CreatedAt      *time.Time `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	DefaultRoute   string     `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkAddress string     `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	NetworkMaskLen int        `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Internet       struct {
		BandWidthMbps int `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

/*
type Subnet struct {
	*Resource        // ID
	propServiceClass // サービスクラス
	propCreatedAt    // 作成日時

	DefaultRoute   string       `json:",omitempty"` // デフォルトルート
	IPAddresses    []*IPAddress `json:",omitempty"` // IPv4アドレス範囲
	NetworkAddress string       `json:",omitempty"` // ネットワークアドレス
	NetworkMaskLen int          `json:",omitempty"` // ネットワークマスク長
	ServiceID      int64        `json:",omitempty"` // サービスID
	StaticRoute    string       `json:",omitempty"` // スタティックルート
	NextHop        string       `json:",omitempty"` // ネクストホップ
	Switch         *Switch      `json:",omitempty"` // スイッチ
	Internet       *Internet    `json:",omitempty"` // ルーター
}
*/
