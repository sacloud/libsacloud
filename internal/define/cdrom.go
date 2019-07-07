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
	cdromAPIName     = "CDROM"
	cdromAPIPathName = "cdrom"
)

var cdromAPI = &schema.Resource{
	Name:       cdromAPIName,
	PathName:   cdromAPIPathName,
	PathSuffix: schema.CloudAPISuffix,
	Operations: schema.Operations{
		// find
		ops.Find(cdromAPIName, cdromNakedType, findParameter, cdromView),

		// create
		{
			ResourceName: cdromAPIName,
			Name:         "Create",
			PathFormat:   schema.DefaultPathFormat,
			Method:       http.MethodPost,
			RequestEnvelope: schema.RequestEnvelope(&schema.EnvelopePayloadDesc{
				Name: names.ResourceFieldName(cdromAPIName, schema.PayloadForms.Singular),
				Type: cdromNakedType,
			}),
			ResponseEnvelope: schema.ResponseEnvelope(
				&schema.EnvelopePayloadDesc{
					Name: names.ResourceFieldName(cdromAPIName, schema.PayloadForms.Singular),
					Type: cdromNakedType,
				},
				&schema.EnvelopePayloadDesc{
					Name: models.ftpServer().Name,
					Type: meta.Static(naked.OpeningFTPServer{}),
				},
			),
			Arguments: schema.Arguments{
				schema.ArgumentZone,
				schema.MappableArgument("param", cdromCreateParam, names.ResourceFieldName(cdromAPIName, schema.PayloadForms.Singular)),
			},
			Results: schema.Results{
				{
					SourceField: names.ResourceFieldName(cdromAPIName, schema.PayloadForms.Singular),
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

	cdromView = &schema.Model{
		Fields: []*schema.FieldDesc{
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

	cdromCreateParam = &schema.Model{
		Name: "CDROMCreateRequest",
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
		NakedType: cdromNakedType,
	}

	cdromUpdateParam = &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}
)
