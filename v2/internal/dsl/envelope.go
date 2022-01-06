// Copyright 2016-2022 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dsl

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
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

// EnvelopePayloadDesc エンベロープに含まれるペイロードの情報
type EnvelopePayloadDesc struct {
	Name string    // ペイロードのフィールド名
	Type meta.Type // ペイロードの型情報
	Tags *FieldTags
}

// TypeName ペイロードの型定義
func (d *EnvelopePayloadDesc) TypeName() string {
	return d.Type.GoTypeSourceCode()
}

// TagString タグの文字列表現
func (d *EnvelopePayloadDesc) TagString() string {
	if d.Tags == nil {
		d.Tags = &FieldTags{
			JSON: ",omitempty",
		}
	}
	tags := d.Tags.String()
	if tags == "" {
		return tags
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

	ret.Payloads = append(ret.Payloads, descs...)

	return ret
}

// RequestEnvelopeFromModel モデル定義からリクエストのエンベロープを作成する
func RequestEnvelopeFromModel(model *Model) *EnvelopeType {
	var descs []*EnvelopePayloadDesc
	for _, field := range model.Fields {
		t := field.Type
		if m, ok := t.(*Model); ok {
			if m.HasNakedType() {
				t = m.NakedType
			}
		}
		descs = append(descs, &EnvelopePayloadDesc{
			Name: field.Name,
			Type: t,
			Tags: field.Tags,
		})
	}
	ret := &EnvelopeType{
		Form: PayloadForms.Singular,
	}

	ret.Payloads = append(ret.Payloads, descs...)

	return ret
}

// ResponseEnvelope エンベロープから抽出するレスポンス定義の追加
func ResponseEnvelope(sourceFields ...*EnvelopePayloadDesc) *EnvelopeType {
	return responseEnvelope(PayloadForms.Singular, sourceFields...)
}

// ResponseEnvelopePlural エンベロープから抽出するレスポンス定義の追加(複数形)
func ResponseEnvelopePlural(sourceFields ...*EnvelopePayloadDesc) *EnvelopeType {
	return responseEnvelope(PayloadForms.Plural, sourceFields...)
}

func responseEnvelope(form PayloadForm, sourceFields ...*EnvelopePayloadDesc) *EnvelopeType {
	et := &EnvelopeType{
		Form: form,
	}
	for _, sf := range sourceFields {
		if sf.Tags == nil {
			sf.Tags = &FieldTags{
				JSON: ",omitempty",
			}
		}
		et.Payloads = append(et.Payloads, sf)
	}
	return et
}
