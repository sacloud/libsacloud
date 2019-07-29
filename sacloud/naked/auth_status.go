package naked

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// AuthStatus 現在の認証状態
type AuthStatus struct {
	Account            *Account                 // アカウント
	Member             *Member                  // 会員情報
	AuthClass          types.EAuthClass         `json:",omitempty" yaml:"auth_class,omitempty" structs:",omitempty"`          // 認証クラス
	AuthMethod         types.EAuthMethod        `json:",omitempty" yaml:"auth_method,omitempty" structs:",omitempty"`         // 認証方法
	ExternalPermission types.ExternalPermission `json:",omitempty" yaml:"external_permission,omitempty" structs:",omitempty"` // 他サービスへのアクセス権
	IsAPIKey           bool                     `json:",omitempty" yaml:"is_api_key,omitempty" structs:",omitempty"`          // APIキーでのアクセスフラグ
	OperationPenalty   types.EOperationPenalty  `json:",omitempty" yaml:"operation_penalty,omitempty" structs:",omitempty"`   // オペレーションペナルティ
	Permission         types.EPermission        `json:",omitempty" yaml:"permission,omitempty" structs:",omitempty"`          // 権限
}
