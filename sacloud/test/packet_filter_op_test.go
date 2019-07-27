package test

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestPacketFilterOp_CRUD(t *testing.T) {
	testutil.Run(t, &testutil.CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &testutil.CRUDTestFunc{
			Func: testPacketFilterCreate,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createPacketFilterExpected,
				IgnoreFields: packetFilterIgnoreFields,
			}),
		},
		Read: &testutil.CRUDTestFunc{
			Func: testPacketFilterRead,
			CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
				ExpectValue:  createPacketFilterExpected,
				IgnoreFields: packetFilterIgnoreFields,
			}),
		},
		Updates: []*testutil.CRUDTestFunc{
			{
				Func: testPacketFilterUpdate,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updatePacketFilterExpected,
					IgnoreFields: packetFilterIgnoreFields,
				}),
			},
			{
				Func: testPacketFilterUpdateToMin,
				CheckFunc: testutil.AssertEqualWithExpected(&testutil.CRUDTestExpect{
					ExpectValue:  updatePacketFilterToMinExpected,
					IgnoreFields: packetFilterIgnoreFields,
				}),
			},
		},
		Delete: &testutil.CRUDTestDeleteFunc{
			Func: testPacketFilterDelete,
		},
	})
}

var (
	packetFilterIgnoreFields = []string{
		"ID",
		"CreatedAt",
		"RequiredHostVersion",
		"ExpressionHash",
	}

	createPacketFilterParam = &sacloud.PacketFilterCreateRequest{
		Name:        "libsacloud-packet-filter",
		Description: "desc",
		Expression: []*sacloud.PacketFilterExpression{
			{
				Protocol:      types.Protocols.TCP,
				SourceNetwork: types.PacketFilterNetwork("192.168.0.1"),
				SourcePort:    types.PacketFilterPort("3000-3100"),
				Action:        types.Actions.Allow,
			},
			{
				Protocol: types.Protocols.IP,
				Action:   types.Actions.Deny,
			},
		},
	}
	createPacketFilterExpected = &sacloud.PacketFilter{
		Name:        createPacketFilterParam.Name,
		Description: createPacketFilterParam.Description,
		Expression:  createPacketFilterParam.Expression,
	}
	updatePacketFilterParam = &sacloud.PacketFilterUpdateRequest{
		Name:        "libsacloud-packet-filter-upd",
		Description: "desc-upd",
		Expression: []*sacloud.PacketFilterExpression{
			{
				Protocol:        types.Protocols.TCP,
				SourceNetwork:   types.PacketFilterNetwork("192.168.0.2"),
				DestinationPort: types.PacketFilterPort("4000-41000"),
				Action:          types.Actions.Allow,
			},
			{
				Protocol:        types.Protocols.UDP,
				SourceNetwork:   types.PacketFilterNetwork("192.168.0.3"),
				DestinationPort: types.PacketFilterPort("5000-5100"),
				Action:          types.Actions.Allow,
			},
			{
				Protocol: types.Protocols.IP,
				Action:   types.Actions.Deny,
			},
		},
	}
	updatePacketFilterExpected = &sacloud.PacketFilter{
		Name:        updatePacketFilterParam.Name,
		Description: updatePacketFilterParam.Description,
		Expression:  updatePacketFilterParam.Expression,
	}

	updatePacketFilterToMinParam = &sacloud.PacketFilterUpdateRequest{
		Name: "libsacloud-packet-filter-to-min",
	}
	updatePacketFilterToMinExpected = &sacloud.PacketFilter{
		Name: updatePacketFilterToMinParam.Name,
	}
)

func testPacketFilterCreate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Create(ctx, testZone, createPacketFilterParam)
}

func testPacketFilterRead(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testPacketFilterUpdate(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updatePacketFilterParam)
}

func testPacketFilterUpdateToMin(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updatePacketFilterToMinParam)
}

func testPacketFilterDelete(ctx *testutil.CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}
