package define

import (
	"github.com/sacloud/libsacloud/v2/internal/define/names"
	"github.com/sacloud/libsacloud/v2/internal/define/ops"
	"github.com/sacloud/libsacloud/v2/internal/dsl"
	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud/naked"
)

const (
	licenseAPIName     = "License"
	licenseAPIPathName = "license"
)

var licenseAPI = &dsl.Resource{
	Name:       licenseAPIName,
	PathName:   licenseAPIPathName,
	PathSuffix: dsl.CloudAPISuffix,
	IsGlobal:   true,
	Operations: dsl.Operations{
		// find
		ops.Find(licenseAPIName, licenseNakedType, findParameter, licenseView),

		// create
		ops.Create(licenseAPIName, licenseNakedType, licenseCreateParam, licenseView),

		// read
		ops.Read(licenseAPIName, licenseNakedType, licenseView),

		// update
		ops.Update(licenseAPIName, licenseNakedType, licenseUpdateParam, licenseView),

		// patch
		ops.Patch(licenseAPIName, licenseNakedType, patchModel(licenseUpdateParam), licenseView),

		// delete
		ops.Delete(licenseAPIName),
	},
}

var (
	licenseNakedType = meta.Static(naked.License{})

	licenseView = &dsl.Model{
		Name:      licenseAPIName,
		NakedType: licenseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.LicenseInfoID(),
			fields.LicenseInfoName(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	licenseCreateParam = &dsl.Model{
		Name:      names.CreateParameterName(licenseAPIName),
		NakedType: licenseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
			fields.LicenseInfoID(),
		},
	}

	licenseUpdateParam = &dsl.Model{
		Name:      names.UpdateParameterName(licenseAPIName),
		NakedType: licenseNakedType,
		Fields: []*dsl.FieldDesc{
			fields.Name(),
		},
	}
)
