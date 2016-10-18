package api

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testDiskName = "libsacloud_test_disk_name"

func TestCRUDByDiskAPI(t *testing.T) {
	diskAPI := client.Disk
	//CREATE : empty disk
	disk := &sacloud.Disk{
		Name:       testDiskName,
		Plan:       sacloud.DiskPlanSSD,
		Connection: sacloud.DiskConnectionVirtio,
		SizeMB:     20480,
	}

	res, err := diskAPI.Create(disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	diskID := res.ID

	//wait
	err = diskAPI.SleepWhileCopying(diskID, 5+time.Minute) //日によって時間がかかることもあるため5分待つ
	assert.NoError(t, err)                                 //timeoutしたらerrに値が格納されている

	//READ
	disk, err = diskAPI.Read(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.NotEmpty(t, disk.ID)
	assert.True(t, disk.IsAvailable())
	assert.Equal(t, disk.Connection, sacloud.DiskConnectionVirtio)

	//UPDATE
	diskUpdateValue := &sacloud.Disk{
		Connection: sacloud.DiskConnectionIDE,
	}

	disk, err = diskAPI.Update(diskID, diskUpdateValue)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.Equal(t, disk.Connection, sacloud.DiskConnectionIDE)

	//DELETE
	disk, err = diskAPI.Delete(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
}

func TestCreateDiskFromSource(t *testing.T) {
	diskAPI := client.Disk

	archive, err := client.Archive.FindLatestStableCentOS()
	assert.NoError(t, err)
	archiveID := archive.ID

	//CREATE : empty disk
	disk := &sacloud.Disk{
		Name:       testDiskName,
		Plan:       sacloud.DiskPlanSSD,
		Connection: sacloud.DiskConnectionVirtio,
		SizeMB:     20480,
	}
	disk.SetSourceArchive(archiveID) //ソースアーカイブはIDだけ指定する

	res, err := diskAPI.Create(disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	diskID := res.ID

	//wait
	err = diskAPI.SleepWhileCopying(diskID, 5*time.Minute) //日によって時間がかかることもあるため5分待つ
	assert.NoError(t, err)                                 //timeoutしたらerrに値が格納されている

	createdDisk, err := diskAPI.Read(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdDisk)
	assert.Equal(t, createdDisk.SourceArchive.ID, archiveID)

	//DELETE
	disk, err = diskAPI.Delete(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
}

//func TestCanEditDisk(t *testing.T) {
//	api := client.Disk
//	client.Zone = "is1a"
//
//	// CentOS
//	res, err := api.CanEditDisk(123456789012)
//	assert.NoError(t, err)
//	assert.True(t, res)
//
//	// SourceDisk/Archive not found
//	res, err = api.CanEditDisk(123456789012)
//	assert.Error(t, err)
//	assert.False(t, res)
//
//	// Blank
//	res, err = api.CanEditDisk(123456789012)
//	assert.NoError(t, err)
//	assert.False(t, res)
//
//	// windows
//	res, err = api.CanEditDisk(123456789012)
//	assert.NoError(t, err)
//	assert.False(t, res)
//	// windows-child
//	res, err = api.CanEditDisk(123456789012)
//	assert.NoError(t, err)
//	assert.False(t, res)
//
//}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupTestDisk)
	testTearDownHandlers = append(testTearDownHandlers, cleanupTestDisk)
}

func cleanupTestDisk() {
	diskAPI := client.Disk
	res, err := diskAPI.withNameLike(testDiskName).Find()
	if err == nil && res.Count > 0 {
		diskAPI.Delete(res.Disks[0].ID)
	}
}
