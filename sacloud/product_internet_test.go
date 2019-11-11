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

	"github.com/stretchr/testify/assert"
)

var testInternetPlanJSON = `
{
	"Index": 0,
	"ID": 100,
	"Name": "100Mbps\u5171\u6709",
	"BandWidthMbps": 100,
	"ServiceClass": "cloud\/internet\/router\/100m",
	"Availability": "available"
}
`

func TestMarshalProductInternetJSON(t *testing.T) {
	var productInternet ProductInternet
	err := json.Unmarshal([]byte(testInternetPlanJSON), &productInternet)

	assert.NoError(t, err)
	assert.NotEmpty(t, productInternet)

	assert.NotEmpty(t, productInternet.ID)
}
