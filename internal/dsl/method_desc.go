package dsl

import "github.com/sacloud/libsacloud/v2/internal/dsl/meta"

// MethodDesc モデルにメソッドを持たせるための定義
type MethodDesc struct {
	// Name メソッドの名称
	Name string

	// AccessorFuncName sacloud/accessor配下に定義されている、かつAccessorTypeNameで指定したaccessorを実装するオブジェクトを
	// 第1引数にとる、exportされているfuncの名称
	//
	// 省略した場合はNameが利用される
	AccessorFuncName string

	// Description 拡張アクセサのgodoc用コメント
	Description string

	// AccessorTypeName sacloud/accessorパッケージ配下のインターフェース名
	AccessorTypeName string

	// Arguments メソッド引数 省略可能
	Arguments Arguments

	// ResultTypes 戻り値
	ResultTypes []meta.Type
}
