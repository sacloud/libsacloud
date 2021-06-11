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

package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	vpcRouterAPIName     = "VPCRouter"
	vpcRouterAPIPathName = "appliance"
)

var vpcRouterAPI = &dsl.Resource{
	Name:       vpcRouterAPIName,
	PathName:   vpcRouterAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.FindAppliance(vpcRouterAPIName, vpcRouterNakedType, findParameter, vpcRouterView),

		// create
		ops.CreateAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterCreateParam, vpcRouterView),

		// read
		ops.ReadAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterView),

		// update
		ops.UpdateAppliance(vpcRouterAPIName, vpcRouterNakedType, vpcRouterUpdateParam, vpcRouterView),
		// updateSettings
		ops.UpdateApplianceSettings(vpcRouterAPIName, vpcRouterUpdateSettingsNakedType, vpcRouterUpdateSettingsParam, vpcRouterView),

		// delete
		ops.Delete(vpcRouterAPIName),

		// config
		ops.Config(vpcRouterAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(vpcRouterAPIName),
		ops.Shutdown(vpcRouterAPIName),
		ops.Reset(vpcRouterAPIName),

		// connect to switch
		ops.WithIDAction(
			vpcRouterAPIName, "ConnectToSwitch", http.MethodPut, "interface/{{.nicIndex}}/to/switch/{{.switchID}}",
			&dsl.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
			&dsl.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		// disconnect from switch
		ops.WithIDAction(
			vpcRouterAPIName, "DisconnectFromSwitch", http.MethodDelete, "interface/{{.nicIndex}}/to/switch",
			&dsl.Argument{
				Name: "nicIndex",
				Type: meta.TypeInt,
			},
		),

		// monitor
		ops.MonitorChildBy(vpcRouterAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),

		// status
		{
			ResourceName: vpcRouterAPIName,
			Name:         "Status",
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			PathFormat: dsl.IDAndSuffixPathFormat("status"),
			Method:     http.MethodGet,
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.VPCRouterStatus{}),
				Name: "Router",
			}),
			Results: dsl.Results{
				{
					SourceField: "Router",
					DestField:   vpcRouterStatusView.Name,
					Model:       vpcRouterStatusView,
				},
			},
		},
	},
}

var (
	vpcRouterNakedType               = meta.Static(naked.VPCRouter{})
	vpcRouterUpdateSettingsNakedType = meta.Static(naked.VPCRouterSettingsUpdate{})

	vpcRouterView = &dsl.Model{
		Name:      vpcRouterAPIName,
		NakedType: vpcRouterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.IconID(),
			fields.CreatedAt(),
			// plan
			fields.AppliancePlanID(),
			// version
			fields.ApplianceVPCRouterVersion(),
			// settings
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
			fields.SettingsHash(),

			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// interfaces
			fields.VPCRouterInterfaces(),
			fields.RemarkZoneID(),
		},
	}

	vpcRouterCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(vpcRouterAPIName),
		NakedType: vpcRouterNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"vpcrouter"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.PlanID(),

			// nic
			{
				Name: "Switch",
				Type: &dsl.Model{
					Name: "ApplianceConnectedSwitch",
					Fields: []*dsl.FieldDesc{
						fields.ID(),
						fields.Scope(),
					},
					NakedType: meta.Static(naked.ConnectedSwitch{}),
				},
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: "Remark.Switch,recursive",
				},
			},

			// TODO remarkとsettings.Interfaces両方に設定する必要がある。うまい方法が思いつかないため当面は利用者側で両方に設定する方法としておく
			fields.ApplianceIPAddresses(),

			// version
			fields.ApplianceVPCRouterVersion(),

			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}

	vpcRouterUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(vpcRouterAPIName),
		NakedType: vpcRouterNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
			// settings hash
			fields.SettingsHash(),
		},
	}

	vpcRouterUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(vpcRouterAPIName),
		NakedType: vpcRouterNakedType,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Settings",
				Type: models.vpcRouterSetting(),
				Tags: &dsl.FieldTags{
					MapConv: ",omitempty,recursive",
				},
			},
			// settings hash
			fields.SettingsHash(),
		},
	}

	vpcRouterStatusView = &dsl.Model{
		Name:      "VPCRouterStatus",
		NakedType: meta.Static(naked.VPCRouterStatus{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("FirewallReceiveLogs", meta.TypeStringSlice),
			fields.Def("FirewallSendLogs", meta.TypeStringSlice),
			fields.Def("VPNLogs", meta.TypeStringSlice),
			fields.Def("SessionCount", meta.TypeInt),
			{
				Name: "WireGuard",
				Type: &dsl.Model{
					Name: "WireGuardStatus",
					Fields: []*dsl.FieldDesc{
						fields.Def("PublicKey", meta.TypeString),
					},
				},
			},
			{
				Name: "DHCPServerLeases",
				Type: &dsl.Model{
					Name:    "VPCRouterDHCPServerLease",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("IPAddress", meta.TypeString),
						fields.Def("MACAddress", meta.TypeString),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]DHCPServerLeases,recursive",
				},
			},
			{
				Name: "L2TPIPsecServerSessions",
				Type: &dsl.Model{
					Name:    "VPCRouterL2TPIPsecServerSession",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("User", meta.TypeString),
						fields.Def("IPAddress", meta.TypeString),
						fields.Def("TimeSec", meta.TypeInt),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]L2TPIPsecServerSessions,recursive",
				},
			},
			{
				Name: "PPTPServerSessions",
				Type: &dsl.Model{
					Name:    "VPCRouterPPTPServerSession",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("User", meta.TypeString),
						fields.Def("IPAddress", meta.TypeString),
						fields.Def("TimeSec", meta.TypeInt),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]PPTPServerSessions,recursive",
				},
			},
			{
				Name: "SiteToSiteIPsecVPNPeers",
				Type: &dsl.Model{
					Name:    "VPCRouterSiteToSiteIPsecVPNPeer",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("Status", meta.TypeString),
						fields.Def("Peer", meta.TypeString),
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]SiteToSiteIPsecVPNPeers,recursive",
				},
			},
		},
	}
)
