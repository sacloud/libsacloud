// Copyright 2016-2020 The Libsacloud Authors
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
