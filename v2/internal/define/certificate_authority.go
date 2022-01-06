// Copyright 2016-2022 The Libsacloud Authors
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

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	certificateAuthorityAPIName     = "CertificateAuthority"
	certificateAuthorityAPIPathName = "commonserviceitem"
)

var certificateAuthorityAPI = &dsl.Resource{
	Name:       certificateAuthorityAPIName,
	PathName:   certificateAuthorityAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(certificateAuthorityAPIName, certificateAuthorityNakedType, findParameter, certificateAuthorityView),

		// create
		ops.CreateCommonServiceItem(certificateAuthorityAPIName, certificateAuthorityNakedType, certificateAuthorityCreateParam, certificateAuthorityView),

		// read
		ops.ReadCommonServiceItem(certificateAuthorityAPIName, certificateAuthorityNakedType, certificateAuthorityView),

		// update
		ops.UpdateCommonServiceItem(certificateAuthorityAPIName, certificateAuthorityNakedType, certificateAuthorityUpdateParam, certificateAuthorityView),

		// delete
		ops.Delete(certificateAuthorityAPIName),

		// ca detail
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "Detail",
			Method:       http.MethodGet,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "CertificateAuthority",
				Type: meta.Static(naked.CertificateAuthorityDetail{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					IsPlural:    false,
					Model:       certificateAuthorityDetailModel,
				},
			},
		},

		/*
		 * Client Certificates
		 */

		// create/add
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "AddClient",
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients"),
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.CertificateAuthorityAddClientParameter{}),
				Name: "CertificateAuthority",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", certificateAuthorityAddClientParam, "CertificateAuthority.Status"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.CertificateAuthorityAddClientOrServerResult{}),
				Name: "CertificateAuthority",
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					IsPlural:    false,
					Model:       certificateAuthorityAddClientOrServerResult,
				},
			},
		},
		// list
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ListClients",
			Method:       http.MethodGet,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Name: "CertificateAuthority",
				Type: meta.Static(naked.CertificateAuthorityClientDetail{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					IsPlural:    true,
					Model:       certificateAuthorityClientModel,
				},
			},
			UseWrappedResult: true,
		},
		// read
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ReadClient",
			Method:       http.MethodGet,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients/{{.clientID}}"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "clientID",
					Type: meta.TypeString,
				},
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "CertificateAuthority",
				Type: meta.Static(naked.CertificateAuthorityClientDetail{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					Model:       certificateAuthorityClientModel,
				},
			},
		},

		// revoke
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "RevokeClient",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients/{{.clientID}}/revoke"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "clientID",
					Type: meta.TypeString,
				},
			},
		},
		// hold
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "HoldClient",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients/{{.clientID}}/hold"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "clientID",
					Type: meta.TypeString,
				},
			},
		},

		// resume
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ResumeClient",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients/{{.clientID}}/resume"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "clientID",
					Type: meta.TypeString,
				},
			},
		},

		// deny
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "DenyClient",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/clients/{{.clientID}}/deny"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "clientID",
					Type: meta.TypeString,
				},
			},
		},

		/*
		 * Sercver Certificates
		 */
		// create/add
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "AddServer",
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers"),
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.CertificateAuthorityAddServerParameter{}),
				Name: "CertificateAuthority",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", certificateAuthorityAddServerParam, "CertificateAuthority.Status"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static(naked.CertificateAuthorityAddClientOrServerResult{}),
				Name: "CertificateAuthority",
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					IsPlural:    false,
					Model:       certificateAuthorityAddClientOrServerResult,
				},
			},
		},
		// list
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ListServers",
			Method:       http.MethodGet,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Name: "CertificateAuthority",
				Type: meta.Static(naked.CertificateAuthorityServerDetail{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					IsPlural:    true,
					Model:       certificateAuthorityServerModel,
				},
			},
			UseWrappedResult: true,
		},
		// read server certs
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ReadServer",
			Method:       http.MethodGet,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers/{{.serverID}}"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "serverID",
					Type: meta.TypeString,
				},
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Name: "CertificateAuthority",
				Type: meta.Static(naked.CertificateAuthorityServerDetail{}),
			}),
			Results: dsl.Results{
				{
					SourceField: "CertificateAuthority",
					DestField:   "CertificateAuthority",
					Model:       certificateAuthorityServerModel,
				},
			},
		},
		// revoke
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "RevokeServer",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers/{{.serverID}}/revoke"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "serverID",
					Type: meta.TypeString,
				},
			},
		},
		// hold
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "HoldServer",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers/{{.serverID}}/hold"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "serverID",
					Type: meta.TypeString,
				},
			},
		},
		// resume
		{
			ResourceName: certificateAuthorityAPIName,
			Name:         "ResumeServer",
			Method:       http.MethodPut,
			PathFormat:   dsl.IDAndSuffixPathFormat("certificateauthority/servers/{{.serverID}}/resume"),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				{
					Name: "serverID",
					Type: meta.TypeString,
				},
			},
		},
	},
}

var (
	certificateAuthorityNakedType = meta.Static(naked.CertificateAuthority{})

	certificateAuthorityView = &dsl.Model{
		Name:      certificateAuthorityAPIName,
		NakedType: certificateAuthorityNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),

			// status
			fields.CAStatusCountry(),
			fields.CAStatusOrganization(),
			fields.CAStatusOrganizationUnit(),
			fields.CAStatusCommonName(),
			fields.CAStatusNotAfter(),
			fields.CAStatusSubject(),
		},
	}

	certificateAuthorityCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(certificateAuthorityAPIName),
		NakedType: certificateAuthorityNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"certificateauthority"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),

			// status
			fields.CAStatusCountry(),
			fields.CAStatusOrganization(),
			fields.CAStatusOrganizationUnit(),
			fields.CAStatusCommonName(),
			fields.CAStatusNotAfter(),
		},
	}

	certificateAuthorityUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(certificateAuthorityAPIName),
		NakedType: certificateAuthorityNakedType,
		Fields: []*dsl.FieldDesc{
			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	certificateAuthorityCertificateData = &dsl.Model{
		Name:      "CertificateData",
		NakedType: meta.Static(naked.CertificateData{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("CertificatePEM", meta.TypeString),
			fields.Def("Subject", meta.TypeString),
			fields.Def("SerialNumber", meta.TypeString),
			fields.Def("NotBefore", meta.TypeTime),
			fields.Def("NotAfter", meta.TypeTime),
		},
	}
	certificateAuthorityDetailModel = &dsl.Model{
		Name:      "CertificateAuthorityDetail",
		NakedType: meta.Static(naked.CertificateAuthorityDetail{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Subject", meta.TypeString),
			fields.Def("CertificateData", certificateAuthorityCertificateData, &dsl.FieldTags{MapConv: ",recursive"}),
		},
	}

	certificateAuthorityAddClientParam = &dsl.Model{
		Name:      "CertificateAuthorityAddClientParam",
		NakedType: meta.Static(naked.CertificateAuthorityAddClientParameterBody{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Country", meta.TypeString),
			fields.Def("Organization", meta.TypeString),
			fields.Def("OrganizationUnit", meta.TypeStringSlice),
			fields.Def("CommonName", meta.TypeString),
			fields.Def("NotAfter", meta.TypeTime),
			fields.Def("IssuanceMethod", meta.Static(types.ECertificateAuthorityIssuanceMethod(""))),

			fields.Def("EMail", meta.TypeString),
			fields.Def("CertificateSigningRequest", meta.TypeString),
			fields.Def("PublicKey", meta.TypeString),
		},
	}

	certificateAuthorityClientModel = &dsl.Model{
		Name:      "CertificateAuthorityClient",
		NakedType: meta.Static(naked.CertificateAuthorityClientDetail{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("ID", meta.TypeString),
			fields.Def("Subject", meta.TypeString),
			fields.Def("EMail", meta.TypeString),
			fields.Def("IssuanceMethod", meta.Static(types.ECertificateAuthorityIssuanceMethod(""))),
			fields.Def("IssueState", meta.TypeString),
			fields.Def("URL", meta.TypeString),
			fields.Def("CertificateData", certificateAuthorityCertificateData, &dsl.FieldTags{MapConv: ",recursive"}),
		},
	}

	certificateAuthorityAddServerParam = &dsl.Model{
		Name:      "CertificateAuthorityAddServerParam",
		NakedType: meta.Static(naked.CertificateAuthorityAddClientParameterBody{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("Country", meta.TypeString),
			fields.Def("Organization", meta.TypeString),
			fields.Def("OrganizationUnit", meta.TypeStringSlice),
			fields.Def("CommonName", meta.TypeString),
			fields.Def("NotAfter", meta.TypeTime),
			fields.Def("SANs", meta.TypeStringSlice),

			fields.Def("CertificateSigningRequest", meta.TypeString),
			fields.Def("PublicKey", meta.TypeString),
		},
	}
	certificateAuthorityServerModel = &dsl.Model{
		Name:      "CertificateAuthorityServer",
		NakedType: meta.Static(naked.CertificateAuthorityServerDetail{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("ID", meta.TypeString),
			fields.Def("Subject", meta.TypeString),
			fields.Def("SANs", meta.TypeStringSlice),
			fields.Def("EMail", meta.TypeString),
			fields.Def("IssueState", meta.TypeString),
			fields.Def("CertificateData", certificateAuthorityCertificateData, &dsl.FieldTags{MapConv: ",recursive"}),
		},
	}

	certificateAuthorityAddClientOrServerResult = &dsl.Model{
		Name:      "CertificateAuthorityAddClientOrServerResult",
		NakedType: meta.Static(naked.CertificateAuthorityAddClientOrServerResult{}),
		Fields: []*dsl.FieldDesc{
			fields.Def("ID", meta.TypeString),
		},
	}
)
