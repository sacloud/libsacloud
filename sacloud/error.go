package sacloud

import (
	"fmt"
	"net/url"
)

// APIErrorResponse APIエラー型
type APIErrorResponse struct {
	IsFatal      bool   `json:"is_fatal,omitempty"`   // IsFatal
	Serial       string `json:"serial,omitempty"`     // Serial
	Status       string `json:"status,omitempty"`     // Status
	ErrorCode    string `json:"error_code,omitempty"` // ErrorCode
	ErrorMessage string `json:"error_msg,omitempty"`  // ErrorMessage
}

// APIError APIコール時のエラー情報
type APIError interface {
	// errorインターフェースを内包
	error

	// エラー発生時のレスポンスコード
	ResponseCode() int

	// エラーコード
	Code() string

	// エラー発生時のメッセージ
	Message() string

	// エラー追跡用シリアルコード
	Serial() string

	// エラー(オリジナル)
	OrigErr() *APIErrorResponse
}

// NewAPIError APIコール時のエラー情報
func NewAPIError(requestMethod string, requestURL *url.URL, requestBody string, responseCode int, err *APIErrorResponse) APIError {
	return &apiError{
		responseCode: responseCode,
		method:       requestMethod,
		url:          requestURL,
		body:         requestBody,
		origErr:      err,
	}
}

type apiError struct {
	responseCode int
	method       string
	url          *url.URL
	body         string
	origErr      *APIErrorResponse
}

// Error errorインターフェース
func (e *apiError) Error() string {
	return fmt.Sprintf("Error in response: %#v", e.origErr)
}

// ResponseCode エラー発生時のレスポンスコード
func (e *apiError) ResponseCode() int {
	return e.responseCode
}

// Code エラーコード
func (e *apiError) Code() string {
	if e.origErr != nil {
		return e.origErr.ErrorCode
	}
	return ""
}

// Message エラー発生時のメッセージ(
func (e *apiError) Message() string {
	if e.origErr != nil {
		return e.origErr.ErrorMessage
	}
	return ""
}

// Serial エラー追跡用シリアルコード
func (e *apiError) Serial() string {
	if e.origErr != nil {
		return e.origErr.Serial
	}
	return ""
}

// OrigErr エラー(オリジナル)
func (e *apiError) OrigErr() *APIErrorResponse {
	return e.origErr
}
