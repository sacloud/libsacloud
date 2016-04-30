package api

import (
	"github.com/stretchr/testify/assert"
	sakura "github.com/yamamoto-febc/libsacloud/resources"
	"testing"
	"time"
)

const testDiskName = "libsacloud_test_disk_name"

func TestCRUDByDiskAPI(t *testing.T) {
	diskAPI := client.Disk

	//CREATE : empty disk
	disk := &sakura.Disk{
		Name:       testDiskName,
		Plan:       sakura.DiskPlanSSD,
		Connection: sakura.DiskConnectionVirtio,
		SizeMB:     20480,
	}

	res, err := diskAPI.Create(disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	diskID := res.ID

	//wait
	err = diskAPI.WaitForAvailable(diskID, 5+time.Minute) //日によって時間がかかることもあるため5分待つ
	assert.NoError(t, err)                                //timeoutしたらerrに値が格納されている

	//READ
	disk, err = diskAPI.Read(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.NotEmpty(t, disk.ID)
	assert.True(t, disk.IsAvailable())
	assert.Equal(t, disk.Connection, sakura.DiskConnectionVirtio)

	//UPDATE
	diskUpdateValue := &sakura.Disk{
		Connection: sakura.DiskConnectionIDE,
	}

	disk, err = diskAPI.Update(diskID, diskUpdateValue)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.Equal(t, disk.Connection, sakura.DiskConnectionIDE)

	//DELETE
	disk, err = diskAPI.Delete(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
}

func TestCreateDiskFromSource(t *testing.T) {
	diskAPI := client.Disk

	archiveID, err := client.Archive.GetUbuntuArchiveID()
	assert.NoError(t, err)

	//CREATE : empty disk
	disk := &sakura.Disk{
		Name:       testDiskName,
		Plan:       sakura.DiskPlanSSD,
		Connection: sakura.DiskConnectionVirtio,
		SizeMB:     20480,
	}
	disk.SetSourceArchive(archiveID) //ソースアーカイブはIDだけ指定する

	res, err := diskAPI.Create(disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	diskID := res.ID

	//wait
	err = diskAPI.WaitForAvailable(diskID, 5*time.Minute) //日によって時間がかかることもあるため5分待つ
	assert.NoError(t, err)                                //timeoutしたらerrに値が格納されている

	createdDisk, err := diskAPI.Read(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdDisk)
	assert.Equal(t, createdDisk.SourceArchive.ID, archiveID)

	//DELETE
	disk, err = diskAPI.Delete(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupTestDisk)
	testTearDownHandlers = append(testTearDownHandlers, cleanupTestDisk)
}

func cleanupTestDisk() {
	diskAPI := client.Disk
	req := &sakura.Request{}
	req.AddFilter("Name", testDiskName)
	res, err := diskAPI.Find(req)
	if err == nil && res.Count > 0 {
		diskAPI.Delete(res.Disks[0].ID)
	}
}
