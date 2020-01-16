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

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPublicPriceJSON = `
{
	"DisplayName": "\u30d7\u30e9\u30f311(\u30c7\u30a3\u30b9\u30af\u306a\u3057)",
	"IsPublic": true,
	"Price": {
		"Base": 0,
		"Daily": 3270,
		"Hourly": 326,
		"Monthly": 65397,
		"Zone": "is1b"
	},
	"ServiceClassID": 50060,
	"ServiceClassName": "plan\/11",
	"ServiceClassPath": "cloud\/plan\/11"
}
`

func TestMarshalPublicPriceJSON(t *testing.T) {
	var publicPrice PublicPrice
	err := json.Unmarshal([]byte(testPublicPriceJSON), &publicPrice)

	assert.NoError(t, err)
	assert.NotEmpty(t, publicPrice)

	assert.NotEmpty(t, publicPrice.DisplayName)
	assert.NotEmpty(t, publicPrice.Price)
}
