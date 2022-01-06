// Copyright 2016-2022 The Libsacloud Authors
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

package fake

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type SourceStruct struct {
	Field1  string
	Field1e string
	Field2  string
	Field2e string
	Field3  int
	Field3e int
	Field4  []string
	Field4e []string
	Field5  interface{}
	Field5e interface{}
}

type DestStruct struct {
	Field1 string
	Field2 string
	Field3 int
	Field4 []string
	Field5 interface{}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		input  *SourceStruct
		output *DestStruct
	}{
		{
			input: &SourceStruct{
				Field1:  "field1",
				Field1e: "field1",
				Field2:  "field2",
				Field2e: "field2",
				Field3:  3,
				Field3e: 3,
				Field4:  []string{"f", "i", "e", "l", "d", "4"},
				Field4e: []string{"f", "i", "e", "l", "d", "4"},
				Field5:  "field5",
				Field5e: "field5",
			},
			output: &DestStruct{
				Field1: "field1",
				Field2: "field2",
				Field3: 3,
				Field4: []string{"f", "i", "e", "l", "d", "4"},
				Field5: "field5",
			},
		},
	}

	for _, tt := range tests {
		dest := &DestStruct{}
		copySameNameField(tt.input, dest)
		require.Equal(t, tt.output, dest)
	}
}
