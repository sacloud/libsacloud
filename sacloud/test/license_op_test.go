package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestLicenseOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testLicenseCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createLicenseExpected,
				IgnoreFields: ignoreLicenseFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testLicenseRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createLicenseExpected,
				IgnoreFields: ignoreLicenseFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testLicenseUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateLicenseExpected,
					IgnoreFields: ignoreLicenseFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testLicenseDelete,
		},
	})
}

var (
	ignoreLicenseFields = []string{
		"ID",
		"CreatedAt",
		"ModifiedAt",
	}

	createLicenseParam = &sacloud.LicenseCreateRequest{
		Name:          "libsacloud-license",
		LicenseInfoID: types.ID(10001),
	}
	createLicenseExpected = &sacloud.License{
		Name:            createLicenseParam.Name,
		LicenseInfoID:   createLicenseParam.LicenseInfoID,
		LicenseInfoName: "Windows RDS SAL",
	}
	updateLicenseParam = &sacloud.LicenseUpdateRequest{
		Name: "libsacloud-license-upd",
	}
	updateLicenseExpected = &sacloud.License{
		Name:            updateLicenseParam.Name,
		LicenseInfoID:   createLicenseParam.LicenseInfoID,
		LicenseInfoName: "Windows RDS SAL",
	}
)

func testLicenseCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Create(ctx, createLicenseParam)
}

func testLicenseRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Read(ctx, ctx.ID)
}

func testLicenseUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewLicenseOp(caller)
	return client.Update(ctx, ctx.ID, updateLicenseParam)
}

func testLicenseDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewLicenseOp(caller)
	return client.Delete(ctx, ctx.ID)
}
