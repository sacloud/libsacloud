package accessor

import (
	"net/url"
	"strings"
	"time"
)

// Filter 検索フィルタ
type Filter interface {
	GetFilter() map[string]interface{}
	SetFilter(v map[string]interface{})
}

// FilterOperator フィルター演算子
type FilterOperator int

const (
	// OpEqual =
	OpEqual FilterOperator = iota
	// OpGreaterThan >
	OpGreaterThan
	// OpGreaterEqual >=
	OpGreaterEqual
	// OpLessThan <
	OpLessThan
	// OpLessEqual <=
	OpLessEqual
)

// ClearFilter 設定されたフィルタをクリアします
func ClearFilter(f Filter) {
	f.SetFilter(map[string]interface{}{})
}

// SetANDFilterWithPartialMatch 指定キーの値に中間一致するフィルタを設定 複数指定した場合はAND条件となる
//
// patternsには文字列のみ指定可能
// keyが存在していた場合は上書きされる
func SetANDFilterWithPartialMatch(f Filter, key string, patterns []string) {
	for i := range patterns {
		patterns[i] = convertToValidFilterValue(patterns[i]).(string)
	}

	filter := f.GetFilter()
	if filter == nil {
		filter = map[string]interface{}{}
	}
	filter[key] = strings.Join(patterns, "%20")
	f.SetFilter(filter)
}

// SetORFilterWithExactMatch 指定キーの値に完全一致するフィルタを設定 複数指定した場合はor条件となる
//
// patternsには文字列のみ指定可能
// keyが存在していた場合は上書きされる
func SetORFilterWithExactMatch(f Filter, key string, patterns []string) {
	for i := range patterns {
		patterns[i] = convertToValidFilterValue(patterns[i]).(string)
	}

	filter := f.GetFilter()
	if filter == nil {
		filter = map[string]interface{}{}
	}
	filter[key] = patterns
	f.SetFilter(filter)
}

// SetNumericFilter 数値型フィルタの設定
func SetNumericFilter(f Filter, key string, op FilterOperator, value int64) {
	filter := f.GetFilter()
	if filter == nil {
		filter = map[string]interface{}{}
	}
	filter[filterKeyWithOp(key, op)] = value
	f.SetFilter(filter)
}

// SetTimeFilter Time型フィルタの設定
func SetTimeFilter(f Filter, key string, op FilterOperator, value time.Time) {
	filter := f.GetFilter()
	if filter == nil {
		filter = map[string]interface{}{}
	}
	filter[filterKeyWithOp(key, op)] = convertToValidFilterValue(value)
	f.SetFilter(filter)
}

func filterKeyWithOp(key string, op FilterOperator) string {
	switch op {
	case OpGreaterThan:
		return key + ">"
	case OpGreaterEqual:
		return key + ">="
	case OpLessThan:
		return key + "<"
	case OpLessEqual:
		return key + "<="
	}
	return key
}

func convertToValidFilterValue(v interface{}) interface{} {
	switch v := v.(type) {
	case time.Time:
		return v.Format(time.RFC3339)
	case string:
		return escapeFilterString(v)
	}
	return v
}

func escapeFilterString(s string) string {
	//HACK さくらのクラウド側でqueryStringでの+エスケープに対応していないため、
	// %20にエスケープされるurl.Pathを利用する。
	// http://qiita.com/shibukawa/items/c0730092371c0e243f62
	u := &url.URL{Path: s}
	return u.String()
}
