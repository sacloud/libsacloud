package accessor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type dummyFilterAccessor struct {
	filter map[string]interface{}
}

func (d *dummyFilterAccessor) GetFilter() map[string]interface{} {
	return d.filter
}

func (d *dummyFilterAccessor) SetFilter(v map[string]interface{}) {
	d.filter = v
}

func TestFilter_SetANDFilterWithPartialMatch(t *testing.T) {
	cases := []struct {
		input  []string
		expect interface{}
	}{
		{
			input:  []string{""},
			expect: "",
		},
		{
			input:  []string{"v1", "v2"},
			expect: "v1%20v2",
		},
		{
			input:  []string{"v1", "v2", "v3"},
			expect: "v1%20v2%20v3",
		},
		{
			input:  []string{"v1", " ", "v2"},
			expect: "v1%20%20%20v2",
		},
	}

	for _, tc := range cases {
		df := &dummyFilterAccessor{}
		SetANDFilterWithPartialMatch(df, "key", tc.input)
		require.Equal(t, tc.expect, df.filter["key"])
	}
}

func TestFilter_SetORFilterWithExactMatch(t *testing.T) {
	cases := []struct {
		input  []string
		expect interface{}
	}{
		{
			input:  []string{""},
			expect: []string{""},
		},
		{
			input:  []string{"tk1a", "is1b"},
			expect: []string{"tk1a", "is1b"},
		},
	}

	for _, tc := range cases {
		df := &dummyFilterAccessor{}
		SetORFilterWithExactMatch(df, "key", tc.input)
		require.Equal(t, tc.expect, df.filter["key"])
	}
}

func TestFilter_SetNumericFilter(t *testing.T) {
	cases := []struct {
		input       int64
		operator    FilterOperator
		expectKey   string
		expectValue interface{}
	}{
		{
			input:       int64(0),
			operator:    OpEqual,
			expectKey:   "key",
			expectValue: int64(0),
		},
		{
			input:       int64(0),
			operator:    OpGreaterThan,
			expectKey:   "key>",
			expectValue: int64(0),
		},
		{
			input:       int64(0),
			operator:    OpGreaterEqual,
			expectKey:   "key>=",
			expectValue: int64(0),
		},
		{
			input:       int64(0),
			operator:    OpLessThan,
			expectKey:   "key<",
			expectValue: int64(0),
		},
		{
			input:       int64(0),
			operator:    OpLessEqual,
			expectKey:   "key<=",
			expectValue: int64(0),
		},
		{
			input:       int64(1),
			operator:    OpEqual,
			expectKey:   "key",
			expectValue: int64(1),
		},
	}

	for _, tc := range cases {
		df := &dummyFilterAccessor{}
		SetNumericFilter(df, "key", tc.operator, tc.input)
		actual, ok := df.filter[tc.expectKey]
		require.True(t, ok)
		require.Equal(t, tc.expectValue, actual)
	}
}

func TestFilter_SetTimeFilter(t *testing.T) {
	targetTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	targetTimeString := targetTime.Format(time.RFC3339)

	cases := []struct {
		input       time.Time
		operator    FilterOperator
		expectKey   string
		expectValue interface{}
	}{
		{
			input:       targetTime,
			operator:    OpEqual,
			expectKey:   "key",
			expectValue: targetTimeString,
		},
		{
			input:       targetTime,
			operator:    OpGreaterThan,
			expectKey:   "key>",
			expectValue: targetTimeString,
		},
		{
			input:       targetTime,
			operator:    OpGreaterEqual,
			expectKey:   "key>=",
			expectValue: targetTimeString,
		},
		{
			input:       targetTime,
			operator:    OpLessThan,
			expectKey:   "key<",
			expectValue: targetTimeString,
		},
		{
			input:       targetTime,
			operator:    OpLessEqual,
			expectKey:   "key<=",
			expectValue: targetTimeString,
		},
		{
			input:       targetTime,
			operator:    OpEqual,
			expectKey:   "key",
			expectValue: targetTimeString,
		},
	}

	for _, tc := range cases {
		df := &dummyFilterAccessor{}
		SetTimeFilter(df, "key", tc.operator, tc.input)
		actual, ok := df.filter[tc.expectKey]
		require.True(t, ok)
		require.Equal(t, tc.expectValue, actual)
	}
}

func TestFilter_ClearFilter(t *testing.T) {
	df := &dummyFilterAccessor{}
	SetNumericFilter(df, "key", OpEqual, 0)

	require.Len(t, df.GetFilter(), 1)

	ClearFilter(df)
	require.Len(t, df.GetFilter(), 0)
}
