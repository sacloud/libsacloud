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
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
			},
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.MappableArgument("param", webAccelCertUpdateParam, "Certificate"),
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
				dsl.ArgumentZone,
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
				dsl.ArgumentZone,
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
			fields.New("DomainType", meta.Static(types.EWebAccelDomainType(""))),
			fields.New("Domain", meta.TypeString),
			fields.New("Subdomain", meta.TypeString),
			fields.New("ASCIIDomain", meta.TypeString),
			fields.New("Origin", meta.TypeString),
			fields.New("HostHeader", meta.TypeString),
			fields.New("Status", meta.Static(types.EWebAccelStatus(""))),
			fields.New("HasCertificate", meta.TypeFlag),
			fields.New("HasOldCertificate", meta.TypeFlag),
			fields.New("GibSentInLastWeek", meta.TypeInt64),
			fields.New("CertValidNotBefore", meta.TypeInt64),
			fields.New("CertValidNotAfter", meta.TypeInt64),
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
		fields.New("SiteID", meta.TypeID),
		fields.New("CertificateChain", meta.TypeString),
		fields.New("Key", meta.TypeString),
		fields.New("CreatedAt", meta.TypeTime),
		fields.New("UpdatedAt", meta.TypeTime),
		fields.New("SerialNumber", meta.TypeString),
		fields.New("NotBefore", meta.TypeInt64),
		fields.New("NotAfter", meta.TypeInt64),
		{
			Name: "Issuer",
			Type: &dsl.Model{
				Name: "WebAccelCertIssuer",
				Fields: []*dsl.FieldDesc{
					fields.New("Country", meta.TypeString),
					fields.New("Organization", meta.TypeString),
					fields.New("OrganizationalUnit", meta.TypeString),
					fields.New("CommonName", meta.TypeString),
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
					fields.New("Country", meta.TypeString),
					fields.New("Organization", meta.TypeString),
					fields.New("OrganizationalUnit", meta.TypeString),
					fields.New("Locality", meta.TypeString),
					fields.New("Province", meta.TypeString),
					fields.New("StreetAddress", meta.TypeString),
					fields.New("PostalCode", meta.TypeString),
					fields.New("SerialNumber", meta.TypeString),
					fields.New("CommonName", meta.TypeString),
				},
			},
			Tags: &dsl.FieldTags{
				MapConv: ",recursive",
			},
		},
		fields.New("DNSNames", meta.TypeStringSlice),
		fields.New("SHA256Fingerprint", meta.TypeString),
	}

	webAccelCertUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(webAccelAPIName + "Cert"),
		NakedType: webAccelCertNakedType,
		Fields: []*dsl.FieldDesc{
			fields.New("CertificateChain", meta.TypeString),
			fields.New("Key", meta.TypeString),
		},
	}

	webAccelDeleteAllCacheParam = &dsl.Model{
		Name:      names.RequestParameterName(webAccelAPIName, "DeleteAllCache"),
		NakedType: webAccelSiteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.New("Domain", meta.TypeString),
		},
	}

	webAccelDeleteCacheParam = &dsl.Model{
		Name:      names.RequestParameterName(webAccelAPIName, "DeleteCache"),
		NakedType: webAccelSiteNakedType,
		Fields: []*dsl.FieldDesc{
			fields.New("URL", meta.TypeStringSlice),
		},
	}

	webAccelDeleteCacheResult = &dsl.Model{
		Name:      webAccelAPIName + "DeleteCacheResult",
		NakedType: meta.Static(naked.WebAccelDeleteCacheResult{}),
		IsArray:   true,
		Fields: []*dsl.FieldDesc{
			fields.New("URL", meta.TypeString),
			fields.New("Status", meta.TypeInt),
			fields.New("Result", meta.TypeString),
		},
	}
)