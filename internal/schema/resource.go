package schema

import (
	"net/http"
	"strings"

	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
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

// Def リソースの定義
func (r *Resources) Def(rs *Resource) {
	if *r == nil {
		rr := Resources{}
		*r = rr
	}
	if rs.OperationsDefineFunc != nil {
		rs.AddOperations(rs.OperationsDefineFunc(rs)...)
	}
	*r = append(*r, rs)
}

// Define リソースの定義(for fluent API)
func (r *Resources) Define(name string) *Resource {
	if *r == nil {
		rr := Resources{}
		*r = rr
	}
	rs := &Resource{
		Name: name,
	}
	*r = append(*r, rs)
	return rs
}

// DefineWith リソースの定義 & 定義したリソースを利用するfuncの実施
func (r *Resources) DefineWith(name string, f func(*Resource)) *Resource {
	rs := r.Define(name)
	f(rs)
	return rs
}

// Models モデル一覧を取得
func (r Resources) Models() Models {
	ms := Models{}
	for _, res := range r {
		for _, o := range res.Operations() {
			ms = append(ms, o.Models()...)
		}
	}
	return ms.UniqByName()
}

// OperationsDefineFunc リソースに対するオペレーション定義用Func
type OperationsDefineFunc func(r *Resource) []*Operation

// Resource APIで操作する対象のリソース
type Resource struct {
	Name                 string               // リソース名 e.g.: Server
	PathName             string               // リソースのパス名 APIのURLで利用される e.g.: server 省略した場合はNameを小文字にしたものとなる
	PathSuffix           string               // APIのURLで利用されるプレフィックス e.g.: api/cloud/1.1
	IsGlobal             bool                 // 全ゾーンで共通リソース(グローバルリソース)
	operations           []*Operation         // このリソースに対する操作、OperationsDefineFuncが設定されている場合はそちらを呼び出して設定される
	OperationsDefineFunc OperationsDefineFunc // このリソースに対する操作を定義するFunc
}

// GetPathName リソースのパス名 APIのエンドポイントURLの算出で利用される 例: server
//
// 省略した場合はNameをスネークケース(小文字+アンダーバー)に変換したものが利用される
func (r *Resource) GetPathName() string {
	if r.PathName != "" {
		return r.PathName
	}
	return toSnakeCaseName(r.Name)
}

// GetPathSuffix PathSuffixの取得
func (r *Resource) GetPathSuffix() string {
	if r.PathSuffix != "" {
		return r.PathSuffix
	}
	return CloudAPISuffix
}

// Operations リソースに対する操作の定義を取得
func (r *Resource) Operations() []*Operation {
	return r.operations
}

// AddOperation リソースに対する操作の定義を追加
func (r *Resource) AddOperation(op *Operation) {
	r.operations = append(r.operations, op)
}

// AddOperations リソースに対する操作の定義を追加
func (r *Resource) AddOperations(ops ...*Operation) {
	for _, op := range ops {
		r.AddOperation(op)
	}
}

// DefineOperation リソースに対する操作の定義
func (r *Resource) DefineOperation(name string) *Operation {
	return &Operation{
		resource: r,
		name:     name,
	}
}

func (r *Resource) defineOperationFind(nakedType meta.Type, findParam, result *Model, payloadName string) *Operation {
	if findParam.Name == "" {
		findParam.Name = "FindCondition"
	}

	if result.Name == "" {
		result.Name = r.Name
	}

	if result.NakedType == nil {
		result.NakedType = nakedType
	}

	return r.DefineOperation("Find").
		Method(http.MethodGet).
		PathFormat(DefaultPathFormat).
		Argument(ArgumentZone).
		PassthroughModelArgumentWithEnvelope("conditions", findParam).
		ResultPluralFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}, "")
}

// DefineOperationFind Find操作を定義
func (r *Resource) DefineOperationFind(nakedType meta.Type, findParam, result *Model) *Operation {
	return r.defineOperationFind(nakedType, findParam, result, "")
}

// DefineOperationApplianceFind Find操作を定義
func (r *Resource) DefineOperationApplianceFind(nakedType meta.Type, findParam, result *Model) *Operation {
	return r.defineOperationFind(nakedType, findParam, result, "Appliances")
}

// DefineOperationCommonServiceItemFind Find操作を定義
func (r *Resource) DefineOperationCommonServiceItemFind(nakedType meta.Type, findParam, result *Model) *Operation {
	return r.defineOperationFind(nakedType, findParam, result, "CommonServiceItems")
}

func (r *Resource) defineOperationCreate(nakedType meta.Type, createParam, result *Model, payloadName string) *Operation {
	if createParam.Name == "" {
		createParam.Name = r.Name + "CreateRequest"
	}
	if result.Name == "" {
		result.Name = r.Name
	}

	if createParam.NakedType == nil {
		createParam.NakedType = nakedType
	}
	if result.NakedType == nil {
		result.NakedType = nakedType
	}

	return r.DefineOperation("Create").
		Method(http.MethodPost).
		PathFormat(DefaultPathFormat).
		RequestEnvelope(&EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}).
		Argument(ArgumentZone).
		MappableArgument("param", createParam).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}, "")
}

// DefineOperationCreate Create操作を定義
func (r *Resource) DefineOperationCreate(nakedType meta.Type, createParam, result *Model) *Operation {
	return r.defineOperationCreate(nakedType, createParam, result, "")
}

// DefineOperationApplianceCreate Create操作を定義
func (r *Resource) DefineOperationApplianceCreate(nakedType meta.Type, createParam, result *Model) *Operation {
	return r.defineOperationCreate(nakedType, createParam, result, "Appliance")
}

// DefineOperationCommonServiceItemCreate Create操作を定義
func (r *Resource) DefineOperationCommonServiceItemCreate(nakedType meta.Type, createParam, result *Model) *Operation {
	return r.defineOperationCreate(nakedType, createParam, result, "CommonServiceItem")
}

func (r *Resource) defineOperationRead(nakedType meta.Type, result *Model, payloadName string) *Operation {
	if result.Name == "" {
		result.Name = r.Name
	}

	if result.NakedType == nil {
		result.NakedType = nakedType
	}

	return r.DefineOperation("Read").
		Method(http.MethodGet).
		PathFormat(DefaultPathFormatWithID).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}, "")
}

// DefineOperationRead Read操作を定義
func (r *Resource) DefineOperationRead(nakedType meta.Type, result *Model) *Operation {
	return r.defineOperationRead(nakedType, result, "")
}

// DefineOperationApplianceRead Read操作を定義
func (r *Resource) DefineOperationApplianceRead(nakedType meta.Type, result *Model) *Operation {
	return r.defineOperationRead(nakedType, result, "Appliance")
}

// DefineOperationCommonServiceItemRead Read操作を定義
func (r *Resource) DefineOperationCommonServiceItemRead(nakedType meta.Type, result *Model) *Operation {
	return r.defineOperationRead(nakedType, result, "CommonServiceItem")
}

func (r *Resource) defineOperationUpdate(nakedType meta.Type, updateParam, result *Model, payloadName string) *Operation {
	if updateParam.Name == "" {
		updateParam.Name = r.Name + "UpdateRequest"
	}
	if result.Name == "" {
		result.Name = r.Name
	}

	if updateParam.NakedType == nil {
		updateParam.NakedType = nakedType
	}
	if result.NakedType == nil {
		result.NakedType = nakedType
	}

	return r.DefineOperation("Update").
		Method(http.MethodPut).
		PathFormat(DefaultPathFormatWithID).
		RequestEnvelope(&EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		MappableArgument("param", updateParam).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: payloadName,
		}, "")
}

// DefineOperationUpdate Update操作を定義
func (r *Resource) DefineOperationUpdate(nakedType meta.Type, updateParam, result *Model) *Operation {
	return r.defineOperationUpdate(nakedType, updateParam, result, "")
}

// DefineOperationApplianceUpdate Update操作を定義
func (r *Resource) DefineOperationApplianceUpdate(nakedType meta.Type, updateParam, result *Model) *Operation {
	return r.defineOperationUpdate(nakedType, updateParam, result, "Appliance")
}

// DefineOperationCommonServiceItemUpdate Update操作を定義
func (r *Resource) DefineOperationCommonServiceItemUpdate(nakedType meta.Type, updateParam, result *Model) *Operation {
	return r.defineOperationUpdate(nakedType, updateParam, result, "CommonServiceItem")
}

// DefineOperationDelete Delete操作を定義
func (r *Resource) DefineOperationDelete() *Operation {
	return r.DefineOperation("Delete").
		Method(http.MethodDelete).
		PathFormat(DefaultPathFormatWithID).
		Argument(ArgumentZone).
		Argument(ArgumentID)
}

// DefineOperationCRUD リソースに対する基本的なCRUDを定義
func (r *Resource) DefineOperationCRUD(nakedType meta.Type, findParam, createParam, updateParam, result *Model) []*Operation {
	var ops []*Operation
	ops = append(ops, r.DefineOperationFind(nakedType, findParam, result))
	ops = append(ops, r.DefineOperationCreate(nakedType, createParam, result))
	ops = append(ops, r.DefineOperationRead(nakedType, result))
	ops = append(ops, r.DefineOperationUpdate(nakedType, updateParam, result))
	ops = append(ops, r.DefineOperationDelete())
	return ops
}

// DefineOperationConfig Config操作を定義
func (r *Resource) DefineOperationConfig() *Operation {
	return r.DefineOperation("Config").
		Method(http.MethodPut).
		PathFormat(IDAndSuffixPathFormat("config")).
		Argument(ArgumentZone).
		Argument(ArgumentID)
}

// DefineOperationBoot リソースに対するBoot操作を定義
func (r *Resource) DefineOperationBoot() *Operation {
	return r.DefineOperation("Boot").
		Method(http.MethodPut).
		PathFormat(IDAndSuffixPathFormat("power")).
		Argument(ArgumentZone).
		Argument(ArgumentID)
}

// DefineOperationShutdown リソースに対するシャットダウン操作を定義
func (r *Resource) DefineOperationShutdown() *Operation {
	return r.DefineOperation("Shutdown").
		Method(http.MethodDelete).
		PathFormat(IDAndSuffixPathFormat("power")).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		PassthroughModelArgumentWithEnvelope("shutdownOption", &Model{
			Name: "ShutdownOption",
			Fields: []*FieldDesc{
				{
					Name: "Force",
					Type: meta.TypeFlag,
				},
			},
		})
}

// DefineOperationReset リソースに対するリセット操作を定義
func (r *Resource) DefineOperationReset() *Operation {
	return r.DefineOperation("Reset").
		Method(http.MethodPut).
		PathFormat(IDAndSuffixPathFormat("reset")).
		Argument(ArgumentZone).
		Argument(ArgumentID)
}

// DefineOperationPowerManagement リソースに対する電源管理操作を定義
func (r *Resource) DefineOperationPowerManagement() []*Operation {
	return []*Operation{
		r.DefineOperationBoot(),
		r.DefineOperationShutdown(),
		r.DefineOperationReset(),
	}
}

// DefineOperationStatus ステータス取得操作を定義
func (r *Resource) DefineOperationStatus(nakedType meta.Type, result *Model) *Operation {
	return r.DefineOperation("Status").
		Method(http.MethodGet).
		PathFormat(IDAndSuffixPathFormat("status")).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		ResultPluralFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: meta.Static(naked.LoadBalancerStatus{}),
			PayloadName: r.FieldName(PayloadForms.Singular),
		}, "Status")
}

// DefineOperationOpenFTP FTPオープン操作を定義
func (r *Resource) DefineOperationOpenFTP(openParam, result *Model) *Operation {
	o := r.DefineOperation("OpenFTP").
		Method(http.MethodPut).
		PathFormat(IDAndSuffixPathFormat("ftp")).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadName: result.Name,
			PayloadType: meta.Static(naked.OpeningFTPServer{}),
		}, result.Name)
	if openParam != nil {
		o.PassthroughModelArgumentWithEnvelope("openOption", openParam)
	}
	return o
}

// DefineOperationCloseFTP FTPクローズ操作を定義
func (r *Resource) DefineOperationCloseFTP() *Operation {
	return r.DefineOperation("CloseFTP").
		Method(http.MethodDelete).
		PathFormat(IDAndSuffixPathFormat("ftp")).
		Argument(ArgumentZone).
		Argument(ArgumentID)
}

// DefineSimpleOperation ID+αのみを引数にとるシンプルなオペレーションを定義
func (r *Resource) DefineSimpleOperation(opName, method, pathSuffix string, arguments ...*Argument) *Operation {
	o := r.DefineOperation(opName).
		Method(method).
		PathFormat(IDAndSuffixPathFormat(pathSuffix)).
		Argument(ArgumentZone).
		Argument(ArgumentID)
	if len(arguments) > 0 {
		o.Arguments(arguments)
	}
	return o
}

// DefineOperationMonitor アクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitor(monitorParam, result *Model) *Operation {
	return r.DefineOperation("Monitor").
		Method(http.MethodGet).
		PathFormat(IDAndSuffixPathFormat("monitor")).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		PassthroughModelArgumentWithEnvelope("condition", monitorParam).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: meta.Static(naked.MonitorValues{}),
			PayloadName: "Data",
		}, result.Name)
}

// DefineOperationMonitorChild アクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitorChild(funcNameSuffix, childResourceName string, monitorParam, result *Model) *Operation {
	return r.DefineOperation("Monitor"+funcNameSuffix).
		Method(http.MethodGet).
		PathFormat(IDAndSuffixPathFormat(childResourceName+"/monitor")).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		PassthroughModelArgumentWithEnvelope("condition", monitorParam).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: meta.Static(naked.MonitorValues{}),
			PayloadName: "Data",
		}, result.Name)
}

// DefineOperationMonitorChildBy アプライアンスなどでの内部リソースインデックスを持つアクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitorChildBy(funcNameSuffix, childResourceName string, monitorParam, result *Model) *Operation {

	pathSuffix := childResourceName + "/{{if eq .index 0}}{{.index}}{{end}}/monitor"

	return r.DefineOperation("Monitor"+funcNameSuffix).
		Method(http.MethodGet).
		PathFormat(IDAndSuffixPathFormat(pathSuffix)).
		Argument(ArgumentZone).
		Argument(ArgumentID).
		Argument(&Argument{
			Name: "index",
			Type: meta.TypeInt,
		}).
		PassthroughModelArgumentWithEnvelope("condition", monitorParam).
		ResultFromEnvelope(result, &EnvelopePayloadDesc{
			PayloadType: meta.Static(naked.MonitorValues{}),
			PayloadName: "Data",
		}, result.Name)
}

// FileSafeName スネークケースにしたResourceの名前、コード生成時の保存先ファイル名に利用される
func (r *Resource) FileSafeName() string {
	return toSnakeCaseName(r.Name)
}

// TypeName 型名を返す、コード生成時の型定義などで利用される
func (r *Resource) TypeName() string {
	return r.Name
}

// ImportStatements コード生成時に利用するimport文を生成する
func (r *Resource) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)
	for _, o := range r.Operations() {
		ss = append(ss, o.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// FieldName ペイロードなどで利用される場合のフィールド名を返す
func (r *Resource) FieldName(form PayloadForm) string {
	switch {
	case form.IsSingular():
		return r.Name
	case form.IsPlural():
		// TODO とりあえずワードで例外指定
		switch {
		case r.Name == "NFS":
			return r.Name
		case strings.HasSuffix(r.Name, "ch"):
			return r.Name + "es"
		default:
			return r.Name + "s"
		}
	default:
		return ""
	}
}
