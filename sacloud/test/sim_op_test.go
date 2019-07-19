package test

import (
	"context"
	"os"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestSIMOpCRUD(t *testing.T) {

	PreCheckEnvsFunc("SAKURACLOUD_SIM_ICCID", "SAKURACLOUD_SIM_PASSCODE")(t)

	initSIMVariables()

	Run(t, &CRUDTestCase{
		Parallel:          true,
		IgnoreStartupWait: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testSIMCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSIMExpected,
				IgnoreFields: ignoreSIMFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testSIMRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createSIMExpected,
				IgnoreFields: ignoreSIMFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testSIMUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateSIMExpected,
					IgnoreFields: ignoreSIMFields,
				}),
			},
			// activate
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.Activate(context.Background(), sacloud.APIDefaultZone, testContext.ID); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertNotNilFunc(t, simInfo, "SIMInfo"),
						AssertTrueFunc(t, simInfo.Activated, "SIMInfo.Activated"),
					)
				},
				SkipExtractID: true,
			},
			// deactivate
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.Deactivate(context.Background(), sacloud.APIDefaultZone, testContext.ID); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertNotNilFunc(t, simInfo, "SIMInfo"),
						AssertFalseFunc(t, simInfo.Activated, "SIMInfo.Activated"),
					)
				},
				SkipExtractID: true,
			},
			// IMEI lock
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.IMEILock(context.Background(), sacloud.APIDefaultZone, testContext.ID, &sacloud.SIMIMEILockRequest{
						IMEI: "123456789012345",
					}); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertTrueFunc(t, simInfo.IMEILock, "SIMInfo.IMEILock"),
					)
				},
				SkipExtractID: true,
			},
			// IMEI unlock
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					if err := client.IMEIUnlock(context.Background(), sacloud.APIDefaultZone, testContext.ID); err != nil {
						return nil, err
					}
					return client.Status(context.Background(), sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					simInfo := v.(*sacloud.SIMInfo)
					return DoAsserts(
						AssertFalseFunc(t, simInfo.IMEILock, "SIMInfo.IMEILock"),
					)
				},
				SkipExtractID: true,
			},
			// network operator
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					client := sacloud.NewSIMOp(caller)
					configs := []*sacloud.SIMNetworkOperatorConfig{
						{
							Name:  "SoftBank",
							Allow: true,
						},
					}
					if err := client.SetNetworkOperator(context.Background(), sacloud.APIDefaultZone, testContext.ID, configs); err != nil {
						return nil, err
					}
					return client.GetNetworkOperator(context.Background(), sacloud.APIDefaultZone, testContext.ID)
				},
				CheckFunc: func(t TestT, testContext *CRUDTestContext, v interface{}) error {
					config := v.([]*sacloud.SIMNetworkOperatorConfig)
					return DoAsserts(
						AssertNotEmptyFunc(t, config, "NetworkOperatorConfig"),
					)
				},
				SkipExtractID: true,
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testSIMDelete,
		},
	})
}

func TestSIMOp_Logs(t *testing.T) {
	if !isAccTest() {
		t.Skip("TestSIMOp_Logs only exec at Acceptance Test")
	}
	PreCheckEnvsFunc("SAKURACLOUD_SIM_ID")(t)
	id := types.StringID(os.Getenv("SAKURACLOUD_SIM_ID"))

	client := sacloud.NewSIMOp(singletonAPICaller())
	logs, err := client.Logs(context.Background(), sacloud.APIDefaultZone, id)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)

}

func initSIMVariables() {

	iccid := os.Getenv("SAKURACLOUD_SIM_ICCID")
	passcode := os.Getenv("SAKURACLOUD_SIM_PASSCODE")

	createSIMParam = &sacloud.SIMCreateRequest{
		Name:        "libsacloud-sim",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		ICCID:       iccid,
		PassCode:    passcode,
	}
	createSIMExpected = &sacloud.SIM{
		Name:         createSIMParam.Name,
		Description:  createSIMParam.Description,
		Tags:         createSIMParam.Tags,
		Availability: types.Availabilities.Available,
		ICCID:        createSIMParam.ICCID,
	}
	updateSIMParam = &sacloud.SIMUpdateRequest{
		Name:        "libsacloud-sim-upd",
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
	}
	updateSIMExpected = &sacloud.SIM{
		Name:         updateSIMParam.Name,
		Description:  updateSIMParam.Description,
		Tags:         updateSIMParam.Tags,
		Availability: types.Availabilities.Available,
		ICCID:        createSIMParam.ICCID,
	}
}

var (
	ignoreSIMFields = []string{
		"ID",
		"Class",
		"IconID",
		"Info",
		"CreatedAt",
		"ModifiedAt",
	}
	createSIMParam    *sacloud.SIMCreateRequest
	createSIMExpected *sacloud.SIM
	updateSIMParam    *sacloud.SIMUpdateRequest
	updateSIMExpected *sacloud.SIM
)

func testSIMCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Create(context.Background(), sacloud.APIDefaultZone, createSIMParam)
}

func testSIMRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Read(context.Background(), sacloud.APIDefaultZone, testContext.ID)
}

func testSIMUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSIMOp(caller)
	return client.Update(context.Background(), sacloud.APIDefaultZone, testContext.ID, updateSIMParam)
}

func testSIMDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSIMOp(caller)
	return client.Delete(context.Background(), sacloud.APIDefaultZone, testContext.ID)
}
