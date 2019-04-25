package schema

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/structs"
)

// ValueType パラメータタイプ
type ValueType int

const (
	// TypeInt int
	TypeInt ValueType = iota
	// TypeInt64 int64
	TypeInt64
	// TypeString string
	TypeString
	// TypeError error
	TypeError
	// TypePtr ポインタ
	TypePtr
)

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (v ValueType) ZeroValueOnSource() string {
	switch v {
	case TypeInt:
		return "0"
	case TypeInt64:
		return "0"
	case TypeString:
		return `""`
	case TypeError:
		return "nil"
	case TypePtr:
		return "nil"
	default:
		log.Fatalf("invalid ValueType: %v", v)
		return ""
	}
}

var (
	// ValueError 戻り値でのerrorを示すValue
	ValueError = &ResultValue{
		Value: &Value{
			Name: "err",
			Type: TypeError,
		},
	}

	// ValueID 引数でのIDを示すValue
	ValueID = &Value{
		Name: "id",
		Type: TypeInt64,
	}

	// ValueZone 引数でのゾーンを示すValue
	ValueZone = &Value{
		Name: "zone",
		Type: TypeString,
	}
)

// Value パラメータの型情報
type Value struct {
	Name       string      // パラメータ名、引数名に利用される
	Type       ValueType   // パラメータのデータ型、TypePtrの場合はStruct/ImportPath/Packageを適切に設定する必要がある
	Struct     interface{} // TypeがTypePtrの場合のポインタが指すデータ型
	ImportPath string      // Structの所属するパッケージのインポートパス
	Package    string      // インポートパスの別名をつける場合に指定する(省略化)
	FieldName  string      // リクエストパラメータを作成する場合の対象フィールド名
}

// ResultValue 戻り値の型情報
type ResultValue struct {
	*Value
	NakedTypeName string
	TagName       string
	Comment       string
}

// ImportStatement コード生成時に利用するimport文を生成する
func (v *Value) ImportStatement() string {
	if v.ImportPath == "" {
		return ""
	}
	prefix := ""
	if v.Package != "" {
		prefix = v.Package + " "
	}
	return fmt.Sprintf(`%s"%s"`, prefix, v.ImportPath)
}

// PackageName インポートパスからパッケージ名を取得する。Packageが設定されている場合はそちらが利用される。
func (v *Value) PackageName() string {
	if v.ImportPath == "" {
		return ""
	}

	if v.Package != "" {
		return v.Package
	}

	parts := strings.Split(v.ImportPath, "/")
	return parts[len(parts)-1]
}

// TypeName 型名の文字列表現、コード生成で利用される
func (v *Value) TypeName() string {
	switch v.Type {
	case TypeInt:
		return "int"
	case TypeInt64:
		return "int64"
	case TypeString:
		return "string"
	case TypeError:
		return "error"
	case TypePtr:
		prefix := "*"
		packageName := v.PackageName()
		if packageName != "" {
			prefix = prefix + packageName + "."
		}
		structName := structs.Name(v.Struct)
		return prefix + structName
	}
	log.Fatalf("invalid value type: %#v", v)
	return ""
}

// ZeroInitializer 値を0初期化する文のコードの文字列表現、コード生成で利用される
func (v *Value) ZeroInitializer() string {
	switch v.Type {
	case TypeInt:
		return "0"
	case TypeInt64:
		return "int64(0)"
	case TypeString:
		return `""`
	case TypeError:
		log.Fatalf("invalid value type: %#v - ZeroInitializer func is not support TypeError", v)
		return ""
	case TypePtr:
		prefix := "&"
		packageName := v.PackageName()
		if packageName != "" {
			prefix = prefix + packageName + "."
		}
		structName := structs.Name(v.Struct)
		return prefix + structName + "{}"
	}
	log.Fatalf("invalid value type: %#v", v)
	return ""
}

// ZeroValueOnSource コード上でのゼロ値の文字列表現。コード生成時に利用する
func (v *Value) ZeroValueOnSource() string {
	return v.Type.ZeroValueOnSource()
}
