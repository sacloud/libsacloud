package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Account さくらのクラウド アカウント
type Account struct {
	ID    types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name  string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Class string   `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"` // リソースクラス
	Code  string   `json:",omitempty" yaml:"code,omitempty" structs:",omitempty"`  // アカウントコード
}
