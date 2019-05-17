package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedServer := meta.Static(naked.Server{})

	server := &schema.Model{
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

			{
				Name: "Interfaces",
				Type: &schema.Model{
					Name:      models.interfaceModel().Name,
					Fields:    models.interfaceModel().Fields,
					IsArray:   true,
					NakedType: meta.Static(naked.Interface{}),
				},
				Tags: &schema.FieldTags{
					JSON:    ",omitempty",
					MapConv: "[]Interfaces,recursive",
				},
			},

			fields.PrivateHostID(),
			fields.PrivateHostName(),

			fields.BundleInfo(),
			fields.Storage(),

			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	createParam := &schema.Model{
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

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	changePlanParam := &schema.Model{
		Name: "ServerChangePlanRequest",
		Fields: []*schema.FieldDesc{
			fields.CPU(),
			fields.MemoryMB(),
			fields.Generation(),
			fields.ServerPlanCommitment(),
		},
		NakedType: meta.Static(naked.ServerPlan{}),
	}

	Resources.DefineWith("Server", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedServer, findParameter, server),

			// create
			r.DefineOperationCreate(nakedServer, createParam, server),

			// read
			r.DefineOperationRead(nakedServer, server),

			// update
			r.DefineOperationUpdate(nakedServer, updateParam, server),

			// delete
			r.DefineOperationDelete(),

			// change plan
			r.DefineOperation("ChangePlan").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("plan")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				PassthroughModelArgumentWithEnvelope("plan", changePlanParam).
				ResultFromEnvelope(server, &schema.EnvelopePayloadDesc{
					PayloadName: server.Name,
					PayloadType: meta.Static(naked.Server{}),
				}),

			// insert cdrom
			r.DefineOperation("InsertCDROM").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("cdrom")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.CDROM{}),
					PayloadName: "CDROM",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.MappableArgument{
					Name: "insertParam",
					Model: &schema.Model{
						Name: "InsertCDROMRequest",
						Fields: []*schema.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					Destination: "CDROM",
				}),

			// eject cdrom
			r.DefineOperation("EjectCDROM").
				Method(http.MethodDelete).
				PathFormat(schema.IDAndSuffixPathFormat("cdrom")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: meta.Static(naked.CDROM{}),
					PayloadName: "CDROM",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.MappableArgument{
					Name: "insertParam",
					Model: &schema.Model{
						Name: "EjectCDROMRequest",
						Fields: []*schema.FieldDesc{
							fields.ID(),
						},
						NakedType: meta.Static(naked.CDROM{}),
					},
					Destination: "CDROM",
				}),

			// power management(boot/shutdown/reset)
			r.DefineOperationBoot(),
			r.DefineOperationShutdown(),
			r.DefineOperationReset(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.cpuTimeModel()),
		)
	})
}
