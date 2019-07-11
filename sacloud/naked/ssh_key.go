package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// SSHKey 公開鍵
type SSHKey struct {
	ID          types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string     `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string     `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	CreatedAt   *time.Time `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	PublicKey   string     `json:",omitempty" yaml:"public_key,omitempty" structs:",omitempty"`  // 公開鍵
	PrivateKey  string     `json:",omitempty" yaml:"public_key,omitempty" structs:",omitempty"`  // 秘密鍵、API側での鍵生成時のみセットされる
	Fingerprint string     `json:",omitempty" yaml:"fingerprint,omitempty" structs:",omitempty"` // フィンガープリント

	GenerateFormat string `json:",omitempty" yaml:"generate_format,omitempty" structs:",omitempty"` // 鍵生成時のみ利用(openssh固定)
	PassPhrase     string `json:",omitempty" yaml:"pass_phrase,omitempty" structs:",omitempty"`     // 鍵生成時のみ利用
}
