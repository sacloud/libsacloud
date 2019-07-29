package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// HealthCheck ヘルスチェック
type HealthCheck struct {
	Protocol types.Protocol     `json:",omitempty" yaml:"protocol,omitempty" structs:""` // プロトコル
	Host     string             `json:",omitempty" yaml:"host,omitempty" structs:""`     // 対象ホスト
	Path     string             `json:",omitempty" yaml:"path,omitempty" structs:""`     // HTTP/HTTPSの場合のリクエストパス
	Status   types.StringNumber `json:",omitempty" yaml:"status,omitempty" structs:""`   // 期待するステータスコード
	Port     types.StringNumber `json:",omitempty" yaml:"port,omitempty" structs:""`     // ポート番号
}
