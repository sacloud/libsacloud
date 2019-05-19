package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.Archive{})

	archive := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.DisplayOrder(),
			fields.Availability(),
			fields.Scope(),
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
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.OriginalArchiveID(),
			fields.SourceInfo(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	createBlankParam := &schema.Model{
		Name: "ArchiveCreateBlankRequest",
		Fields: []*schema.FieldDesc{
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

	Resources.DefineWith("Archive", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, archive),

			// create
			r.DefineOperationCreate(nakedType, createParam, archive),

			// CreateBlank
			r.DefineOperationCreate(nakedType, createBlankParam, archive).
				ResultFromEnvelope(models.ftpServer(), &schema.EnvelopePayloadDesc{
					PayloadName: models.ftpServer().Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}).Name("CreateBlank"),
			// TODO 他ゾーンからの転送コピー作成

			// read
			r.DefineOperationRead(nakedType, archive),

			// update
			r.DefineOperationUpdate(nakedType, updateParam, archive),

			// delete
			r.DefineOperationDelete(),

			// openFTP
			r.DefineOperationOpenFTP(models.ftpServerOpenParameter(), models.ftpServer()),

			// closeFTP
			r.DefineOperationCloseFTP(),
		)
	})
}
