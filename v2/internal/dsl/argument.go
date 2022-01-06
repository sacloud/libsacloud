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

// Arguments Operationへの引数リスト
type Arguments []*Argument

var (
	// ArgumentID 引数でのIDを示すValue
	ArgumentID = &Argument{
		Name: "id",
		Type: meta.TypeID,
	}
)

// Argument 引数の型情報
type Argument struct {
	Name            string    // パラメータ名、引数名に利用される
	Type            meta.Type // パラメータの型情報
	PathFormatAlias string    // リクエストパス組み立て時に利用するパラメータ名のエイリアス 省略時はNameとなる
	MapConvTag      string
}

// ImportStatements コード生成時に利用するimport文を生成する
func (a *Argument) ImportStatements() []string {
	if a.Type.GoPkg() == "" {
		return []string{}
	}
	return wrapByDoubleQuote(a.Type.GoImportPath())
}

// PackageName インポートパスからパッケージ名を取得する
func (a *Argument) PackageName() string {
	return a.Type.GoPkg()
}

// ArgName 引数の変数名、コード生成で利用される
func (a *Argument) ArgName() string {
	return a.Name
}

// TypeName 型名の文字列表現、コード生成で利用される
func (a *Argument) TypeName() string {
	return a.Type.GoTypeSourceCode()
}

// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
func (a *Argument) ZeroInitializer() string {
	return a.Type.ZeroInitializeSourceCode()
}

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (a *Argument) ZeroValueOnSource() string {
	return a.Type.ZeroValueSourceCode()
}

// MapConvTagSrc コード上でのmapconvタグの文字列表現。コード生成時に利用する
func (a *Argument) MapConvTagSrc() string {
	if a.MapConvTag == "" {
		return ""
	}
	return fmt.Sprintf("`mapconv:\"%s\"`", a.MapConvTag)
}

// PathFormatName リクエストパス組み立て時に利用するパラメータ名の
func (a *Argument) PathFormatName() string {
	if a.PathFormatAlias != "" {
		return a.PathFormatAlias
	}
	return a.Name
}

// MappableArgument 引数定義の追加
func MappableArgument(name string, model *Model, destField string) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: fmt.Sprintf("%s,recursive", destField),
	}
}

// PassthroughModelArgument 引数定義の追加
func PassthroughModelArgument(name string, model *Model) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: ",squash",
	}
}
