package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testArchiveName = "libsacloud_test_archive"

func TestGetUbuntuArchiveID(t *testing.T) {
	archiveAPI := client.Archive
	id, err := archiveAPI.GetUbuntuArchiveID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	t.Logf("ubuntu archive ID : %s", id)
}

func TestFindState(t *testing.T) {
	api := client.Archive

	api.Reset().WithNameLike("hoge").FilterBy("Fuga", "fuga").Limit(10).Offset(1).Include("inc").Exclude("enc")

	state := api.state

	assert.NotEmpty(t, state)
	assert.Equal(t, state.Filter["Name"], "hoge")
	assert.Equal(t, state.Filter["Fuga"], "fuga")
	assert.Equal(t, state.Count, 10)
	assert.Equal(t, state.From, 1)
	assert.Equal(t, state.Include[0], "inc")
	assert.Equal(t, state.Exclude[0], "enc")

	//clear state
	api.Reset()
	state = api.state
	assert.Empty(t, state.Filter)
	assert.Empty(t, state.Count)
	assert.Empty(t, state.From)
	assert.Empty(t, state.Include)
	assert.Empty(t, state.Exclude)

	res, err := api.withNameLike("CentOS").limit(1).Find()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, res.Count, 1)
	assert.Contains(t, res.Archives[0].Name, "CentOS")
}

func TestCRUDAndFTP(t *testing.T) {
	api := client.Archive

	//CREATE
	newArchive := api.New()
	newArchive.Name = testArchiveName
	newArchive.Description = "hoge"
	newArchive.SizeMB = 20480

	archive, err := api.Create(newArchive)

	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	archiveID := archive.ID

	//READ
	archive, err = api.Read(archiveID)
	assert.NoError(t, err)
	assert.NotEmpty(t, archive)

	//Open
	ftpServer, err := api.OpenFTP(archive.ID, false)
	assert.NoError(t, err)
	assert.NotEmpty(t, ftpServer.Password)

	password := ftpServer.Password

	//Close
	res, err := api.CloseFTP(archiveID)
	assert.NoError(t, err)
	assert.True(t, res)

	//Re-Open(password not changed)
	//ftpServer, err = api.OpenFTP(archive.ID, false)
	//assert.NoError(t, err)
	//assert.Equal(t, ftpServer.Password, password)
	//
	////Close
	//api.CloseFTP(archiveID)

	//Re-Open(will password change)
	ftpServer, err = api.OpenFTP(archive.ID, false)
	assert.NoError(t, err)
	assert.NotEqual(t, ftpServer.Password, password)

	//Delete
	_, err = api.Delete(archiveID)
	assert.NoError(t, err)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupArchive)
	testTearDownHandlers = append(testTearDownHandlers, cleanupArchive)
}

func cleanupArchive() {
	items, _ := client.Archive.Reset().WithNameLike(testArchiveName).Find()
	for _, item := range items.Archives {
		client.Archive.Delete(item.ID)
	}
}
