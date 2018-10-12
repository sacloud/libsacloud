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
