// Copyright 2016-2019 The Libsacloud Authors
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
	"strings"
)

// Operations リソースへの操作(スライス)
type Operations []*Operation

// Operation リソースへの操作
type Operation struct {
	ResourceName        string
	Name                string        // 操作名、メソッド名となる
	Method              string        // HTTPリクエストメソッド GET/POST/PUT/DELETE
	PathFormat          string        // パスのフォーマット、省略した場合はDefaultPathFormatが設定される
	Arguments           Arguments     // 引数の定義
	Results             Results       // レスポンス
	RequestEnvelope     *EnvelopeType // リクエスト時のエンベロープ
	ResponseEnvelope    *EnvelopeType // レスポンス時のエンベロープ
	UseWrappedResult    bool          // trueの場合APIからの戻り値としてラッパー型(xxxResult)を返す
	LockLevel           LockLevel     // APIコール時のロックレベル
	LockKeyCustomFormat string        // ロックキーのgoテンプレートフォーマット(PathFormatと同じパラメータが利用可能)
	// IsPatch Patchメソッドであるかのフラグ
	//
	// この値に応じてコード生成の挙動を切り替える
	// [REMARK] 他にもコード生成のカスタマイズが必要になったらこの項目の持ち方を再考する
	IsPatch bool
}

// LockKeyFormat ロックキーのフォーマット、ロックなしの場合空になる
func (o *Operation) LockKeyFormat() string {
	if o.LockLevel == LockLevelNone {
		return ""
	}
	if o.LockKeyCustomFormat != "" {
		return o.LockKeyCustomFormat
	}
	switch o.LockLevel {
	case LockLevelResource:
		return o.GetPathFormat()
	case LockLevelAPI:
		return fmt.Sprintf("%s.%s.%s", o.ResourceName, o.Name, o.Method)
	case LockLevelGlobal:
		return "GlobalLock"
	default:
		return ""
	}
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
		if o.UseWrappedResult {
			return "nil, err"
		}
		var ret string
		for range o.Results {
			ret += "nil,"
		}
		ret += "err"
		return ret
	}
	return "err"
}

// ReturnStatement コード生成時に利用するreturn部分を生成する
func (o *Operation) ReturnStatement() string {
	if !o.HasResults() {
		return "err"
	}
	if o.UseWrappedResult {
		return "results, err"
	}
	var ret string
	for _, r := range o.Results {
		ret += fmt.Sprintf("results.%s,", r.DestField)
	}
	ret += "nil"
	return ret
}

// ResultsStatement 戻り値定義部のソースを出力
func (o *Operation) ResultsStatement() string {
	if !o.HasResults() {
		return "error"
	}
	if o.UseWrappedResult {
		return fmt.Sprintf("(%s, error)", o.resultType().GoTypeSourceCode())
	}
	var ret string
	for _, r := range o.Results {
		ret += r.GoTypeSourceCode() + ","
	}
	return fmt.Sprintf("(%s error)", ret)
}

// ResultsTypeInfo 戻り値の型情報(error型を含まない)
func (o *Operation) ResultsTypeInfo() []*ResultTypeInfo {
	var info []*ResultTypeInfo
	if !o.HasResults() {
		return info
	}
	if o.UseWrappedResult {
		info = append(info, &ResultTypeInfo{
			VarName:   "result",
			FieldName: "Result",
			Type:      o.resultType().Type(),
		})
		return info
	}
	for _, r := range o.Results {
		info = append(info, &ResultTypeInfo{
			VarName:   "result" + r.DestField,
			FieldName: r.DestField,
			Type:      r.Type(),
		})
	}
	return info
}

// ResultsAssignStatement API戻り値を変数にアサインする場合の変数部分のソースを出力、主にtraceで利用する
func (o *Operation) ResultsAssignStatement() string {
	if !o.HasResults() {
		return "err"
	}
	if o.UseWrappedResult {
		return "result, err"
	}
	var ret string
	for _, r := range o.Results {
		ret += fmt.Sprintf(",result%s", r.DestField)
	}
	return fmt.Sprintf("%s, err", ret)
}

// RequestEnvelopeStructName エンベロープのstruct名(camel-case)
func (o *Operation) RequestEnvelopeStructName() string {
	return fmt.Sprintf("%s%sRequestEnvelope", firstRuneToLower(o.ResourceName), o.Name)
}

// ResponseEnvelopeStructName エンベロープのstruct名(camel-case)
func (o *Operation) ResponseEnvelopeStructName() string {
	return fmt.Sprintf("%s%sResponseEnvelope", firstRuneToLower(o.ResourceName), o.Name)
}

// ResultTypeName API戻り値の型名
func (o *Operation) ResultTypeName() string {
	if o.UseWrappedResult {
		return o.resultType().GoType()
	}
	return firstRuneToLower(o.resultType().GoType())

}

// HasResults 戻り値が定義されているかを取得
func (o *Operation) HasResults() bool {
	return len(o.Results) > 0
}

// StubFieldDefines スタブ生成時のフィールド定義文を全フィールド分出力
func (o *Operation) StubFieldDefines() []string {
	if !o.HasResults() {
		return nil
	}
	if o.UseWrappedResult {
		return []string{fmt.Sprintf("Values %s", o.resultType().GoTypeSourceCode())}
	}
	var rets []string
	for _, r := range o.Results {
		rets = append(rets, fmt.Sprintf("%s %s", r.DestField, r.GoTypeSourceCode()))
	}
	return rets
}

// StubReturnStatement スタブ生成時のreturn文
func (o *Operation) StubReturnStatement(receiverName string) string {
	if !o.HasResults() {
		return fmt.Sprintf("return %s.%sStubResult.Err", receiverName, o.MethodName())
	}
	var rets []string
	if o.UseWrappedResult {
		rets = append(rets, fmt.Sprintf("%s.%sStubResult.Values", receiverName, o.MethodName()))
	} else {
		for _, r := range o.Results {
			rets = append(rets, fmt.Sprintf("%s.%sStubResult.%s", receiverName, o.MethodName(), r.DestField))
		}
	}

	rets = append(rets, fmt.Sprintf("%s.%sStubResult.Err", receiverName, o.MethodName()))
	return fmt.Sprintf("return %s", strings.Join(rets, ","))
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

func (o *Operation) resultType() *ResultType {
	return &ResultType{
		resourceName: o.ResourceName,
		operation:    o,
		results:      o.Results,
	}
}

// PatchArgument Patchメソッドで操作の対象とする引数
//
// ID以外で最初に現れた引数を対象とする。
// (Patchメソッドは複数の引数を取らない前提での実装)
func (o *Operation) PatchArgument() *Argument {
	for _, arg := range o.Arguments {
		if arg != ArgumentID {
			if _, ok := arg.Type.(*Model); ok {
				return arg
			}
		}
	}
	panic(fmt.Errorf("operation %q doesn't have the Patch Argument", o.Name))
}

// PatchArgumentModel Patchメソッドで操作の対象とする引数のモデル
//
// PatchArgumentのコメントも参照すること
func (o *Operation) PatchArgumentModel() *Model {
	arg := o.PatchArgument()
	m, ok := arg.Type.(*Model)
	if !ok {
		panic(fmt.Errorf("operation %q doesn't have the Patch Argument", o.Name))
	}
	return m
}
