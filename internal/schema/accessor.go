package schema

import "github.com/sacloud/libsacloud/internal/schema/meta"

// Accessor モデルに拡張アクセサを持たせるための定義
type Accessor struct {
	Name             string
	Description      string
	AccessorTypeName string // accessorパッケージ配下のインターフェース名
	ResultType       meta.Type
}
