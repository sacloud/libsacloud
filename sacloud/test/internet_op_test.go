package test

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestInternetOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testInternetCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createInternetExpected,
				IgnoreFields: ignoreInternetFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testInternetRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createInternetExpected,
				IgnoreFields: ignoreInternetFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testInternetUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateInternetExpected,
					IgnoreFields: ignoreInternetFields,
				}),
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testInternetDelete,
		},
	})
}

var (
	ignoreInternetFields = []string{
		"ID",
		"IconID",
		"CreatedAt",
		"Switch",
	}
	createInternetParam = &sacloud.InternetCreateRequest{
		Name:           "libsacloud-internet",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
	}
	createInternetExpected = &sacloud.Internet{
		Name:           createInternetParam.Name,
		Description:    createInternetParam.Description,
		Tags:           createInternetParam.Tags,
		NetworkMaskLen: createInternetParam.NetworkMaskLen,
		BandWidthMbps:  createInternetParam.BandWidthMbps,
	}
	updateInternetParam = &sacloud.InternetUpdateRequest{
		Name:        "libsacloud-internet-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
	}
	updateInternetExpected = &sacloud.Internet{
		Name:           updateInternetParam.Name,
		Description:    updateInternetParam.Description,
		Tags:           updateInternetParam.Tags,
		NetworkMaskLen: createInternetParam.NetworkMaskLen,
		BandWidthMbps:  createInternetParam.BandWidthMbps,
	}
)

func testInternetCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	return client.Create(context.Background(), testZone, createInternetParam)
}

func testInternetRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testInternetUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateInternetParam)
}

func testInternetDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInternetOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestInternetOp_Subnet(t *testing.T) {
	client := sacloud.NewInternetOp(singletonAPICaller())
	var minIP, maxIP string
	var subnetID types.ID

	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				ctx := context.Background()

				var internet *sacloud.Internet
				internet, err := client.Create(ctx, testZone, createInternetParam)
				if err != nil {
					return nil, err
				}
				waiter := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
					return client.Read(ctx, testZone, internet.ID)
				}, 100)
				if _, err := waiter.WaitForState(context.TODO()); err != nil {
					t.Error("WaitForUp is failed: ", err)
					return nil, err
				}

				internet, err = client.Read(ctx, testZone, internet.ID)
				if err != nil {
					return nil, err
				}

				swOp := sacloud.NewSwitchOp(singletonAPICaller())
				sw, err := swOp.Read(ctx, testZone, internet.Switch.ID)
				if err != nil {
					return nil, err
				}
				minIP = sw.Subnets[0].AssignedIPAddressMin
				maxIP = sw.Subnets[0].AssignedIPAddressMax

				return internet, nil
			},
		},
		Read: &CRUDTestFunc{
			Func: testInternetRead,
		},
		Updates: []*CRUDTestFunc{
			// add subnet
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// add subnet
					subnet, err := client.AddSubnet(context.Background(), testZone, testContext.ID, &sacloud.InternetAddSubnetRequest{
						NetworkMaskLen: 28,
						NextHop:        minIP,
					})
					if err != nil {
						return nil, err
					}

					if !assert.Len(t, subnet.IPAddresses, 16) {
						return nil, errors.New("unexpected state: Subnet.IPAddresses")
					}
					if !assert.Equal(t, minIP, subnet.NextHop) {
						return nil, errors.New("unexpected state: Subnet.NextHop")
					}
					subnetID = subnet.ID
					return subnet, nil
				},
				SkipExtractID: true,
			},
			// update subnet
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					subnet, err := client.UpdateSubnet(context.Background(), testZone, testContext.ID, subnetID, &sacloud.InternetUpdateSubnetRequest{
						NextHop: maxIP,
					})
					if err != nil {
						return nil, err
					}

					if !assert.Len(t, subnet.IPAddresses, 16) {
						return nil, errors.New("unexpected state: Subnet.IPAddresses")
					}
					if !assert.Equal(t, maxIP, subnet.NextHop) {
						return nil, errors.New("unexpected state: Subnet.NextHop")
					}
					return subnet, nil
				},
				SkipExtractID: true,
			},
			// delete subnet
			{
				Func: func(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					return nil, client.DeleteSubnet(context.Background(), testZone, testContext.ID, subnetID)
				},
				SkipExtractID: true,
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testInternetDelete,
		},
	})
}
