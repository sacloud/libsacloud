// Copyright 2016-2021 The Libsacloud Authors
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
	"reflect"
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

func TestTransformer_transformSetCertificatesArgs(t *testing.T) {
	op := &ProxyLBOp{}

	cert := &ProxyLBPrimaryCert{
		ServerCertificate:       "aaa",
		IntermediateCertificate: "bbb",
		PrivateKey:              "ccc",
	}

	ret, err := op.transformSetCertificatesArgs(types.ID(1),
		&ProxyLBSetCertificatesRequest{
			PrimaryCerts:    cert,
			AdditionalCerts: []*ProxyLBAdditionalCert{},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	if ret.ProxyLB.PrimaryCert == nil {
		t.Fatal("ProxyLB.PrimaryCert is nil")
	}
	if ret.ProxyLB.PrimaryCert.ServerCertificate != cert.ServerCertificate {
		t.Fatal("ProxyLB.PrimaryCert has unexpected value: ServerCertificate")
	}
	if ret.ProxyLB.PrimaryCert.IntermediateCertificate != cert.IntermediateCertificate {
		t.Fatal("ProxyLB.PrimaryCert has unexpected value: IntermediateCertificate")
	}
	if ret.ProxyLB.PrimaryCert.PrivateKey != cert.PrivateKey {
		t.Fatal("ProxyLB.PrimaryCert has unexpected value: PrivateKey")
	}
}

func TestTransformer_transformDatabaseGetParameterResult(t *testing.T) {
	op := &DatabaseOp{}
	results, err := op.transformGetParameterResults([]byte(databaseParameterJSON))
	if err != nil {
		t.Fatal(err)
	}
	v := results.DatabaseParameter
	require.EqualValues(t, map[string]interface{}{
		"MariaDB/server.cnf/mysqld/event_scheduler": float64(100),
	}, v.Settings)
	require.Len(t, v.MetaInfo, 1)
	require.EqualValues(t, "string", v.MetaInfo[0].Type)
	require.EqualValues(t, "MariaDB/server.cnf/mysqld/event_scheduler", v.MetaInfo[0].Name)
	require.EqualValues(t, "event_scheduler", v.MetaInfo[0].Label)
	require.EqualValues(t, "65536", v.MetaInfo[0].Example)
	require.EqualValues(t, 1024, v.MetaInfo[0].Min)
	require.EqualValues(t, 2147483647, v.MetaInfo[0].Max)
	require.EqualValues(t, 10, v.MetaInfo[0].MaxLen)
	require.EqualValues(t, "dynamic", v.MetaInfo[0].Reboot)
}

func TestTransformer_transformDatabaseGetParameterResult_minimum(t *testing.T) {
	op := &DatabaseOp{}
	results, err := op.transformGetParameterResults([]byte(databaseParameterJSONMinimum))
	if err != nil {
		t.Fatal(err)
	}
	v := results.DatabaseParameter
	require.Empty(t, v.Settings)
	require.Empty(t, v.MetaInfo)
}

const (
	databaseParameterJSON = `
{
  "Database": {
    "Parameter": {
      "NoteID": "123456789012",
      "Attr": {
			"MariaDB/server.cnf/mysqld/event_scheduler": 100
		}
    },
    "Remark": {
      "Settings": [],
      "Form": [
        {
          "type": "radios",
          "name": "MariaDB/server.cnf/mysqld/event_scheduler",
          "label": "event_scheduler",
          "options": {
            "ex": "65536",
            "min": 1024,
            "max": 2147483647,
            "maxlen": 10,
            "text": "イベントスケジュールの有効無効を設定します。",
            "reboot": "dynamic",
            "type": "string"
          },
          "items": [
            ["ON","ON"],
            ["OFF","OFF"]
          ]
        }
      ]
    }
  }
}
`
	databaseParameterJSONMinimum = `
{
  "Database": {
    "Parameter": {
      "NoteID": "123456789012",
      "Attr":[]
    },
    "Remark": {
      "Settings": [],
      "Form": []
    }
  }
}
`
)

func TestTransformer_transformDatabaseSetParameterArgs(t *testing.T) {
	op := &DatabaseOp{}
	result, err := op.transformSetParameterArgs(types.ID(0), map[string]interface{}{
		"foo": "bar",
	})
	if err != nil {
		t.Fatal(err)
	}
	require.NotNil(t, result.Parameter)
	require.NotEmpty(t, result.Parameter.Attr)
	require.EqualValues(t, "bar", result.Parameter.Attr["foo"])
}

func TestServerOp_transformChangePlanArgs(t *testing.T) {
	type fields struct {
		Client     APICaller
		PathSuffix string
		PathName   string
	}
	type args struct {
		id   types.ID
		plan *ServerChangePlanRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *serverChangePlanRequestEnvelope
		wantErr bool
	}{
		{
			name: "standard",
			args: args{
				id: 1,
				plan: &ServerChangePlanRequest{
					CPU:      2,
					MemoryMB: 4,
				},
			},
			want: &serverChangePlanRequestEnvelope{
				CPU:                  2,
				MemoryMB:             4,
				ServerPlanCommitment: types.Commitments.Standard,
			},
			wantErr: false,
		},
		{
			name: "dedicated-cpu",
			args: args{
				id: 1,
				plan: &ServerChangePlanRequest{
					CPU:                  2,
					MemoryMB:             4,
					ServerPlanGeneration: types.PlanGenerations.G200,
					ServerPlanCommitment: types.Commitments.DedicatedCPU,
				},
			},
			want: &serverChangePlanRequestEnvelope{
				CPU:                  2,
				MemoryMB:             4,
				ServerPlanGeneration: types.PlanGenerations.G200,
				ServerPlanCommitment: types.Commitments.DedicatedCPU,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ServerOp{
				Client:     tt.fields.Client,
				PathSuffix: tt.fields.PathSuffix,
				PathName:   tt.fields.PathName,
			}
			got, err := o.transformChangePlanArgs(tt.args.id, tt.args.plan)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformChangePlanArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformChangePlanArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
