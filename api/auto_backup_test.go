package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testAutoBackupName = "test_libsakuracloud_ab"

func TestAutoBackupCRUD(t *testing.T) {

	defer initAutoBackup()()

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1b"
	api := client.AutoBackup

	disk := client.Disk.New()
	disk.Name = testAutoBackupName
	disk, err := client.Disk.Create(disk)

	assert.NoError(t, err)
	assert.NotEmpty(t, disk)

	//CREATE
	ab := api.New(testAutoBackupName, disk.ID)

	ab.Description = "before"
	ab.SetBackupMaximumNumberOfArchives(2)
	ab.SetBackupSpanWeekdays([]string{"mon", "tue", "wed"})

	item, err := client.AutoBackup.Create(ab)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.Name, testAutoBackupName)
	assert.Equal(t, item.Description, "before")
	assert.Equal(t, item.Settings.Autobackup.MaximumNumberOfArchives, 2)
	assert.Equal(t, item.Settings.Autobackup.BackupSpanWeekdays, []string{"mon", "tue", "wed"})

	id := item.ID

	//READ
	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	//UPDATE
	item.Description = "after"
	item.SetBackupMaximumNumberOfArchives(3)
	item.SetBackupSpanWeekdays([]string{"mon", "tue", "sat", "sun"})

	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")
	assert.Equal(t, item.Settings.Autobackup.MaximumNumberOfArchives, 3)
	assert.Equal(t, item.Settings.Autobackup.BackupSpanWeekdays, []string{"mon", "tue", "sat", "sun"})

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func initAutoBackup() func() {
	cleanupAutoBackup()
	return cleanupArchive
}

func cleanupAutoBackup() {
	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1b"

	items, _ := client.AutoBackup.Reset().WithNameLike(testAutoBackupName).Find()
	if items.CommonServiceAutoBackupItems != nil {
		for _, item := range items.CommonServiceAutoBackupItems {
			client.AutoBackup.Delete(item.ID)
		}
	}

	disks, _ := client.Disk.Reset().WithNameLike(testAutoBackupName).Find()
	for _, disk := range disks.Disks {
		client.Disk.Delete(disk.ID)
	}
}
