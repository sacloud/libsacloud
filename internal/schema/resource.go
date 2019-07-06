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

	// TODO あとで直す
	o := &Operation{
		Resource:        r,
		Name:            "Find",
		PathFormat:      DefaultPathFormat,
		Method:          http.MethodGet,
		RequestEnvelope: RequestEnvelopeFromModel(findParam),
		Arguments: Arguments{
			ArgumentZone,
			PassthroughModelArgument("conditions", findParam),
		},
	}
	o.ResponseEnvelope = ResultPluralFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: nakedType,
		PayloadName: payloadName,
	}, "")
	return o
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

	destField := payloadName
	if destField == "" {
		destField = r.FieldName(PayloadForms.Singular)
	}

	// TODO あとで直す
	o := &Operation{
		Resource:   r,
		Name:       "Create",
		PathFormat: DefaultPathFormat,
		Method:     http.MethodPost,
		RequestEnvelope: RequestEnvelope(&EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: destField,
		}),
		Arguments: Arguments{
			ArgumentZone,
			MappableArgument("param", createParam, destField),
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: nakedType,
		PayloadName: payloadName,
	}, "")
	return o
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

	// TODO あとで直す
	o := &Operation{
		Resource:   r,
		Name:       "Read",
		PathFormat: DefaultPathFormatWithID,
		Method:     http.MethodGet,
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: nakedType,
		PayloadName: payloadName,
	}, "")
	return o
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
	destField := payloadName
	if destField == "" {
		destField = r.FieldName(PayloadForms.Singular)
	}
	// TODO あとで直す
	o := &Operation{
		Resource:   r,
		Name:       "Update",
		PathFormat: DefaultPathFormatWithID,
		Method:     http.MethodPut,
		RequestEnvelope: RequestEnvelope(&EnvelopePayloadDesc{
			PayloadType: nakedType,
			PayloadName: destField,
		}),
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
			MappableArgument("param", updateParam, destField),
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: nakedType,
		PayloadName: payloadName,
	}, "")
	return o
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
	return &Operation{
		Resource: r,
		Name:     "Delete",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: DefaultPathFormatWithID,
		Method:     http.MethodDelete,
	}
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
	return &Operation{
		Resource: r,
		Name:     "Config",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("config"),
		Method:     http.MethodPut,
	}
}

// DefineOperationBoot リソースに対するBoot操作を定義
func (r *Resource) DefineOperationBoot() *Operation {
	return &Operation{
		Resource: r,
		Name:     "Boot",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("power"),
		Method:     http.MethodPut,
	}
}

// DefineOperationShutdown リソースに対するシャットダウン操作を定義
func (r *Resource) DefineOperationShutdown() *Operation {
	param := &Model{
		Name: "ShutdownOption",
		Fields: []*FieldDesc{
			{
				Name: "Force",
				Type: meta.TypeFlag,
			},
		},
	}
	o := &Operation{
		Resource:        r,
		Name:            "Shutdown",
		PathFormat:      IDAndSuffixPathFormat("power"),
		Method:          http.MethodDelete,
		RequestEnvelope: RequestEnvelopeFromModel(param),
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
			PassthroughModelArgument("shutdownOption", param),
		},
	}

	return o
}

// DefineOperationReset リソースに対するリセット操作を定義
func (r *Resource) DefineOperationReset() *Operation {
	return &Operation{
		Resource: r,
		Name:     "Reset",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("reset"),
		Method:     http.MethodPut,
	}
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
	// TODO あとで直す
	o := &Operation{
		Resource: r,
		Name:     "Status",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("status"),
		Method:     http.MethodGet,
	}
	o.ResponseEnvelope = ResultPluralFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: meta.Static(naked.LoadBalancerStatus{}),
		PayloadName: r.FieldName(PayloadForms.Singular),
	}, "Status")
	return o
}

// DefineOperationOpenFTP FTPオープン操作を定義
func (r *Resource) DefineOperationOpenFTP(openParam, result *Model) *Operation {
	// TODO あとで直す
	o := &Operation{
		Resource: r,
		Name:     "OpenFTP",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("ftp"),
		Method:     http.MethodPut,
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadName: result.Name,
		PayloadType: meta.Static(naked.OpeningFTPServer{}),
	}, result.Name)

	if openParam != nil {
		o.Arguments = append(o.Arguments, PassthroughModelArgument("openOption", openParam))
		o.RequestEnvelope = RequestEnvelopeFromModel(openParam)
	}
	return o
}

// DefineOperationCloseFTP FTPクローズ操作を定義
func (r *Resource) DefineOperationCloseFTP() *Operation {
	return &Operation{
		Resource: r,
		Name:     "CloseFTP",
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat("ftp"),
		Method:     http.MethodDelete,
	}
}

// DefineSimpleOperation ID+αのみを引数にとるシンプルなオペレーションを定義
func (r *Resource) DefineSimpleOperation(opName, method, pathSuffix string, arguments ...*Argument) *Operation {
	// TODO あとで直す
	o := &Operation{
		Resource: r,
		Name:     opName,
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
		},
		PathFormat: IDAndSuffixPathFormat(pathSuffix),
		Method:     method,
	}
	o.Arguments = append(o.Arguments, arguments...)
	return o
}

// DefineOperationMonitor アクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitor(monitorParam, result *Model) *Operation {
	// TODO あとで直す
	o := &Operation{
		Resource:        r,
		Name:            "Monitor",
		PathFormat:      IDAndSuffixPathFormat("monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: RequestEnvelopeFromModel(monitorParam),
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
			PassthroughModelArgument("condition", monitorParam),
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: meta.Static(naked.MonitorValues{}),
		PayloadName: "Data",
	}, result.Name)
	return o
}

// DefineOperationMonitorChild アクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitorChild(funcNameSuffix, childResourceName string, monitorParam, result *Model) *Operation {
	// TODO あとで直す
	o := &Operation{
		Resource:        r,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      IDAndSuffixPathFormat(childResourceName + "/monitor"),
		Method:          http.MethodGet,
		RequestEnvelope: RequestEnvelopeFromModel(monitorParam),
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
			PassthroughModelArgument("condition", monitorParam),
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: meta.Static(naked.MonitorValues{}),
		PayloadName: "Data",
	}, result.Name)
	return o
}

// DefineOperationMonitorChildBy アプライアンスなどでの内部リソースインデックスを持つアクティビティモニタ取得操作を定義
func (r *Resource) DefineOperationMonitorChildBy(funcNameSuffix, childResourceName string, monitorParam, result *Model) *Operation {

	pathSuffix := childResourceName + "/{{if eq .index 0}}{{.index}}{{end}}/monitor"

	o := &Operation{
		Resource:        r,
		Name:            "Monitor" + funcNameSuffix,
		PathFormat:      IDAndSuffixPathFormat(pathSuffix),
		Method:          http.MethodGet,
		RequestEnvelope: RequestEnvelopeFromModel(monitorParam),
		Arguments: Arguments{
			ArgumentZone,
			ArgumentID,
			{
				Name: "index",
				Type: meta.TypeInt,
			},
			PassthroughModelArgument("condition", monitorParam),
		},
	}
	o.ResponseEnvelope = ResultFromEnvelope(o, result, &EnvelopePayloadDesc{
		PayloadType: meta.Static(naked.MonitorValues{}),
		PayloadName: "Data",
	}, result.Name)
	return o
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
