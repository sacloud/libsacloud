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
	serverAPIName     = "Server"
	serverAPIPathName = "server"
)

var serverAPI = &schema.Resource{
	Name:       serverAPIName,
	PathName:   serverAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(serverAPIName, serverNakedType, findParameter, serverView),

		// create
		ops.Create(serverAPIName, serverNakedType, serverCreateParam, serverView),

		// read
		ops.Read(serverAPIName, serverNakedType, serverView),

		// update
		ops.Update(serverAPIName, serverNakedType, serverUpdateParam, serverView),

		// delete
		ops.Delete(serverAPIName),

		// change plan
		{
			ResourceName:    serverAPIName,
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
				Name: serverAPIName,
				Type: meta.Static(naked.Server{}),
			}),
			Results: schema.Results{
				{
					SourceField: names.ResourceFieldName(serverAPIName, schema.PayloadForms.Singular),
					DestField:   names.ResourceFieldName(serverAPIName, schema.PayloadForms.Singular),
					IsPlural:    false,
					Model:       serverView,
				},
			},
		},

		// insert cdrom
		{
			ResourceName: serverAPIName,
			Name:         "InsertCDROM",
			PathFormat:   schema.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodPut,
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
			ResourceName: serverAPIName,
			Name:         "EjectCDROM",
			PathFormat:   schema.IDAndSuffixPathFormat("cdrom"),
			Method:       http.MethodDelete,
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
		ops.Boot(serverAPIName),
		ops.Shutdown(serverAPIName),
		ops.Reset(serverAPIName),

		// monitor
		ops.Monitor(serverAPIName, monitorParameter, monitors.cpuTimeModel()),
	},
}

var (
	serverNakedType = meta.Static(naked.Server{})

	serverView = &schema.Model{
		Name:      serverAPIName,
		NakedType: serverNakedType,
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
		Name:      names.CreateParameterName(serverAPIName),
		NakedType: serverNakedType,
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
		Name:      names.UpdateParameterName(serverAPIName),
		NakedType: serverNakedType,
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
