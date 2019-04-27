package schema

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
)

// Operation リソースへの操作
type Operation struct {
	resource         *Resource
	name             string        // 操作名、メソッド名となる
	method           string        // HTTPリクエストメソッド GET/POST/PUT/DELETE
	pathFormat       string        // パスのフォーマット、省略した場合はDefaultPathFormatが設定される
	arguments        Arguments     // 引数の定義
	results          Results       // レスポンス
	requestEnvelope  *EnvelopeType // リクエスト時のエンベロープ
	responseEnvelope *EnvelopeType // レスポンス時のエンベロープ
}

// Name 操作名、メソッド名となる
func (o *Operation) Name(name string) *Operation {
	o.name = name
	return o
}

// Method HTTPリクエストメソッド GET/POST/PUT/DELETE
func (o *Operation) Method(method string) *Operation {
	o.method = method
	return o
}

// GetMethod HTTPリクエストメソッドの取得
func (o *Operation) GetMethod() string {
	return o.method
}

// PathFormat パスのフォーマット、省略した場合はDefaultPathFormatが設定される
func (o *Operation) PathFormat(pathFormat string) *Operation {
	o.pathFormat = pathFormat
	return o
}

// Argument 引数定義の追加(単数)
func (o *Operation) Argument(arg Argument) *Operation {
	o.arguments = append(o.arguments, arg)
	return o
}

// MappableArgument 引数定義の追加
func (o *Operation) MappableArgument(name string, model *Model) *Operation {
	var destField string
	if o.requestEnvelope != nil {
		destField = o.requestEnvelope.PayloadName()
		if destField == "" {
			destField = o.resource.FieldName(o.requestEnvelope.Form)
		}
	}
	return o.Argument(&MappableArgument{
		Name:        name,
		Model:       model,
		Destination: destField,
	})
}

// Arguments 引数定義の追加(複数)
func (o *Operation) Arguments(args []Argument) *Operation {
	o.arguments = append(o.arguments, args...)
	return o
}

// Result レスポンス定義の追加
func (o *Operation) Result(m *Model) *Operation {
	sourceField := ""
	if o.responseEnvelope != nil {
		sourceField = o.responseEnvelope.PayloadName()
		if sourceField == "" {
			sourceField = o.resource.FieldName(o.responseEnvelope.Form)
		}
	}
	return o.ResultWithSourceField(sourceField, m)
}

// ResultWithSourceField レスポンス定義の追加
func (o *Operation) ResultWithSourceField(sourceField string, m *Model) *Operation {
	o.results = append(o.results, &Result{Model: m, SourceField: sourceField})
	return o
}

// RequestEnvelope リクエストのエンベロープを指定する
func (o *Operation) RequestEnvelope(t meta.Type) *Operation {
	o.requestEnvelope = &EnvelopeType{
		Form: PayloadForms.Singular,
		Payload: &EnvelopePayloadDesc{
			PayloadType: t,
		},
	}
	return o
}

// RequestEnvelopePlural リクエストのエンベロープを複数形として指定する
func (o *Operation) RequestEnvelopePlural(t meta.Type) *Operation {
	o.requestEnvelope = &EnvelopeType{
		Form: PayloadForms.Plural,
		Payload: &EnvelopePayloadDesc{
			PayloadType: t,
		},
	}
	return o
}

// ResponseEnvelope レスポンスのエンベロープを指定する
func (o *Operation) ResponseEnvelope(t meta.Type) *Operation {
	o.responseEnvelope = &EnvelopeType{
		Form: PayloadForms.Singular,
		Payload: &EnvelopePayloadDesc{
			PayloadType: t,
		},
	}
	return o
}

// ResponseEnvelopePlural レスポンスのエンベロープを複数形として指定する
func (o *Operation) ResponseEnvelopePlural(t meta.Type) *Operation {
	o.responseEnvelope = &EnvelopeType{
		Form: PayloadForms.Plural,
		Payload: &EnvelopePayloadDesc{
			PayloadType: t,
		},
	}
	return o
}

// ResponseEnvelopeManual レスポンスのエンベロープを詳細指定する
func (o *Operation) ResponseEnvelopeManual(form PayloadForm, payloadType meta.Type, additionalPayloads []*EnvelopePayloadDesc) *Operation {
	o.responseEnvelope = &EnvelopeType{
		Form: form,
		Payload: &EnvelopePayloadDesc{
			PayloadType: payloadType,
		},
		AdditionalPayloads: additionalPayloads,
	}
	return o
}

// DefineResult 操作に対するレスポンスの定義
func (o *Operation) DefineResult() *Model {
	return &Model{}
}

// GetPathFormat パスのフォーマット
func (o *Operation) GetPathFormat() string {
	if o.pathFormat != "" {
		return o.pathFormat
	}
	return DefaultPathFormat
}

// ImportStatements コード生成時に利用するimport文を生成する
func (o *Operation) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)

	for _, arg := range o.arguments {
		ss = append(ss, arg.ImportStatements()...)
	}

	for _, m := range o.results {
		ss = append(ss, m.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// MethodName コード生成時に利用する、メソッド名を出力する
func (o *Operation) MethodName() string {
	return o.name
}

// ReturnErrorStatement コード生成時に利用する、エラーをreturnする文を生成する
func (o *Operation) ReturnErrorStatement() string {
	ss := make([]string, len(o.results))
	for i, res := range o.results {
		s := res.ZeroValueSourceCode()
		ss[i] = s
	}
	ss = append(ss, "err")
	return strings.Join(ss, ",")
}

// HasMapDestinationDeciders エンベロープへのパラメータマッピングを行う必要のなる引数を持つか
func (o *Operation) HasMapDestinationDeciders() bool {
	return len(o.arguments.MapDestinationDeciders()) > 0
}

// MapDestinationDeciders Argumentsのうち、MapDestDeciderであるもののリストを返す
func (o *Operation) MapDestinationDeciders() []MapDestinationDecider {
	return o.arguments.MapDestinationDeciders()
}

// RequestEnvelopeStructName エンベロープのstruct名
func (o *Operation) RequestEnvelopeStructName() string {
	return fmt.Sprintf("%s%sRequestEnvelope", o.resource.name, o.name)
}

// ResponseEnvelopeStructName エンベロープのstruct名
func (o *Operation) ResponseEnvelopeStructName() string {
	return fmt.Sprintf("%s%sResponseEnvelope", o.resource.name, o.name)
}

// AllArguments 設定されている全てのArgumentを取得
func (o *Operation) AllArguments() Arguments {
	return o.arguments
}

// HasResults 戻り値が定義されているかを取得
func (o *Operation) HasResults() bool {
	return len(o.results) > 0
}

// AllResults 戻り値
func (o *Operation) AllResults() Results {
	return o.results
}

// ResultsStatement 戻り値定義部のソースを出力
func (o *Operation) ResultsStatement() string {
	if len(o.results) == 0 {
		return "error"
	}
	rs := "(%s)"
	var strResults []string
	for _, res := range o.results {
		strResults = append(strResults, res.Type().GoTypeSourceCode())
	}
	strResults = append(strResults, "error")
	return fmt.Sprintf(rs, strings.Join(strResults, ","))
}

// Models オペレーション配下の(Nameで)ユニークなモデル一覧を取得
func (o *Operation) Models() Models {
	ms := o.results.Models()
	for _, arg := range o.MapDestinationDeciders() {
		ms = append(ms, arg.DestinationModel())
	}

	return Models(ms).UniqByName()
}

// HasRequestEnvelope リクエストエンベロープが設定されているか
func (o *Operation) HasRequestEnvelope() bool {
	return o.requestEnvelope != nil
}

// RequestEnvelopePayloadName リクエストエンベロープでのペイロードのフィールド名を取得
func (o *Operation) RequestEnvelopePayloadName() string {
	if !o.HasRequestEnvelope() {
		return ""
	}
	return o.envelopePayloadName(o.requestEnvelope)
}

// RequestPayloadTypeName リクエストペイロードの型名
func (o *Operation) RequestPayloadTypeName() string {
	if !o.HasRequestEnvelope() {
		return ""
	}
	return o.requestEnvelope.PayloadType().GoTypeSourceCode()
}

// HasResponseEnvelope レスポンスエンベロープが設定されているか
func (o *Operation) HasResponseEnvelope() bool {
	return o.responseEnvelope != nil
}

// ResponseEnvelopePayloadName レスポンスエンベロープでのペイロードフィールド名を取得
func (o *Operation) ResponseEnvelopePayloadName() string {
	if !o.HasResponseEnvelope() {
		return ""
	}
	return o.envelopePayloadName(o.responseEnvelope)
}

// ResponsePayloadTypeName レスポンスペイロードの型名
func (o *Operation) ResponsePayloadTypeName() string {
	if !o.HasResponseEnvelope() {
		return ""
	}
	return o.responseEnvelope.PayloadType().GoTypeSourceCode()
}

// ResponseAdditionalPayloads レスポンスペイロードの追加フィールド
func (o *Operation) ResponseAdditionalPayloads() []*EnvelopePayloadDesc {
	if !o.HasResponseEnvelope() {
		return nil
	}
	return o.responseEnvelope.AdditionalPayloads
}

// IsResponseSingular レスポンスが単数系か
func (o *Operation) IsResponseSingular() bool {
	if o.HasResponseEnvelope() {
		return o.responseEnvelope.IsSingular()
	}
	return false
}

// IsResponsePlural レスポンスが複数形か
func (o *Operation) IsResponsePlural() bool {
	if o.HasResponseEnvelope() {
		return o.responseEnvelope.IsPlural()
	}
	return false
}

func (o *Operation) envelopePayloadName(evl *EnvelopeType) string {
	if evl == nil {
		return ""
	}

	if evl.PayloadName() == "" {
		return o.resource.FieldName(evl.Form)
	}
	return evl.PayloadName()
}
