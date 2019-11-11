// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
