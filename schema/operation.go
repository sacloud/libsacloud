package schema

import (
	"net/http"
	"strings"
)

// ResponseType Operationで呼び出すAPIからのレスポンスタイプ
type ResponseType int

const (
	// ResponseTypeSingular 単数形レスポンス
	ResponseTypeSingular ResponseType = iota
	// ResponseTypePlural 複数形レスポンス
	ResponseTypePlural
	// ResponseTypeCustom カスタムレスポンス
	ResponseTypeCustom
)

// Operation リソースへの操作
type Operation struct {
	Name         string         // 操作名、メソッド名となる
	Method       string         // HTTPリクエストメソッド GET/POST/PUT/DELETE
	PathFormat   string         // パスのフォーマット、省略した場合はDefaultPathFormatが設定される
	Arguments    []*Value       // パラメータ
	Results      []*ResultValue // レスポンス
	ResponseType ResponseType
}

// IsSingular レスポンスが単数形データか
func (o *Operation) IsSingular() bool {
	return o.ResponseType == ResponseTypeSingular
}

// IsPlural レスポンスが複数形データか
func (o *Operation) IsPlural() bool {
	return o.ResponseType == ResponseTypePlural
}

// GetPathFormat パスのフォーマット
func (o *Operation) GetPathFormat() string {
	if o.PathFormat != "" {
		return o.PathFormat
	}
	return DefaultPathFormat
}

// ImportStatements コード生成時に利用するimport文を生成する
func (o *Operation) ImportStatements() []string {
	ss := []string{`"context"`}
	var values []*Value
	values = append(values, o.Arguments...)
	for _, v := range o.Results {
		values = append(values, v.Value)
	}
	for _, value := range values {
		s := value.ImportStatement()
		if s != "" {
			ss = append(ss, s)
		}
	}

	return uniqStrings(ss)
}

// ReturnErrorStatement コード生成時に利用する、エラーをreturnする文を生成する
func (o *Operation) ReturnErrorStatement() string {
	ss := make([]string, len(o.Results))
	for i, res := range o.Results {
		s := res.ZeroValueOnSource()
		if res.Type == TypeError {
			s = res.Name
		}
		ss[i] = s
	}
	return strings.Join(ss, ",")
}

// ResultsWithoutError Operationに設定されているResultsのうち、エラー型以外を返す
func (o *Operation) ResultsWithoutError() []*ResultValue {
	var rs []*ResultValue
	for _, r := range o.Results {
		if r.Type != TypeError {
			rs = append(rs, r)
		}
	}
	return rs
}

// TypePtrArguments TypePtrな引数のリストを返す
func (o *Operation) TypePtrArguments() []*Value {
	var rs []*Value
	for _, r := range o.Arguments {
		if r.Type == TypePtr {
			rs = append(rs, r)
		}
	}
	return rs
}

// HasTypePtrArguments TypePtrな引数を持っているか
func (o *Operation) HasTypePtrArguments() bool {
	return len(o.TypePtrArguments()) > 0
}

// OperationType 操作タイプ、この値により生成されるコードが変わる
type OperationType int

const (
	// OperationTypeFind Find操作、検索条件をパラメータに取り、複数リソースを返す
	OperationTypeFind OperationType = iota
	// OperationTypeCreate Create操作 作成内容をパラメータに取り、作成後のリソースを返す
	OperationTypeCreate
	// OperationTypeRead Read操作、IDをパラメータに取り、単一リソースを返す
	OperationTypeRead
	// OperationTypeUpdate Update操作、IDと変更内容をパラメータに取り、変更後のリソースを返す
	OperationTypeUpdate
	// OperationTypeDelete Delete操作、IDをパラメータにとる
	OperationTypeDelete
	// OperationTypeManual マニュアル操作、雛形だけコード生成され、具体的な処理は各自で実装する
	OperationTypeManual
)

// CreateOperationParam CreateOperation()に渡すパラメータ
type CreateOperationParam struct {
	FieldName       string
	ParameterStruct interface{}
	ResponseStruct  interface{}
}

// CreateOperation Create(POST)操作を示すOperationを作成して返す
func CreateOperation(p *CreateOperationParam) *Operation {
	return &Operation{
		Name:       "Create",
		Method:     http.MethodPost,
		PathFormat: DefaultPathFormat,
		Arguments: []*Value{
			ValueZone,
			{
				Name:      "param",
				Type:      TypePtr,
				Struct:    p.ParameterStruct,
				FieldName: p.FieldName,
			},
		},
		Results: []*ResultValue{
			{
				Value: &Value{
					Type:      TypePtr,
					Struct:    p.ResponseStruct,
					FieldName: p.FieldName,
				},
				NakedTypeName: p.FieldName,
				Comment:       "対象リソース",
			},
			ValueError,
		},
	}
}

// ReadOperationParam ReadOperation()に渡すパラメータ
type ReadOperationParam struct {
	FieldName      string
	ResponseStruct interface{}
}

// ReadOperation Read(GET)操作を示すOperationを作成して返す
func ReadOperation(p *ReadOperationParam) *Operation {
	return &Operation{
		Name:       "Read",
		Method:     http.MethodGet,
		PathFormat: DefaultPathFormat + "/{{.id}}",
		Arguments: []*Value{
			ValueZone,
			ValueID,
		},
		Results: []*ResultValue{
			{
				Value: &Value{
					Type:      TypePtr,
					Struct:    p.ResponseStruct,
					FieldName: p.FieldName,
				},
				NakedTypeName: p.FieldName,
				Comment:       "対象リソース",
			},
			ValueError,
		},
	}
}

// UpdateOperationParam UpdateOperation()に渡すパラメータ
type UpdateOperationParam struct {
	FieldName       string
	ParameterStruct interface{}
	ResponseStruct  interface{}
}

// UpdateOperation Update(PUT)操作を示すOperationを作成して返す
func UpdateOperation(p *UpdateOperationParam) *Operation {
	return &Operation{
		Name:       "Update",
		Method:     http.MethodPut,
		PathFormat: DefaultPathFormat + "/{{.id}}",
		Arguments: []*Value{
			ValueZone,
			ValueID,
			{
				Name:      "param",
				Type:      TypePtr,
				Struct:    p.ParameterStruct,
				FieldName: p.FieldName,
			},
		},
		Results: []*ResultValue{
			{
				Value: &Value{
					Type:      TypePtr,
					Struct:    p.ResponseStruct,
					FieldName: p.FieldName,
				},
				NakedTypeName: p.FieldName,
				Comment:       "対象リソース",
			},
			ValueError,
		},
	}
}

// DeleteOperation Delete(DELETE)操作を示すOperationを作成して返す
func DeleteOperation() *Operation {
	return &Operation{
		Name:       "Delete",
		Method:     http.MethodDelete,
		PathFormat: DefaultPathFormat + "/{{.id}}",
		Arguments: []*Value{
			ValueZone,
			ValueID,
		},
		Results: []*ResultValue{
			ValueError,
		},
	}
}
