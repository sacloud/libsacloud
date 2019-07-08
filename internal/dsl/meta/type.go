package meta

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Type 型情報
type Type interface {
	// GoType 型名
	GoType() string
	// GoPkg パッケージ名
	GoPkg() string
	// GoImportPath インポートパス
	GoImportPath() string
	// GoTypeSourceCode ソースコードでの型表現
	GoTypeSourceCode() string
	// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
	ZeroInitializeSourceCode() string
	// ZeroValueSourceCode 型に応じたzero値コード
	ZeroValueSourceCode() string
}

// StaticType あらかじめ静的参照できる型
type StaticType struct {
	goType       string
	goPkg        string
	goImportPath string
	reflectKind  reflect.Kind
}

// GoType 型名
func (t *StaticType) GoType() string { return t.goType }

// GoPkg パッケージ名
func (t *StaticType) GoPkg() string { return t.goPkg }

// GoImportPath インポートパス
func (t *StaticType) GoImportPath() string { return t.goImportPath }

// GoTypeSourceCode ソースコードでの型表現
func (t *StaticType) GoTypeSourceCode() string {
	if t.goPkg != "" && t.reflectKind == reflect.Struct && t.goType != "time.Time" {
		return fmt.Sprintf("*%s", t.goType)
	}
	return t.goType
}

// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
func (t *StaticType) ZeroInitializeSourceCode() string {
	format := "%s"
	if t.goPkg != "" {
		switch t.reflectKind {
		case reflect.Bool, reflect.Int, reflect.Int64,
			reflect.Float32, reflect.Float64, reflect.String:
			format = t.goType + "(%s)"
		}
	}
	switch t.reflectKind {
	case reflect.Bool:
		return fmt.Sprintf(format, "false")
	case reflect.Int:
		return fmt.Sprintf(format, "0")
	case reflect.Int64:
		return fmt.Sprintf(format, "int64(0)")
	case reflect.Float32:
		return fmt.Sprintf(format, "float32(0)")
	case reflect.Float64:
		return fmt.Sprintf(format, "float64(0)")
	case reflect.Interface, reflect.Map, reflect.Slice:
		return fmt.Sprintf(format, t.goType+"{}")
	case reflect.Struct:
		if t.goType == "time.Time" {
			return fmt.Sprintf(format, t.goType+"{}")
		}
		return fmt.Sprintf(format, "&"+t.goType+"{}")
	case reflect.String:
		return fmt.Sprintf(format, `""`)
	default:
		log.Panicf("unsupported Kind: %s", t.reflectKind)
		return ""
	}
}

// ZeroValueSourceCode 型に応じたzero値コード
func (t *StaticType) ZeroValueSourceCode() string {
	format := "%s"
	if t.goPkg != "" {
		switch t.reflectKind {
		case reflect.Bool, reflect.Int, reflect.Int64,
			reflect.Float32, reflect.Float64, reflect.String:
			format = t.goType + "(%s)"
		}
	}
	switch t.reflectKind {
	case reflect.Bool:
		return fmt.Sprintf(format, "false")
	case reflect.Int:
		return fmt.Sprintf(format, "0")
	case reflect.Int64:
		return fmt.Sprintf(format, "int64(0)")
	case reflect.Float32:
		return fmt.Sprintf(format, "float32(0)")
	case reflect.Float64:
		return fmt.Sprintf(format, "float64(0)")
	case reflect.Interface, reflect.Map, reflect.Slice, reflect.Struct:
		if t.goType == "time.Time" {
			return fmt.Sprintf(format, t.goType+"{}")
		}
		return fmt.Sprintf(format, "nil")
	case reflect.String:
		return fmt.Sprintf(format, `""`)
	default:
		log.Panicf("unsupported Kind: %s", t.reflectKind)
		return ""
	}
}

// Static 型情報を受け取りTypeを返す
func Static(v interface{}) *StaticType {
	t := reflect.TypeOf(v)
	pkgName := ""
	pkgPath := t.PkgPath()
	if len(pkgPath) > 0 {
		pathes := strings.Split(t.PkgPath(), "/")
		pkgName = pathes[len(pathes)-1]
	}
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Interface,
		reflect.Map, reflect.Slice, reflect.Struct, reflect.String:
		// noop
	default:
		log.Panicf("unsupported Kind: %s", t.Kind())
		return nil
	}
	return &StaticType{
		goType:       t.String(),
		goPkg:        pkgName,
		goImportPath: t.PkgPath(),
		reflectKind:  t.Kind(),
	}
}
