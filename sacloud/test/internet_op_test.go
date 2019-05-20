package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestInternetOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testInternetCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createInternetExpected,
				IgnoreFields: ignoreInternetFields,
			},
		},
		Read: &CRUDTestFunc{
			Func: testInternetRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createInternetExpected,
				IgnoreFields: ignoreInternetFields,
			},
		},
		Update: &CRUDTestFunc{
			Func: testInternetUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateInternetExpected,
				IgnoreFields: ignoreInternetFields,
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
		Name:           "libsacloud-v2-internet",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		NetworkMaskLen: 24,
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
		Name:        "libsacloud-v2-internet-upd",
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
	return readInternet(testContext.ID, caller)
}

func testInternetUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateInternetParam)
}

func testInternetDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInternetOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func readInternet(id types.ID, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	// TODO スイッチ+ルータ作成後しばらくは404が返ってくる問題にどう対応するか?
	max := 30
	for {
		if max == 0 {
			break
		}
		res, err := client.Read(context.Background(), testZone, id)
		if err != nil || sacloud.IsNotFoundError(err) {
			max--
			time.Sleep(3 * time.Second)
			continue
		}
		return res, err
	}
	return nil, fmt.Errorf("internet[%s] is not found", id)
}

func TestInternetOp_Subnets(t *testing.T) {

	client := sacloud.NewInternetOp(singletonAPICaller())
	ctx := context.Background()

	// prepare
	internet, err := client.Create(ctx, testZone, createInternetParam)
	require.NoError(t, err)

	v, err := readInternet(internet.ID, singletonAPICaller())
	require.NoError(t, err)
	internet = v.(*sacloud.Internet)

	swOp := sacloud.NewSwitchOp(singletonAPICaller())
	sw, err := swOp.Read(ctx, testZone, internet.Switch.ID)
	require.NoError(t, err)

	// add subnet
	subnet, err := client.AddSubnet(ctx, testZone, internet.ID, &sacloud.InternetAddSubnetRequest{
		NetworkMaskLen: 28,
		NextHop:        sw.Subnets[0].AssignedIPAddressMin,
	})
	require.NoError(t, err)
	require.Len(t, subnet.IPAddresses, 16)
	require.Equal(t, sw.Subnets[0].AssignedIPAddressMin, subnet.NextHop)

	// update
	updSubnet, err := client.UpdateSubnet(ctx, testZone, internet.ID, subnet.ID, &sacloud.InternetUpdateSubnetRequest{
		NextHop: sw.Subnets[0].AssignedIPAddressMax,
	})
	require.NoError(t, err)
	require.Len(t, updSubnet.IPAddresses, 16)
	require.Equal(t, sw.Subnets[0].AssignedIPAddressMax, updSubnet.NextHop)

	// delete
	err = client.DeleteSubnet(ctx, testZone, internet.ID, subnet.ID)
	require.NoError(t, err)

	// delete internet
	err = client.Delete(ctx, testZone, internet.ID)
	require.NoError(t, err)

}
