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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestTransformer_transformLoadBalancerCreateArgs(t *testing.T) {
	op := &LoadBalancerOp{}

	ret, err := op.transformCreateArgs(&LoadBalancerCreateRequest{
		VirtualIPAddresses: LoadBalancerVirtualIPAddresses{
			{
				VirtualIPAddress: "192.168.0.1",
				Servers: LoadBalancerServers{
					{
						IPAddress: "192.168.0.11",
						Port:      80,
						Enabled:   true,
						HealthCheck: &LoadBalancerServerHealthCheck{
							Protocol:     types.LoadBalancerHealthCheckProtocols.HTTP,
							Path:         "/",
							ResponseCode: 200,
						},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	v := ret.Appliance.Settings.LoadBalancer[0].Servers[0].HealthCheck.Status
	if v != types.StringNumber(200) {
		t.Fatal("unexpected value:", v)
	}
}

func TestTransformer_transformSIMReadResults(t *testing.T) {
	op := &SIMOp{}
	data := []string{simJSONWithString, simJSONWithNumber}
	for _, d := range data {
		result, err := op.transformReadResults([]byte(d))
		require.NoError(t, err)
		require.Equal(t, int64(10101010), result.SIM.Info.TrafficBytesOfCurrentMonth.UplinkBytes)
		require.Equal(t, int64(20202020), result.SIM.Info.TrafficBytesOfCurrentMonth.DownlinkBytes)
	}
}

const simJSONWithString = `
{
  "CommonServiceItem": {
    "ID": 123456789012,
    "Name": "dummy",
    "Description": "dummy",
    "Settings": null,
    "SettingsHash": null,
    "Status": {
      "ICCID": "1111111111111111111",
      "sim": {
        "iccid": "1111111111111111111",
        "session_status": "DOWN",
        "imei_lock": false,
        "registered": true,
        "activated": true,
        "resource_id": "123456789012",
        "registered_date": "2020-10-01T05:39:17+00:00",
        "activated_date": "2020-10-01T05:39:17+00:00",
        "deactivated_date": "2020-10-01T04:48:39+00:00",
        "traffic_bytes_of_current_month": {
          "uplink_bytes": "10101010",
          "downlink_bytes": "20202020"
        }
      }
    },
    "ServiceClass": "cloud/sim/1",
    "Availability": "available",
    "CreatedAt": "2020-10-01T14:39:17+09:00",
    "ModifiedAt": "2020-10-01T14:39:17+09:00",
    "Provider": {
      "ID": 8000001,
      "Class": "sim",
      "Name": "sakura-sim",
      "ServiceClass": "cloud/sim"
    },
    "Icon": null
  },
  "is_ok": true
}
`

const simJSONWithNumber = `
{
  "CommonServiceItem": {
    "ID": 123456789012,
    "Name": "dummy",
    "Description": "dummy",
    "Settings": null,
    "SettingsHash": null,
    "Status": {
      "ICCID": "1111111111111111111",
      "sim": {
        "iccid": "1111111111111111111",
        "session_status": "DOWN",
        "imei_lock": false,
        "registered": true,
        "activated": true,
        "resource_id": "123456789012",
        "registered_date": "2020-10-01T05:39:17+00:00",
        "activated_date": "2020-10-01T05:39:17+00:00",
        "deactivated_date": "2020-10-01T04:48:39+00:00",
        "traffic_bytes_of_current_month": {
          "uplink_bytes": 10101010,
          "downlink_bytes": 20202020
        }
      }
    },
    "ServiceClass": "cloud/sim/1",
    "Availability": "available",
    "CreatedAt": "2020-10-01T14:39:17+09:00",
    "ModifiedAt": "2020-10-01T14:39:17+09:00",
    "Provider": {
      "ID": 8000001,
      "Class": "sim",
      "Name": "sakura-sim",
      "ServiceClass": "cloud/sim"
    },
    "Icon": null
  },
  "is_ok": true
}
`
