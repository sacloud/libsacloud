package dsl

import "github.com/sacloud/libsacloud/v2/internal/dsl/meta"

// ResultTypeInfo 戻り値の型情報
//
// 主にtraceで利用される
type ResultTypeInfo struct {
	VarName   string    // 変数名
	FieldName string    // トレース時の見出し
	Type      meta.Type // 型
}
