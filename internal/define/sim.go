package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	simAPIName     = "SIM"
	simAPIPathName = "commonserviceitem"
)

var simAPI = &schema.Resource{
	Name:       simAPIName,
	PathName:   simAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.FindCommonServiceItem(simAPIName, simNakedType, findParameter, simView),

		// create
		ops.CreateCommonServiceItem(simAPIName, simNakedType, simCreateParam, simView),

		// read
		ops.ReadCommonServiceItem(simAPIName, simNakedType, simView),

		// update
		ops.UpdateCommonServiceItem(simAPIName, simNakedType, simUpdateParam, simView),

		// delete
		ops.Delete(simAPIName),

		// activate
		ops.WithIDAction(simAPIName, "Activate", http.MethodPut, "sim/activate"),
		// deactivate
		ops.WithIDAction(simAPIName, "Deactivate", http.MethodPut, "sim/deactivate"),

		// assignIP
		{
			ResourceName: simAPIName,
			Name:         "AssignIP",
			PathFormat:   schema.IDAndSuffixPathFormat("sim/ip"),
			Method:       http.MethodPut,
			RequestEnvelope: schema.RequestEnvelope(
				&schema.EnvelopePayloadDesc{
					Name: "SIM",
					Type: meta.Static(naked.SIMAssignIPRequest{}),
					Tags: &schema.FieldTags{
						JSON: "sim",
					},
				},
			),
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.ArgumentID,
				schema.MappableArgument("param", simAssignIPParam, "SIM"),
			},
		},

		// clearIP
		ops.WithIDAction(simAPIName, "ClearIP", http.MethodDelete, "sim/ip"),

		// IMEILock
		{
			ResourceName: simAPIName,
			Name:         "IMEILock",
			PathFormat:   schema.IDAndSuffixPathFormat("sim/imeilock"),
			Method:       http.MethodPut,
			RequestEnvelope: schema.RequestEnvelope(
				&schema.EnvelopePayloadDesc{
					Name: "SIM",
					Type: meta.Static(naked.SIMIMEILockRequest{}),
					Tags: &schema.FieldTags{
						JSON: "sim",
					},
				},
			),
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.ArgumentID,
				schema.MappableArgument("param", simIMEILockParam, "SIM"),
			},
		},

		// IMEIUnlock
		ops.WithIDAction(simAPIName, "IMEIUnlock", http.MethodDelete, "sim/imeilock"),

		// Logs
		{
			ResourceName: simAPIName,
			PathFormat:   schema.IDAndSuffixPathFormat("sim/sessionlog"),
			Method:       http.MethodGet,
			Name:         "Logs",
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.ArgumentID,
			},
			ResponseEnvelope: schema.ResponseEnvelopePlural(&schema.EnvelopePayloadDesc{
				Name: "Logs",
				Type: meta.Static(naked.SIMLog{}),
			}),
			Results: schema.Results{
				{
					SourceField: "Logs",
					DestField:   "Logs",
					IsPlural:    true,
					Model:       simLogView,
				},
			},
		},

		// GetNetworkOperator
		{
			ResourceName: simAPIName,
			Name:         "GetNetworkOperator",
			PathFormat:   schema.IDAndSuffixPathFormat("sim/network_operator_config"),
			Method:       http.MethodGet,
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.ArgumentID,
			},
			ResponseEnvelope: schema.ResponseEnvelopePlural(&schema.EnvelopePayloadDesc{
				Name: "NetworkOperationConfigs",
				Type: meta.Static(naked.SIMNetworkOperatorConfig{}),
			}),
			Results: schema.Results{
				{
					SourceField: "NetworkOperationConfigs",
					DestField:   "Configs",
					IsPlural:    true,
					Model:       simNetworkOperatorConfigView,
				},
			},
		},

		// SetNetworkOperator
		{
			ResourceName:    simAPIName,
			Name:            "SetNetworkOperator",
			PathFormat:      schema.IDAndSuffixPathFormat("sim/network_operator_config"),
			Method:          http.MethodPut,
			RequestEnvelope: schema.RequestEnvelopeFromModel(simNetworkOperatorsConfigView),
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.ArgumentID,
				schema.PassthroughModelArgument("configs", simNetworkOperatorsConfigView),
			},
		},

		// monitor
		ops.MonitorChild(simAPIName, "SIM", "sim/metrics",
			monitorParameter, monitors.linkModel()),
	},
}

var (
	simNakedType = meta.Static(naked.SIM{})

	simView = &schema.Model{
		Name:      simAPIName,
		NakedType: simNakedType,
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.Class(),
			fields.SIMICCID(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	simCreateParam = &schema.Model{
		Name:      names.CreateParameterName(simAPIName),
		NakedType: simNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.SIMProviderClass(),
			fields.SIMICCID(),
			fields.SIMPassCode(),
		},
	}

	simUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(simAPIName),
		NakedType: simNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	simAssignIPParam = &schema.Model{
		Name: "SIMAssignIPRequest",
		Fields: []*schema.FieldDesc{
			{
				Name: "IP",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMAssignIPRequest{}),
	}
	simIMEILockParam = &schema.Model{
		Name: "SIMIMEILockRequest",
		Fields: []*schema.FieldDesc{
			{
				Name: "IMEI",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMIMEILockRequest{}),
	}

	simLogView = &schema.Model{
		Name: "SIMLog",
		Fields: []*schema.FieldDesc{
			fields.New("Date", meta.TypeTime),
			fields.New("SessionStatus", meta.TypeString),
			fields.New("ResourceID", meta.TypeString),
			fields.New("IMEI", meta.TypeString),
			fields.New("IMSI", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMLog{}),
	}
	simNetworkOperatorConfigView = &schema.Model{
		Name: "SIMNetworkOperatorConfig",
		Fields: []*schema.FieldDesc{
			fields.New("Allow", meta.TypeFlag),
			fields.New("CountryCode", meta.TypeString),
			fields.New("Name", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMNetworkOperatorConfig{}),
	}
	simNetworkOperatorsConfigView = &schema.Model{
		Name: "SIMNetworkOperatorConfigs",
		Fields: []*schema.FieldDesc{
			{
				Name: "NetworkOperatorConfigs",
				Type: &schema.Model{
					Name: "SIMNetworkOperatorConfig",
					Fields: []*schema.FieldDesc{
						fields.New("Allow", meta.TypeFlag),
						fields.New("CountryCode", meta.TypeString),
						fields.New("Name", meta.TypeString),
					},
					NakedType: meta.Static(naked.SIMNetworkOperatorConfig{}),
					IsArray:   true,
				},
			},
		},
		NakedType: meta.Static(naked.SIMNetworkOperatorConfigs{}),
	}
)
