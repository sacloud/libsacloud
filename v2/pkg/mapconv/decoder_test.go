// Copyright 2016-2020 The Libsacloud Authors
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
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapConvDecoder_inheritsDecoderConfig(t *testing.T) {
	in := &Nest1{
		Nest2: &Nest2{Field: "foo"},
	}
	expect := &Dest{
		Field: "FOO",
	}

	decoder := &Decoder{
		Config: &DecoderConfig{
			TagName: DefaultMapConvTag,
			FilterFuncs: map[string]FilterFunc{
				"test": func(v interface{}) (interface{}, error) {
					s := v.(string)
					return strings.ToUpper(s), nil
				},
			},
		},
	}

	out := &Dest{}
	if err := decoder.ConvertTo(in, out); err != nil {
		t.Fatal(err)
	}
	require.EqualValues(t, expect, out)
}

type Nest1 struct {
	Nest2 *Nest2 `mapconv:",squash"`
}

type Nest2 struct {
	Field string `mapconv:",filters=test"`
}

type Dest struct {
	Field string
}
