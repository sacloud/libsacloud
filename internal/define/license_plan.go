package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	licenseInfoAPIName     = "LicenseInfo"
	licenseInfoAPIPathName = "product/license"
)

var licenseInfoAPI = &dsl.Resource{
	Name:       licenseInfoAPIName,
	PathName:   licenseInfoAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		ops.Find(licenseInfoAPIName, licenseInfoNakedType, findParameter, licenseInfoView),
		ops.Read(licenseInfoAPIName, licenseInfoNakedType, licenseInfoView),
	},
}

var (
	licenseInfoNakedType = meta.Static(naked.LicenseInfo{})
	licenseInfoView      = &dsl.Model{
		Name:      licenseInfoAPIName,
		NakedType: licenseInfoNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
			fields.Def("TermsOfUse", meta.TypeString),
		},
	}
)
