package schema

import (
	"github.com/sacloud/libsacloud/internal/schema/meta"
)

// Results Resultの配列
type Results []*Result

// Models Resultsに登録されているModelを返す
func (r *Results) Models() Models {
	ms := Models{}
	for _, res := range *r {
		ms = append(ms, res.Model)
		ms = append(ms, res.Model.FieldModels()...)
	}
	return ms
}

// Result Operationでの戻り値定義
type Result struct {
	SourceField string // エンベロープでの宛先ペイロード名
	Model       *Model // パラメータの型情報
}

// ImportStatements コード生成時に利用するimport文を生成する
func (r *Result) ImportStatements(additionalImports ...string) []string {
	return r.Model.ImportStatementsForModelDef(additionalImports...)
}

// Type モデルの型情報
func (r *Result) Type() meta.Type {
	return r
}

// GoType 型名
func (r *Result) GoType() string {
	return r.Model.Name
}

// GoPkg パッケージ名
func (r *Result) GoPkg() string {
	return r.Model.GoPkg()
}

// GoImportPath インポートパス
func (r *Result) GoImportPath() string {
	return r.Model.GoImportPath()
}

// GoTypeSourceCode ソースコードでの型表現
func (r *Result) GoTypeSourceCode() string {
	return r.Model.GoTypeSourceCode()
}

// ZeroInitializeSourceCode 型に応じたzero値での初期化コード
func (r *Result) ZeroInitializeSourceCode() string {
	return r.Model.ZeroInitializeSourceCode()
}

// ZeroValueSourceCode 型に応じたzero値コード
func (r *Result) ZeroValueSourceCode() string {
	return r.Model.ZeroValueSourceCode()
}
