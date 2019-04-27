package schema

import (
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
)

// Arguments Operationへの引数リスト
type Arguments []Argument

// MapDestinationDeciders Argumentsのうち、MapDestDeciderであるもののリストを返す
func (a Arguments) MapDestinationDeciders() []MapDestinationDecider {
	var deciders = make([]MapDestinationDecider, 0)
	for _, arg := range a {
		if v, ok := arg.(MapDestinationDecider); ok {
			deciders = append(deciders, v)
		}
	}
	return deciders
}

// Argument Operationへの引数を表す
type Argument interface {
	// ImportStatements コード生成時に利用するimport文を生成する
	ImportStatements() []string
	// PackageName インポートパスからパッケージ名を取得する
	PackageName() string
	// ArgName 引数の変数名、コード生成で利用される
	ArgName() string
	// TypeName 型名の文字列表現、コード生成で利用される
	TypeName() string
	// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
	ZeroInitializer() string
	// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
	ZeroValueOnSource() string
}

// MapDestinationDecider マッピング先のフィールド名を決定する
type MapDestinationDecider interface {
	Argument
	// DestinationFieldName マッピング先となるフィールド名を取得
	DestinationFieldName() string
	DestinationModel() *Model
}

var (
	// ArgumentID 引数でのIDを示すValue
	ArgumentID Argument = &SimpleArgument{
		Name: "id",
		Type: meta.TypeID,
	}

	// ArgumentZone 引数でのゾーンを示すValue
	ArgumentZone Argument = &SimpleArgument{
		Name: "zone",
		Type: meta.TypeString,
	}
)

// SimpleArgument 引数の型情報
type SimpleArgument struct {
	Name string    // パラメータ名、引数名に利用される
	Type meta.Type // パラメータの型情報
}

// ImportStatements コード生成時に利用するimport文を生成する
func (a *SimpleArgument) ImportStatements() []string {
	if a.Type.GoPkg() == "" {
		return []string{}
	}
	return wrapByDoubleQuote(a.Type.GoImportPath())
}

// PackageName インポートパスからパッケージ名を取得する
func (a *SimpleArgument) PackageName() string {
	return a.Type.GoPkg()
}

// ArgName 引数の変数名、コード生成で利用される
func (a *SimpleArgument) ArgName() string {
	return a.Name
}

// TypeName 型名の文字列表現、コード生成で利用される
func (a *SimpleArgument) TypeName() string {
	return a.Type.GoTypeSourceCode()
}

// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
func (a *SimpleArgument) ZeroInitializer() string {
	return a.Type.ZeroInitializeSourceCode()
}

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (a *SimpleArgument) ZeroValueOnSource() string {
	return a.Type.ZeroValueSourceCode()
}

// MappableArgument 引数の型情報、APIパラメータへのマッピングが行われる
type MappableArgument struct {
	Name        string // パラメータ名、引数名に利用される
	Destination string // エンベロープでの宛先ペイロード名
	Model       *Model // パラメータの型情報
}

// ImportStatements コード生成時に利用するimport文を生成する
func (a *MappableArgument) ImportStatements() []string {
	p := a.Model.GoImportPath()
	if p == "" {
		return nil
	}
	return []string{p}
}

// PackageName インポートパスからパッケージ名を取得する
func (a *MappableArgument) PackageName() string {
	return a.Model.GoPkg()
}

// ArgName 引数の変数名、コード生成で利用される
func (a *MappableArgument) ArgName() string {
	return a.Name
}

// TypeName 型名の文字列表現、コード生成で利用される
func (a *MappableArgument) TypeName() string {
	return a.Model.GoTypeSourceCode()
}

// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
func (a *MappableArgument) ZeroInitializer() string {
	return a.Model.ZeroInitializeSourceCode()
}

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (a *MappableArgument) ZeroValueOnSource() string {
	return a.Model.ZeroValueSourceCode()
}

// DestinationFieldName マッピング先となるフィールド名を取得
func (a *MappableArgument) DestinationFieldName() string {
	return a.Destination
}

// DestinationModel マッピング先のModelを取得
func (a *MappableArgument) DestinationModel() *Model {
	return a.Model
}
