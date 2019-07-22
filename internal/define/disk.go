package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	diskAPIName     = "Disk"
	diskAPIPathName = "disk"
)

var diskAPI = &dsl.Resource{
	Name:       diskAPIName,
	PathName:   diskAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(diskAPIName, diskNakedType, findParameter, diskModel),

		// create
		{
			ResourceName: diskAPIName,
			Name:         "Create",
			PathFormat:   dsl.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				},
				&dsl.EnvelopePayloadDesc{
					Type: diskDistantFromType,
					Name: "DistantFrom",
				},
			),
			Arguments: dsl.Arguments{
				{
					Name:       "createParam",
					MapConvTag: "Disk,recursive",
					Type:       diskCreateParam,
				},
				{
					Name:       "distantFrom",
					MapConvTag: "DistantFrom",
					Type:       diskDistantFromType,
				},
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: diskNakedType,
				Name: "Disk",
			}),
			Results: dsl.Results{
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
			ResourceName:    diskAPIName,
			Name:            "Config",
			PathFormat:      dsl.IDAndSuffixPathFormat("config"),
			Method:          http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(diskEditParam),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.PassthroughModelArgument("edit", diskEditParam),
			},
		},

		// create with config(DiskEdit)
		{
			ResourceName: diskAPIName,
			Name:         "CreateWithConfig",
			PathFormat:   dsl.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				},
				&dsl.EnvelopePayloadDesc{
					Type: diskEditNakedType,
					Name: "Config",
				},
				&dsl.EnvelopePayloadDesc{
					Type: meta.TypeFlag,
					Name: "BootAtAvailable",
				},
				&dsl.EnvelopePayloadDesc{
					Type: diskDistantFromType,
					Name: "DistantFrom",
				},
			),
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: diskNakedType,
				Name: "Disk",
			}),
			Arguments: dsl.Arguments{
				{
					Name:       "createParam",
					MapConvTag: "Disk,recursive",
					Type:       diskCreateParam,
				},
				{
					Name:       "editParam",
					MapConvTag: "Config,recursive",
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
			Results: dsl.Results{
				{
					SourceField: "Disk",
					DestField:   diskModel.Name,
					IsPlural:    false,
					Model:       diskModel,
				},
			},
		},

		// to blank
		ops.WithIDAction(diskAPIName, "ToBlank", http.MethodPut, "to/blank"),

		// resize partition
		ops.WithIDAction(diskAPIName, "ResizePartition", http.MethodPut, "resize-partition"),

		// connect to server
		ops.WithIDAction(diskAPIName, "ConnectToServer", http.MethodPut, "to/server/{{.serverID}}",
			&dsl.Argument{
				Name: "serverID",
				Type: meta.TypeID,
			},
		),

		// disconnect from server
		ops.WithIDAction(diskAPIName, "DisconnectFromServer", http.MethodDelete, "to/server"),

		// install
		{
			ResourceName: diskAPIName,
			Name:         "Install",
			PathFormat:   dsl.IDAndSuffixPathFormat("install"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(
				&dsl.EnvelopePayloadDesc{
					Type: diskNakedType,
					Name: "Disk",
				},
				&dsl.EnvelopePayloadDesc{
					Type: diskDistantFromType,
					Name: "DistantFrom",
				},
			),
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: diskNakedType,
				Name: "Disk",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name:       "installParam",
					MapConvTag: "Disk,recursive",
					Type:       diskInstallParam,
				},
				{
					Name:       "distantFrom",
					MapConvTag: "DistantFrom",
					Type:       diskDistantFromType,
				},
			},
			Results: dsl.Results{
				{
					SourceField: "Disk",
					DestField:   diskModel.Name,
					IsPlural:    false,
					Model:       diskModel,
				},
			},
		},

		// read
		ops.Read(diskAPIName, diskNakedType, diskModel),

		// update
		ops.Update(diskAPIName, diskNakedType, diskUpdateParam, diskModel),

		// delete
		ops.Delete(diskAPIName),

		// monitor
		ops.Monitor(diskAPIName, monitorParameter, monitors.diskModel()),
	},
}

var (
	diskNakedType       = meta.Static(naked.Disk{})
	diskEditNakedType   = meta.Static(naked.DiskEdit{})
	diskDistantFromType = meta.Static([]types.ID{})

	diskModel = &dsl.Model{
		Name:      diskAPIName,
		NakedType: diskNakedType,
		Fields: []*dsl.FieldDesc{
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

	diskCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(diskAPIName),
		NakedType: diskNakedType,
		Fields: []*dsl.FieldDesc{
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

	diskUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(diskAPIName),
		NakedType: diskNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.DiskConnection(),
		},
	}

	diskEditParam = models.diskEdit()

	diskInstallParam = &dsl.Model{
		Name: "DiskInstallRequest",
		Fields: []*dsl.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.SizeMB(),
		},
		NakedType: diskNakedType,
	}
)
