// Copyright 2016-2022 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dsl

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
)

// FieldDesc フィールド記述
type FieldDesc struct {
	Name         string
	Tags         *FieldTags
	Type         meta.Type
	Description  string // TODO 現在は未使用
	Methods      []*MethodDesc
	DefaultValue string // デフォルト値、コード生成時にソースコードに直接転記される
}

// CloneFieldWithTags 指定のフィールドをクローンし指定のタグを設定して返す
func CloneFieldWithTags(f *FieldDesc, tags ...*FieldTags) *FieldDesc {
	var tag *FieldTags
	for _, t := range tags {
		if t != nil {
			tag = t
			break
		}
	}
	if tag == nil {
		tag = &FieldTags{}
	}
	return &FieldDesc{
		Name:         f.Name,
		Tags:         tag,
		Type:         f.Type,
		Description:  f.Description,
		Methods:      f.Methods,
		DefaultValue: f.DefaultValue,
	}
}

// CloneUpdateFieldWithTags 指定のフィールドのポインタ型をクローンし指定のタグを設定して返す
func CloneUpdateFieldWithTags(f *FieldDesc, tags ...*FieldTags) *FieldDesc {
	return CloneFieldWithTags(f.ToPtrType(), tags...)
}

// ToPtrType ポインタを受け取る型に変換したFieldDescを返す
func (f *FieldDesc) ToPtrType() *FieldDesc {
	return &FieldDesc{
		Name:         f.Name,
		Tags:         f.Tags,
		Type:         f.Type.ToPtrType(),
		Description:  f.Description,
		Methods:      f.Methods,
		DefaultValue: f.DefaultValue,
	}
}

// HasTag タグの定義がなされているか
func (f *FieldDesc) HasTag() bool {
	return f.Tags != nil && !f.Tags.Empty()
}

// SetTags タグの設定
func (f *FieldDesc) SetTags(t *FieldTags) {
	f.Tags = t
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

// IsRequired 必須項目であるかを判定
func (f *FieldDesc) IsRequired() bool {
	if f.Tags == nil {
		return false
	}
	return strings.Contains(f.Tags.Validate, "required")
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
	// Request requestタグ(service向けmapconvタグ)
	Request string
}

func (f *FieldTags) Empty() bool {
	return f.JSON == "" && f.YAML == "" && f.Structs == "" && f.MapConv == "" && f.Validate == "" && f.Request == ""
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
	if f.Request != "" {
		tags = append(tags, fmt.Sprintf(`request:"%s"`, f.Request))
	}
	if f.Validate != "" {
		tags = append(tags, fmt.Sprintf(`validate:"%s"`, f.Validate))
	}
	return strings.Join(tags, " ")
}

func ValidateRequiredTag() *FieldTags {
	return ValidateTag("required")
}

func ValidateTag(t string, args ...interface{}) *FieldTags {
	tag := t
	if len(args) > 0 {
		tag = fmt.Sprintf(t, args...)
	}
	return &FieldTags{Validate: tag}
}

func RequestOmitEmptyTag() *FieldTags {
	return RequestTag(",omitempty")
}

func RequestTag(t string, args ...interface{}) *FieldTags {
	tag := t
	if len(args) > 0 {
		tag = fmt.Sprintf(t, args...)
	}
	return &FieldTags{Request: tag}
}
