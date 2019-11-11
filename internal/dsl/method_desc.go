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

// MethodDesc モデルにメソッドを持たせるための定義
type MethodDesc struct {
	// Name メソッドの名称
	Name string

	// AccessorFuncName sacloud/accessor配下に定義されている、accessorを実装するオブジェクトを
	// 第1引数にとる、exportされているfuncの名称
	//
	// 省略した場合はNameが利用される
	AccessorFuncName string

	// Description 拡張アクセサのgodoc用コメント
	Description string

	// Arguments メソッド引数 省略可能
	Arguments Arguments

	// ResultTypes 戻り値
	ResultTypes []meta.Type
}
