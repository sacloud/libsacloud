package test

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func TestDatabaseOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,
		Setup:              setupSwitchFunc("db", createDatabaseParam, createDatabaseExpected, updateDatabaseExpected),
		Create: &CRUDTestFunc{
			Func: testDatabaseCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDatabaseExpected,
				IgnoreFields: ignoreDatabaseFields,
			}),
		},
		Read: &CRUDTestFunc{
			Func: testDatabaseRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDatabaseExpected,
				IgnoreFields: ignoreDatabaseFields,
			}),
		},
		Updates: []*CRUDTestFunc{
			{
				Func: testDatabaseUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateDatabaseExpected,
					IgnoreFields: ignoreDatabaseFields,
				}),
			},
		},
		Shutdown: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewDatabaseOp(caller)
			return client.Shutdown(context.Background(), testZone, testContext.ID, &sacloud.ShutdownOption{Force: true})
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testDatabaseDelete,
		},

		Cleanup: cleanupSwitchFunc("db"),
	})
}

var (
	ignoreDatabaseFields = []string{
		"ID",
		"Class",
		"Tags", // Create(POST)時は指定したタグが返ってくる。その後利用可能になったらデータベースの種類に応じて@MariaDBxxxのようなタグが付与される
		"Availability",
		"InstanceStatus",
		"InstanceHostName",
		"InstanceHostInfoURL",
		"InstanceStatusChangedAt",
		"Interfaces",
		"Switch",
		"ZoneID",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
		"SettingsHash",
	}

	createDatabaseParam = &sacloud.DatabaseCreateRequest{
		PlanID:         types.ID(10),
		IPAddresses:    []string{"192.168.0.11"},
		NetworkMaskLen: 24,
		DefaultRoute:   "192.168.0.1",
		Name:           "libsacloud-db",
		Description:    "desc",
		Tags:           []string{"tag1", "tag2"},

		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     "MariaDB",
			DatabaseVersion:  "10.3",
			DatabaseRevision: "10.3.15",
			DefaultUser:      "exa.mple",
			UserPassword:     "LibsacloudExamplePassword01",
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort:  5432,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword01",
		},
	}
	createDatabaseExpected = &sacloud.Database{
		Name:           createDatabaseParam.Name,
		Description:    createDatabaseParam.Description,
		Availability:   types.Availabilities.Available,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		PlanID:         createDatabaseParam.PlanID,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting:  createDatabaseParam.CommonSetting,
	}
	updateDatabaseParam = &sacloud.DatabaseUpdateRequest{
		Name:        "libsacloud-db-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
		Description: "desc-upd",
		CommonSetting: &sacloud.DatabaseSettingCommonUpdate{
			ServicePort:  5432,
			DefaultUser:  "exa.mple",
			UserPassword: "LibsacloudExamplePassword01",
		},
	}
	updateDatabaseExpected = &sacloud.Database{
		Name:           updateDatabaseParam.Name,
		Description:    updateDatabaseParam.Description,
		Availability:   types.Availabilities.Available,
		PlanID:         createDatabaseParam.PlanID,
		InstanceStatus: types.ServerInstanceStatuses.Up,
		DefaultRoute:   createDatabaseParam.DefaultRoute,
		NetworkMaskLen: createDatabaseParam.NetworkMaskLen,
		IPAddresses:    createDatabaseParam.IPAddresses,
		Conf:           createDatabaseParam.Conf,
		CommonSetting:  createDatabaseParam.CommonSetting,
	}
)

func testDatabaseCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Create(context.Background(), testZone, createDatabaseParam)
}

func testDatabaseRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testDatabaseUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDatabaseOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateDatabaseParam)
}

func testDatabaseDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDatabaseOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
