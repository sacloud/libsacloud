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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
	mobileGatewayNakedType = meta.Static(naked.MobileGateway{})

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
			// plan
			fields.AppliancePlanID(),
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
			fields.New("Settings", models.mobileGatewaySetting(), mapConvTag(",omitempty,recursive")),
			fields.SettingsHash(),
		},
	}

	mobileGatewayCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(mobileGatewayAPIName),
		NakedType: mobileGatewayNakedType,
		Fields: []*dsl.FieldDesc{
			fields.MobileGatewayClass(),
			fields.ApplianceSwitchShared(),
			fields.AppliancePlanID(),
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
		},
	}

	mobileGatewayDNSModel = &dsl.Model{
		Name:      "MobileGatewayDNSSetting",
		NakedType: meta.Static(naked.MobileGatewaySIMGroup{}),
		Fields: []*dsl.FieldDesc{
			fields.New("DNS1", meta.TypeString),
			fields.New("DNS2", meta.TypeString),
		},
	}

	mobileGatewaySIMRouteParam = &dsl.Model{
		Name:      "MobileGatewaySIMRouteParam",
		NakedType: meta.Static(naked.MobileGatewaySIMRoute{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.New("ResourceID", meta.TypeString),
			fields.New("Prefix", meta.TypeString),
		},
	}
	mobileGatewaySIMRouteView = &dsl.Model{
		Name:      "MobileGatewaySIMRoute",
		NakedType: meta.Static(naked.MobileGatewaySIMRoute{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.New("ResourceID", meta.TypeString),
			fields.New("Prefix", meta.TypeString),
			fields.New("ICCID", meta.TypeString),
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
			fields.New("Date", meta.TypeTime),
			fields.New("SessionStatus", meta.TypeString),
			fields.New("ResourceID", meta.TypeString),
			fields.New("IMEI", meta.TypeString),
			fields.New("IMSI", meta.TypeString),
		},
	}

	mobileGatewayTrafficConfigModel = &dsl.Model{
		Name:      mobileGatewayAPIName + "TrafficControl",
		NakedType: meta.Static(naked.TrafficMonitoringConfig{}),
		Fields: []*dsl.FieldDesc{
			fields.New("TrafficQuotaInMB", meta.TypeInt),
			fields.New("BandWidthLimitInKbps", meta.TypeInt),
			fields.New("EmailNotifyEnabled", meta.TypeFlag, mapConvTag("EMailConfig.Enabled")),
			fields.New("SlackNotifyEnabled", meta.TypeFlag, mapConvTag("SlackConfig.Enabled")),
			fields.New("SlackNotifyWebhooksURL", meta.TypeString, mapConvTag("SlackConfig.IncomingWebhooksURL")),
			fields.New("AutoTrafficShaping", meta.TypeFlag),
		},
	}

	mobileGatewayTrafficStatusModel = &dsl.Model{
		Name:      mobileGatewayAPIName + "TrafficStatus",
		NakedType: meta.Static(naked.TrafficStatus{}),
		Fields: []*dsl.FieldDesc{
			fields.New("UplinkBytes", meta.TypeStringNumber),
			fields.New("DownlinkBytes", meta.TypeStringNumber),
			fields.New("TrafficShaping", meta.TypeFlag),
		},
	}
)
