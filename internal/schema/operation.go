package schema

import (
	"fmt"
	"strings"
)

// Operation リソースへの操作
type Operation struct {
	Resource         *Resource
	Name             string        // 操作名、メソッド名となる
	Method           string        // HTTPリクエストメソッド GET/POST/PUT/DELETE
	PathFormat       string        // パスのフォーマット、省略した場合はDefaultPathFormatが設定される
	Arguments        Arguments     // 引数の定義
	Results          Results       // レスポンス
	RequestEnvelope  *EnvelopeType // リクエスト時のエンベロープ
	ResponseEnvelope *EnvelopeType // レスポンス時のエンベロープ
}

// MappableArgument 引数定義の追加
func MappableArgument(o *Operation, name string, model *Model) *Argument {
	var destField string
	if o.RequestEnvelope != nil {
		destField = o.RequestEnvelope.PayloadName()
		if destField == "" {
			destField = o.Resource.FieldName(o.RequestEnvelope.Form)
		}
	}
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: fmt.Sprintf("%s,recursive", destField),
	}
}

// PassthroughModelArgumentWithEnvelope 引数定義の追加、ペイロードの定義も同時に行われる
func PassthroughModelArgumentWithEnvelope(o *Operation, name string, model *Model) *Argument {
	var descs []*EnvelopePayloadDesc
	for _, field := range model.Fields {
		descs = append(descs, &EnvelopePayloadDesc{
			PayloadName: field.Name,
			PayloadType: field.Type,
		})
	}
	o.RequestEnvelope = RequestEnvelope(o, descs...)
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: ",squash",
	}
}

// RequestEnvelope リクエストのエンベロープを追加する
func RequestEnvelope(o *Operation, descs ...*EnvelopePayloadDesc) *EnvelopeType {
	ret := &EnvelopeType{
		Form: PayloadForms.Singular,
	}

	for _, desc := range descs {
		if desc.PayloadName == "" {
			desc.PayloadName = o.Resource.FieldName(ret.Form)
		}
		ret.Payloads = append(ret.Payloads, desc)
	}

	return ret
}

// ResultFromEnvelope エンベロープから抽出するレスポンス定義の追加
func ResultFromEnvelope(o *Operation, m *Model, sourceField *EnvelopePayloadDesc, destFieldName string) *EnvelopeType {
	responseEnvelope := &EnvelopeType{
		Form: PayloadForms.Singular,
	}
	if sourceField.PayloadName == "" {
		sourceField.PayloadName = o.Resource.FieldName(responseEnvelope.Form)
	}
	if destFieldName == "" {
		destFieldName = o.Resource.FieldName(responseEnvelope.Form)
	}
	responseEnvelope.Payloads = append(responseEnvelope.Payloads, sourceField)
	if sourceField.Tags == nil {
		sourceField.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	// TODO あとで直す
	o.resultFromEnvelope(sourceField.PayloadName, destFieldName, false, m)
	return responseEnvelope
}

// ResultPluralFromEnvelope エンベロープから抽出するレスポンス定義の追加(複数形)
func ResultPluralFromEnvelope(o *Operation, m *Model, sourceField *EnvelopePayloadDesc, destFieldName string) *EnvelopeType {
	responseEnvelope := &EnvelopeType{
		Form: PayloadForms.Plural,
	}
	if sourceField.PayloadName == "" {
		sourceField.PayloadName = o.Resource.FieldName(responseEnvelope.Form)
	}
	if destFieldName == "" {
		destFieldName = o.Resource.FieldName(responseEnvelope.Form)
	}

	responseEnvelope.Payloads = append(responseEnvelope.Payloads, sourceField)
	if sourceField.Tags == nil {
		sourceField.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	// TODO あとで直す
	o.resultFromEnvelope(sourceField.PayloadName, destFieldName, true, m)
	return responseEnvelope
}

// resultWithDestField レスポンス定義の追加
func (o *Operation) resultFromEnvelope(sourceField, destField string, isPlural bool, m *Model) *Operation {
	if destField == "" {
		destField = m.Name
	}
	result := &Result{
		Model:       m,
		SourceField: sourceField,
		DestField:   destField,
		IsPlural:    isPlural,
	}
	o.Results = append(o.Results, result)
	return o
}

// GetPathFormat パスのフォーマット
func (o *Operation) GetPathFormat() string {
	if o.PathFormat != "" {
		return o.PathFormat
	}
	return DefaultPathFormat
}

// ImportStatements コード生成時に利用するimport文を生成する
func (o *Operation) ImportStatements(additionalImports ...string) []string {
	ss := wrapByDoubleQuote(additionalImports...)

	for _, arg := range o.Arguments {
		ss = append(ss, arg.ImportStatements()...)
	}

	for _, m := range o.Results {
		ss = append(ss, m.ImportStatements()...)
	}

	return uniqStrings(ss)
}

// MethodName コード生成時に利用する、メソッド名を出力する
func (o *Operation) MethodName() string {
	return o.Name
}

// ReturnErrorStatement コード生成時に利用する、エラーをreturnする文を生成する
func (o *Operation) ReturnErrorStatement() string {
	if o.HasResults() {
		return "nil, err"
	}
	return "err"
}

// RequestEnvelopeStructName エンベロープのstruct名(camel-case)
func (o *Operation) RequestEnvelopeStructName() string {
	return fmt.Sprintf("%s%sRequestEnvelope", toCamelWithFirstLower(o.Resource.Name), o.Name)
}

// ResponseEnvelopeStructName エンベロープのstruct名(camel-case)
func (o *Operation) ResponseEnvelopeStructName() string {
	return fmt.Sprintf("%s%sResponseEnvelope", toCamelWithFirstLower(o.Resource.Name), o.Name)
}

// ResultTypeName API戻り値の型名
func (o *Operation) ResultTypeName() string {
	return o.resultType().GoType()
}

// HasResults 戻り値が定義されているかを取得
func (o *Operation) HasResults() bool {
	return len(o.Results) > 0
}

// ResultsStatement 戻り値定義部のソースを出力
func (o *Operation) ResultsStatement() string {
	if !o.HasResults() {
		return "error"
	}
	return fmt.Sprintf("(%s, error)", o.resultType().GoTypeSourceCode())
}

// StubFieldDefines スタブ生成時のフィールド定義文を全フィールド分出力
func (o *Operation) StubFieldDefines() []string {
	if len(o.Results) == 0 {
		return nil
	}
	return []string{fmt.Sprintf("Values %s", o.resultType().GoTypeSourceCode())}
}

// StubReturnStatement スタブ生成時のreturn文
func (o *Operation) StubReturnStatement(receiverName string) string {
	if len(o.Results) == 0 {
		return fmt.Sprintf("return %s.%sStubResult.Err", receiverName, o.MethodName())
	}
	var strResults []string
	strResults = append(strResults, fmt.Sprintf("%s.%sStubResult.Values", receiverName, o.MethodName()))
	strResults = append(strResults, fmt.Sprintf("%s.%sStubResult.Err", receiverName, o.MethodName()))
	return fmt.Sprintf("return %s", strings.Join(strResults, ","))
}

// Models オペレーション配下の(Nameで)ユニークなモデル一覧を取得
func (o *Operation) Models() Models {
	ms := o.Results.Models()
	for _, arg := range o.Arguments {
		m, ok := arg.Type.(*Model)
		if ok {
			ms = append(ms, m)
			ms = append(ms, m.FieldModels()...)
		}

	}
	return Models(ms).UniqByName()
}

// HasRequestEnvelope リクエストエンベロープが設定されているか
func (o *Operation) HasRequestEnvelope() bool {
	return o.RequestEnvelope != nil
}

// RequestPayloads リクエストペイロードを取得
func (o *Operation) RequestPayloads() []*EnvelopePayloadDesc {
	if o.HasRequestEnvelope() {
		return o.RequestEnvelope.Payloads
	}
	return nil
}

// HasResponseEnvelope レスポンスエンベロープが設定されているか
func (o *Operation) HasResponseEnvelope() bool {
	return o.ResponseEnvelope != nil
}

// ResponsePayloads レスポンスペイロードを取得
func (o *Operation) ResponsePayloads() []*EnvelopePayloadDesc {
	if o.HasResponseEnvelope() {
		return o.ResponseEnvelope.Payloads
	}
	return nil
}

// IsRequestSingular リクエストが単数系か
func (o *Operation) IsRequestSingular() bool {
	if o.HasRequestEnvelope() {
		return o.RequestEnvelope.IsSingular()
	}
	return false
}

// IsRequestPlural リクエストが複数形か
func (o *Operation) IsRequestPlural() bool {
	if o.HasRequestEnvelope() {
		return o.RequestEnvelope.IsPlural()
	}
	return false
}

// IsResponseSingular レスポンスが単数系か
func (o *Operation) IsResponseSingular() bool {
	if o.HasResponseEnvelope() {
		return o.ResponseEnvelope.IsSingular()
	}
	return false
}

// IsResponsePlural レスポンスが複数形か
func (o *Operation) IsResponsePlural() bool {
	if o.HasResponseEnvelope() {
		return o.ResponseEnvelope.IsPlural()
	}
	return false
}

// FileSafeName スネークケースにしたResourceの名前、コード生成時の保存先ファイル名に利用される
func (o *Operation) FileSafeName() string {
	return toSnakeCaseName(o.Name)
}

// ResourceTypeName リソースの名称を取得
func (o *Operation) ResourceTypeName() string {
	return o.Resource.TypeName()
}

// ResourceIsGlobal リソースがグローバルリソースか
func (o *Operation) ResourceIsGlobal() bool {
	return o.Resource.IsGlobal
}

func (o *Operation) resultType() *ResultType {
	return &ResultType{
		resource:  o.Resource,
		operation: o,
		results:   o.Results,
	}
}
