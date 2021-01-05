// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mapconv

import (
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

type dummyFrom struct {
	A          string `mapconv:"ValueA.A"`
	B          string `mapconv:"ValueA.ValueB.B"`
	C          string `mapconv:"ValueA.ValueB.ValueC.C"`
	Ignore     string `mapconv:"-"`
	Pointer    *time.Time
	Slice      []string
	NoTag      string
	Bool       bool
	unexported string
}

type dummyTo struct {
	ValueA *struct {
		A      string
		ValueB *struct {
			B      string
			ValueC *struct {
				C string
			}
		}
	}
	Ignore  string
	Pointer *time.Time
	Slice   []string
	NoTag   string
	Bool    bool
}

func TestConvertTo(t *testing.T) {
	zeroTime := time.Unix(0, 0)
	tests := []struct {
		input  *dummyFrom
		output *dummyTo
		err    error
	}{
		{
			input: &dummyFrom{
				A:          "A",
				B:          "B",
				C:          "C",
				Ignore:     "ignored",
				Pointer:    &zeroTime,
				Slice:      []string{"a", "b", "c"},
				NoTag:      "NoTag",
				Bool:       true,
				unexported: "unexported",
			},
			output: &dummyTo{
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
				Bool:    true,
			},
		},
	}

	for _, tt := range tests {
		output := &dummyTo{}
		err := ConvertTo(tt.input, output)
		require.Equal(t, tt.err, err)
		if err == nil {
			require.EqualValues(t, tt.output.ValueA, output.ValueA)
			require.EqualValues(t, tt.output.Pointer.String(), output.Pointer.String())
			require.EqualValues(t, tt.output.Slice, output.Slice)
			require.EqualValues(t, tt.output.NoTag, output.NoTag)
		}
	}
}

func TestConvertFrom(t *testing.T) {
	tests := []struct {
		output *dummyFrom
		input  *dummyTo
		err    error
	}{
		{
			output: &dummyFrom{
				A:     "A",
				B:     "B",
				C:     "C",
				NoTag: "NoTag",
				Bool:  true,
			},
			input: &dummyTo{
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
				Bool:  true,
			},
		},
	}

	for _, tt := range tests {
		output := &dummyFrom{}
		err := ConvertFrom(tt.input, output)
		require.Equal(t, tt.err, err)
		if err == nil {
			require.Equal(t, tt.output, output)
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
	tests := []struct {
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

	for _, tt := range tests {
		output := &dummyExtractInnerSlice{}
		err := ConvertFrom(tt.input, output)

		require.NoError(t, err)
		require.Equal(t, tt.expect, output)
	}
}

func TestInsertInnerSlice(t *testing.T) {
	tests := []struct {
		input  *dummyExtractInnerSlice
		output *dummySlice
	}{
		{
			input: &dummyExtractInnerSlice{
				Values:       []string{"value1", "value2", "value3"},
				NestedValues: []string{"value4", "value5"},
			},
			output: &dummySlice{
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

	for _, tt := range tests {
		output := &dummySlice{}
		err := ConvertTo(tt.input, output)

		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type hasDefaultSource struct {
	Field string `mapconv:"Field,default=default-value"`
}

type hasDefaultDest struct {
	Field string
}

func TestDefaultValue(t *testing.T) {
	tests := []struct {
		input  *hasDefaultSource
		output *hasDefaultDest
	}{
		{
			input: &hasDefaultSource{},
			output: &hasDefaultDest{
				Field: "default-value",
			},
		},
	}

	for _, tt := range tests {
		output := &hasDefaultDest{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type multipleSource struct {
	Field string `mapconv:"Field1/Field2"`
}

type multipleDest struct {
	Field1 string
	Field2 string
}

func TestMultipleDestination(t *testing.T) {
	tests := []struct {
		input  *multipleSource
		output *multipleDest
	}{
		{
			input: &multipleSource{
				Field: "value",
			},
			output: &multipleDest{
				Field1: "value",
				Field2: "value",
			},
		},
	}

	for _, tt := range tests {
		output := &multipleDest{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)
	}
}

type recursiveSource struct {
	Field *recursiveSourceChild `mapconv:",recursive"`
}

type recursiveSourceChild struct {
	Field1 string `mapconv:"Dest1,omitempty"`
	Field2 string `mapconv:"Dest2,omitempty"`
}

type recursiveDest struct {
	Field *recursiveDestChild
}

type recursiveDestChild struct {
	Dest1 string
	Dest2 string
}

type recursiveSourceSlice struct {
	Fields []*recursiveSourceChild `mapconv:"[]Slice,recursive"`
}

type recursiveDestSlice struct {
	Slice []*recursiveDestChild
}

func TestRecursive(t *testing.T) {
	tests := []struct {
		input  *recursiveSource
		expect *recursiveDest
	}{
		{
			input: &recursiveSource{
				Field: &recursiveSourceChild{
					Field1: "value1",
					Field2: "value2",
				},
			},
			expect: &recursiveDest{
				Field: &recursiveDestChild{
					Dest1: "value1",
					Dest2: "value2",
				},
			},
		},
	}

	for _, tt := range tests {
		dest := &recursiveDest{}
		err := ConvertTo(tt.input, dest)
		require.NoError(t, err)
		require.Equal(t, tt.expect, dest)

		// reverse
		source := &recursiveSource{}
		err = ConvertFrom(tt.expect, source)
		require.NoError(t, err)
		require.Equal(t, tt.input, source)
	}
}

func TestRecursiveSlice(t *testing.T) {
	tests := []struct {
		input  *recursiveSourceSlice
		output *recursiveDestSlice
	}{
		{
			input: &recursiveSourceSlice{
				Fields: []*recursiveSourceChild{
					{
						Field1: "value1",
						Field2: "value2",
					},
					{
						Field1: "value3",
						Field2: "value4",
					},
				},
			},
			output: &recursiveDestSlice{
				Slice: []*recursiveDestChild{
					{
						Dest1: "value1",
						Dest2: "value2",
					},
					{
						Dest1: "value3",
						Dest2: "value4",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		output := &recursiveDestSlice{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)

		// reverse
		source := &recursiveSourceSlice{}
		err = ConvertFrom(tt.output, source)
		require.NoError(t, err)
		require.Equal(t, tt.input, source)
	}
}

func TestRecursiveSliceMerging(t *testing.T) {
	tests := []struct {
		src    *recursiveSourceSlice
		dest   *recursiveDestSlice
		expect *recursiveDestSlice
	}{
		{
			src: &recursiveSourceSlice{
				Fields: []*recursiveSourceChild{
					{
						Field1: "value1-upd",
					},
					{
						Field2: "value4-upd",
					},
				},
			},
			dest: &recursiveDestSlice{
				Slice: []*recursiveDestChild{
					{
						Dest1: "value1",
						Dest2: "value2",
					},
					{
						Dest1: "value3",
						Dest2: "value4",
					},
				},
			},
			expect: &recursiveDestSlice{
				Slice: []*recursiveDestChild{
					{
						Dest1: "value1-upd",
						Dest2: "value2",
					},
					{
						Dest1: "value3",
						Dest2: "value4-upd",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		err := ConvertTo(tc.src, tc.dest)
		require.NoError(t, err)
		require.EqualValues(t, tc.expect, tc.dest)
	}
}

type sourceSquash struct {
	Field *sourceSquashChild `mapconv:",squash"`
}

type sourceSquashChild struct {
	Field1 string
	Field2 string
}

type destSquash struct {
	Field1 string
	Field2 string
}

func TestSquash(t *testing.T) {
	tests := []struct {
		input  *sourceSquash
		output *destSquash
	}{
		{
			input: &sourceSquash{
				Field: &sourceSquashChild{
					Field1: "f1",
					Field2: "f2",
				},
			},
			output: &destSquash{
				Field1: "f1",
				Field2: "f2",
			},
		},
	}

	for _, tt := range tests {
		output := &destSquash{}
		err := ConvertTo(tt.input, output)
		require.NoError(t, err)
		require.Equal(t, tt.output, output)

		// reverse
		source := &sourceSquash{}
		err = ConvertFrom(tt.output, source)
		require.Error(t, err)
	}
}

func testDecoder() *Decoder {
	strToNumFilter := func(v interface{}) (interface{}, error) {
		return strconv.ParseInt(v.(string), 10, 64)
	}
	toUpperFilter := func(v interface{}) (interface{}, error) {
		// to upper
		return strings.ToUpper(v.(string)), nil
	}
	numToIDFilter := func(v interface{}) (interface{}, error) {
		return types.ID(v.(int64)), nil
	}
	errorFilter := func(v interface{}) (interface{}, error) {
		return nil, errors.New("foobar")
	}

	return &Decoder{Config: &DecoderConfig{
		TagName: DefaultMapConvTag,
		FilterFuncs: map[string]FilterFunc{
			"toUpper":  toUpperFilter,
			"strToNum": strToNumFilter,
			"numToID":  numToIDFilter,
			"error":    errorFilter,
		},
	}}
}

func TestFiltersWithConvertTo(t *testing.T) {
	decoder := testDecoder()

	cases := []struct {
		in     interface{}
		dest   interface{}
		expect interface{}
		err    error
	}{
		{
			in: &struct {
				Field string `mapconv:",filters=toUpper"`
			}{Field: "foo"},
			dest:   &struct{ Field string }{},
			expect: &struct{ Field string }{Field: "FOO"},
		},
		{
			in: &struct {
				Field string `mapconv:",filters=strToNum numToID"`
			}{Field: "1"},
			dest:   &struct{ Field types.ID }{},
			expect: &struct{ Field types.ID }{Field: types.ID(1)},
		},
		{
			in: &struct {
				Field string `mapconv:",filters=error"`
			}{Field: "1"},
			dest: &struct{ Field types.ID }{},
			err:  errors.New("failed to apply the filter: foobar"),
		},
		{
			in: &struct {
				Field string `mapconv:",filters=strToNum numToID error"`
			}{Field: "1"},
			dest: &struct{ Field types.ID }{},
			err:  errors.New("failed to apply the filter: foobar"),
		},
	}

	for _, tc := range cases {
		err := decoder.ConvertTo(tc.in, tc.dest)
		require.Equal(t, tc.err, err)
		if err == nil {
			require.EqualValues(t, tc.expect, tc.dest)
		}
	}
}

func TestFiltersWithConvertFrom(t *testing.T) {
	decoder := testDecoder()

	cases := []struct {
		in     interface{}
		dest   interface{}
		expect interface{}
		err    error
	}{
		{
			in: &struct{ Field string }{Field: "foo"},
			dest: &struct {
				Field string `mapconv:",filters=toUpper"`
			}{},
			expect: &struct {
				Field string `mapconv:",filters=toUpper"`
			}{Field: "FOO"},
		},
		{
			in: &struct{ Field string }{Field: "1"},
			dest: &struct {
				Field types.ID `mapconv:",filters=strToNum numToID"`
			}{},
			expect: &struct {
				Field types.ID `mapconv:",filters=strToNum numToID"`
			}{Field: types.ID(1)},
		},
		{
			in: &struct{ Field string }{Field: "1"},
			dest: &struct {
				Field types.ID `mapconv:",filters=error"`
			}{},
			err: errors.New("failed to apply the filter: foobar"),
		},
		{
			in: &struct{ Field string }{Field: "1"},
			dest: &struct {
				Field types.ID `mapconv:",filters=strToNum numToID error"`
			}{},
			err: errors.New("failed to apply the filter: foobar"),
		},
	}

	for _, tc := range cases {
		err := decoder.ConvertFrom(tc.in, tc.dest)
		require.Equal(t, tc.err, err)
		if err == nil {
			require.EqualValues(t, tc.expect, tc.dest)
		}
	}
}

type recursiveMerge struct {
	Nest *recursiveMergeNest `mapconv:",recursive"`
}

type recursiveMergeNest struct {
	Field1 string `mapconv:",omitempty"`
	Field2 string `mapconv:",omitempty"`
}

func TestOverwrite(t *testing.T) {
	cases := []struct {
		src    *recursiveMergeNest
		dest   *recursiveMergeNest
		expect *recursiveMergeNest
	}{
		{
			src: &recursiveMergeNest{
				Field1: "field1-upd",
				Field2: "",
			},
			dest: &recursiveMergeNest{
				Field1: "field1",
				Field2: "field2",
			},
			expect: &recursiveMergeNest{
				Field1: "field1-upd",
				Field2: "field2",
			},
		},
	}
	for _, tc := range cases {
		if err := ConvertTo(tc.src, tc.dest); err != nil {
			t.Fatal(err)
		}
		require.EqualValues(t, tc.expect, tc.dest)
	}
}

func TestRecursiveMerge(t *testing.T) {
	cases := []struct {
		src    *recursiveMerge
		dest   *recursiveMerge
		expect *recursiveMerge
	}{
		{
			src: &recursiveMerge{
				Nest: &recursiveMergeNest{
					Field1: "field1-upd",
					Field2: "",
				},
			},
			dest: &recursiveMerge{
				Nest: &recursiveMergeNest{
					Field1: "field1",
					Field2: "field2",
				},
			},
			expect: &recursiveMerge{
				Nest: &recursiveMergeNest{
					Field1: "field1-upd",
					Field2: "field2",
				},
			},
		},
	}
	for _, tc := range cases {
		if err := ConvertTo(tc.src, tc.dest); err != nil {
			t.Fatal(err)
		}
		require.EqualValues(t, tc.expect, tc.dest)
	}
}
