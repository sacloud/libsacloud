package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKey(t *testing.T) {
	cases := []struct {
		input  Key
		expect string
	}{
		{
			input: Key{
				Field: "field",
				Op:    OpEqual,
			},
			expect: "field",
		},
		{
			input: Key{
				Field: "field",
				Op:    OpGreaterThan,
			},
			expect: "field>",
		},
		{
			input: Key{
				Field: "field",
				Op:    OpGreaterEqual,
			},
			expect: "field>=",
		},
		{
			input: Key{
				Field: "field",
				Op:    OpLessThan,
			},
			expect: "field<",
		},
		{
			input: Key{
				Field: "field",
				Op:    OpLessEqual,
			},
			expect: "field<=",
		},
		{
			input: Key{
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
