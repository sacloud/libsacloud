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

var testLicenseInfoJSON = `
{
	"ID": 10001,
	"Name": "Windows RDS SAL",
	"ServiceClass": "cloud\/os\/windows\/rds-sal",
	"TermsOfUse": "1\u30e9\u30a4\u30bb\u30f3\u30b9\u306b\u3064\u304d\u30011\u4eba\u306e\u30e6\u30fc\u30b6\u304c\u5229\u7528\u3067\u304d\u307e\u3059\u3002",
	"CreatedAt": "2013-11-27T10:07:52+09:00",
	"ModifiedAt": "2013-11-27T10:07:52+09:00"
}
`

func TestMarshalProductLicenseJSON(t *testing.T) {
	var productLicense ProductLicense
	err := json.Unmarshal([]byte(testLicenseInfoJSON), &productLicense)

	assert.NoError(t, err)
	assert.NotEmpty(t, productLicense)

	assert.NotEmpty(t, productLicense.ID)
}
