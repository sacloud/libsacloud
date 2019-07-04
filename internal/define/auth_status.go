package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/schema"
	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	authStatusAPIName  = "AuthStatus"
	authStatusPathName = "auth-status"
)

var authStatusAPI = &schema.Resource{
	Name:       authStatusAPIName,
	PathName:   authStatusPathName,
	PathSuffix: schema.CloudAPISuffix,
	IsGlobal:   true,
	Operations: schema.Operations{
		{
			ResourceName: authStatusAPIName,
			Name:         "Read",
			Method:       http.MethodGet,
			PathFormat:   schema.DefaultPathFormat,
			Arguments:    schema.Arguments{schema.ArgumentZone},
			ResponseEnvelope: schema.ResponseEnvelope(
				&schema.EnvelopePayloadDesc{
					Name: authStatusAPIName,
					Type: archiveNakedType,
				},
			),
			Results: schema.Results{
				{
					SourceField: authStatusAPIName,
					DestField:   authStatusView.Name,
					IsPlural:    false,
					Model:       authStatusView,
				},
			},
		},
	},
}

var (
	authStatusNakedType = meta.Static(naked.AuthStatus{})
	authStatusView      = &schema.Model{
		Name:      "AuthStatus",
		NakedType: authStatusNakedType,
		Fields: []*schema.FieldDesc{
			fields.AccountID(),
			fields.AccountName(),
			fields.AccountCode(),
			fields.AccountClass(),
			fields.MemberCode(),
			fields.MemberClass(),
			fields.New("AuthClass", meta.Static(types.EAuthClass(""))),
			fields.New("AuthMethod", meta.Static(types.EAuthMethod(""))),
			fields.New("IsAPIKey", meta.TypeFlag),
			fields.New("ExternalPermission", meta.Static(types.ExternalPermission(""))),
			fields.New("OperationPenalty", meta.Static(types.EOperationPenalty(""))),
			fields.New("Permission", meta.Static(types.EPermission(""))),
		},
	}
)
