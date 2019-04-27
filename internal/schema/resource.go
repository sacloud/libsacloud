package schema

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
)

// Resources []*Resourceのエイリアス
type Resources []*Resource

// ImportStatements コード生成時に利用するimport文を生成する
func (r Resources) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)

	for _, re := range r {
		ss = append(ss, re.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// ImportStatementsForModelDef Resources配下に含まれる全てのモデルのフィールドを含めたimport文を生成する
func (r Resources) ImportStatementsForModelDef(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, m := range r.Models() {
		ss = append(ss, m.ImportStatementsForModelDef()...)
	}
	return uniqStrings(ss)
}

// Define リソースの定義(for fluent API)
func (r *Resources) Define(name string) *Resource {
	if *r == nil {
		rr := Resources{}
		*r = rr
	}
	rs := &Resource{
		name: name,
	}
	*r = append(*r, rs)
	return rs
}

// Models モデル一覧を取得
func (r Resources) Models() Models {
	ms := Models{}
	for _, res := range r {
		for _, o := range res.operations {
			ms = append(ms, o.Models()...)
		}
	}
	return ms.UniqByName()
}

// Resource APIで操作する対象のリソース
type Resource struct {
	name       string       // リソース名 e.g.: Server
	pathName   string       // リソースのパス名 APIのURLで利用される e.g.: server 省略した場合はNameを小文字にしたものとなる
	pathSuffix string       // APIのURLで利用されるプレフィックス e.g.: api/cloud/1.1
	operations []*Operation // このリソースに対する操作
}

// Name リソース名 例: Server
func (r *Resource) Name(name string) *Resource {
	r.name = name
	return r
}

// PathName リソースのパス名 APIのエンドポイントURLの算出で利用される 例: server
//
// 省略した場合はNameをスネークケース(小文字+アンダーバー)に変換したものが利用される
func (r *Resource) PathName(pathName string) *Resource {
	r.pathName = pathName
	return r
}

// GetPathName リソースのパス名 APIのエンドポイントURLの算出で利用される 例: server
//
// 省略した場合はNameをスネークケース(小文字+アンダーバー)に変換したものが利用される
func (r *Resource) GetPathName() string {
	if r.pathName != "" {
		return r.pathSuffix
	}
	return toSnakeCaseName(r.name)
}

// PathSuffix URLでのパス部分のサフィックス APIのエンドポイントURLの算出で利用される 例: api/cloud/1.1
func (r *Resource) PathSuffix(pathSuffix string) *Resource {
	r.pathSuffix = pathSuffix
	return r
}

// GetPathSuffix PathSuffixの取得
func (r *Resource) GetPathSuffix() string {
	if r.pathSuffix != "" {
		return r.pathSuffix
	}
	return CloudAPISuffix
}

// Operation リソースに対する操作の定義を追加
func (r *Resource) Operation(op *Operation) *Resource {
	r.operations = append(r.operations, op)
	return r
}

// Operations リソースに対する操作の定義を追加
func (r *Resource) Operations(ops ...*Operation) *Resource {
	for _, op := range ops {
		r.Operation(op)
	}
	return r
}

// AllOperations 定義されている操作を取得
func (r *Resource) AllOperations() []*Operation {
	return r.operations
}

// DefineOperation リソースに対する操作の定義
func (r *Resource) DefineOperation(name string) *Operation {
	return &Operation{
		resource: r,
		name:     name,
	}
}

// OperationCRUD リソースに対する基本的なCRUDを定義
func (r *Resource) OperationCRUD(nakedType meta.Type, createParam, updateParam, result *Model) *Resource {
	if createParam.Name == "" {
		createParam.Name = r.name + "CreateRequest"
	}
	if updateParam.Name == "" {
		updateParam.Name = r.name + "UpdateRequest"
	}
	if result.Name == "" {
		result.Name = r.name
	}

	if createParam.NakedType == nil {
		createParam.NakedType = nakedType
	}
	if updateParam.NakedType == nil {
		updateParam.NakedType = nakedType
	}
	if result.NakedType == nil {
		result.NakedType = nakedType
	}

	r.Operations(
		// create
		r.DefineOperation("Create").
			Method(http.MethodPost).
			PathFormat(DefaultPathFormat).
			RequestEnvelope(nakedType).
			ResponseEnvelope(nakedType).
			Argument(ArgumentZone).
			MappableArgument("param", createParam).
			Result(result),

		// read
		r.DefineOperation("Read").
			Method(http.MethodGet).
			PathFormat(DefaultPathFormatWithID).
			ResponseEnvelope(nakedType).
			Argument(ArgumentZone).
			Argument(ArgumentID).
			Result(result),

		// update
		r.DefineOperation("Update").
			Method(http.MethodPut).
			PathFormat(DefaultPathFormatWithID).
			RequestEnvelope(nakedType).
			ResponseEnvelope(nakedType).
			Argument(ArgumentZone).
			Argument(ArgumentID).
			MappableArgument("param", updateParam).
			Result(result),

		// delete
		r.DefineOperation("Delete").
			Method(http.MethodDelete).
			PathFormat(DefaultPathFormatWithID).
			Argument(ArgumentZone).
			Argument(ArgumentID),
	)

	return r
}

// FileSafeName スネークケースにしたResourceの名前、コード生成時の保存先ファイル名に利用される
func (r *Resource) FileSafeName() string {
	return toSnakeCaseName(r.name)
}

// TypeName 型名を返す、コード生成時の型定義などで利用される
func (r *Resource) TypeName() string {
	return r.name
}

// ImportStatements コード生成時に利用するimport文を生成する
func (r *Resource) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, o := range r.operations {
		ss = append(ss, o.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// FieldName ペイロードなどで利用される場合のフィールド名を返す
func (r *Resource) FieldName(form PayloadForm) string {
	switch {
	case form.IsSingular():
		return r.name
	case form.IsPlural():
		return r.name + "s"
	default:
		return ""
	}
}
