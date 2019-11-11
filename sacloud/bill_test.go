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

package sacloud

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testBillDetailJSON = map[string]string{
	"global": `{
		    "ContractID": 123456789001,
		    "Description": "foobar",
		    "Index": 1,
		    "ServiceClassID": 50000
		}`,
	"normal": `{
		    "ContractID": 123456789002,
		    "Amount": 1620,
		    "Description": "foobar",
		    "Index": 2,
		    "ServiceClassID": 50072,
		    "Usage": 1296000,
		    "Zone": "is1b"
		}`,
	"ended": `{
		    "ContractID": 123456789003,
		    "Amount": 1677,
		    "ContractEndAt": "2017-01-01T10:10:10+09:00",
		    "Description": "foobar",
		    "Index": 3,
		    "ServiceClassID": 50226,
		    "Usage": 1117137,
		    "Zone": "is1b"
		}`,
}

func TestMarshalBillDetailJSON(t *testing.T) {
	for _, v := range testBillDetailJSON {
		var d BillDetail
		err := json.Unmarshal([]byte(v), &d)

		assert.NoError(t, err)
		assert.NotEmpty(t, d)
	}
}

func TestMarshalBillDetailContractEnded(t *testing.T) {
	for k, v := range testBillDetailJSON {
		var d BillDetail
		json.Unmarshal([]byte(v), &d)

		res := false
		if k == "ended" {
			res = true
		}

		tm := time.Now()
		assert.Equal(t, d.IsContractEnded(tm), res)
	}
}
