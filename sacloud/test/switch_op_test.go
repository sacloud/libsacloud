package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestSwitchOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testSwitchCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Read: &CRUDTestFunc{
			Func: testSwitchRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Update: &CRUDTestFunc{
			Func: testSwitchUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateSwitchExpected,
				IgnoreFields: ignoreSwitchFields,
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testSwitchDelete,
		},
	})
}

var (
	ignoreSwitchFields = []string{
		"ID",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
	}
	createSwitchParam = &sacloud.SwitchCreateRequest{
		Name:           "libsacloud-switch",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},
		DefaultRoute:   "192.168.0.1",
		NetworkMaskLen: 24,
	}
	createSwitchExpected = &sacloud.Switch{
		Name:           createSwitchParam.Name,
		Description:    createSwitchParam.Description,
		Tags:           createSwitchParam.Tags,
		DefaultRoute:   createSwitchParam.DefaultRoute,
		NetworkMaskLen: createSwitchParam.NetworkMaskLen,
		Scope:          types.Scopes.User,
	}
	updateSwitchParam = &sacloud.SwitchUpdateRequest{
		Name:           "libsacloud-switch-upd",
		Tags:           []string{"tag1-upd", "tag2-upd"},
		Description:    "desc-upd",
		DefaultRoute:   "192.168.0.2",
		NetworkMaskLen: 28,
	}
	updateSwitchExpected = &sacloud.Switch{
		Name:           updateSwitchParam.Name,
		Description:    updateSwitchParam.Description,
		Tags:           updateSwitchParam.Tags,
		DefaultRoute:   updateSwitchParam.DefaultRoute,
		NetworkMaskLen: updateSwitchParam.NetworkMaskLen,
		Scope:          createSwitchExpected.Scope,
	}
)

func testSwitchCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	res, err := client.Create(context.Background(), testZone, createSwitchParam)
	if err != nil {
		return nil, err
	}
	return res.Switch, nil
}

func testSwitchRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.Switch, nil
}

func testSwitchUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewSwitchOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateSwitchParam)
	if err != nil {
		return nil, err
	}
	return res.Switch, nil
}

func testSwitchDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewSwitchOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestSwitchOp_BridgeConnection(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	caller := singletonAPICaller()

	swOp := sacloud.NewSwitchOp(caller)
	bridgeOp := sacloud.NewBridgeOp(caller)

	// create switch
	swCreateResult, err := swOp.Create(ctx, testZone, &sacloud.SwitchCreateRequest{
		Name: "libsacloud-switch-for-bridge",
	})
	require.NoError(t, err)
	sw := swCreateResult.Switch

	bridgeCreateResult, err := bridgeOp.Create(ctx, testZone, &sacloud.BridgeCreateRequest{
		Name: "libsacloud-bridge",
	})
	require.NoError(t, err)
	bridge := bridgeCreateResult.Bridge

	// connect to bridge
	err = swOp.ConnectToBridge(ctx, testZone, sw.ID, bridge.ID)
	require.NoError(t, err)

	// confirm
	swReadResult, err := swOp.Read(ctx, testZone, sw.ID)
	require.NoError(t, err)
	sw = swReadResult.Switch
	require.Equal(t, bridge.ID, sw.BridgeID)

	bridgeReadResult, err := bridgeOp.Read(ctx, testZone, bridge.ID)
	require.NoError(t, err)
	bridge = bridgeReadResult.Bridge

	require.Equal(t, sw.ID, bridge.SwitchInZone.ID)
	require.Len(t, bridge.BridgeInfo, 0) // 他ゾーンのスイッチのみ

	// disconnect
	err = swOp.DisconnectFromBridge(ctx, testZone, sw.ID)
	require.NoError(t, err)

	// confirm
	swReadResult, err = swOp.Read(ctx, testZone, sw.ID)
	require.NoError(t, err)
	sw = swReadResult.Switch

	require.True(t, sw.BridgeID.IsEmpty())

	bridgeReadResult, err = bridgeOp.Read(ctx, testZone, bridge.ID)
	require.NoError(t, err)
	bridge = bridgeReadResult.Bridge

	require.Nil(t, bridge.SwitchInZone)
	require.Len(t, bridge.BridgeInfo, 0)

	// delete
	err = swOp.Delete(ctx, testZone, sw.ID)
	require.NoError(t, err)

	err = bridgeOp.Delete(ctx, testZone, bridge.ID)
	require.NoError(t, err)
}
