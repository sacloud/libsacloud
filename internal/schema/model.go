package schema

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
)

// Models APIのリクエスト/レスポンスなどのデータ型を示すモデル
type Models []*Model

// ImportStatements コード生成時に利用するimport文を生成する
func (m Models) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, model := range m {
		ss = append(ss, model.ImportStatements()...)
	}
	return uniqStrings(ss)
}

// IsEmpty 空であるか判定
func (m Models) IsEmpty() bool {
	return len(m) == 0
}

// UniqByName NameでユニークなModelの一覧を返す
func (m Models) UniqByName() Models {
	models := Models{}
	isUniq := func(name string) bool {
		for _, model := range models {
			if model.Name == name {
				return false
			}
		}
		return true
	}
	for _, model := range m {
		if isUniq(model.Name) {
			models = append(models, model)
		}
	}
	return models
}

// Model APIのリクエスト/レスポンスなどのデータ型を示すモデル
type Model struct {
	Name      string       // 型名
	Fields    []*FieldDesc // フィールド定義
	Accessors []*Accessor  // アクセサ
	NakedType meta.Type    // 対応するnaked型の情報
	IsArray   bool
}

// HasNakedType 対応するnaked型の情報が登録されているか
func (m *Model) HasNakedType() bool {
	return m.NakedType != nil
}

// Type モデルの型情報
func (m *Model) Type() meta.Type {
	return m
}

// GoType 型名
func (m *Model) GoType() string {
	prefix := ""
	if m.IsArray {
		prefix = "[]"
	}

	name := m.Name
	if IsOutOfSacloudPackage {
		name = "sacloud." + name
	}

	return prefix + name
}

// GoPkg パッケージ名
func (m *Model) GoPkg() string {
	if IsOutOfSacloudPackage {
		return "sacloud"
	}
	return ""
}

// GoImportPath インポートパス
func (m *Model) GoImportPath() string {
	if IsOutOfSacloudPackage {
		return "github.com/sacloud/libsacloud/v2/sacloud"
	}
	return ""
}

// GoTypeSourceCode ソースコードでの型表現
func (m *Model) GoTypeSourceCode() string {
	name := m.Name
	if IsOutOfSacloudPackage {
		name = "sacloud." + name
	}
	if m.IsArray {
		return fmt.Sprintf("[]*%s", name)
	}
	return fmt.Sprintf("*%s", name)
}

// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
func (m *Model) ZeroInitializeSourceCode() string {
	name := m.Name
	if IsOutOfSacloudPackage {
		name = "sacloud." + name
	}
	if m.IsArray {
		return fmt.Sprintf("[]*%s{}", name)
	}
	return fmt.Sprintf("&%s{}", name)
}

// ZeroValueSourceCode 型に応じたzero値コード
func (m *Model) ZeroValueSourceCode() string {
	return "nil"
}

// ImportStatements コード生成時に利用するimport文を生成する
func (m *Model) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	return uniqStrings(ss)
}

// ImportStatementsForModelDef モデルのフィールドを含めたimport文を生成する
func (m *Model) ImportStatementsForModelDef(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, f := range m.Fields {
		s := f.Type.GoImportPath()
		if s != "" {
			ss = append(ss, wrapByDoubleQuote(s)[0])
		}
	}
	return uniqStrings(ss)
}

// FieldModels フィールド定義に含まれる*Model(FieldDesc.Type)を取得
func (m *Model) FieldModels() []*Model {
	var ms []*Model
	for _, f := range m.Fields {
		if f.Type == nil {
			continue
		}
		if m, ok := f.Type.(*Model); ok {
			ms = append(ms, m)
			ms = append(ms, m.FieldModels()...)
		}
	}
	return ms
}
