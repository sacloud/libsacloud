package test

import (
	"errors"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestDiskOpBlankDiskCRUD(t *testing.T) {
	Run(t, &CRUDTestCase{
		Parallel: true,

		SetupAPICallerFunc: singletonAPICaller,

		Create: &CRUDTestFunc{
			Func: testDiskCreate,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			}),
		},

		Read: &CRUDTestFunc{
			Func: testDiskRead,
			CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
				ExpectValue:  createDiskExpected,
				IgnoreFields: ignoreDiskFields,
			}),
		},

		Updates: []*CRUDTestFunc{
			{
				Func: testDiskUpdate,
				CheckFunc: AssertEqualWithExpected(&CRUDTestExpect{
					ExpectValue:  updateDiskExpected,
					IgnoreFields: ignoreDiskFields,
				}),
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

func testDiskCreate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Create(ctx, testZone, createDiskParam, nil)
}

func testDiskRead(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Read(ctx, testZone, ctx.ID)
}

func testDiskUpdate(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
	client := sacloud.NewDiskOp(caller)
	return client.Update(ctx, testZone, ctx.ID, updateDiskParam)
}

func testDiskDelete(ctx *CRUDTestContext, caller sacloud.APICaller) error {
	client := sacloud.NewDiskOp(caller)
	return client.Delete(ctx, testZone, ctx.ID)
}

func TestDiskOp_Config(t *testing.T) {

	// source archive
	var archiveID types.ID

	Run(t, &CRUDTestCase{
		Parallel:           true,
		SetupAPICallerFunc: singletonAPICaller,
		Setup: func(ctx *CRUDTestContext, caller sacloud.APICaller) error {
			archiveName := "CentOS"
			client := sacloud.NewArchiveOp(singletonAPICaller())
			searched, err := client.Find(ctx, testZone, &sacloud.FindCondition{
				Filter: search.Filter{
					search.Key("Name"): search.PartialMatch(archiveName),
				},
			})
			if !assert.NoError(t, err) {
				return err
			}
			if searched.Count == 0 {
				return errors.New("archive is not found")
			}
			archiveID = searched.Archives[0].ID
			return nil
		},
		Create: &CRUDTestFunc{
			Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
				client := sacloud.NewDiskOp(singletonAPICaller())
				disk, err := client.Create(ctx, testZone, &sacloud.DiskCreateRequest{
					Name:            "libsacloud-disk-edit",
					DiskPlanID:      types.ID(4),
					SizeMB:          20 * 1024,
					SourceArchiveID: archiveID,
				}, nil)
				if err != nil {
					return nil, err
				}
				if _, err = sacloud.WaiterForReady(func() (interface{}, error) {
					return client.Read(ctx, testZone, disk.ID)
				}).WaitForState(ctx); err != nil {
					return disk, err
				}

				return disk, nil
			},
		},
		Read: &CRUDTestFunc{
			Func: testDiskRead,
		},
		Updates: []*CRUDTestFunc{
			{
				Func: func(ctx *CRUDTestContext, caller sacloud.APICaller) (interface{}, error) {
					// edit disk
					client := sacloud.NewDiskOp(singletonAPICaller())
					err := client.Config(ctx, testZone, ctx.ID, &sacloud.DiskEditRequest{
						Password: "password",
						SSHKeys: []*sacloud.DiskEditSSHKey{
							{
								PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC4LDQuDiKecOJDPY9InS7EswZ2fPnoRZXc48T1EqyRLyJhgEYGSDWaBiMDs2R/lWgA81Hp37qhrNqZPjFHUkBr93FOXxt9W0m1TNlkNepK0Uyi+14B2n0pdoeqsKEkb3sTevWF0ztxxWrwUd7Mems2hf+wFODITHYye9RlDAKLKPCFRvlQ9xQj4bBWOogQwoaXMSK1znMPjudcm1tRry4KIifLdXmwVKU4qDPGxoXfqs44Dgsikk43UVBStQ7IFoqPgAqcJFSGHLoMS7tPKdTvY9+GME5QidWK84gl69piAkgIdwd+JTMUOc/J+9DXAt220HqZ6l3yhWG5nIgi0x8n",
							},
						},
						DisablePWAuth: true,
						EnableDHCP:    true,
						HostName:      "hostname",
						UserIPAddress: "192.2.0.11",
						UserSubnet: &sacloud.DiskEditUserSubnet{
							DefaultRoute:   "192.2.0.1",
							NetworkMaskLen: 24,
						},
					})
					return nil, err
				},
			},
		},
		Delete: &CRUDTestDeleteFunc{
			Func: testDiskDelete,
		},
	})

}
