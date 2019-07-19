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
	couponAPIName  = "Coupon"
	couponPathName = "coupon"
)

var couponAPI = &dsl.Resource{
	Name:       couponAPIName,
	PathName:   couponPathName,
	PathSuffix: dsl.BillingAPISuffix, // 請求情報向けエンドポイント
	IsGlobal:   true,
	Operations: dsl.Operations{
		{
			ResourceName:     couponAPIName,
			Name:             "Find",
			PathFormat:       couponPathFormat,
			Method:           http.MethodGet,
			UseWrappedResult: true,
			Arguments: dsl.Arguments{
				dsl.ArgumentZone,
				couponArgAccountID,
			},
			ResponseEnvelope: dsl.ResponseEnvelopePlural(&dsl.EnvelopePayloadDesc{
				Type: couponNakedType,
				Name: couponAPIName,
			}),
			Results: dsl.Results{
				{
					SourceField: couponAPIName,
					DestField:   names.ResourceFieldName(couponAPIName, dsl.PayloadForms.Plural),
					IsPlural:    true,
					Model:       couponView,
				},
			},
		},
	},
}

var (
	couponArgAccountID = &dsl.Argument{
		Name: "accountID",
		Type: meta.TypeID,
	}
	couponPathFormat = fmt.Sprintf("{{.rootURL}}/{{.zone}}/{{.pathSuffix}}/{{.pathName}}/{{.%s}}", couponArgAccountID.ArgName())
)

var (
	couponNakedType = meta.Static(naked.Coupon{})
	couponView      = &dsl.Model{
		Name:      "Coupon",
		NakedType: couponNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.New("MemberID", meta.TypeString),
			fields.New("ContractID", meta.TypeID),
			fields.New("ServiceClassID", meta.TypeID),
			fields.New("Discount", meta.TypeInt64),
			fields.New("AppliedAt", meta.TypeTime),
			fields.New("UntilAt", meta.TypeTime),
		},
	}
)