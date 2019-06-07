package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

var diskModel = &schema.Model{
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

func init() {
	nakedDisk := meta.Static(naked.Disk{})
	nakedDiskEdit := meta.Static(naked.DiskEdit{})

	createParam := &schema.Model{
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

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
			fields.DiskConnection(),
		},
	}

	diskEditParam := models.diskEdit()

	installParam := &schema.Model{
		Name: "DiskInstallRequest",
		Fields: []*schema.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.SizeMB(),
		},
		NakedType: nakedDisk,
	}

	distantFromType := meta.Static([]types.ID{})

	diskAPI := &schema.Resource{
		Name:       "Disk",
		PathName:   "disk",
		PathSuffix: schema.CloudAPISuffix,
	}

	diskAPI.Operations = []*schema.Operation{
		// find
		diskAPI.DefineOperationFind(nakedDisk, findParameter, diskModel),

		// create
		diskAPI.DefineOperationCreate(nakedDisk, createParam, diskModel),

		// create distantly
		diskAPI.DefineOperation("CreateDistantly").
			Method(http.MethodPost).
			PathFormat(schema.DefaultPathFormat).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: distantFromType,
				PayloadName: "DistantFrom",
			}).
			Argument(schema.ArgumentZone).
			Argument(&schema.Argument{
				Name:       "createParam",
				MapConvTag: "Disk",
				Type:       createParam,
			}).
			Argument(&schema.Argument{
				Name:       "distantFrom",
				MapConvTag: "DistantFrom",
				Type:       distantFromType,
			}).
			ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}),

		// config(DiskEdit)
		diskAPI.DefineOperation("Config").
			Method(http.MethodPut).
			PathFormat(schema.IDAndSuffixPathFormat("config")).
			Argument(schema.ArgumentZone).
			Argument(schema.ArgumentID).
			PassthroughModelArgumentWithEnvelope("edit", diskEditParam),

		// create with config(DiskEdit)
		diskAPI.DefineOperation("CreateWithConfig").
			Method(http.MethodPost).
			PathFormat(schema.DefaultPathFormat).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDiskEdit,
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
				Type:       createParam,
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
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}),

		diskAPI.DefineOperation("CreateWithConfigDistantly").
			Method(http.MethodPost).
			PathFormat(schema.DefaultPathFormat).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDiskEdit,
				PayloadName: "Config",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: meta.TypeFlag,
				PayloadName: "BootAtAvailable",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: distantFromType,
				PayloadName: "DistantFrom",
			}).
			Argument(schema.ArgumentZone).
			Argument(&schema.Argument{
				Name:       "createParam",
				MapConvTag: "Disk",
				Type:       createParam,
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
				Type:       distantFromType,
				MapConvTag: "DistantFrom",
			}).
			ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}),

		// to blank
		diskAPI.DefineSimpleOperation("ToBlank", http.MethodPut, "to/blank"),

		// resize partition
		diskAPI.DefineSimpleOperation("ResizePartition", http.MethodPut, "resize-partition"),

		// connect to server
		diskAPI.DefineSimpleOperation("ConnectToServer", http.MethodPut, "to/server/{{.serverID}}",
			&schema.Argument{
				Name: "serverID",
				Type: meta.TypeID,
			},
		),

		// disconnect from server
		diskAPI.DefineSimpleOperation("DisconnectFromServer", http.MethodDelete, "to/server"),

		// install
		diskAPI.DefineOperation("InstallDistantFrom").
			Method(http.MethodPut).
			PathFormat(schema.IDAndSuffixPathFormat("install")).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: distantFromType,
				PayloadName: "DistantFrom",
			}).
			Argument(schema.ArgumentZone).
			Argument(schema.ArgumentID).
			Argument(&schema.Argument{
				Name:       "installParam",
				MapConvTag: "Disk",
				Type:       installParam,
			}).
			Argument(&schema.Argument{
				Name:       "distantFrom",
				MapConvTag: "DistantFrom",
				Type:       distantFromType,
			}).
			ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}),

		diskAPI.DefineOperation("Install").
			Method(http.MethodPut).
			PathFormat(schema.IDAndSuffixPathFormat("install")).
			RequestEnvelope(&schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}).
			Argument(schema.ArgumentZone).
			Argument(schema.ArgumentID).
			Argument(&schema.Argument{
				Name:       "installParam",
				MapConvTag: "Disk",
				Type:       installParam,
			}).
			ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
				PayloadType: nakedDisk,
				PayloadName: "Disk",
			}),

		// read
		diskAPI.DefineOperationRead(nakedDisk, diskModel),

		// update
		diskAPI.DefineOperationUpdate(nakedDisk, updateParam, diskModel),

		// delete
		diskAPI.DefineOperationDelete(),

		// monitor
		diskAPI.DefineOperationMonitor(monitorParameter, monitors.diskModel()),
	}
	Resources.Def(diskAPI)
}
