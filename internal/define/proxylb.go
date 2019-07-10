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
	proxyLBAPIName     = "ProxyLB"
	proxyLBAPIPathName = "commonserviceitem"
)

var proxyLBAPI = &dsl.Resource{
	Name:       proxyLBAPIName,
	PathName:   proxyLBAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.FindCommonServiceItem(proxyLBAPIName, proxyLBNakedType, findParameter, proxyLBView),

		// create
		ops.CreateCommonServiceItem(proxyLBAPIName, proxyLBNakedType, proxyLBCreateParam, proxyLBView),

		// read
		ops.ReadCommonServiceItem(proxyLBAPIName, proxyLBNakedType, proxyLBView),

		// update
		ops.UpdateCommonServiceItem(proxyLBAPIName, proxyLBNakedType, proxyLBUpdateParam, proxyLBView),

		// delete
		ops.Delete(proxyLBAPIName),

		// change plan
		{
			ResourceName: proxyLBAPIName,
			Name:         "ChangePlan",
			PathFormat:   dsl.DefaultPathFormatWithID,
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: proxyLBNakedType,
				Name: "CommonServiceItem",
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.MappableArgument("param", proxyLBChangePlanParam, "CommonServiceItem"),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: proxyLBNakedType,
				Name: "CommonServiceItem",
			}),
			Results: dsl.Results{
				{
					SourceField: "CommonServiceItem",
					DestField:   proxyLBView.Name,
					IsPlural:    false,
					Model:       proxyLBView,
				},
			},
		},

		// get certificates
		{
			ResourceName: proxyLBAPIName,
			Name:         "GetCertificates",
			PathFormat:   dsl.IDAndSuffixPathFormat("proxylb/sslcertificate"),
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: proxyLBCertificatesNakedType,
				Name: proxyLBAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: proxyLBAPIName,
					DestField:   proxyLBCertificateView.Name,
					IsPlural:    false,
					Model:       proxyLBCertificateView,
				},
			},
		},

		// set certificates
		{
			ResourceName: proxyLBAPIName,
			Name:         "SetCertificates",
			PathFormat:   dsl.IDAndSuffixPathFormat("proxylb/sslcertificate"),
			Method:       http.MethodPut,
			RequestEnvelope: dsl.RequestEnvelope(&dsl.EnvelopePayloadDesc{
				Type: proxyLBCertificatesNakedType,
				Name: proxyLBAPIName,
			}),
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
				dsl.MappableArgument("param", proxyLBCertificateSetParam, proxyLBAPIName),
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: proxyLBCertificatesNakedType,
				Name: proxyLBAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: proxyLBAPIName,
					DestField:   proxyLBCertificateView.Name,
					IsPlural:    false,
					Model:       proxyLBCertificateView,
				},
			},
		},

		// delete certificates
		{
			ResourceName: proxyLBAPIName,
			Name:         "DeleteCertificates",
			PathFormat:   dsl.IDAndSuffixPathFormat("proxylb/sslcertificate"),
			Method:       http.MethodDelete,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				dsl.ArgumentID,
			},
		},

		// renew Let's Encrypt certificates
		ops.WithIDAction(proxyLBAPIName, "RenewLetsEncryptCert", http.MethodPut, "proxylb/letsencryptrenew"),

		// Health
		ops.HealthStatus(proxyLBAPIName, meta.Static(naked.ProxyLBHealth{}), proxyLBHealth),
	},
}

var (
	proxyLBNakedType             = meta.Static(naked.ProxyLB{})
	proxyLBCertificatesNakedType = meta.Static(naked.ProxyLBCertificates{})

	proxyLBView = &dsl.Model{
		Name:      proxyLBAPIName,
		NakedType: proxyLBNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.Availability(),
			fields.IconID(),
			fields.CreatedAt(),
			fields.ModifiedAt(),

			fields.ProxyLBProviderClass(),
			fields.ProxyLBPlan(),

			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),
			fields.SettingsHash(),

			// status
			fields.ProxyLBUseVIPFailover(),
			fields.ProxyLBRegion(),
			fields.ProxyLBProxyNetworks(),
			fields.ProxyLBFQDN(),
		},
	}

	proxyLBCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(proxyLBAPIName),
		NakedType: proxyLBNakedType,
		Fields: []*dsl.FieldDesc{
			// required
			fields.ProxyLBProviderClass(),
			fields.ProxyLBPlan(),

			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),

			// status
			fields.ProxyLBUseVIPFailover(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	proxyLBUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(proxyLBAPIName),
		NakedType: proxyLBNakedType,
		Fields: []*dsl.FieldDesc{

			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	proxyLBChangePlanParam = &dsl.Model{
		Name:      proxyLBAPIName + "ChangePlanRequest",
		NakedType: proxyLBNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ProxyLBPlan(),
		},
	}

	proxyLBCertificateView = &dsl.Model{
		Name:      proxyLBAPIName + "Certificates",
		NakedType: meta.Static(naked.ProxyLBCertificates{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "ServerCertificate",
				Type: meta.TypeString,
			},
			{
				Name: "IntermediateCertificate",
				Type: meta.TypeString,
			},
			{
				Name: "PrivateKey",
				Type: meta.TypeString,
			},
			{
				Name: "CertificateEndDate",
				Type: meta.TypeTime,
			},
			{
				Name: "CertificateCommonName",
				Type: meta.TypeString,
			},
			{
				Name: "AdditionalCerts",
				Type: &dsl.Model{
					Name:    proxyLBAPIName + "AdditionalCert",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						{
							Name: "ServerCertificate",
							Type: meta.TypeString,
						},
						{
							Name: "IntermediateCertificate",
							Type: meta.TypeString,
						},
						{
							Name: "PrivateKey",
							Type: meta.TypeString,
						},
						{
							Name: "CertificateEndDate",
							Type: meta.TypeString,
						},
						{
							Name: "CertificateCommonName",
							Type: meta.TypeString,
						},
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]AdditionalCerts, recursive",
				},
			},
		},
	}

	proxyLBCertificateSetParam = &dsl.Model{
		Name:      proxyLBAPIName + "SetCertificatesRequest",
		NakedType: meta.Static(naked.ProxyLBCertificates{}),
		Fields: []*dsl.FieldDesc{
			{
				Name: "ServerCertificate",
				Type: meta.TypeString,
			},
			{
				Name: "IntermediateCertificate",
				Type: meta.TypeString,
			},
			{
				Name: "PrivateKey",
				Type: meta.TypeString,
			},
			{
				Name: "AdditionalCerts",
				Type: &dsl.Model{
					Name:    proxyLBAPIName + "AdditionalCert",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						{
							Name: "ServerCertificate",
							Type: meta.TypeString,
						},
						{
							Name: "IntermediateCertificate",
							Type: meta.TypeString,
						},
						{
							Name: "PrivateKey",
							Type: meta.TypeString,
						},
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]AdditionalCerts, recursive",
				},
			},
		},
	}

	proxyLBHealth = &dsl.Model{
		Name: "ProxyLBHealth",
		Fields: []*dsl.FieldDesc{
			{
				Name: "ActiveConn",
				Type: meta.TypeInt,
			},
			{
				Name: "CPS",
				Type: meta.TypeInt,
			},
			{
				Name: "CurrentVIP",
				Type: meta.TypeString,
			},
			{
				Name: "Servers",
				Type: &dsl.Model{
					Name:    "LoadBalancerServerStatus",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						{
							Name: "ActiveConn",
							Type: meta.TypeInt,
						},
						{
							Name: "Status",
							Type: meta.TypeInstanceStatus,
						},
						{
							Name: "IPAddress",
							Type: meta.TypeString,
						},
						{
							Name: "Port",
							Type: meta.TypeStringNumber,
						},
						{
							Name: "CPS",
							Type: meta.TypeInt,
						},
					},
				},
				Tags: &dsl.FieldTags{
					MapConv: "[]Servers,recursive",
				},
			},
		},
		NakedType: meta.Static(naked.ProxyLBHealth{}),
	}
)
