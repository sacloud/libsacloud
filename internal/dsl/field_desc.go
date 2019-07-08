package dsl

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
)

// FieldDesc フィールド記述
type FieldDesc struct {
	Name            string
	Tags            *FieldTags
	Type            meta.Type
	Description     string            // TODO 現在は未使用
	ExtendAccessors []*ExtendAccessor // 拡張アクセッサ、Get+指定の名前、Set+指定の名前で拡張アクセッサコードが生成される
}

// HasTag タグの定義がなされているか
func (f *FieldDesc) HasTag() bool {
	return f.Tags != nil
}

// TypeName フィールドの型を返す、コード生成で利用される
func (f *FieldDesc) TypeName() string {
	return f.Type.GoTypeSourceCode()
}

// TagString タグのソースコード上での表現を返す
func (f *FieldDesc) TagString() string {
	if !f.HasTag() {
		return ""
	}
	return f.Tags.String()
}

// FieldTags フィールドに付与するタグ
type FieldTags struct {
	// JSON jsonタグ
	JSON string
	// YAML yamlタグ
	YAML string
	// Structs structsタグ
	Structs string
	// MapConv mapconvタグ
	MapConv string
	// Validate validateタグ
	Validate string
}

// String FieldTagsの文字列表現
func (f *FieldTags) String() string {
	var tags []string
	if f.JSON != "" {
		tags = append(tags, fmt.Sprintf(`json:"%s"`, f.JSON))
	}
	if f.YAML != "" {
		tags = append(tags, fmt.Sprintf(`yaml:"%s"`, f.YAML))
	}
	if f.Structs != "" {
		tags = append(tags, fmt.Sprintf(`structs:"%s"`, f.Structs))
	}
	if f.MapConv != "" {
		tags = append(tags, fmt.Sprintf(`mapconv:"%s"`, f.MapConv))
	}
	if f.Validate != "" {
		tags = append(tags, fmt.Sprintf(`validate:"%s"`, f.Validate))
	}
	return strings.Join(tags, " ")
}

// ExtendAccessor 拡張アクセッサ
type ExtendAccessor struct {
	// Name 拡張アクセッサ名、Get+Name,Set+Nameなfuncが生成される
	Name string
	// AvoidGetter trueの場合Getの生成を抑制
	AvoidGetter bool
	// AvoidSetter trueの場合Setの生成を抑制
	AvoidSetter bool
	// Type 引数の型(省略可能)
	Type meta.Type
}

// HasType Typeが指定されているか
func (a *ExtendAccessor) HasType() bool {
	return a.Type != nil
}

// TypeName フィールドの型を返す、コード生成で利用される
func (a *ExtendAccessor) TypeName() string {
	if a.HasType() {
		return a.Type.GoTypeSourceCode()
	}
	return ""
}
