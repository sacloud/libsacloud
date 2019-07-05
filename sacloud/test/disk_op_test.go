package test

import (
	"context"
	"strings"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDiskOpBlankDiskCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

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
		"IconID",
		"CreatedAt",
		"ModifiedAt",
	}

	createDiskParam = &sacloud.DiskCreateRequest{
		DiskPlanID:  types.ID(4), //SSD
		Connection:  types.DiskConnections.VirtIO,
		Name:        "libsacloud-disk",
		Description: "desc",
		Tags:        []string{"tag1", "tag2"},
		SizeMB:      20 * 1024,
	}
	createDiskExpected = &sacloud.Disk{
		Name:        createDiskParam.Name,
		Description: createDiskParam.Description,
		Tags:        createDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
	}
	updateDiskParam = &sacloud.DiskUpdateRequest{
		Name:        "libsacloud-disk-upd",
		Description: "desc-upd",
		Tags:        []string{"tag1-upd", "tag2-upd"},
	}
	updateDiskExpected = &sacloud.Disk{
		Name:        updateDiskParam.Name,
		Description: updateDiskParam.Description,
		Tags:        updateDiskParam.Tags,
		DiskPlanID:  createDiskParam.DiskPlanID,
		Connection:  createDiskParam.Connection,
	}
)

func testDiskCreate(_ *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	res, err := client.Create(context.Background(), testZone, createDiskParam)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func testDiskRead(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	res, err := client.Read(context.Background(), testZone, testContext.ID)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil
}

func testDiskUpdate(testContext *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	res, err := client.Update(context.Background(), testZone, testContext.ID, updateDiskParam)
	if err != nil {
		return nil, err
	}
	return res.Disk, nil

}

func testDiskDelete(testContext *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDiskOp(caller)
	return client.Delete(context.Background(), testZone, testContext.ID)
}

func TestDiskOp_Config(t *testing.T) {
	t.Parallel()

	archiveClient := sacloud.NewArchiveOp(singletonAPICaller())
	client := sacloud.NewDiskOp(singletonAPICaller())
	ctx := context.Background()

	// find source public archive
	var archiveID types.ID
	archiveFindResult, err := archiveClient.Find(ctx, testZone, nil)
	require.NoError(t, err)
	for _, a := range archiveFindResult.Archives {
		if strings.HasPrefix(a.Name, "CentOS 7") {
			archiveID = a.ID
			break
		}
	}
	if archiveID.IsEmpty() {
		t.Fatal("archive is not found")
	}

	// create
	diskCreateResult, err := client.Create(ctx, testZone, &sacloud.DiskCreateRequest{
		Name:            "libsacloud-disk-edit",
		DiskPlanID:      types.ID(4),
		SizeMB:          20 * 1024,
		SourceArchiveID: archiveID,
	})
	require.NoError(t, err)
	disk := diskCreateResult.Disk

	// wait for ready
	waiter := sacloud.WaiterForReady(func() (interface{}, error) {
		res, err := client.Read(ctx, testZone, disk.ID)
		if err != nil {
			return nil, err
		}
		return res.Disk, nil
	})
	_, err = waiter.WaitForState(ctx)
	require.NoError(t, err)

	defer func() {
		// cleanup
		if err := client.Delete(ctx, testZone, disk.ID); err != nil {
			t.Fatal(err)
		}
	}()

	// edit disk
	err = client.Config(context.Background(), testZone, disk.ID, &sacloud.DiskEditRequest{
		Password: "password",
		SSHKeys: []*sacloud.DiskEditSSHKey{
			{
				PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4LDQuDiKecOJDPY9InS7EswZ2fPnoRZXc48T1EqyRLyJhgEYGSDWaBiMDs2R/lWgA81Hp37qhrNqZPjFHUkBr93FOXxt9W0m1TNlkNepK0Uyi+14B2n0pdoeqsKEkb3sTevWF0ztxxWrwUd7Mems2hf+wFODITHYye9RlDAKLKPCFRvlQ9xQj4bBWOogQwoaXMSK1znMPjudcm1tRry4KIifLdXmwVKU4qDPGxoXfqs44Dgsikk43UVBStQ7IFoqPgAqcJFSGHLoMS7tPKdTvY9+GME5QidWK84gl69piAkgIdwd+JTMUOc/J+9DXAt220HqZ6l3yhWG5nIgi0x8n",
			},
		},
		DisablePWAuth: true,
		EnableDHCP:    true,
		HostName:      "hostname",
		//Notes: []*DiskEditNote{
		//	{
		//		ID: types.ID(123456789012),
		//	},
		//},
		UserIPAddress: "192.2.0.11",
		UserSubnet: &sacloud.DiskEditUserSubnet{
			DefaultRoute:   "192.2.0.1",
			NetworkMaskLen: 24,
		},
	})
	require.NoError(t, err)

}
