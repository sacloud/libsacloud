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
	mobileGatewayAPIName     = "MobileGateway"
	mobileGatewayAPIPathName = "appliance"
)

var mobileGatewayAPI = &dsl.Resource{
	Name:       mobileGatewayAPIName,
	PathName:   mobileGatewayAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.FindAppliance(mobileGatewayAPIName, mobileGatewayNakedType, findParameter, mobileGatewayView),

		// create
		ops.CreateAppliance(mobileGatewayAPIName, mobileGatewayNakedType, mobileGatewayCreateParam, mobileGatewayView),

		// read
		ops.ReadAppliance(mobileGatewayAPIName, mobileGatewayNakedType, mobileGatewayView),

		// update
		ops.UpdateAppliance(mobileGatewayAPIName, mobileGatewayNakedType, mobileGatewayUpdateParam, mobileGatewayView),
		// updateSettings
		ops.UpdateApplianceSettings(mobileGatewayAPIName, mobileGatewayUpdateSettingsNakedType, mobileGatewayUpdateSettingsParam, mobileGatewayView),

		// delete
		ops.Delete(mobileGatewayAPIName),

		// config
		ops.Config(mobileGatewayAPIName),

		// power management(boot/shutdown/reset)
		ops.Boot(mobileGatewayAPIName),
		ops.Shutdown(mobileGatewayAPIName),
		ops.Reset(mobileGatewayAPIName),

		// connect to switch
		ops.WithIDAction(
			mobileGatewayAPIName, "ConnectToSwitch", http.MethodPut, "interface/1/to/switch/{{.switchID}}",
			&dsl.Argument{
				Name: "switchID",
				Type: meta.TypeID,
			},
		),

		// disconnect from switch
		ops.WithIDAction(
			mobileGatewayAPIName, "DisconnectFromSwitch", http.MethodDelete, "interface/1/to/switch",
		),

		// DNS
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "GetDNS",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/dnsresolver"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "SIMGroup",
				Type: meta.Static(naked.MobileGatewaySIMGroup{}),
				Tags: &dsl.FieldTags{
					JSON: "sim_group",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "SIMGroup",
					DestField:   "SIMGroup",
					IsPlural:    false,
					Model:       mobileGatewayDNSModel,
				},
			},
		},
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "SetDNS",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/dnsresolver"),
			Method:       http.MethodPut,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", mobileGatewayDNSModel, "SIMGroup"),
			},
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.MobileGatewaySIMGroup{}),
				Name: "SIMGroup",
				Tags: &dsl.FieldTags{
					JSON: "sim_group",
				},
			}),
		},

		// SIM Route
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "GetSIMRoutes",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/simroutes"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "SIMRoutes",
				Type: meta.Static([]*naked.MobileGatewaySIMRoute{}),
				Tags: &dsl.FieldTags{
					JSON: "sim_routes",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "SIMRoutes",
					DestField:   "SIMRoutes",
					IsPlural:    true,
					Model:       mobileGatewaySIMRouteView,
				},
			},
		},
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "SetSIMRoutes",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/simroutes"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static([]*naked.MobileGatewaySIMRoute{}),
				Name: "SIMRoutes",
				Tags: &dsl.FieldTags{
					JSON: "sim_routes",
				},
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", mobileGatewaySIMRouteParam, "[]SIMRoutes"),
			},
		},

		// list SIM
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "ListSIM",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/sims"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "SIM",
				Type: meta.Static([]*naked.SIMInfo{}),
				Tags: &dsl.FieldTags{
					JSON: "sim",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "SIM",
					DestField:   "SIM",
					IsPlural:    true,
					Model:       models.simInfoList(),
				},
			},
		},
		// add SIM
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "AddSIM",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/sims"),
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.SIMInfo{}),
				Name: "SIM",
				Tags: &dsl.FieldTags{
					JSON: "sim",
				},
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", mobileGatewayAddSIMParam, "SIM"),
			},
		},
		// delete SIM
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "DeleteSIM",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/sims/{{.simID}}"),
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "simID",
					Type: meta.TypeID,
				},
			},
		},

		// session logs
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "Logs",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/sessionlog"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "Logs",
				Type: meta.Static([]*naked.SIMLog{}),
				Tags: &dsl.FieldTags{
					JSON: "logs",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "Logs",
					DestField:   "Logs",
					IsPlural:    true,
					Model:       mobileGatewaySIMLogsModel,
				},
			},
		},

		// get traffic config
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "GetTrafficConfig",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/traffic_monitoring"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "TrafficMonitoring",
				Type: meta.Static(naked.TrafficMonitoringConfig{}),
				Tags: &dsl.FieldTags{
					JSON: "traffic_monitoring_config",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "TrafficMonitoring",
					DestField:   "TrafficMonitoring",
					IsPlural:    false,
					Model:       mobileGatewayTrafficConfigModel,
				},
			},
		},
		// set traffic config
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "SetTrafficConfig",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/traffic_monitoring"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "TrafficMonitoring",
				Type: meta.Static(naked.TrafficMonitoringConfig{}),
				Tags: &dsl.FieldTags{
					JSON: "traffic_monitoring_config",
				},
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", mobileGatewayTrafficConfigModel, "TrafficMonitoring"),
			},
		},
		// delete SIM
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "DeleteTrafficConfig",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/traffic_monitoring"),
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
		},

		// traffic status
		{
			ResourceName: mobileGatewayAPIName,
			Name:         "TrafficStatus",
			PathFormat:   dsl.IDAndSuffixPathFormat("mobilegateway/traffic_status"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "TrafficStatus",
				Type: meta.Static(naked.TrafficStatus{}),
				Tags: &dsl.FieldTags{
					JSON: "traffic_status",
				},
			}),
			Results: dsl.Results{
				{
					SourceField: "TrafficStatus",
					DestField:   "TrafficStatus",
					IsPlural:    false,
					Model:       mobileGatewayTrafficStatusModel,
				},
			},
		},

		// monitor
		ops.MonitorChildBy(mobileGatewayAPIName, "Interface", "interface",
			monitorParameter, monitors.interfaceModel()),
	},
}

var (
	mobileGatewayNakedType               = meta.Static(naked.MobileGateway{})
	mobileGatewayUpdateSettingsNakedType = meta.Static(naked.MobileGatewaySettingsUpdate{})

	mobileGatewayView = &dsl.Model{
		Name:      mobileGatewayAPIName,
		NakedType: mobileGatewayNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.IconID(),
			fields.CreatedAt(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceStatusChangedAt(),
			// interfaces
			fields.MobileGatewayInterfaces(),
			// remark
			fields.RemarkZoneID(),
			// settings
			fields.Def("Settings", models.mobileGatewaySetting(), mapConvTag(",omitempty,recursive")),
			fields.SettingsHash(),
		},
	}

	mobileGatewayCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(mobileGatewayAPIName),
		NakedType: mobileGatewayNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name:  "Class",
				Type:  meta.TypeString,
				Value: `"mobilegateway"`,
			},
			{
				Name: "PlanID",
				Tags: &dsl.FieldTags{
					MapConv: "Remark.Plan.ID/Plan.ID",
				},
				Type:  meta.TypeID,
				Value: `types.ID(2)`,
			},
			{
				Name: "SwitchID",
				Tags: &dsl.FieldTags{
					MapConv: "Remark.Switch.Scope",
				},
				Type:  meta.TypeString,
				Value: `"shared"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "Settings",
				Type: models.mobileGatewaySettingCreate(),
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",omitempty,recursive",
				},
			},
		},
	}

	mobileGatewayUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(mobileGatewayAPIName),
		NakedType: mobileGatewayNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "Settings",
				Type: models.mobileGatewaySetting(),
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",omitempty,recursive",
				},
			},
			// settings hash
			fields.SettingsHash(),
		},
	}
	mobileGatewayUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(mobileGatewayAPIName),
		NakedType: mobileGatewayNakedType,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Settings",
				Type: models.mobileGatewaySetting(),
				Tags: &dsl.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",omitempty,recursive",
				},
			},
			// settings hash
			fields.SettingsHash(),
		},
	}

	mobileGatewayDNSModel = &dsl.Model{
		Name:      "MobileGatewayDNSSetting",
		NakedType: meta.Static(naked.MobileGatewaySIMGroup{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("DNS1", meta.TypeString),
			fields.Def("DNS2", meta.TypeString),
		},
	}

	mobileGatewaySIMRouteParam = &dsl.Model{
		Name:      "MobileGatewaySIMRouteParam",
		NakedType: meta.Static(naked.MobileGatewaySIMRoute{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("ResourceID", meta.TypeString),
			fields.Def("Prefix", meta.TypeString),
		},
	}
	mobileGatewaySIMRouteView = &dsl.Model{
		Name:      "MobileGatewaySIMRoute",
		NakedType: meta.Static(naked.MobileGatewaySIMRoute{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("ResourceID", meta.TypeString),
			fields.Def("Prefix", meta.TypeString),
			fields.Def("ICCID", meta.TypeString),
		},
	}

	mobileGatewayAddSIMParam = &dsl.Model{
		Name:      mobileGatewayAPIName + "AddSIMRequest",
		NakedType: meta.Static(naked.SIMInfo{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "SIMID",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "ResourceID",
					JSON:    "resource_id",
				},
			},
		},
	}

	mobileGatewaySIMLogsModel = &dsl.Model{
		Name:      mobileGatewayAPIName + "SIMLogs",
		NakedType: meta.Static(naked.SIMLog{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("Date", meta.TypeTime),
			fields.Def("SessionStatus", meta.TypeString),
			fields.Def("ResourceID", meta.TypeString),
			fields.Def("IMEI", meta.TypeString),
			fields.Def("IMSI", meta.TypeString),
		},
	}

	mobileGatewayTrafficConfigModel = &dsl.Model{
		Name:      mobileGatewayAPIName + "TrafficControl",
		NakedType: meta.Static(naked.TrafficMonitoringConfig{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("TrafficQuotaInMB", meta.TypeInt),
			fields.Def("BandWidthLimitInKbps", meta.TypeInt),
			fields.Def("EmailNotifyEnabled", meta.TypeFlag, mapConvTag("EMailConfig.Enabled")),
			fields.Def("SlackNotifyEnabled", meta.TypeFlag, mapConvTag("SlackConfig.Enabled")),
			fields.Def("SlackNotifyWebhooksURL", meta.TypeString, mapConvTag("SlackConfig.IncomingWebhooksURL")),
			fields.Def("AutoTrafficShaping", meta.TypeFlag),
		},
	}

	mobileGatewayTrafficStatusModel = &dsl.Model{
		Name:      mobileGatewayAPIName + "TrafficStatus",
		NakedType: meta.Static(naked.TrafficStatus{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("UplinkBytes", meta.TypeStringNumber),
			fields.Def("DownlinkBytes", meta.TypeStringNumber),
			fields.Def("TrafficShaping", meta.TypeFlag),
		},
	}
)
