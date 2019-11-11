package define

import (
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

const (
	webAccelAPIName = "WebAccel"
)

var webaccelAPI = &dsl.Resource{
	Name:       webAccelAPIName,
	PathName:   "", // その他のリソースとURLパターンが異なるため各オペレーションでPathFormatを指定する
	PathSuffix: dsl.WebAccelAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// list
		{
			ResourceName:     webAccelAPIName,
			Name:             "List",
			PathFormat:       "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site",
			Method:           http.MethodGet,
			UseWrappedResult: true,
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: webAccelSiteNakedType,
				Name: "Sites",
			}),
			Results: dsl.Results{
				{
					SourceField: "Sites",
					DestField:   names.ResourceFieldName(webAccelSiteView.Name, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       webAccelSiteView,
				},
			},
		},
		// read
		{
			ResourceName: webAccelAPIName,
			Name:         "Read",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site/{{.id}}",
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelSiteNakedType,
				Name: "Site",
			}),
			Results: dsl.Results{
				{
					SourceField: "Site",
					DestField:   names.ResourceFieldName(webAccelSiteView.Name, dsl.PayloadForms.Singular),
					IsPlural:    false,
					Model:       webAccelSiteView,
				},
			},
		},
		// read certificate
		{
			ResourceName: webAccelAPIName,
			Name:         "ReadCertificate",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site/{{.id}}/certificate",
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelCertsNakedType,
				Name: "Certificate",
			}),
			Results: dsl.Results{
				{
					SourceField: "Certificate",
					DestField:   "Certificate",
					IsPlural:    false,
					Model:       webAccelCertsView,
				},
			},
		},
		// create certificate
		{
			ResourceName: webAccelAPIName,
			Name:         "CreateCertificate",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site/{{.id}}/certificate",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelCertNakedType,
				Name: "Certificate",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", webAccelCertParam, "Certificate"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelCertsNakedType,
				Name: "Certificate",
			}),
			Results: dsl.Results{
				{
					SourceField: "Certificate",
					DestField:   "Certificate",
					IsPlural:    false,
					Model:       webAccelCertsView,
				},
			},
		},
		// update certificate
		{
			ResourceName: webAccelAPIName,
			Name:         "UpdateCertificate",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site/{{.id}}/certificate",
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelCertNakedType,
				Name: "Certificate",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
				dsl.MappableArgument("param", webAccelCertParam, "Certificate"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelCertsNakedType,
				Name: "Certificate",
			}),
			Results: dsl.Results{
				{
					SourceField: "Certificate",
					DestField:   "Certificate",
					IsPlural:    false,
					Model:       webAccelCertsView,
				},
			},
		},
		// delete certificate
		{
			ResourceName: webAccelAPIName,
			Name:         "DeleteCertificate",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/site/{{.id}}/certificate",
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
		},
		// delete all cache
		{
			ResourceName: webAccelAPIName,
			Name:         "DeleteAllCache",
			PathFormat:   "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/deleteallcache",
			Method:       http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: webAccelSiteNakedType,
				Name: "Site",
			}),
			Arguments: dsl.Arguments{
				dsl.MappableArgument("param", webAccelDeleteAllCacheParam, "Site"),
			},
		},

		// delete cache
		{
			ResourceName:    webAccelAPIName,
			Name:            "DeleteCache",
			PathFormat:      "{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/deletecache",
			Method:          http.MethodPost,
			RequestEnvelope: dsl.RequestEnvelopeFromModel(webAccelDeleteCacheParam),
			Arguments: dsl.Arguments{
				dsl.PassthroughModelArgument("param", webAccelDeleteCacheParam),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: meta.Static([]*naked.WebAccelDeleteCacheResult{}),
				Name: "Results",
			}),
			Results: dsl.Results{
				{
					SourceField: "Results",
					DestField:   "Results",
					IsPlural:    true,
					Model:       webAccelDeleteCacheResult,
				},
			},
		},
	},
}

var (
	webAccelSiteNakedType  = meta.Static(naked.WebAccelSite{})
	webAccelCertNakedType  = meta.Static(naked.WebAccelCert{})
	webAccelCertsNakedType = meta.Static(naked.WebAccelCerts{})

	webAccelSiteView = &dsl.Model{
		Name:      webAccelAPIName,
		NakedType: webAccelSiteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Def("DomainType", meta.Static(types.EWebAccelDomainType(""))),
			fields.Def("Domain", meta.TypeString),
			fields.Def("Subdomain", meta.TypeString),
			fields.Def("ASCIIDomain", meta.TypeString),
			fields.Def("Origin", meta.TypeString),
			fields.Def("HostHeader", meta.TypeString),
			fields.Def("Status", meta.Static(types.EWebAccelStatus(""))),
			fields.Def("HasCertificate", meta.TypeFlag),
			fields.Def("HasOldCertificate", meta.TypeFlag),
			fields.Def("GibSentInLastWeek", meta.TypeInt64),
			fields.Def("CertValidNotBefore", meta.TypeInt64),
			fields.Def("CertValidNotAfter", meta.TypeInt64),
			fields.CreatedAt(),
		},
	}

	webAccelCertsView = &dsl.Model{
		Name:      webAccelAPIName + "Certs",
		NakedType: webAccelCertsNakedType,
		Fields: []*dsl.FieldDesc{
			{
				Name: "Current",
				Type: &dsl.Model{
					Name:      webAccelAPIName + "CurrentCert",
					NakedType: meta.Static(naked.WebAccelCert{}),
					IsArray:   false,
					Fields:    webAccelCertViewFields,
				},
			},
			{
				Name: "Old",
				Type: &dsl.Model{
					Name:      webAccelAPIName + "OldCerts",
					NakedType: meta.Static(naked.WebAccelCert{}),
					IsArray:   true,
					Fields:    webAccelCertViewFields,
				},
			},
		},
	}

	webAccelCertViewFields = []*dsl.FieldDesc{
		fields.ID(),
		fields.Def("SiteID", meta.TypeID),
		fields.Def("CertificateChain", meta.TypeString),
		fields.Def("Key", meta.TypeString),
		fields.Def("CreatedAt", meta.TypeTime),
		fields.Def("UpdatedAt", meta.TypeTime),
		fields.Def("SerialNumber", meta.TypeString),
		fields.Def("NotBefore", meta.TypeInt64),
		fields.Def("NotAfter", meta.TypeInt64),
		{
			Name: "Issuer",
			Type: &dsl.Model{
				Name: "WebAccelCertIssuer",
				Fields: []*dsl.FieldDesc{
					fields.Def("Country", meta.TypeString),
					fields.Def("Organization", meta.TypeString),
					fields.Def("OrganizationalUnit", meta.TypeString),
					fields.Def("CommonName", meta.TypeString),
				},
			},
			Tags: &dsl.FieldTags{
				MapConv: ",recursive",
			},
		},
		{
			Name: "Subject",
			Type: &dsl.Model{
				Name: "WebAccelCertSubject",
				Fields: []*dsl.FieldDesc{
					fields.Def("Country", meta.TypeString),
					fields.Def("Organization", meta.TypeString),
					fields.Def("OrganizationalUnit", meta.TypeString),
					fields.Def("Locality", meta.TypeString),
					fields.Def("Province", meta.TypeString),
					fields.Def("StreetAddress", meta.TypeString),
					fields.Def("PostalCode", meta.TypeString),
					fields.Def("SerialNumber", meta.TypeString),
					fields.Def("CommonName", meta.TypeString),
				},
			},
			Tags: &dsl.FieldTags{
				MapConv: ",recursive",
			},
		},
		fields.Def("DNSNames", meta.TypeStringSlice),
		fields.Def("SHA256Fingerprint", meta.TypeString),
	}

	webAccelCertParam = &dsl.Model{
		Name:      webAccelAPIName + "CertRequest",
		NakedType: webAccelCertNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("CertificateChain", meta.TypeString),
			fields.Def("Key", meta.TypeString),
		},
	}

	webAccelDeleteAllCacheParam = &dsl.Model{
		Name:      names.RequestParameterName(webAccelAPIName, "DeleteAllCache"),
		NakedType: webAccelSiteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("Domain", meta.TypeString),
		},
	}

	webAccelDeleteCacheParam = &dsl.Model{
		Name:      names.RequestParameterName(webAccelAPIName, "DeleteCache"),
		NakedType: webAccelSiteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("URL", meta.TypeStringSlice),
		},
	}

	webAccelDeleteCacheResult = &dsl.Model{
		Name:      webAccelAPIName + "DeleteCacheResult",
		NakedType: meta.Static(naked.WebAccelDeleteCacheResult{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.Def("URL", meta.TypeString),
			fields.Def("Status", meta.TypeInt),
			fields.Def("Result", meta.TypeString),
		},
	}
)
