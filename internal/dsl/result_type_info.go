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

import "github.com/sacloud/libsacloud/v2/internal/dsl/meta"

// ResultTypeInfo 戻り値の型情報
//
// 主にtraceで利用される
type ResultTypeInfo struct {
	VarName   string    // 変数名
	FieldName string    // トレース時の見出し
	Type      meta.Type // 型
}
