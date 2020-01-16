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

var testIconJSON = `
{
	"ID": 123456789012,
	"URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/123456789012.png",
	"Name": "TEST",
	"Scope": "shared"
}
`

func TestMarshalIconJSON(t *testing.T) {
	var icon Icon
	err := json.Unmarshal([]byte(testIconJSON), &icon)

	assert.NoError(t, err)
	assert.NotEmpty(t, icon)

	assert.NotEmpty(t, icon.ID)
}
