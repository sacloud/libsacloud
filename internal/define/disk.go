package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var diskAPI = &schema.Resource{
	Name:       "Disk",
	PathName:   "disk",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(diskNakedType, findParameter, diskModel),

			// create
			r.DefineOperationCreate(diskNakedType, diskCreateParam, diskModel),

			// TODO あとで直す
			// create distantly
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "CreateDistantly",
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						{
							Name:       "createParam",
							MapConvTag: "Disk",
							Type:       diskCreateParam,
						},
						{
							Name:       "distantFrom",
							MapConvTag: "DistantFrom",
							Type:       diskDistantFromType,
						},
					},
					PathFormat: schema.DefaultPathFormat,
					Method:     http.MethodPost,
				}
				o.ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name)
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: diskNakedType,
						PayloadName: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: diskDistantFromType,
						PayloadName: "DistantFrom",
					},
				)
				return o
			}(),

			// TODO あとで直す
			// config(DiskEdit)
			func() *schema.Operation {
				o := &schema.Operation{
					Resource:   r,
					Name:       "Config",
					PathFormat: schema.IDAndSuffixPathFormat("config"),
					Method:     http.MethodPut,
				}
				o.Arguments = schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgumentWithEnvelope(o, "edit", diskEditParam),
				}
				return o
			}(),

			// TODO あとで直す
			// create with config(DiskEdit)
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "CreateWithConfig",
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						{
							Name:       "createParam",
							MapConvTag: "Disk",
							Type:       diskCreateParam,
						},
						{
							Name:       "editParam",
							MapConvTag: "Config",
							Type:       diskEditParam,
						},
						{
							Name:       "bootAtAvailable",
							Type:       meta.TypeFlag,
							MapConvTag: "BootAtAvailable",
						},
					},
					PathFormat: schema.DefaultPathFormat,
					Method:     http.MethodPost,
				}
				o.ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name)
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: diskNakedType,
						PayloadName: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: diskEditNakedType,
						PayloadName: "Config",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: meta.TypeFlag,
						PayloadName: "BootAtAvailable",
					},
				)
				return o
			}(),

			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "CreateWithConfigDistantly",
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						{
							Name:       "createParam",
							MapConvTag: "Disk",
							Type:       diskCreateParam,
						},
						{
							Name:       "editParam",
							MapConvTag: "Config",
							Type:       diskEditParam,
						},
						{
							Name:       "bootAtAvailable",
							Type:       meta.TypeFlag,
							MapConvTag: "BootAtAvailable",
						},
						{
							Name:       "distantFrom",
							Type:       diskDistantFromType,
							MapConvTag: "DistantFrom",
						},
					},
					PathFormat: schema.DefaultPathFormat,
					Method:     http.MethodPost,
				}
				o.ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name)
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: diskNakedType,
						PayloadName: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: diskEditNakedType,
						PayloadName: "Config",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: meta.TypeFlag,
						PayloadName: "BootAtAvailable",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: diskDistantFromType,
						PayloadName: "DistantFrom",
					},
				)
				return o
			}(),

			// to blank
			r.DefineSimpleOperation("ToBlank", http.MethodPut, "to/blank"),

			// resize partition
			r.DefineSimpleOperation("ResizePartition", http.MethodPut, "resize-partition"),

			// connect to server
			r.DefineSimpleOperation("ConnectToServer", http.MethodPut, "to/server/{{.serverID}}",
				&schema.Argument{
					Name: "serverID",
					Type: meta.TypeID,
				},
			),

			// disconnect from server
			r.DefineSimpleOperation("DisconnectFromServer", http.MethodDelete, "to/server"),

			// install
			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "InstallDistantFrom",
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						schema.ArgumentID,
						{
							Name:       "installParam",
							MapConvTag: "Disk",
							Type:       diskInstallParam,
						},
						{
							Name:       "distantFrom",
							MapConvTag: "DistantFrom",
							Type:       diskDistantFromType,
						},
					},
					PathFormat: schema.IDAndSuffixPathFormat("install"),
					Method:     http.MethodPut,
				}
				o.ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name)
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: diskNakedType,
						PayloadName: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						PayloadType: diskDistantFromType,
						PayloadName: "DistantFrom",
					},
				)
				return o
			}(),

			// TODO あとで直す
			func() *schema.Operation {
				o := &schema.Operation{
					Resource: r,
					Name:     "Install",
					Arguments: schema.Arguments{
						schema.ArgumentZone,
						schema.ArgumentID,
						{
							Name:       "installParam",
							MapConvTag: "Disk",
							Type:       diskInstallParam,
						},
					},
					PathFormat: schema.IDAndSuffixPathFormat("install"),
					Method:     http.MethodPut,
				}
				o.ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name)
				o.RequestEnvelope = schema.RequestEnvelope(o,
					&schema.EnvelopePayloadDesc{
						PayloadType: diskNakedType,
						PayloadName: "Disk",
					},
				)
				return o
			}(),

			// read
			r.DefineOperationRead(diskNakedType, diskModel),

			// update
			r.DefineOperationUpdate(diskNakedType, diskUpdateParam, diskModel),

			// delete
			r.DefineOperationDelete(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.diskModel()),
		}
	},
}

var (
	diskNakedType       = meta.Static(naked.Disk{})
	diskEditNakedType   = meta.Static(naked.DiskEdit{})
	diskDistantFromType = meta.Static([]types.ID{})

	diskModel = &schema.Model{
		Name: "Disk",
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.DiskConnection(),
			fields.DiskConnectionOrder(),
			fields.DiskReinstallCount(),
			fields.SizeMB(),
			fields.MigratedMB(),
			fields.DiskPlanID(),
			fields.DiskPlanName(),
			fields.DiskPlanStorageClass(),
			fields.SourceDiskID(),
			fields.SourceDiskAvailability(),
			fields.SourceArchiveID(),
			fields.SourceArchiveAvailability(),
			fields.BundleInfo(),
			fields.Storage(),
			fields.ServerID(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	diskCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.DiskPlanID(),
			fields.DiskConnection(),
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.ServerID(),
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	diskUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.DiskConnection(),
		},
	}

	diskEditParam = models.diskEdit()

	diskInstallParam = &schema.Model{
		Name: "DiskInstallRequest",
		Fields: []*schema.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.SizeMB(),
		},
		NakedType: diskNakedType,
	}
)
