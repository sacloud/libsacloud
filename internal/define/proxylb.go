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
		// updateSettings
		ops.UpdateCommonServiceItemSettings(proxyLBAPIName, proxyLBUpdateSettingsNakedType, proxyLBUpdateSettingsParam, proxyLBView),

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
				dsl.ArgumentID,
			},
		},

		// renew Let's Encrypt certificates
		ops.WithIDAction(proxyLBAPIName, "RenewLetsEncryptCert", http.MethodPut, "proxylb/letsencryptrenew"),

		// Health
		ops.HealthStatus(proxyLBAPIName, meta.Static(naked.ProxyLBHealth{}), proxyLBHealth),

		// Monitor
		ops.MonitorChild(proxyLBAPIName, "Connection", "activity/proxylb",
			monitorParameter, monitors.connectionModel()),
	},
}

var (
	proxyLBNakedType               = meta.Static(naked.ProxyLB{})
	proxyLBUpdateSettingsNakedType = meta.Static(naked.ProxyLBSettingsUpdate{})
	proxyLBCertificatesNakedType   = meta.Static(naked.ProxyLBCertificates{})

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

			fields.ProxyLBPlan(),

			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),
			fields.ProxyLBTimeout(),
			fields.SettingsHash(),

			// status
			fields.ProxyLBUseVIPFailover(),
			fields.ProxyLBRegion(),
			fields.ProxyLBProxyNetworks(),
			fields.ProxyLBFQDN(),
			fields.ProxyLBVIP(),
		},
	}

	proxyLBCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(proxyLBAPIName),
		NakedType: proxyLBNakedType,
		ConstFields: []*dsl.ConstFieldDesc{
			{
				Name: "Class",
				Type: meta.TypeString,
				Tags: &dsl.FieldTags{
					MapConv: "Provider.Class",
				},
				Value: `"proxylb"`,
			},
		},
		Fields: []*dsl.FieldDesc{
			// required
			fields.ProxyLBPlan(),

			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),
			fields.ProxyLBTimeout(),

			// status
			fields.ProxyLBUseVIPFailover(),
			fields.ProxyLBRegion(),

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
			fields.ProxyLBTimeout(),
			// settings hash
			fields.SettingsHash(),

			// common fields
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	proxyLBUpdateSettingsParam = &dsl.Model{
		Name:      names.UpdateSettingsParameterName(proxyLBAPIName),
		NakedType: proxyLBNakedType,
		Fields: []*dsl.FieldDesc{
			// settings
			fields.ProxyLBHealthCheck(),
			fields.ProxyLBSorryServer(),
			fields.ProxyLBBindPorts(),
			fields.ProxyLBServers(),
			fields.ProxyLBLetsEncrypt(),
			fields.ProxyLBStickySession(),
			fields.ProxyLBTimeout(),
			// settings hash
			fields.SettingsHash(),
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
				Name: "PrimaryCert",
				Type: &dsl.Model{
					Name: proxyLBAPIName + "PrimaryCert",
					Fields: []*dsl.FieldDesc{
						fields.Def("ServerCertificate", meta.TypeString),
						fields.Def("IntermediateCertificate", meta.TypeString),
						fields.Def("PrivateKey", meta.TypeString),
						fields.Def("CertificateEndDate", meta.TypeTime),
						fields.Def("CertificateCommonName", meta.TypeString),
					},
				},
			},
			{
				Name: "AdditionalCerts",
				Type: &dsl.Model{
					Name:    proxyLBAPIName + "AdditionalCert",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("ServerCertificate", meta.TypeString),
						fields.Def("IntermediateCertificate", meta.TypeString),
						fields.Def("PrivateKey", meta.TypeString),
						fields.Def("CertificateEndDate", meta.TypeTime),
						fields.Def("CertificateCommonName", meta.TypeString),
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
				Name: "PrimaryCerts",
				Type: &dsl.Model{
					Name: proxyLBAPIName + "PrimaryCert",
					Fields: []*dsl.FieldDesc{
						fields.Def("ServerCertificate", meta.TypeString),
						fields.Def("IntermediateCertificate", meta.TypeString),
						fields.Def("PrivateKey", meta.TypeString),
					},
				},
			},
			{
				Name: "AdditionalCerts",
				Type: &dsl.Model{
					Name:    proxyLBAPIName + "AdditionalCert",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("ServerCertificate", meta.TypeString),
						fields.Def("IntermediateCertificate", meta.TypeString),
						fields.Def("PrivateKey", meta.TypeString),
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
			fields.Def("ActiveConn", meta.TypeInt),
			fields.Def("CPS", meta.TypeInt),
			fields.Def("CurrentVIP", meta.TypeString),
			{
				Name: "Servers",
				Type: &dsl.Model{
					Name:    "LoadBalancerServerStatus",
					IsArray: true,
					Fields: []*dsl.FieldDesc{
						fields.Def("ActiveConn", meta.TypeInt),
						fields.Def("Status", meta.TypeInstanceStatus),
						fields.Def("IPAddress", meta.TypeString),
						fields.Def("Port", meta.TypeStringNumber),
						fields.Def("CPS", meta.TypeInt),
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
