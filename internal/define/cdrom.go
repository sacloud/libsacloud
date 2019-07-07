package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

var cdromAPI = &schema.Resource{
	Name:       "CDROM",
	PathName:   "cdrom",
	PathSuffix: schema.CloudAPISuffix,
	OperationsDefineFunc: func(r *schema.Resource) []*schema.Operation {
		return []*schema.Operation{
			// find
			r.DefineOperationFind(cdromNakedType, findParameter, cdromView),

			// create
			{
				Resource:   r,
				Name:       "Create",
				PathFormat: schema.DefaultPathFormat,
				Method:     http.MethodPost,
				RequestEnvelope: schema.RequestEnvelope(&schema.EnvelopePayloadDesc{
					Name: r.FieldName(schema.PayloadForms.Singular),
					Type: cdromNakedType,
				}),
				ResponseEnvelope: schema.ResponseEnvelope(
					&schema.EnvelopePayloadDesc{
						Name: r.FieldName(schema.PayloadForms.Singular),
						Type: cdromNakedType,
					},
					&schema.EnvelopePayloadDesc{
						Name: models.ftpServer().Name,
						Type: meta.Static(naked.OpeningFTPServer{}),
					},
				),
				Arguments: schema.Arguments{
					schema.ArgumentZone,
					schema.MappableArgument("param", cdromCreateParam, r.FieldName(schema.PayloadForms.Singular)),
				},
				Results: schema.Results{
					{
						SourceField: r.FieldName(schema.PayloadForms.Singular),
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
			r.DefineOperationRead(cdromNakedType, cdromView),

			// update
			r.DefineOperationUpdate(cdromNakedType, cdromUpdateParam, cdromView),

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
