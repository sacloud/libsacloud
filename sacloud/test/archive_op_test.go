package test

import (
	"context"
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestArchiveOpCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Setup: func(testContext *CRUDTestContext, caller sacloud.APICaller) error {
			client := sacloud.NewArchiveOp(caller)
			findResult, err := client.Find(context.Background(), testZone, nil)
			if err != nil {
				return err
			}
			archives := findResult.Archives
			for _, a := range archives {
				if a.GetSizeGB() == 20 && a.Availability.IsAvailable() {
					testContext.Values["archive"] = a.ID
					createArchiveParam.SourceArchiveID = a.ID
					createArchiveExpected.SourceArchiveID = a.ID
					createArchiveExpected.SourceArchiveAvailability = a.Availability
					updateArchiveExpected.SourceArchiveID = a.ID
					updateArchiveExpected.SourceArchiveAvailability = a.Availability

					return nil
				}
			}
			return errors.New("valid archive is not found")
		},

		Create: &CRUDTestFunc{
			Func: testArchiveCreate,
			Expect: &CRUDTestExpect{
				ExpectValue:  createArchiveExpected,
				IgnoreFields: ignoreArchiveFields,
			},
		},

		Read: &CRUDTestFunc{
			Func: testArchiveRead,
			Expect: &CRUDTestExpect{
				ExpectValue:  createArchiveExpected,
				IgnoreFields: ignoreArchiveFields,
			},
		},

		Update: &CRUDTestFunc{
			Func: testArchiveUpdate,
			Expect: &CRUDTestExpect{
				ExpectValue:  updateArchiveExpected,
				IgnoreFields: ignoreArchiveFields,
			},
		},

		Delete: &CRUDTestDeleteFunc{
			Func: testArchiveDelete,
		},
	})
}

var (
	ignoreArchiveFields = []string{
		"ID",
		"DisplayOrder",
		"Availability",
		"DiskPlanName",
		"SizeMB",
		"MigratedMB",
		"DiskPlanStorageClass",
		"SourceDiskID",
		"SourceDiskAvailability",
		"BundleInfo",
		"Storage",
		"IconID",
		"CreatedAt",
		"ModifiedAt",
		"OriginalArchiveID",
		"SourceInfo",
	}

	createArchiveParam = &sacloud.ArchiveCreateRequest{
		Name:        "libsacloud-archive",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
	}
	createArchiveExpected = &sacloud.Archive{
		Name:        createArchiveParam.Name,
		Description: createArchiveParam.Description,
		Tags:        createArchiveParam.Tags,
		Scope:       types.Scopes.User,
		DiskPlanID:  types.ID(2),
	}
	updateArchiveParam = &sacloud.ArchiveUpdateRequest{
		Name:        "libsacloud-archive-upd",
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
	}
	updateArchiveExpected = &sacloud.Archive{
		Name:        updateArchiveParam.Name,
		Description: updateArchiveParam.Description,
		Tags:        updateArchiveParam.Tags,
		Scope:       types.Scopes.User,
		DiskPlanID:  types.ID(2),
	}
)

func testArchiveCreate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	res, err := client.Create(context.Background(), testZone, createArchiveParam)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func testArchiveRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func testArchiveUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewArchiveOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateArchiveParam)
	if err != nil {
		return nil, err
	}
	return res.Archive, nil
}

func testArchiveDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewArchiveOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestArchiveOp_CreateBlank(t *testing.T) {
	t.Parallel()

	client := sacloud.NewArchiveOp(singletonAPICaller())

	createResult, err := client.CreateBlank(context.Background(), testZone, &sacloud.ArchiveCreateBlankRequest{
		SizeMB: 20 * 1024,
		Name:   "libsacloud-archive-blank",
	})
	require.NoError(t, err)
	require.NotNil(t, createResult.Archive)
	require.NotNil(t, createResult.FTPServer)
	defer func() {
		client.Delete(context.Background(), testZone, createResult.Archive.ID) // nolint ignore error
	}()
}
