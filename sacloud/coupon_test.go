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
      "CouponID": "xxxxxxxxxxxxxxxx",
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
