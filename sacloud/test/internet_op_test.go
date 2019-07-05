package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestInternetOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
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
		Name:           "libsacloud-internet",
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
	vpcRouterDeleted = false
)

func testInternetCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	res, err := client.Create(context.Background(), testZone, createInternetParam)
	if err != nil {
		return nil, err
	}
	return res.Internet, nil
}

func testInternetRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	res, err := readInternet(testContext.ID, caller)
	if err != nil {
		return nil, err
	}
	return res.Internet, nil
}

func testInternetUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewInternetOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateInternetParam)
	if err != nil {
		return nil, err
	}
	return res.Internet, nil
}

func testInternetDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewInternetOp(caller)
	err := client.Delete(context.Background(), testZone, testContext.ID)
	if err == nil {
		vpcRouterDeleted = true
	}
	return err
}

func readInternet(id types.ID, caller sacloud.APICaller) (*sacloud.InternetReadResult, error) {
	client := sacloud.NewInternetOp(caller)
	if vpcRouterDeleted {
		return client.Read(context.Background(), testZone, id)
	}
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
	t.Parallel()

	client := sacloud.NewInternetOp(singletonAPICaller())
	ctx := context.Background()

	// prepare
	var internet *sacloud.Internet
	createResult, err := client.Create(ctx, testZone, createInternetParam)
	require.NoError(t, err)
	internet = createResult.Internet

	readResult, err := readInternet(internet.ID, singletonAPICaller())
	require.NoError(t, err)
	internet = readResult.Internet

	swOp := sacloud.NewSwitchOp(singletonAPICaller())
	swReadResult, err := swOp.Read(ctx, testZone, internet.Switch.ID)
	require.NoError(t, err)
	sw := swReadResult.Switch

	// add subnet
	addSubnetResult, err := client.AddSubnet(ctx, testZone, internet.ID, &sacloud.InternetAddSubnetRequest{
		NetworkMaskLen: 28,
		NextHop:        sw.Subnets[0].AssignedIPAddressMin,
	})
	require.NoError(t, err)
	subnet := addSubnetResult.Subnet

	require.Len(t, subnet.IPAddresses, 16)
	require.Equal(t, sw.Subnets[0].AssignedIPAddressMin, subnet.NextHop)

	// update
	updateSubnetResult, err := client.UpdateSubnet(ctx, testZone, internet.ID, subnet.ID, &sacloud.InternetUpdateSubnetRequest{
		NextHop: sw.Subnets[0].AssignedIPAddressMax,
	})
	require.NoError(t, err)
	updSubnet := updateSubnetResult.Subnet

	require.Len(t, updSubnet.IPAddresses, 16)
	require.Equal(t, sw.Subnets[0].AssignedIPAddressMax, updSubnet.NextHop)

	// delete
	err = client.DeleteSubnet(ctx, testZone, internet.ID, subnet.ID)
	require.NoError(t, err)

	// delete internet
	err = client.Delete(ctx, testZone, internet.ID)
	require.NoError(t, err)

}
