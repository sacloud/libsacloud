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

			// create distantly
			{
				Resource:   r,
				Name:       "CreateDistantly",
				PathFormat: schema.DefaultPathFormat,
				Method:     http.MethodPost,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: diskNakedType,
						Name: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						Type: diskDistantFromType,
						Name: "DistantFrom",
					},
				),
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
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				}),
				Results: schema.Results{
					{
						SourceField: "Disk",
						DestField:   diskModel.Name,
						IsPlural:    false,
						Model:       diskModel,
					},
				},
			},

			// config(DiskEdit)
			{
				Resource:        r,
				Name:            "Config",
				PathFormat:      schema.IDAndSuffixPathFormat("config"),
				Method:          http.MethodPut,
				RequestEnvelope: schema.RequestEnvelopeFromModel(diskEditParam),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					schema.PassthroughModelArgument("edit", diskEditParam),
				},
			},

			// create with config(DiskEdit)
			{
				Resource:   r,
				Name:       "CreateWithConfig",
				PathFormat: schema.DefaultPathFormat,
				Method:     http.MethodPost,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: diskNakedType,
						Name: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						Type: diskEditNakedType,
						Name: "Config",
					},
					&schema.EnvelopePayloadDesc{
						Type: meta.TypeFlag,
						Name: "BootAtAvailable",
					},
				),
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
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				}),
				Results: schema.Results{
					{
						SourceField: "Disk",
						DestField:   diskModel.Name,
						IsPlural:    false,
						Model:       diskModel,
					},
				},
			},

			{
				Resource:   r,
				Name:       "CreateWithConfigDistantly",
				PathFormat: schema.DefaultPathFormat,
				Method:     http.MethodPost,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: diskNakedType,
						Name: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						Type: diskEditNakedType,
						Name: "Config",
					},
					&schema.EnvelopePayloadDesc{
						Type: meta.TypeFlag,
						Name: "BootAtAvailable",
					},
					&schema.EnvelopePayloadDesc{
						Type: diskDistantFromType,
						Name: "DistantFrom",
					},
				),
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
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				}),
				Results: schema.Results{
					{
						SourceField: "Disk",
						DestField:   diskModel.Name,
						IsPlural:    false,
						Model:       diskModel,
					},
				},
			},

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
			{
				Resource:   r,
				Name:       "InstallDistantFrom",
				PathFormat: schema.IDAndSuffixPathFormat("install"),
				Method:     http.MethodPut,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: diskNakedType,
						Name: "Disk",
					},
					&schema.EnvelopePayloadDesc{
						Type: diskDistantFromType,
						Name: "DistantFrom",
					},
				),
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
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				}),
				Results: schema.Results{
					{
						SourceField: "Disk",
						DestField:   diskModel.Name,
						IsPlural:    false,
						Model:       diskModel,
					},
				},
			},

			{
				Resource:   r,
				Name:       "Install",
				PathFormat: schema.IDAndSuffixPathFormat("install"),
				Method:     http.MethodPut,
				RequestEnvelope: schema.RequestEnvelope(
					&schema.EnvelopePayloadDesc{
						Type: diskNakedType,
						Name: "Disk",
					},
				),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.ArgumentID,
					{
						Name:       "installParam",
						MapConvTag: "Disk",
						Type:       diskInstallParam,
					},
				},
				ResponseEnvelope: schema.ResponseEnvelope(&schema.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				}),
				Results: schema.Results{
					{
						SourceField: "Disk",
						DestField:   diskModel.Name,
						IsPlural:    false,
						Model:       diskModel,
					},
				},
			},

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
