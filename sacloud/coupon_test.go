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

var testCouponJSON = `
    {
      "AppliedAt": "2019-01-10T11:12:13+09:00",
      "ContractID": 111111111111,
      "CouponID": "123456789012",
      "Discount": 999999,
      "MemberID": "abc99999",
      "ServiceClassID": 50122,
      "UntilAt": "2019-03-31T23:59:59+09:00"
    }
`

func TestMarshalCouponJSON(t *testing.T) {
	var d Coupon
	err := json.Unmarshal([]byte(testCouponJSON), &d)

	assert.NoError(t, err)
	assert.NotEmpty(t, d)
}
