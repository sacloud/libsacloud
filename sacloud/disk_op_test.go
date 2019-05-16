package sacloud

import (
	"context"
	"strings"
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestDiskOpBlankDiskCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
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

func TestDiskOp_Config(t *testing.T) {
	if !isAccTest() {
		t.Skip("TESTACC is not set. skip")
	}

	archiveClient := NewArchiveOp(singletonAPICaller())
	client := NewDiskOp(singletonAPICaller())
	ctx := context.Background()

	// find source public archive
	var archiveID types.ID
	archives, err := archiveClient.Find(ctx, testZone, nil)
	require.NoError(t, err)
	for _, a := range archives {
		if strings.HasPrefix(a.Name, "CentOS 7") {
			archiveID = a.ID
			break
		}
	}
	if archiveID.IsEmpty() {
		t.Fatal("archive is not found")
	}

	// create
	disk, err := client.Create(ctx, testZone, &DiskCreateRequest{
		Name:            "libsacloud-v2-disk-edit",
		DiskPlanID:      types.ID(4),
		SizeMB:          20 * 1024,
		SourceArchiveID: archiveID,
	})
	require.NoError(t, err)

	// wait for ready
	waiter := WaiterForReady(func() (interface{}, error) { return client.Read(ctx, testZone, disk.ID) })
	_, err = waiter.WaitForState(ctx)
	require.NoError(t, err)

	defer func() {
		// cleanup
		if err := client.Delete(ctx, testZone, disk.ID); err != nil {
			t.Fatal(err)
		}
	}()

	// edit disk
	err = client.Config(context.Background(), testZone, disk.ID, &DiskEditRequest{
		Password: "password",
		SSHKeys: []*DiskEditSSHKey{
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
		UserSubnet: &DiskEditUserSubnet{
			DefaultRoute:   "192.2.0.1",
			NetworkMaskLen: 24,
		},
	})
	require.NoError(t, err)

}
