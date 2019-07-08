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
	simAPIName     = "SIM"
	simAPIPathName = "commonserviceitem"
)

var simAPI = &dsl.Resource{
	Name:       simAPIName,
	PathName:   simAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
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
			PathFormat:   dsl.IDAndSuffixPathFormat("sim/ip"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "SIM",
					Type: meta.Static(naked.SIMAssignIPRequest{}),
					Tags: &dsl.FieldTags{
						JSON: "sim",
					},
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.MappableArgument("param", simAssignIPParam, "SIM"),
			},
		},

		// clearIP
		ops.WithIDAction(simAPIName, "ClearIP", http.MethodDelete, "sim/ip"),

		// IMEILock
		{
			ResourceName: simAPIName,
			Name:         "IMEILock",
			PathFormat:   dsl.IDAndSuffixPathFormat("sim/imeilock"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: "SIM",
					Type: meta.Static(naked.SIMIMEILockRequest{}),
					Tags: &dsl.FieldTags{
						JSON: "sim",
					},
				},
			),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.MappableArgument("param", simIMEILockParam, "SIM"),
			},
		},

		// IMEIUnlock
		ops.WithIDAction(simAPIName, "IMEIUnlock", http.MethodDelete, "sim/imeilock"),

		// Logs
		{
			ResourceName:     simAPIName,
			PathFormat:       dsl.IDAndSuffixPathFormat("sim/sessionlog"),
			Method:           http.MethodGet,
			Name:             "Logs",
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Name: "Logs",
				Type: meta.Static(naked.SIMLog{}),
			}),
			Results: dsl.Results{
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
			ResourceName:     simAPIName,
			Name:             "GetNetworkOperator",
			PathFormat:       dsl.IDAndSuffixPathFormat("sim/network_operator_config"),
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Name: "NetworkOperationConfigs",
				Type: meta.Static(naked.SIMNetworkOperatorConfig{}),
			}),
			Results: dsl.Results{
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
			PathFormat:      dsl.IDAndSuffixPathFormat("sim/network_operator_config"),
			Method:          http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(simNetworkOperatorsConfigView),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("configs", simNetworkOperatorsConfigView),
			},
		},

		// monitor
		ops.MonitorChild(simAPIName, "SIM", "sim/metrics",
			monitorParameter, monitors.linkModel()),
	},
}

var (
	simNakedType = meta.Static(naked.SIM{})

	simView = &dsl.Model{
		Name:      simAPIName,
		NakedType: simNakedType,
		Fields: []*dsl.FieldDesc{
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

	simCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(simAPIName),
		NakedType: simNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.SIMProviderClass(),
			fields.SIMICCID(),
			fields.SIMPassCode(),
		},
	}

	simUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(simAPIName),
		NakedType: simNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	simAssignIPParam = &dsl.Model{
		Name: "SIMAssignIPRequest",
		Fields: []*dsl.FieldDesc{
			{
				Name: "IP",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMAssignIPRequest{}),
	}
	simIMEILockParam = &dsl.Model{
		Name: "SIMIMEILockRequest",
		Fields: []*dsl.FieldDesc{
			{
				Name: "IMEI",
				Type: meta.TypeString,
			},
		},
		NakedType: meta.Static(naked.SIMIMEILockRequest{}),
	}

	simLogView = &dsl.Model{
		Name: "SIMLog",
		Fields: []*dsl.FieldDesc{
			fields.New("Date", meta.TypeTime),
			fields.New("SessionStatus", meta.TypeString),
			fields.New("ResourceID", meta.TypeString),
			fields.New("IMEI", meta.TypeString),
			fields.New("IMSI", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMLog{}),
	}
	simNetworkOperatorConfigView = &dsl.Model{
		Name: "SIMNetworkOperatorConfig",
		Fields: []*dsl.FieldDesc{
			fields.New("Allow", meta.TypeFlag),
			fields.New("CountryCode", meta.TypeString),
			fields.New("Name", meta.TypeString),
		},
		NakedType: meta.Static(naked.SIMNetworkOperatorConfig{}),
	}
	simNetworkOperatorsConfigView = &dsl.Model{
		Name: "SIMNetworkOperatorConfigs",
		Fields: []*dsl.FieldDesc{
			{
				Name: "NetworkOperatorConfigs",
				Type: &dsl.Model{
					Name: "SIMNetworkOperatorConfig",
					Fields: []*dsl.FieldDesc{
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
