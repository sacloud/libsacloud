// Copyright 2016-2020 The Libsacloud Authors
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
	"fmt"
	"net/http"

	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	billAPIName     = "Bill"
	billAPIPathName = "bill"
)

var billAPI = &dsl.Resource{
	Name:       billAPIName,
	PathName:   billAPIPathName,
	PathSuffix: dsl.BillingAPISuffix, // 請求情報向けエンドポイント
	IsGlobal:   true,
	Operations: dsl.Operations{
		// by-contract
		{
			ResourceName:     billAPIName,
			Name:             "ByContract",
			PathFormat:       billByContractPath,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				billArgAccountID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: billNakedType,
				Name: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       billView,
				},
			},
		},
		// by-contract/year
		{
			ResourceName:     billAPIName,
			Name:             "ByContractYear",
			PathFormat:       billByContractYearPath,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				billArgAccountID,
				billArgYear,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: billNakedType,
				Name: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       billView,
				},
			},
		},
		// by-contract/year/month
		{
			ResourceName:     billAPIName,
			Name:             "ByContractYearMonth",
			PathFormat:       billByContractYearMonthPath,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				billArgAccountID,
				billArgYear,
				billArgMonth,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: billNakedType,
				Name: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       billView,
				},
			},
		},
		// by id(レスポンスは複数形)
		{
			ResourceName:     billAPIName,
			Name:             "Read",
			PathFormat:       billByIDPath,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: billNakedType,
				Name: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(billAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       billView,
				},
			},
		},
		// details
		{
			ResourceName:     billAPIName,
			Name:             "Details",
			PathFormat:       billDetailPath,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				billArgMemberCode,
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: billDetailNakedType,
				Name: names.ResourceFieldName(billAPIName+"Detail", dsl.PayloadForms.Plural),
			}),
			Results: dsl.Results{
				{
					SourceField: names.ResourceFieldName(billAPIName+"Detail", dsl.PayloadForms.Plural),
					DestField:   names.ResourceFieldName(billAPIName+"Detail", dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       billDetailView,
				},
			},
		},
		// detail csv
		{
			ResourceName: billAPIName,
			Name:         "DetailsCSV",
			PathFormat:   billDetailPath + "/csv",
			Method:       http.MethodGet,
			Arguments: dsl.Arguments{
				billArgMemberCode,
				dsl.ArgumentID,
			},
			ResponseEnvelope: dsl.ResponseEnvelope(&dsl.EnvelopePayloadDesc{
				Type: billDetailCSVNakedType,
				Name: "CSV",
			}),
			Results: dsl.Results{
				{
					SourceField: "CSV",
					DestField:   billDetailCSVView.Name,
					IsPlural:    false,
					Model:       billDetailCSVView,
				},
			},
		},
	},
}

var (
	billArgAccountID = &dsl.Argument{
		Name: "accountID",
		Type: meta.TypeID,
	}
	billArgYear = &dsl.Argument{
		Name: "year",
		Type: meta.TypeInt,
	}
	billArgMonth = &dsl.Argument{
		Name: "month",
		Type: meta.TypeInt,
	}
	billArgMemberCode = &dsl.Argument{
		Name: "MemberCode",
		Type: meta.TypeString,
	}

	billByContractPath          = fmt.Sprintf("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/by-contract/{{.%s}}", billArgAccountID.ArgName())
	billByContractYearPath      = fmt.Sprintf("%s/{{.%s}}", billByContractPath, billArgYear.ArgName())
	billByContractYearMonthPath = fmt.Sprintf("%s/{{.%s}}", billByContractYearPath, billArgMonth.ArgName())
	billByIDPath                = fmt.Sprintf("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/id/{{.%s}}", dsl.ArgumentID.ArgName())

	billDetailPath = fmt.Sprintf(
		"{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}detail/{{.%s}}/{{.%s}}",
		billArgMemberCode.ArgName(),
		dsl.ArgumentID.ArgName(),
	)
)

var (
	billNakedType = meta.Static(naked.Bill{})
	billView      = &dsl.Model{
		Name:      billAPIName,
		NakedType: billNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			billFields.Amount(),
			billFields.Date(),
			billFields.MemberID(),
			billFields.Paid(),
			billFields.PayLimit(),
			billFields.PaymentClassID(),
		},
	}

	billDetailNakedType = meta.Static(naked.BillDetail{})
	billDetailView      = &dsl.Model{
		Name:      billAPIName + "Detail",
		NakedType: billDetailNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			billFields.Amount(),
			fields.Description(),
			billFields.ServiceClassID(),
			billFields.Usage(),
			billFields.Zone(),
			billFields.ContractEndAt(),
		},
	}

	billDetailCSVNakedType = meta.Static(naked.BillDetailCSV{})
	billDetailCSVView      = &dsl.Model{
		Name:      billAPIName + "DetailCSV",
		NakedType: billDetailCSVNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Def("Count", meta.TypeInt),
			fields.Def("ResponsedAt", meta.TypeTime),
			fields.Def("Filename", meta.TypeString),
			fields.Def("RawBody", meta.TypeString),
			fields.Def("HeaderRow", meta.TypeStringSlice),
			fields.Def("BodyRows", meta.Static([][]string{})),
		},
	}
)

type billFieldsDef struct{}

var billFields = &billFieldsDef{}

func (f *billFieldsDef) Amount() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Amount",
		Type: meta.TypeInt64,
	}
}

func (f *billFieldsDef) BillID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
	}
}

func (f *billFieldsDef) Date() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Date",
		Type: meta.TypeTime,
	}
}

func (f *billFieldsDef) MemberID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "MemberID",
		Type: meta.TypeString,
	}
}

func (f *billFieldsDef) Paid() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Paid",
		Type: meta.TypeFlag,
	}
}

func (f *billFieldsDef) PayLimit() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PayLimit",
		Type: meta.TypeTime,
	}
}

func (f *billFieldsDef) PaymentClassID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "PaymentClassID",
		Type: meta.TypeID,
	}
}

func (f *billFieldsDef) ServiceClassID() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ServiceClassID",
		Type: meta.TypeID,
	}
}

func (f *billFieldsDef) Usage() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Usage",
		Type: meta.TypeInt64,
	}
}

func (f *billFieldsDef) Zone() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "Zone",
		Type: meta.TypeString,
	}
}

func (f *billFieldsDef) ContractEndAt() *dsl.FieldDesc {
	return &dsl.FieldDesc{
		Name: "ContractEndAt",
		Type: meta.TypeTime,
	}
}
