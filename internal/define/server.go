package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var serverAPI = &schema.Resource{
	Name:       "Server",
	PathName:   "server",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(serverNakedType, findParameter, serverView),

			// create
			r.DefineOperationCreate(serverNakedType, serverCreateParam, serverView),

			// read
			r.DefineOperationRead(serverNakedType, serverView),

			// update
			r.DefineOperationUpdate(serverNakedType, serverUpdateParam, serverView),

			// delete
			r.DefineOperationDelete(),

			// change plan
			{
				Resource:        r,
				Name:            "ChangePlan",
				PathFormat:      schema.IDAndSuffixPathFormat("plan"),
				Method:          http.MethodPut,
				RequestEnvelope: schema.RequestEnvelopeFromModel(serverChangePlanParam),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgument("plan", serverChangePlanParam),
				},
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Name: r.Name,
					Type: meta.Static(naked.Server{}),
				}),
				Results: schema.Results{
					{
						SourceField: r.FieldName(schema.PayloadForms.Singular),
						DestField:   r.FieldName(schema.PayloadForms.Singular),
						IsPlural:    false,
						Model:       serverView,
					},
				},
			},

			// insert cdrom
			{
				Resource:   r,
				Name:       "InsertCDROM",
				PathFormat: schema.IDAndSuffixPathFormat("cdrom"),
				Method:     http.MethodPut,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: meta.Static(naked.CDROM{}),
						Name: "CDROM",
					},
				),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					{
						Name: "insertParam",
						Type: &schema.Model{
							Name: "InsertCDROMRequest",
							Fields: []*schema.FieldDesc{
								fields.ID(),
							},
							NakedType: meta.Static(naked.CDROM{}),
						},
						MapConvTag: "CDROM",
					},
				},
			},

			// eject cdrom
			{
				Resource:   r,
				Name:       "EjectCDROM",
				PathFormat: schema.IDAndSuffixPathFormat("cdrom"),
				Method:     http.MethodDelete,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: meta.Static(naked.CDROM{}),
						Name: "CDROM",
					},
				),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					{
						Name: "insertParam",
						Type: &schema.Model{
							Name: "EjectCDROMRequest",
							Fields: []*schema.FieldDesc{
								fields.ID(),
							},
							NakedType: meta.Static(naked.CDROM{}),
						},
						MapConvTag: "CDROM",
					},
				},
			},

			// power management(boot/shutdown/reset)
			r.DefineOperationBoot(),
			r.DefineOperationShutdown(),
			r.DefineOperationReset(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.cpuTimeModel()),
		}
	},
}

var (
	serverNakedType = meta.Static(naked.Server{})

	serverView = &schema.Model{
		Name: "Server",
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.HostName(),
			fields.InterfaceDriver(),
			// server plan
			fields.ServerPlanID(),
			fields.ServerPlanName(),
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			// zone
			fields.Zone(),
			// instance
			fields.InstanceHostName(),
			fields.InstanceHostInfoURL(),
			fields.InstanceStatus(),
			fields.InstanceBeforeStatus(),
			fields.InstanceStatusChangedAt(),
			fields.InstanceWarnings(),
			fields.InstanceWarningsValue(),

			// disks
			{
				Name: "Disks",
				Type: &schema.Model{
					Name:      diskModel.Name,
					Fields:    diskModel.Fields,
					NakedType: meta.Static(naked.Disk{}),
					IsArray:   true,
				},
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: ",recursive",
				},
			},
			fields.Interfaces(),

			fields.CDROMID(),

			fields.PrivateHostID(),
			fields.PrivateHostName(),

			fields.BundleInfo(),

			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	serverCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			// server plan
			fields.ServerPlanCPU(),
			fields.ServerPlanMemoryMB(),
			fields.ServerPlanCommitment(),
			fields.ServerPlanGeneration(),
			{
				Name: "ConnectedSwitches",
				Type: &schema.Model{
					Name: "ConnectedSwitch",
					Fields: []*schema.FieldDesc{
						fields.ID(),
						fields.Scope(),
					},
					IsArray:   true,
					NakedType: meta.Static(naked.ConnectedSwitch{}),
				},
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "[]ConnectedSwitches,recursive",
				},
			},
			fields.InterfaceDriver(),
			fields.HostName(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			{
				Name: "WaitDiskMigration",
				Type: meta.TypeFlag,
				Tags: &schema.FieldTags{
					MapConv: ",omitempty",
					JSON:    ",omitempty",
				},
			},
		},
	}

	serverUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	serverChangePlanParam = &schema.Model{
		Name: "ServerChangePlanRequest",
		Fields: []*schema.FieldDesc{
			fields.CPU(),
			fields.MemoryMB(),
			fields.Generation(),
			fields.ServerPlanCommitment(),
		},
		NakedType: meta.Static(naked.ServerPlan{}),
	}
)
