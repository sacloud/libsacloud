package mapconv

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type dummyTagged struct {
	A          string `mapconv:"ValueA.A"`
	B          string `mapconv:"ValueA.ValueB.B"`
	C          string `mapconv:"ValueA.ValueB.ValueC.C"`
	Pointer    *time.Time
	Slice      []string
	NoTag      string
	unexported string
}

type dummyNaked struct {
	ValueA *struct {
		A      string
		ValueB *struct {
			B      string
			ValueC *struct {
				C string
			}
		}
	}
	Pointer    *time.Time
	Slice      []string
	NoTag      string
	unexported string
}

func TestToNaked(t *testing.T) {
	zeroTime := time.Unix(0, 0)
	expects := []struct {
		tagged *dummyTagged
		naked  *dummyNaked
		err    error
	}{
		{
			tagged: &dummyTagged{
				A:          "A",
				B:          "B",
				C:          "C",
				Pointer:    &zeroTime,
				Slice:      []string{"a", "b", "c"},
				NoTag:      "NoTag",
				unexported: "unexported",
			},
			naked: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				Pointer: &zeroTime,
				Slice:   []string{"a", "b", "c"},
				NoTag:   "NoTag",
			},
		},
	}

	for _, expect := range expects {
		naked := &dummyNaked{}
		err := ConvertTo(expect.tagged, naked)
		require.Equal(t, expect.err, err)
		if err == nil {
			require.Equal(t, expect.naked, naked)
		}
	}

}

func TestFromNaked(t *testing.T) {

	expects := []struct {
		tagged *dummyTagged
		naked  *dummyNaked
		err    error
	}{
		{
			tagged: &dummyTagged{
				A:     "A",
				B:     "B",
				C:     "C",
				NoTag: "NoTag",
			},
			naked: &dummyNaked{
				ValueA: &struct {
					A      string
					ValueB *struct {
						B      string
						ValueC *struct {
							C string
						}
					}
				}{
					A: "A",
					ValueB: &struct {
						B      string
						ValueC *struct {
							C string
						}
					}{
						B: "B",
						ValueC: &struct {
							C string
						}{
							C: "C",
						},
					},
				},
				NoTag: "NoTag",
			},
		},
	}

	for _, expect := range expects {
		tagged := &dummyTagged{}
		err := ConvertFrom(expect.naked, tagged)
		require.Equal(t, expect.err, err)
		if err == nil {
			require.Equal(t, expect.tagged, tagged)
		}
	}

}

type dummySlice struct {
	Slice []*dummySliceInner `json:",omitempty"`
}

type dummySliceInner struct {
	Value string             `json:",omitempty"`
	Slice []*dummySliceInner `json:",omitempty"`
}

type dummyExtractInnerSlice struct {
	Values       []string `json:",omitempty" mapconv:"[]Slice.Value"`
	NestedValues []string `json:",omitempty" mapconv:"[]Slice.[]Slice.Value"`
}

func TestExtractInnerSlice(t *testing.T) {
	expects := []struct {
		input  *dummySlice
		expect *dummyExtractInnerSlice
	}{
		{
			input: &dummySlice{
				Slice: []*dummySliceInner{
					{Value: "value1"},
					{Value: "value2"},
					{
						Value: "value3",
						Slice: []*dummySliceInner{
							{Value: "value4"},
							{Value: "value5"},
						},
					},
				},
			},
			expect: &dummyExtractInnerSlice{
				Values:       []string{"value1", "value2", "value3"},
				NestedValues: []string{"value4", "value5"},
			},
		},
	}

	for _, tc := range expects {
		dest := &dummyExtractInnerSlice{}
		err := ConvertFrom(tc.input, dest)

		require.NoError(t, err)
		require.Equal(t, tc.expect, dest)
	}
}

func TestInsertInnerSlice(t *testing.T) {
	expects := []struct {
		input  *dummyExtractInnerSlice
		expect *dummySlice
	}{
		{
			input: &dummyExtractInnerSlice{
				Values:       []string{"value1", "value2", "value3"},
				NestedValues: []string{"value4", "value5"},
			},
			expect: &dummySlice{
				Slice: []*dummySliceInner{
					{Value: "value1"},
					{Value: "value2"},
					{Value: "value3"},
					{
						Slice: []*dummySliceInner{
							{Value: "value4"},
						},
					},
					{
						Slice: []*dummySliceInner{
							{Value: "value5"},
						},
					},
				},
			},
		},
	}

	for _, tc := range expects {
		dest := &dummySlice{}
		err := ConvertTo(tc.input, dest)

		require.NoError(t, err)
		require.Equal(t, tc.expect, dest)
	}
}

type hasDefaultSource struct {
	Field string `mapconv:"Field:default-value"`
}

type hasDefaultDest struct {
	Field string
}

func TestDefaultValue(t *testing.T) {
	expects := []struct {
		input  *hasDefaultSource
		expect *hasDefaultDest
	}{
		{
			input: &hasDefaultSource{},
			expect: &hasDefaultDest{
				Field: "default-value",
			},
		},
	}

	for _, tc := range expects {
		dest := &hasDefaultDest{}
		err := ConvertTo(tc.input, dest)
		require.NoError(t, err)
		require.Equal(t, tc.expect, dest)
	}
}

type multipleSource struct {
	Field string `mapconv:"Field1,Field2"`
}

type multipleDest struct {
	Field1 string
	Field2 string
}

func TestMultipleDestination(t *testing.T) {
	expects := []struct {
		input  *multipleSource
		expect *multipleDest
	}{
		{
			input: &multipleSource{
				Field: "value",
			},
			expect: &multipleDest{
				Field1: "value",
				Field2: "value",
			},
		},
	}

	for _, tc := range expects {
		dest := &multipleDest{}
		err := ConvertTo(tc.input, dest)
		require.NoError(t, err)
		require.Equal(t, tc.expect, dest)
	}
}
