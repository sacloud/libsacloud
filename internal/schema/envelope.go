package schema

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
)

// EnvelopeType Modelを用いてAPIとやりとりする際のリクエスト/レスポンスのエンベロープ
type EnvelopeType struct {
	Form     PayloadForm // ペイロードの形体(単数/複数)
	Payloads []*EnvelopePayloadDesc
}

// IsSingular 単数系か判定
func (f *EnvelopeType) IsSingular() bool {
	return f.Form.IsSingular()
}

// IsPlural 複数形か判定
func (f *EnvelopeType) IsPlural() bool {
	return f.Form.IsPlural()
}

// PayloadName ペイロード名を取得
func (f *EnvelopeType) PayloadName() string {
	if len(f.Payloads) > 0 {
		return f.Payloads[0].PayloadName
	}
	return ""
}

// PayloadType ペイロードの型情報を取得
func (f *EnvelopeType) PayloadType() meta.Type {
	if len(f.Payloads) > 0 {
		return f.Payloads[0].PayloadType
	}
	return nil
}

// EnvelopePayloadDesc エンベロープに含まれるペイロードの情報
type EnvelopePayloadDesc struct {
	PayloadName string    // ペイロードのフィールド名
	PayloadType meta.Type // ペイロードの型情報
	Tags        *FieldTags
}

// TypeName ペイロードの型定義
func (d *EnvelopePayloadDesc) TypeName() string {
	return d.PayloadType.GoTypeSourceCode()
}

// TagString タグの文字列表現
func (d *EnvelopePayloadDesc) TagString() string {
	if d.Tags == nil {
		d.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	return fmt.Sprintf("`%s`", d.Tags.String())
}

// PayloadForm ペイロードの形体
type PayloadForm int

// PayloadForms ペイロードの形体
var PayloadForms = struct {
	Singular PayloadForm
	Plural   PayloadForm
}{
	Singular: PayloadForm(0),
	Plural:   PayloadForm(1),
}

// IsSingular 単数系か判定
func (f PayloadForm) IsSingular() bool {
	return int(f) == int(PayloadForms.Singular)
}

// IsPlural 複数形か判定
func (f PayloadForm) IsPlural() bool {
	return int(f) == int(PayloadForms.Plural)
}

// RequestEnvelope リクエストのエンベロープを作成する
func RequestEnvelope(descs ...*EnvelopePayloadDesc) *EnvelopeType {
	ret := &EnvelopeType{
		Form: PayloadForms.Singular,
	}

	for _, desc := range descs {
		ret.Payloads = append(ret.Payloads, desc)
	}

	return ret
}

// RequestEnvelopeFromModel モデル定義からリクエストのエンベロープを作成する
func RequestEnvelopeFromModel(model *Model) *EnvelopeType {
	var descs []*EnvelopePayloadDesc
	for _, field := range model.Fields {
		descs = append(descs, &EnvelopePayloadDesc{
			PayloadName: field.Name,
			PayloadType: field.Type,
		})
	}
	ret := &EnvelopeType{
		Form: PayloadForms.Singular,
	}

	for _, desc := range descs {
		ret.Payloads = append(ret.Payloads, desc)
	}

	return ret
}

// ResponseEnvelope エンベロープから抽出するレスポンス定義の追加
func ResponseEnvelope(sourceField *EnvelopePayloadDesc) *EnvelopeType {
	responseEnvelope := &EnvelopeType{
		Form:     PayloadForms.Singular,
		Payloads: []*EnvelopePayloadDesc{sourceField},
	}
	if sourceField.Tags == nil {
		sourceField.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	return responseEnvelope
}

// ResponseEnvelopePlural エンベロープから抽出するレスポンス定義の追加(複数形)
func ResponseEnvelopePlural(sourceField *EnvelopePayloadDesc) *EnvelopeType {
	responseEnvelope := &EnvelopeType{
		Form:     PayloadForms.Plural,
		Payloads: []*EnvelopePayloadDesc{sourceField},
	}
	if sourceField.Tags == nil {
		sourceField.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	return responseEnvelope
}
