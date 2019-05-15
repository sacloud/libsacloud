package define

import (
	"net/http"

	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedDisk := meta.Static(naked.Disk{})
	nakedDiskEdit := meta.Static(naked.DiskEdit{})

	disk := &schema.Model{
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
			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}
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
		},
	}

	diskEditParam := models.diskEdit()

	Resources.DefineWith("Disk", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedDisk, findParameter, disk),

			// create
			r.DefineOperationCreate(nakedDisk, createParam, disk),

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
				ResultFromEnvelope(disk, &schema.EnvelopePayloadDesc{
					PayloadType: nakedDisk,
					PayloadName: "Disk",
				}),

			// read
			r.DefineOperationRead(nakedDisk, disk),

			// update
			r.DefineOperationUpdate(nakedDisk, updateParam, disk),

			// delete
			r.DefineOperationDelete(),

			// monitor
			r.DefineOperationMonitor(monitorParameter, monitors.diskModel()),
		)
	})
}
