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
			r.DefineOperation("CreateDistantly").
				Method(http.MethodPost).
				PathFormat(schema.DefaultPathFormat).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskDistantFromType,
					PayloadName: "DistantFrom",
				}).
				Argument(schema.ArgumentZone).
				Argument(&schema.Argument{
					Name:       "createParam",
					MapConvTag: "Disk",
					Type:       diskCreateParam,
				}).
				Argument(&schema.Argument{
					Name:       "distantFrom",
					MapConvTag: "DistantFrom",
					Type:       diskDistantFromType,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name),

			// config(DiskEdit)
			r.DefineOperation("Config").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("config")).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				PassthroughModelArgumentWithEnvelope("edit", diskEditParam),

			// create with config(DiskEdit)
			r.DefineOperation("CreateWithConfig").
				Method(http.MethodPost).
				PathFormat(schema.DefaultPathFormat).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskEditNakedType,
					PayloadName: "Config",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: meta.TypeFlag,
					PayloadName: "BootAtAvailable",
				}).
				Argument(schema.ArgumentZone).
				Argument(&schema.Argument{
					Name:       "createParam",
					MapConvTag: "Disk",
					Type:       diskCreateParam,
				}).
				Argument(&schema.Argument{
					Name:       "editParam",
					MapConvTag: "Config",
					Type:       diskEditParam,
				}).
				Argument(&schema.Argument{
					Name:       "bootAtAvailable",
					Type:       meta.TypeFlag,
					MapConvTag: "BootAtAvailable",
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name),

			r.DefineOperation("CreateWithConfigDistantly").
				Method(http.MethodPost).
				PathFormat(schema.DefaultPathFormat).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskEditNakedType,
					PayloadName: "Config",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: meta.TypeFlag,
					PayloadName: "BootAtAvailable",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskDistantFromType,
					PayloadName: "DistantFrom",
				}).
				Argument(schema.ArgumentZone).
				Argument(&schema.Argument{
					Name:       "createParam",
					MapConvTag: "Disk",
					Type:       diskCreateParam,
				}).
				Argument(&schema.Argument{
					Name:       "editParam",
					MapConvTag: "Config",
					Type:       diskEditParam,
				}).
				Argument(&schema.Argument{
					Name:       "bootAtAvailable",
					Type:       meta.TypeFlag,
					MapConvTag: "BootAtAvailable",
				}).
				Argument(&schema.Argument{
					Name:       "distantFrom",
					Type:       diskDistantFromType,
					MapConvTag: "DistantFrom",
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name),

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
			r.DefineOperation("InstallDistantFrom").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("install")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskDistantFromType,
					PayloadName: "DistantFrom",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.Argument{
					Name:       "installParam",
					MapConvTag: "Disk",
					Type:       diskInstallParam,
				}).
				Argument(&schema.Argument{
					Name:       "distantFrom",
					MapConvTag: "DistantFrom",
					Type:       diskDistantFromType,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name),

			r.DefineOperation("Install").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("install")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.Argument{
					Name:       "installParam",
					MapConvTag: "Disk",
					Type:       diskInstallParam,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: diskNakedType,
					PayloadName: "Disk",
				}, diskModel.Name),

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
