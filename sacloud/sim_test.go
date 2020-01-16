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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSIMJSON = `
    {
      "Index": 0,
      "ID": 123456789012,
      "Name": "example-sim",
      "Description": "example-desc",
      "Settings": null,
      "SettingsHash": null,
      "Status": {
        "ICCID": "898104xxxxxxxxxxxxx"
		%s
      },
      "ServiceClass": "cloud/sim/1",
      "Availability": "available",
      "CreatedAt": "2018-01-12T15:33:53+09:00",
      "ModifiedAt": "2018-01-12T15:33:53+09:00",
      "Provider": {
        "ID": 8000001,
        "Class": "sim",
        "Name": "sakura-sim",
        "ServiceClass": "cloud/sim"
      },
      "Icon": null,
      "Tags": []
    }
`

const testSIMStatusBody = `
  "sim": {
    "iccid": "898104xxxxxxxxxxxxx",
    "imsi": [
      "xxxxxxxxxxxxxxx"
    ],
    "session_status": "DOWN",
    "imei_lock": false,
    "registered": true,
    "activated": true,
    "resource_id": "xxxxxxxxxxxx",
    "registered_date": "2018-04-04T07:46:11+00:00",
    "activated_date": "2018-04-04T07:48:48+00:00",
    "deactivated_date": "2018-04-04T07:46:02+00:00",
    "traffic_bytes_of_current_month": {
      "uplink_bytes": "169734354",
      "downlink_bytes": "509606410"
    },
    "connected_imei": "xxxxxxxxxxxxxxx"
  },
  "is_ok": true
`

const testSIMStatusNoTrafficBody = `
{
  "sim": {
    "iccid": "898104xxxxxxxxxxxxx",
    "imsi": [
      "xxxxxxxxxxxxxxx"
    ],
    "session_status": "DOWN",
    "imei_lock": false,
    "registered": true,
    "activated": true,
    "resource_id": "xxxxxxxxxxxx",
    "registered_date": "2018-04-04T07:46:11+00:00",
    "activated_date": "2018-04-04T07:48:48+00:00",
    "deactivated_date": "2018-04-04T07:46:02+00:00",
    "traffic_bytes_of_current_month": [],
    "connected_imei": "xxxxxxxxxxxxxxx"
  },
  "is_ok": true
}
`

const testSIMLogJSON = `
{
    "date": "2018-04-03T08:18:07+00:00",
    "session_status": "Created",
    "resource_id": "xxxxxxxxxxxx",
    "imei": "xxxxxxxxxxxxxxx",
    "imsi": "xxxxxxxxxxxxxxx"
}
`

const testSIMNetworkOperatorConfigsJSON = `
{
  "network_operator_config": [
    {
      "name": "SoftBank",
      "allow": true
    },
    {
      "name": "NTT DOCOMO",
      "allow": true
    },
    {
      "name": "KDDI",
      "allow": true
    }
  ]
}
`

func TestMarshalSIMJSON(t *testing.T) {

	t.Run("SIM no status", func(t *testing.T) {
		var sim SIM
		err := json.Unmarshal([]byte(fmt.Sprintf(testSIMJSON, "")), &sim)
		assert.NoError(t, err)
		assert.NotEmpty(t, sim)
		assert.NotEmpty(t, sim.ID)
		assert.NotEmpty(t, sim.Name)
		assert.NotEmpty(t, sim.Description)
		assert.NotEmpty(t, sim.ServiceClass)
		assert.NotEmpty(t, sim.Availability)
		assert.NotEmpty(t, sim.CreatedAt)
		assert.NotEmpty(t, sim.ModifiedAt)
		assert.NotEmpty(t, sim.Provider)
		assert.Empty(t, sim.Status.SIMInfo)
	})

	t.Run("SIM includes status", func(t *testing.T) {
		var sim SIM
		err := json.Unmarshal([]byte(fmt.Sprintf(testSIMJSON, ","+testSIMStatusBody)), &sim)
		assert.NoError(t, err)
		assert.NotEmpty(t, sim)
		assert.NotEmpty(t, sim.Status.SIMInfo)
		info := sim.Status.SIMInfo

		assert.NotEmpty(t, info.ICCID)
		assert.NotEmpty(t, info.IMSI)
		assert.NotEmpty(t, info.SessionStatus)
		assert.False(t, info.IMEILock)
		assert.True(t, info.Registered)
		assert.True(t, info.Activated)
		assert.NotEmpty(t, info.ResourceID)
		assert.NotEmpty(t, info.RegisteredDate)
		assert.NotEmpty(t, info.ActivatedDate)
		assert.NotEmpty(t, info.DeactivatedDate)
		assert.NotEmpty(t, info.TrafficBytesOfCurrentMonth)
		assert.NotEmpty(t, info.ConnectedIMEI)
	})

	t.Run("SIM status no traffic bytes", func(t *testing.T) {
		var simInfo SIMInfo
		err := json.Unmarshal([]byte(testSIMStatusNoTrafficBody), &simInfo)
		assert.NoError(t, err)
		assert.Empty(t, simInfo.TrafficBytesOfCurrentMonth)
	})

	t.Run("SIM logs", func(t *testing.T) {
		var logs SIMLog
		err := json.Unmarshal([]byte(testSIMLogJSON), &logs)
		assert.NoError(t, err)
		assert.NotEmpty(t, logs)
		assert.NotEmpty(t, logs.Date)
		assert.NotEmpty(t, logs.SessionStatus)
		assert.NotEmpty(t, logs.ResourceID)
		assert.NotEmpty(t, logs.IMEI)
		assert.NotEmpty(t, logs.IMSI)
	})

	t.Run("SIM NetworkOperatorConfig", func(t *testing.T) {
		var configs SIMNetworkOperatorConfigs
		err := json.Unmarshal([]byte(testSIMNetworkOperatorConfigsJSON), &configs)

		assert.NoError(t, err)
		assert.NotEmpty(t, configs)
		assert.Len(t, configs.NetworkOperatorConfigs, 3)

	})
}

func TestCreateNewSIM(t *testing.T) {

	sim := CreateNewSIM("name", "iccID", "passcode")
	assert.NotNil(t, sim)
	assert.Equal(t, "name", sim.Name)
	assert.Equal(t, "iccID", sim.Status.ICCID)
	assert.Equal(t, "passcode", sim.Remark.PassCode)

}
