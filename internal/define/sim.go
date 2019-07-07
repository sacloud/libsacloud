package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var simAPI = &schema.Resource{
	Name:       "SIM",
	PathName:   "commonserviceitem",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationCommonServiceItemFind(simNakedType, findParameter, simView),

			// create
			r.DefineOperationCommonServiceItemCreate(simNakedType, simCreateParam, simView),

			// read
			r.DefineOperationCommonServiceItemRead(simNakedType, simView),

			// update
			r.DefineOperationCommonServiceItemUpdate(simNakedType, simUpdateParam, simView),

			// delete
			r.DefineOperationDelete(),

			// activate
			r.DefineSimpleOperation("Activate", http.MethodPut, "sim/activate"),
			// deactivate
			r.DefineSimpleOperation("Deactivate", http.MethodPut, "sim/deactivate"),

			// assignIP
			{
				Resource:   r,
				Name:       "AssignIP",
				PathFormat: schema.IDAndSuffixPathFormat("sim/ip"),
				Method:     http.MethodPut,
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
			r.DefineSimpleOperation("ClearIP", http.MethodDelete, "sim/ip"),

			// IMEILock
			{
				Resource:   r,
				Name:       "IMEILock",
				PathFormat: schema.IDAndSuffixPathFormat("sim/imeilock"),
				Method:     http.MethodPut,
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
			r.DefineSimpleOperation("IMEIUnlock", http.MethodDelete, "sim/imeilock"),

			// Logs
			{
				Resource:   r,
				PathFormat: schema.IDAndSuffixPathFormat("sim/sessionlog"),
				Method:     http.MethodGet,
				Name:       "Logs",
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
				Resource:   r,
				Name:       "GetNetworkOperator",
				PathFormat: schema.IDAndSuffixPathFormat("sim/network_operator_config"),
				Method:     http.MethodGet,
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
				Resource:        r,
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
			r.DefineOperationMonitorChild("SIM", "sim/metrics",
				monitorParameter, monitors.linkModel()),
		}
	},
}

var (
	simNakedType = meta.Static(naked.SIM{})

	simView = &schema.Model{
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
