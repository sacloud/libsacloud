package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	archiveAPIName     = "Archive"
	archiveAPIPathName = "archive"
)

var archiveAPI = &dsl.Resource{
	Name:       archiveAPIName,
	PathName:   archiveAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(archiveAPIName, archiveNakedType, findParameter, archiveView, []dsl.SearchKeyDesc{fields.Availability()}),

		// create
		ops.Create(archiveAPIName, archiveNakedType, archiveCreateParam, archiveView),

		// CreateBlank
		{
			ResourceName: archiveAPIName,
			Name:         "CreateBlank",
			PathFormat:   dsl.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
				Type: archiveNakedType,
			}),
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
					Type: archiveNakedType,
				},
				&dsl.EnvelopePayloadDesc{
					Name: models.ftpServer().Name,
					Type: meta.Static(naked.OpeningFTPServer{}),
				},
			),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", archiveCreateBlankParam, names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular)),
			},
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(archiveAPIName, dsl.PayloadForms.Singular),
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

	archiveView = &dsl.Model{
		Name:      archiveAPIName,
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
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

	archiveCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SourceDiskID(),
			fields.SourceArchiveID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveCreateBlankParam = &dsl.Model{
		Name:      "ArchiveCreateBlankRequest",
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	archiveUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(archiveAPIName),
		NakedType: archiveNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
