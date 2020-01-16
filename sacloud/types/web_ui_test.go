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

package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebUI_MarshalJSON(t *testing.T) {
	cases := []struct {
		in  WebUI
		out []byte
	}{
		{
			in:  WebUI(""),
			out: []byte(`false`),
		},
		{
			in:  WebUI("true"),
			out: []byte(`true`),
		},
		{
			in:  WebUI("false"),
			out: []byte(`false`),
		},
		{
			in:  WebUI("http://localhost:8080"),
			out: []byte(`"http://localhost:8080"`),
		},
	}

	for _, cc := range cases {
		data, err := json.Marshal(cc.in)
		if err != nil {
			t.Fatal(err, cc)
		}
		require.Equal(t, cc.out, data, "target: %#v", cc)
	}
}

func TestWebUI_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		in      string
		out     WebUI
		enabled bool
	}{
		{
			in:      `""`,
			out:     WebUI(""),
			enabled: false,
		},
		{
			in:      `true`,
			out:     WebUI("true"),
			enabled: true,
		},
		{
			in:      `false`,
			out:     WebUI("false"),
			enabled: false,
		},
		{
			in:      `"http://localhost:8080"`,
			out:     WebUI("http://localhost:8080"),
			enabled: true,
		},
	}

	for _, cc := range cases {
		var v WebUI
		if err := json.Unmarshal([]byte(cc.in), &v); err != nil {
			t.Fatal(err, cc)
		}
		require.Equal(t, cc.enabled, v.Bool(), "target: %#v", cc)
	}
}
