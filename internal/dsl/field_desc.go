// Copyright 2016-2019 The Libsacloud Authors
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
	IgnorePatch  bool   // trueの場合patch時の値のクリアを行えないようにする
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

// IsPatchEmptyParam クリアパラメータか判定
func (f *FieldDesc) IsPatchEmptyParam() bool {
	return strings.HasPrefix(f.Name, "PatchEmptyTo")
}

// IsRequired 必須項目であるかを判定
func (f *FieldDesc) IsRequired() bool {
	if f.Tags == nil {
		return false
	}
	return strings.Contains(f.Tags.Validate, "required")
}

// IsNeedPatchEmpty Patch時の値のクリアが必要な項目かを判定
func (f *FieldDesc) IsNeedPatchEmpty() bool {
	return !f.IsPatchEmptyParam() && !f.IsRequired() && !f.IgnorePatch
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
