package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	authStatusAPIName  = "AuthStatus"
	authStatusPathName = "auth-status"
)

var authStatusAPI = &dsl.Resource{
	Name:       authStatusAPIName,
	PathName:   authStatusPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		{
			ResourceName: authStatusAPIName,
			Name:         "Read",
			Method:       http.MethodGet,
			PathFormat:   dsl.DefaultPathFormat,
			Arguments:    dsl.Arguments{dsl.ArgumentZone},
			ResponseEnvelope: dsl.ResponseEnvelope(
				&dsl.EnvelopePayloadDesc{
					Name: authStatusAPIName,
					Type: authStatusNakedType,
				},
			),
			Results: dsl.Results{
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
	authStatusView      = &dsl.Model{
		Name:      "AuthStatus",
		NakedType: authStatusNakedType,
		Fields: []*dsl.FieldDesc{
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
