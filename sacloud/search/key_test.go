package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKey(t *testing.T) {
	cases := []struct {
		input  FilterKey
		expect string
	}{
		{
			input: FilterKey{
				Field: "field",
				Op:    OpEqual,
			},
			expect: "field",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpGreaterThan,
			},
			expect: "field>",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpGreaterEqual,
			},
			expect: "field>=",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpLessThan,
			},
			expect: "field<",
		},
		{
			input: FilterKey{
				Field: "field",
				Op:    OpLessEqual,
			},
			expect: "field<=",
		},
		{
			input: FilterKey{
				Field: "another-field-name",
				Op:    OpEqual,
			},
			expect: "another-field-name",
		},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expect, tc.input.String())
	}
}
