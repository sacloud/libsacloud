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
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource:   r,
					Name:       "ChangePlan",
					PathFormat: schema.IDAndSuffixPathFormat("plan"),
					Method:     http.MethodPut,
				}
				o.ResultFromEnvelope(serverView, &schema.EnvelopePayloadDesc{
					PayloadName: r.Name,
					PayloadType: meta.Static(naked.Server{}),
				}, "")
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgumentWithEnvelope(o, "plan", serverChangePlanParam),
				}
				return o
			}(),

			// insert cdrom
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "InsertCDROM",
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
					PathFormat: schema.IDAndSuffixPathFormat("cdrom"),
					Method:     http.MethodPut,
				}
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: meta.Static(naked.CDROM{}),
						PayloadName: "CDROM",
					},
				)
				return o
			}(),

			// eject cdrom
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "EjectCDROM",
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
					PathFormat: schema.IDAndSuffixPathFormat("cdrom"),
					Method:     http.MethodDelete,
				}
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: meta.Static(naked.CDROM{}),
						PayloadName: "CDROM",
					},
				)
				return o
			}(),

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
