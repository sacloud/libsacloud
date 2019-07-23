package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Filter 検索系APIでの検索条件
//
// Note: libsacloudではリクエスト時に`X-Sakura-Bigint-As-Int`ヘッダを指定することで
// 文字列で表されているBitintをintとして取得している。
// このため、libsacloud側では数値型に見える項目でもさくらのクラウド側では文字列となっている場合がある。
// これらの項目ではOpEqual以外の演算子は利用できない。
// また、これらの項目でスカラ値を検索条件に与えた場合は部分一致ではなく完全一致となるため注意。
type Filter Criteria

// Criteria 検索条件
type Criteria []Criterion

// MarshalJSON 検索系APIコール時のGETパラメータを出力するためのjson.Marshaler実装
func (f Filter) MarshalJSON() ([]byte, error) {
	var results []string

	for _, item := range f {
		key := item.Key
		expression := item.Value

		if expression == nil {
			continue
		}

		exp := expression
		switch key.Op {
		case OpEqual:
			if _, ok := exp.(*EqualExpression); !ok {
				exp = OrEqual(expression)
			}
		default:
			exp = convertToValidFilterCondition(exp)
		}

		marshaled, err := json.Marshal(map[string]interface{}{key.String(): exp})
		if err != nil {
			return nil, err
		}

		result := strings.Trim(string(marshaled), "{}")
		results = append(results, result)
	}

	return []byte(fmt.Sprintf("{%s}", strings.Join(results, ","))), nil
}

// Add 項目追加
func (f *Filter) Add(item Criterion) {
	*f = append(*f, item)
}

// AddNew 項目を作成して追加
func (f *Filter) AddNew(key FilterKey, value interface{}) {
	f.Add(Criterion{Key: key, Value: value})
}

// Criterion フィルタの項目
type Criterion struct {
	// Key キー
	Key FilterKey
	// Value 検索条件値
	Value interface{}
}

func convertToValidFilterCondition(v interface{}) string {
	switch v := v.(type) {
	case time.Time:
		return v.Format(time.RFC3339)
	case string:
		return escapeFilterString(v)
	}
	return fmt.Sprintf("%v", v)
}

func escapeFilterString(s string) string {
	//HACK さくらのクラウド側でqueryStringでの+エスケープに対応していないため、
	// %20にエスケープされるurl.Pathを利用する。
	// http://qiita.com/shibukawa/items/c0730092371c0e243f62
	u := &url.URL{Path: s}
	return u.String()
}
