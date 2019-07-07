package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	archiveAPIName     = "Archive"
	archiveAPIPathName = "archive"
)

var archiveAPI = &schema.Resource{
	Name:       archiveAPIName,
	PathName:   archiveAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(archiveAPIName, archiveNakedType, findParameter, archiveView),

		// create
		ops.Create(archiveAPIName, archiveNakedType, archiveCreateParam, archiveView),

		// CreateBlank
		{
			ResourceName: archiveAPIName,
			Name:         "CreateBlank",
			PathFormat:   schema.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: schema.RequestEnvelope(&schema.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(archiveAPIName, schema.PayloadForms.Singular),
				Type: archiveNakedType,
			}),
			ResponseEnvelope: schema.ResponseEnvelope(
				&schema.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(archiveAPIName, schema.PayloadForms.Singular),
					Type: archiveNakedType,
				},
				&schema.EnvelopePayloadDesc{
					Name: models.ftpServer().Name,
					Type: meta.Static(naked.OpeningFTPServer{}),
				},
			),
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.MappableArgument("param", archiveCreateBlankParam, names.ResourceFieldName(archiveAPIName, schema.PayloadForms.Singular)),
			},
			Results: schema.Results{
				{
					SourceField: names.ResourceFieldName(archiveAPIName, schema.PayloadForms.Singular),
					DestField:   archiveView.Name,
					IsPlural:    false,
					Model:       archiveView,
				},
				{
					SourceField: models.ftpServer().Name,
					DestField:   models.ftpServer().Name,
					IsPlural:    false,
					Model:       models.ftpServer(),
				},
			},
		},

		// TODO 他ゾーンからの転送コピー作成

		// read
		ops.Read(archiveAPIName, archiveNakedType, archiveView),

		// update
		ops.Update(archiveAPIName, archiveNakedType, archiveUpdateParam, archiveView),

		// delete
		ops.Delete(archiveAPIName),

		// openFTP
		ops.OpenFTP(archiveAPIName, models.ftpServerOpenParameter(), models.ftpServer()),

		// closeFTP
		ops.CloseFTP(archiveAPIName),
	},
}

var (
	archiveNakedType = meta.Static(naked.Archive{})

	archiveView = &schema.Model{
		Name:      archiveAPIName,
		NakedType: archiveNakedType,
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
		Name:      names.CreateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
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
		Name:      "ArchiveCreateBlankRequest",
		NakedType: archiveNakedType,
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveUpdateParam = &schema.Model{
		Name:      names.UpdateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
