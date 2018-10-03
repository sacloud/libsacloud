package api

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testDiskName = "libsacloud-test-disk-name"

func TestCRUDByDiskAPI(t *testing.T) {
	defer initDisk()()

	diskAPI := client.Disk
	//CREATE : empty disk
	disk := diskAPI.New()
	disk.Name = testDiskName

	// HACK 現状ではディスクの存在チェックが行われていないため、ここでテスト可能。
	// 今後仕様変更などの際は切り出してテストする
	disk.DistantFrom = []int64{111111111111}

	res, err := diskAPI.Create(disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.ID)

	diskID := res.ID

	//wait
	err = diskAPI.SleepWhileCopying(diskID, 5*time.Minute) //日によって時間がかかることもあるため5分待つ
	assert.NoError(t, err)                                 //timeoutしたらerrに値が格納されている

	//READ
	disk, err = diskAPI.Read(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.NotEmpty(t, disk.ID)
	assert.True(t, disk.IsAvailable())
	assert.Equal(t, disk.Connection, sacloud.DiskConnectionVirtio)

	//UPDATE
	disk.SetDiskConnection(sacloud.DiskConnectionIDE)

	disk, err = diskAPI.Update(diskID, disk)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
	assert.Equal(t, disk.Connection, sacloud.DiskConnectionIDE)

	//DELETE
	disk, err = diskAPI.Delete(diskID)
	assert.NoError(t, err)
	assert.NotEmpty(t, disk)
}

func TestCreateDiskFromSource(t *testing.T) {
	defer initDisk()()

	diskAPI := client.Disk

	archive, err := client.Archive.FindLatestStableCentOS()
	assert.NoError(t, err)
	archiveID := archive.ID

	//CREATE : empty disk
	disk := diskAPI.New()
	disk.Name = testDiskName
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

func TestCreateDiskWithConfig(t *testing.T) {
	defer initDisk()()

	diskAPI := client.Disk

	// server
	newServer := client.Server.New()
	newServer.Name = testDiskName
	newServer.SetServerPlanByValue(1, 1, sacloud.PlanDefault)
	newServer.AddPublicNWConnectedParam() //公開セグメントに接続
	newServer.WaitDiskMigration = true

	server, err := client.Server.Create(newServer)
	assert.NoError(t, err)
	serverID := server.ID

	archive, err := client.Archive.FindLatestStableCentOS()
	assert.NoError(t, err)
	archiveID := archive.ID

	//CREATE : empty disk
	disk := diskAPI.New()
	disk.Name = testDiskName
	disk.SetServerID(serverID)       // Server.IDをディスクに指定しておくことでディスク作成完了後にサーバを起動してくれる
	disk.SetSourceArchive(archiveID) //ソースアーカイブはIDだけ指定する

	config := diskAPI.NewCondig()
	config.SetPassword("p@ssw0rd")
	config.SetHostName(testDiskName)

	res, err := diskAPI.CreateWithConfig(disk, config, true)
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

	err = client.Server.SleepUntilUp(serverID, 5*time.Minute)
	assert.NoError(t, err)

	createdServer, err := client.Server.Read(serverID)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdServer)
	assert.True(t, createdServer.IsAvailable())
	assert.True(t, createdServer.IsUp())

	//DELETE
	_, err = client.Server.Stop(serverID)
	assert.NoError(t, err)

	err = client.Server.SleepUntilDown(serverID, client.DefaultTimeoutDuration)
	assert.NoError(t, err)

	_, err = client.Server.DeleteWithDisk(serverID, []int64{diskID})
	assert.NoError(t, err)
}

//func TestCanEditDisk(t *testing.T) {
//	api := client.Disk
//	client.Zone = "is1b"
//
//	//// CentOS
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

func TestDiskAPI_FindByFilters(t *testing.T) {
	defer initDisk()()

	api := client.Disk

	ids := []int64{}
	name1 := fmt.Sprintf("libsacloud_test_disk_name%d", 1)
	name2 := fmt.Sprintf("libsacloud_test_disk_name%d", 2)
	name3 := fmt.Sprintf("libsacloud_test_disk_name%d", 3)

	names := []string{name1, name2, name3}
	for _, name := range names {
		disk := api.New()
		disk.Name = name
		d, err := api.Create(disk)
		if !assert.NoError(t, err) {
			return
		}
		ids = append(ids, d.ID)

	}

	res, err := api.Reset().Include("ID").Include("Name").
		FilterBy("Name", "libsacloud_test_disk_name").FilterBy("Name", "ssssss").Find()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res.Disks) >= 3)

	res, err = api.Reset().Include("ID").Include("Name").
		FilterMultiBy("Name", name1).FilterMultiBy("Name", name2).Find()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Disks), 2)

	for _, id := range ids {
		api.Delete(id)
	}

}

func initDisk() func() {
	cleanupDisk()
	return cleanupDisk
}

func cleanupDisk() {
	diskAPI := client.Disk
	res, err := diskAPI.withNameLike(testDiskName).Find()
	if err == nil && res.Count > 0 {
		for _, disk := range res.Disks {
			diskAPI.Delete(disk.ID)
		}
	}

	items, _ := client.Server.Reset().WithNameLike(testDiskName).Find()
	for _, item := range items.Servers {
		client.Server.Delete(item.ID)
	}
}
