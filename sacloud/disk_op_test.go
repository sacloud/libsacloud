package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func TestDiskOpBlankDiskCRUD(t *testing.T) {
	Test(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICaller: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testDiskCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testDiskRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testDiskUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateDiskExpected,
				IgnoreFields: ignoreDiskFields,
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testDiskDelete,
		},
	})
}

var (
	ignoreDiskFields = []string{
		"ID",
		"DisplayOrder",
		"Availability",
		"DiskPlanName",
		"DiskPlanStorageClass",
		"SizeMB",
		"MigratedMB",
		"SourceDiskID",
		"SourceDiskAvailability",
		"SourceArchiveID",
		"SourceArchiveAvailability",
		"BundleInfo",
		"Server",
		"Storage",
		"Icon",
		"CreatedAt",
		"ModifiedAt",
	}

	createDiskParam = &DiskCreateRequest{
		DiskPlanID:  types.ID(4), //SSD
		Connection:  types.DiskConnections.VirtIO,
		Name:        "libsacloud-v2-disk",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		SizeMB:      20 * 1024,
	}
	createDiskExpected = &Disk{
		Name:        createDiskParam.Name,
		Description: createDiskParam.Description,
		Tags:        createDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
	}
	updateDiskParam = &DiskUpdateRequest{
		Name:        "libsacloud-v2-disk-upd",
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
	}
	updateDiskExpected = &Disk{
		Name:        updateDiskParam.Name,
		Description: updateDiskParam.Description,
		Tags:        updateDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
	}
)

func testDiskCreate(_ *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewDiskOp(caller)
	return client.Create(context.Background(), testZone, createDiskParam)
}

func testDiskRead(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewDiskOp(caller)
	return client.Read(context.Background(), testZone, testContext.ID)
}

func testDiskUpdate(testContext *CRUDTestContext, caller APICaller) (interface{}, error) {
	client := NewDiskOp(caller)
	return client.Update(context.Background(), testZone, testContext.ID, updateDiskParam)
}

func testDiskDelete(testContext *CRUDTestContext, caller APICaller) error {
	client := NewDiskOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}
