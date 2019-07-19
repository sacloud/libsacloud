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
	sshKeyAPIName     = "SSHKey"
	sshKeyAPIPathName = "sshkey"
)

var sshKeyAPI = &dsl.Resource{
	Name:       sshKeyAPIName,
	PathName:   sshKeyAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.Find(sshKeyAPIName, sshKeyNakedType, findParameter, sshKeyView),

		// create
		ops.Create(sshKeyAPIName, sshKeyNakedType, sshKeyCreateParam, sshKeyView),

		// generate
		{
			ResourceName: sshKeyAPIName,
			Name:         "Generate",
			PathFormat:   dsl.DefaultPathFormat + "/generate",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: sshKeyNakedType,
				Name: sshKeyAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", sshKeyGenerateParam, sshKeyAPIName),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: sshKeyNakedType,
				Name: sshKeyAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: sshKeyAPIName,
					DestField:   sshKeyGeneratedView.Name,
					IsPlural:    false,
					Model:       sshKeyGeneratedView,
				},
			},
		},

		// read
		ops.Read(sshKeyAPIName, sshKeyNakedType, sshKeyView),

		// update
		ops.Update(sshKeyAPIName, sshKeyNakedType, sshKeyUpdateParam, sshKeyView),

		// delete
		ops.Delete(sshKeyAPIName),
	},
}

var (
	sshKeyNakedType = meta.Static(naked.SSHKey{})

	sshKeyFields = []*dsl.FieldDesc{
		fields.ID(),
		fields.Name(),
		fields.Description(),
		fields.CreatedAt(),
		fields.PublicKey(),
		fields.Fingerprint(),
	}

	sshKeyView = &dsl.Model{
		Name:      sshKeyAPIName,
		NakedType: sshKeyNakedType,
		Fields:    sshKeyFields,
	}

	sshKeyGeneratedView = &dsl.Model{
		Name:      sshKeyAPIName + "Generated",
		NakedType: sshKeyNakedType,
		Fields:    append(sshKeyFields, fields.PrivateKey()),
	}

	sshKeyCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(sshKeyAPIName),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PublicKey(),
		},
	}

	sshKeyGenerateParam = &dsl.Model{
		Name:      names.RequestParameterName(sshKeyAPIName, "Generate"),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.PassPhrase(),
			{
				Name: "GenerateFormat",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: ",default=openssh",
				},
			},
		},
	}

	sshKeyUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(sshKeyAPIName),
		NakedType: sshKeyNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.Description(),
		},
	}
)
