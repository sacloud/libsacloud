package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Filter 検索系APIでの検索条件
//
// Note: libsacloudではリクエスト時に`X-Sakura-Bigint-As-Int`ヘッダを指定することで
// 文字列で表されているBitintをintとして取得している。
// このため、libsacloud側では数値型に見える項目でもさくらのクラウド側では文字列となっている場合がある。
// これらの項目ではOpEqual以外の演算子は利用できない。
// また、これらの項目でスカラ値を検索条件に与えた場合は部分一致ではなく完全一致となるため注意。
type Filter map[FilterKey]interface{}

// MarshalJSON 検索系APIコール時のGETパラメータを出力するためのjson.Marshaler実装
func (f Filter) MarshalJSON() ([]byte, error) {
	result := make(map[string]interface{})

	for key, expression := range f {
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

		result[key.String()] = exp
	}

	return json.Marshal(result)
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
