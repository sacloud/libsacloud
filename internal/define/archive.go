package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

var archiveAPI = &schema.Resource{
	Name:       "Archive",
	PathName:   "archive",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(archiveNakedType, findParameter, archiveView),

			// create
			r.DefineOperationCreate(archiveNakedType, archiveCreateParam, archiveView),

			// CreateBlank
			r.DefineOperationCreate(archiveNakedType, archiveCreateBlankParam, archiveView).
				ResultFromEnvelope(models.ftpServer(), &schema.EnvelopePayloadDesc{
					PayloadName: models.ftpServer().Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}).Name("CreateBlank"),
			// TODO 他ゾーンからの転送コピー作成

			// read
			r.DefineOperationRead(archiveNakedType, archiveView),

			// update
			r.DefineOperationUpdate(archiveNakedType, archiveUpdateParam, archiveView),

			// delete
			r.DefineOperationDelete(),

			// openFTP
			r.DefineOperationOpenFTP(models.ftpServerOpenParameter(), models.ftpServer()),

			// closeFTP
			r.DefineOperationCloseFTP(),
		}
	},
}

var (
	archiveNakedType = meta.Static(naked.Archive{})

	archiveView = &schema.Model{
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

	archiveCreateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveCreateBlankParam = &schema.Model{
		Name: "ArchiveCreateBlankRequest",
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
