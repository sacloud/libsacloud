package define

import (
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
			func() *schema.Operation {
				o := r.DefineOperationCreate(cdromNakedType, cdromCreateParam, cdromView)
				o.ResponseEnvelope.Payloads = append(o.ResponseEnvelope.Payloads, &schema.EnvelopePayloadDesc{
					PayloadName: models.ftpServer().Name,
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				})
				o.Results = append(o.Results, &schema.Result{
					SourceField: models.ftpServer().Name,
					DestField:   models.ftpServer().Name,
					IsPlural:    false,
					Model:       models.ftpServer(),
				})
				return o
			}(),

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
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
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
