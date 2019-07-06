package schema

import "fmt"

// MappableArgument 引数定義の追加
func MappableArgument(name string, model *Model, destField string) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: fmt.Sprintf("%s,recursive", destField),
	}
}

// PassthroughModelArgument 引数定義の追加、ペイロードの定義も同時に行われる
func PassthroughModelArgument(name string, model *Model) *Argument {
	return &Argument{
		Name:       name,
		Type:       model,
		MapConvTag: ",squash",
	}
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
