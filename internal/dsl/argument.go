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

	// ArgumentZone 引数でのゾーンを示すValue
	ArgumentZone = &Argument{
		Name: "zone",
		Type: meta.TypeString,
	}
)

// Argument 引数の型情報
type Argument struct {
	Name       string    // パラメータ名、引数名に利用される
	Type       meta.Type // パラメータの型情報
	MapConvTag string
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

// MappableArgument 引数定義の追加
func MappableArgument(name string, model *Model, destField string) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: fmt.Sprintf("%s,recursive", destField),
	}
}

// PassthroughModelArgument 引数定義の追加、ペイロードの定義も同時に行われる
func PassthroughModelArgument(name string, model *Model) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: ",squash",
	}
}
