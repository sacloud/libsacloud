package schema

import (
	"fmt"
	"strings"
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

// PassthroughArgument 引数定義の追加
func (o *Operation) PassthroughArgument(name string, model *Model) *Operation {
	return o.Argument(&PassthroughArgument{
		Name:  name,
		Model: model,
	})
}

// PassthroughArgumentToPayload 引数定義の追加、ペイロードの定義も同時に行われる
func (o *Operation) PassthroughArgumentToPayload(name string, model *Model) *Operation {
	var descs []*EnvelopePayloadDesc
	for _, field := range model.Fields {
		descs = append(descs, &EnvelopePayloadDesc{
			PayloadName: field.Name,
			PayloadType: field.Type,
		})
	}
	o.RequestEnvelope(descs...)
	return o.Argument(&PassthroughArgument{
		Name:  name,
		Model: model,
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

// ResultFromEnvelope エンベロープから抽出するレスポンス定義の追加
func (o *Operation) ResultFromEnvelope(m *Model, sourceField *EnvelopePayloadDesc) *Operation {
	if o.responseEnvelope == nil {
		o.responseEnvelope = &EnvelopeType{
			Form: PayloadForms.Singular,
		}
	}
	if sourceField.PayloadName == "" {
		sourceField.PayloadName = o.resource.FieldName(o.responseEnvelope.Form)
	}
	o.responseEnvelope.Payloads = append(o.responseEnvelope.Payloads, sourceField)
	return o.ResultWithSourceField(sourceField.PayloadName, m)
}

// ResultPluralFromEnvelope エンベロープから抽出するレスポンス定義の追加(複数形)
func (o *Operation) ResultPluralFromEnvelope(m *Model, sourceField *EnvelopePayloadDesc) *Operation {
	if o.responseEnvelope == nil {
		o.responseEnvelope = &EnvelopeType{
			Form: PayloadForms.Plural,
		}
	}
	if sourceField.PayloadName == "" {
		sourceField.PayloadName = o.resource.FieldName(o.responseEnvelope.Form)
	}
	o.responseEnvelope.Payloads = append(o.responseEnvelope.Payloads, sourceField)
	return o.ResultWithSourceField(sourceField.PayloadName, m)
}

// ResultWithSourceField レスポンス定義の追加
func (o *Operation) ResultWithSourceField(sourceField string, m *Model) *Operation {
	if sourceField == "" {
		sourceField = m.Name
	}
	o.results = append(o.results, &Result{Model: m, SourceField: sourceField})
	return o
}

// RequestEnvelope リクエストのエンベロープを追加する
func (o *Operation) RequestEnvelope(descs ...*EnvelopePayloadDesc) *Operation {
	if o.requestEnvelope == nil {
		o.requestEnvelope = &EnvelopeType{
			Form: PayloadForms.Singular,
		}
	}
	for _, desc := range descs {
		if desc.PayloadName == "" {
			desc.PayloadName = o.resource.FieldName(o.requestEnvelope.Form)
		}
		o.requestEnvelope.Payloads = append(o.requestEnvelope.Payloads, desc)
	}

	return o
}

// RequestEnvelopePlural リクエストのエンベロープを複数形として追加する
func (o *Operation) RequestEnvelopePlural(descs ...*EnvelopePayloadDesc) *Operation {
	if o.requestEnvelope == nil {
		o.requestEnvelope = &EnvelopeType{
			Form: PayloadForms.Plural,
		}
	}
	for _, desc := range descs {
		if desc.PayloadName == "" {
			desc.PayloadName = o.resource.FieldName(o.requestEnvelope.Form)
		}
		o.requestEnvelope.Payloads = append(o.requestEnvelope.Payloads, desc)
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

// HasPassthroughFieldDecider エンベロープへのパラメータマッピングを行う必要のなる引数を持つか
func (o *Operation) HasPassthroughFieldDecider() bool {
	return len(o.arguments.PassthroughFieldDeciders()) > 0
}

// PassthroughFieldDeciders Argumentsのうち、PassthroughFieldDeciderであるもののリストを返す
func (o *Operation) PassthroughFieldDeciders() []PassthroughFieldDecider {
	return o.arguments.PassthroughFieldDeciders()
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
		prefix := ""
		if o.IsResponsePlural() {
			prefix = "[]"
		}
		strResults = append(strResults, prefix+res.Type().GoTypeSourceCode())
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
	for _, arg := range o.PassthroughFieldDeciders() {
		ms = append(ms, arg.DestinationModel())
	}
	for _, res := range o.results {
		ms = append(ms, res.Model)
	}
	return Models(ms).UniqByName()
}

// HasRequestEnvelope リクエストエンベロープが設定されているか
func (o *Operation) HasRequestEnvelope() bool {
	return o.requestEnvelope != nil
}

// RequestPayloads リクエストペイロードを取得
func (o *Operation) RequestPayloads() []*EnvelopePayloadDesc {
	if o.HasRequestEnvelope() {
		return o.requestEnvelope.Payloads
	}
	return nil
}

// HasResponseEnvelope レスポンスエンベロープが設定されているか
func (o *Operation) HasResponseEnvelope() bool {
	return o.responseEnvelope != nil
}

// ResponsePayloads レスポンスペイロードを取得
func (o *Operation) ResponsePayloads() []*EnvelopePayloadDesc {
	if o.HasResponseEnvelope() {
		return o.responseEnvelope.Payloads
	}
	return nil
}

// IsRequestSingular リクエストが単数系か
func (o *Operation) IsRequestSingular() bool {
	if o.HasRequestEnvelope() {
		return o.requestEnvelope.IsSingular()
	}
	return false
}

// IsRequestPlural リクエストが複数形か
func (o *Operation) IsRequestPlural() bool {
	if o.HasRequestEnvelope() {
		return o.requestEnvelope.IsPlural()
	}
	return false
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
