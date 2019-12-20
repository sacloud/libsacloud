// Copyright 2016-2019 The Libsacloud Authors
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

package size

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSizeFunctions(t *testing.T) {
	cases := []struct {
		msg string
		f   func(int) int
		in  int
		out int
	}{
		{
			msg: "MiBToGiB",
			f:   MiBToGiB,
			in:  1024,
			out: 1,
		},
		{
			msg: "MiBToGiB",
			f:   MiBToGiB,
			in:  51200,
			out: 50,
		},
		{
			msg: "MiBToGiB",
			f:   MiBToGiB,
			in:  102400,
			out: 100,
		},
		{
			msg: "ToGiB with Zero",
			f:   MiBToGiB,
			in:  0,
			out: 0,
		},
		{
			msg: "GiBToMiB",
			f:   GiBToMiB,
			in:  1,
			out: 1024,
		},
		{
			msg: "GiBToMiB",
			f:   GiBToMiB,
			in:  100,
			out: 102400,
		},
		{
			msg: "GiBToMiB with Zero",
			f:   GiBToMiB,
			in:  0,
			out: 0,
		},
	}

	for _, tt := range cases {
		got := tt.f(tt.in)
		require.Equal(t, tt.out, got, "%s returns unexpected value: expect: %v actual:", tt.msg, tt.out, got)
	}
}
