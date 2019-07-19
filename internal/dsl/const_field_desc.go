package dsl

import "github.com/sacloud/libsacloud/v2/internal/dsl/meta"

// ConstFieldDesc 固定値フィールド記述
type ConstFieldDesc struct {
	Name        string
	Tags        *FieldTags
	Type        meta.Type
	Description string // TODO 現在は未使用
	Value       string
}

// HasTag タグの定義がなされているか
func (f *ConstFieldDesc) HasTag() bool {
	return f.Tags != nil
}

// TypeName フィールドの型を返す、コード生成で利用される
func (f *ConstFieldDesc) TypeName() string {
	return f.Type.GoTypeSourceCode()
}

// TagString タグのソースコード上での表現を返す
func (f *ConstFieldDesc) TagString() string {
	if !f.HasTag() {
		return ""
	}
	return f.Tags.String()
}
