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

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
)

// ResultType Operationからの戻り値の型情報
type ResultType struct {
	resourceName string
	operation    *Operation
	results      Results
}

// Type モデルの型情報
func (r *ResultType) Type() meta.Type {
	return r
}

// GoType 型名
func (r *ResultType) GoType() string {
	return fmt.Sprintf("%s%sResult", r.resourceName, r.operation.Name)
}

// GoPkg パッケージ名
func (r *ResultType) GoPkg() string {
	if IsOutOfSacloudPackage {
		return "sacloud"
	}
	return ""
}

// GoImportPath インポートパス
func (r *ResultType) GoImportPath() string {
	if IsOutOfSacloudPackage {
		return "github.com/sacloud/libsacloud/v2/sacloud"
	}
	return ""
}

// GoTypeSourceCode ソースコードでの型表現
func (r *ResultType) GoTypeSourceCode() string {
	name := r.GoType()
	prefix := ""
	if IsOutOfSacloudPackage {
		prefix = "sacloud."
	}
	return fmt.Sprintf("*%s%s", prefix, name)
}

// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
func (r *ResultType) ZeroInitializeSourceCode() string {
	name := r.GoType()
	if IsOutOfSacloudPackage {
		name = "sacloud." + name
	}
	return fmt.Sprintf("&%s{}", name)
}

// ZeroValueSourceCode 型に応じたzero値コード
func (r *ResultType) ZeroValueSourceCode() string {
	return "nil"
}

// ToPtrType ポインタ型への変換
func (r *ResultType) ToPtrType() meta.Type {
	return nil // not use
}
