package search

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type inputKeyValue struct {
	key       FilterKey
	condition interface{}
}

func TestFilter(t *testing.T) {

	loc := time.FixedZone("Asia/Tokyo", 9*60*60)

	cases := []struct {
		conditions []*inputKeyValue
		expect     string
	}{
		// default is OpEqual, OR match
		{
			conditions: []*inputKeyValue{
				{
					key:       Key("field"), // same as KeyWithOp("fields", OpEqual)
					condition: "value",
				},
			},
			expect: `{"field":["value"]}`,
		},
		// with comparison operator
		{
			conditions: []*inputKeyValue{
				{
					key:       KeyWithOp("field", OpLessEqual),
					condition: "1",
				},
			},
			expect: `{"field\u003c=":"1"}`,
		},
		// with EqualExpression
		{
			conditions: []*inputKeyValue{
				{
					key:       Key("field"),
					condition: AndEqual("value1", "value2"),
				},
			},
			expect: `{"field":"value1%20value2"}`,
		},
		// multiple keys(AND)
		{
			conditions: []*inputKeyValue{
				{
					key:       Key("field1"),
					condition: "value1",
				},
				{
					key:       Key("field2"),
					condition: "value2",
				},
			},
			expect: `{"field1":["value1"],"field2":["value2"]}`,
		},
		// array values(AND)
		{
			conditions: []*inputKeyValue{
				{
					key:       Key("field1"),
					condition: []string{"value1", "value2"},
				},
			},
			expect: `{"field1":[["value1","value2"]]}`,
		},
		// multiple keys with same key, different operator
		{
			conditions: []*inputKeyValue{
				{
					key:       KeyWithOp("field", OpLessEqual),
					condition: "1",
				},
				{
					key:       KeyWithOp("field", OpGreaterThan),
					condition: "2",
				},
			},
			expect: `{"field\u003c=":"1","field\u003e":"2"}`,
		},
		// example of same as API document - https://developer.sakura.ad.jp/cloud/api/1.1/
		{
			conditions: []*inputKeyValue{
				{
					key:       Key("Name"),
					condition: AndEqual("test", "example"),
				},
				{
					key:       Key("Zone.Name"),
					condition: OrEqual("is1a", "is1b"),
				},
				{
					key:       KeyWithOp("CreatedAt", OpLessThan),
					condition: time.Date(2011, 9, 1, 0, 0, 0, 0, loc),
				},
			},
			expect: `{"CreatedAt\u003c":"2011-09-01T00:00:00+09:00","Name":"test%20example","Zone.Name":["is1a","is1b"]}`,
		},
	}

	for _, tc := range cases {
		filter := Filter{}
		for _, kv := range tc.conditions {
			filter[kv.key] = kv.condition
		}

		data, err := json.Marshal(filter)
		require.NoError(t, err)
		require.Equal(t, tc.expect, string(data))
	}

}
