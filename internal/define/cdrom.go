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
	cdromAPIName     = "CDROM"
	cdromAPIPathName = "cdrom"
)

var cdromAPI = &dsl.Resource{
	Name:       cdromAPIName,
	PathName:   cdromAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	Operations: dsl.Operations{
		// find
		ops.Find(cdromAPIName, cdromNakedType, findParameter, cdromView),

		// create
		{
			ResourceName: cdromAPIName,
			Name:         "Create",
			PathFormat:   dsl.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(cdromAPIName, dsl.PayloadForms.Singular),
				Type: cdromNakedType,
			}),
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(cdromAPIName, dsl.PayloadForms.Singular),
					Type: cdromNakedType,
				},
				&dsl.EnvelopePayloadDesc{
					Name: models.ftpServer().Name,
					Type: meta.Static(naked.OpeningFTPServer{}),
				},
			),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", cdromCreateParam, names.ResourceFieldName(cdromAPIName, dsl.PayloadForms.Singular)),
			},
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(cdromAPIName, dsl.PayloadForms.Singular),
					DestField:   cdromView.Name,
					IsPlural:    false,
					Model:       cdromView,
				},
				{
					SourceField: models.ftpServer().Name,
					DestField:   models.ftpServer().Name,
					IsPlural:    false,
					Model:       models.ftpServer(),
				},
			},
		},

		// read
		ops.Read(cdromAPIName, cdromNakedType, cdromView),

		// update
		ops.Update(cdromAPIName, cdromNakedType, cdromUpdateParam, cdromView),

		// patch
		ops.Patch(cdromAPIName, cdromNakedType, patchModel(cdromUpdateParam), cdromView),

		// delete
		ops.Delete(cdromAPIName),

		// openFTP
		ops.OpenFTP(cdromAPIName, models.ftpServerOpenParameter(), models.ftpServer()),

		// closeFTP
		ops.CloseFTP(cdromAPIName),
	},
}
var (
	cdromNakedType = meta.Static(naked.CDROM{})

	cdromView = &dsl.Model{
		Name:      cdromAPIName,
		NakedType: cdromNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.DisplayOrder(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.Storage(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	cdromCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(cdromAPIName),
		NakedType: cdromNakedType,
		Fields: []*dsl.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	cdromUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(cdromAPIName),
		NakedType: cdromNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
