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

	Resources.DefineWith("Disk", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedDisk, findParameter, diskModel),

			// create
			r.DefineOperationCreate(nakedDisk, createParam, diskModel),

			// create distantly
			r.DefineOperation("CreateDistantly").
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
				Argument(&schema.MappableArgument{
					Name:        "createParam",
					Destination: "Disk",
					Model:       createParam,
				}).
				Argument(&schema.PassthroughSimpleArgument{
					Name:        "distantFrom",
					Destination: "DistantFrom",
					Type:        distantFromType,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

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
				Argument(&schema.MappableArgument{
					Name:        "createParam",
					Destination: "Disk",
					Model:       createParam,
				}).
				Argument(&schema.MappableArgument{
					Name:        "editParam",
					Destination: "Config",
					Model:       diskEditParam,
				}).
				Argument(&schema.PassthroughSimpleArgument{
					Name:        "bootAtAvailable",
					Type:        meta.TypeFlag,
					Destination: "BootAtAvailable",
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

			r.DefineOperation("CreateWithConfigDistantly").
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
				Argument(&schema.MappableArgument{
					Name:        "createParam",
					Destination: "Disk",
					Model:       createParam,
				}).
				Argument(&schema.MappableArgument{
					Name:        "editParam",
					Destination: "Config",
					Model:       diskEditParam,
				}).
				Argument(&schema.PassthroughSimpleArgument{
					Name:        "bootAtAvailable",
					Type:        meta.TypeFlag,
					Destination: "BootAtAvailable",
				}).
				Argument(&schema.PassthroughSimpleArgument{
					Name:        "distantFrom",
					Destination: "DistantFrom",
					Type:        distantFromType,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

			// to blank
			r.DefineSimpleOperation("ToBlank", http.MethodPut, "to/blank"),

			// resize partition
			r.DefineSimpleOperation("ResizePartition", http.MethodPut, "resize-partition"),

			// connect to server
			r.DefineSimpleOperation("ConnectToServer", http.MethodPut, "to/server/{{.serverID}}",
				&schema.SimpleArgument{
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
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: distantFromType,
					PayloadName: "DistantFrom",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.MappableArgument{
					Name:        "installParam",
					Destination: "Disk",
					Model:       installParam,
				}).
				Argument(&schema.PassthroughSimpleArgument{
					Name:        "distantFrom",
					Destination: "DistantFrom",
					Type:        distantFromType,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

			r.DefineOperation("Install").
				Method(http.MethodPut).
				PathFormat(schema.IDAndSuffixPathFormat("install")).
				RequestEnvelope(&schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}).
				Argument(schema.ArgumentZone).
				Argument(schema.ArgumentID).
				Argument(&schema.MappableArgument{
					Name:        "installParam",
					Destination: "Disk",
					Model:       installParam,
				}).
				ResultFromEnvelope(diskModel, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

			// read
			r.DefineOperationRead(nakedDisk, diskModel),

			// update
			r.DefineOperationUpdate(nakedDisk, updateParam, diskModel),

			// delete
			r.DefineOperationDelete(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.diskModel()),
		)
	})
}
