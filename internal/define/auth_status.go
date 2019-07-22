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
			fields.Def("AccountID", meta.TypeID, mapConvTag("Account.ID")),
			fields.Def("AccountName", meta.TypeString, mapConvTag("Account.Name")),
			fields.Def("AccountCode", meta.TypeString, mapConvTag("Account.Code")),
			fields.Def("AccountClass", meta.TypeString, mapConvTag("Account.Class")),
			fields.Def("MemberCode", meta.TypeString, mapConvTag("Member.Code")),
			fields.Def("MemberClass", meta.TypeString, mapConvTag("Member.Class")),
			fields.Def("AuthClass", meta.Static(types.EAuthClass(""))),
			fields.Def("AuthMethod", meta.Static(types.EAuthMethod(""))),
			fields.Def("IsAPIKey", meta.TypeFlag),
			fields.Def("ExternalPermission", meta.Static(types.ExternalPermission(""))),
			fields.Def("OperationPenalty", meta.Static(types.EOperationPenalty(""))),
			fields.Def("Permission", meta.Static(types.EPermission(""))),
		},
	}
)
