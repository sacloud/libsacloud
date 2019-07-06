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
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("AssignIP").
					Argument(schema.ArgumentZone).
					Argument(schema.ArgumentID).
					RequestEnvelope(&schema.EnvelopePayloadDesc{
						PayloadName: "SIM",
						PayloadType: meta.Static(naked.SIMAssignIPRequest{}),
						Tags: &schema.FieldTags{
							JSON: "sim",
						},
					}).
					MappableArgument("param", simAssignIPParam)
				o.PathFormat = schema.IDAndSuffixPathFormat("sim/ip")
				o.Method = http.MethodPut
				return o
			}(),

			// clearIP
			r.DefineSimpleOperation("ClearIP", http.MethodDelete, "sim/ip"),

			// IMEILock
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("IMEILock").
					Argument(schema.ArgumentZone).
					Argument(schema.ArgumentID).
					RequestEnvelope(&schema.EnvelopePayloadDesc{
						PayloadName: "SIM",
						PayloadType: meta.Static(naked.SIMIMEILockRequest{}),
						Tags: &schema.FieldTags{
							JSON: "sim",
						},
					}).
					MappableArgument("param", simIMEILockParam)
				o.PathFormat = schema.IDAndSuffixPathFormat("sim/imeilock")
				o.Method = http.MethodPut
				return o
			}(),

			// IMEIUnlock
			r.DefineSimpleOperation("IMEIUnlock", http.MethodDelete, "sim/imeilock"),

			// Logs
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("Logs").
					Argument(schema.ArgumentZone).
					Argument(schema.ArgumentID).
					ResultPluralFromEnvelope(simLogView, &schema.EnvelopePayloadDesc{
						PayloadName: "Logs",
						PayloadType: meta.Static(naked.SIMLog{}),
					}, "Logs")
				o.PathFormat = schema.IDAndSuffixPathFormat("sim/sessionlog")
				o.Method = http.MethodGet
				return o
			}(),

			// GetNetworkOperator
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("GetNetworkOperator").
					Argument(schema.ArgumentZone).
					Argument(schema.ArgumentID).
					ResultPluralFromEnvelope(simNetworkOperatorConfigView, &schema.EnvelopePayloadDesc{
						PayloadName: "NetworkOperationConfigs",
						PayloadType: meta.Static(naked.SIMNetworkOperatorConfig{}),
					}, "Configs")
				o.PathFormat = schema.IDAndSuffixPathFormat("sim/network_operator_config")
				o.Method = http.MethodGet
				return o
			}(),

			// SetNetworkOperator
			// TODO あとで直す
			func() *schema.Operation {
				o := r.DefineOperation("SetNetworkOperator").
					Argument(schema.ArgumentZone).
					Argument(schema.ArgumentID).
					PassthroughModelArgumentWithEnvelope("configs", simNetworkOperatorsConfigView)
				o.PathFormat = schema.IDAndSuffixPathFormat("sim/network_operator_config")
				o.Method = http.MethodPut
				return o
			}(),

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
