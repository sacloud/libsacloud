package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestPacketFilterOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel:           true,
		IgnoreStartupWait:  true,
		SetupAPICallerFunc: singletonAPICaller,
		Create: &CRUDTestFunc{
			Func: testPacketFilterCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createPacketFilterExpected,
				IgnoreFields: packetFilterIgnoreFields,
			},
		},
		Read: &CRUDTestFunc{
			Func: testPacketFilterRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createPacketFilterExpected,
				IgnoreFields: packetFilterIgnoreFields,
			},
		},
		Update: &CRUDTestFunc{
			Func: testPacketFilterUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updatePacketFilterExpected,
				IgnoreFields: packetFilterIgnoreFields,
			},
		},
		Delete: &CRUDTestDeleteFunc{
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
)

func testPacketFilterCreate(_ *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	res, err := client.Create(context.Background(), testZone, createPacketFilterParam)
	if err != nil {
		return nil, err
	}
	return res.PacketFilter, nil
}

func testPacketFilterRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.PacketFilter, nil
}

func testPacketFilterUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewPacketFilterOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updatePacketFilterParam)
	if err != nil {
		return nil, err
	}
	return res.PacketFilter, nil
}

func testPacketFilterDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewPacketFilterOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
