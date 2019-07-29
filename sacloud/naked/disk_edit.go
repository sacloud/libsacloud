package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// DiskEdit ディスクの修正パラメータ
type DiskEdit struct {
	Password            string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パスワード
	SSHKey              *DiskEditSSHKey   `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // 公開鍵(単体)
	SSHKeys             []*DiskEditSSHKey `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // 公開鍵(複数)
	DisablePWAuth       bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パスワード認証無効化フラグ
	EnableDHCP          bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // DHCPの有効化
	ChangePartitionUUID bool              `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // パーティションのUUID変更
	HostName            string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // ホスト名
	Notes               []DiskEditNote    `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // スタートアップスクリプト
	UserIPAddress       string            `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // IPアドレス
	UserSubnet          *UserSubnet       `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // デフォルトルート/サブネットマスク長
}

// DiskEditSSHKey ディスク修正時のSSHキー
type DiskEditSSHKey struct {
	ID        types.ID `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	PublicKey string   `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// DiskEditNote ディスクの修正で指定するスタートアップスクリプト
type DiskEditNote struct {
	ID        types.ID               `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Variables map[string]interface{} `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}
