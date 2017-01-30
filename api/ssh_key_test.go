package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testSSHKeyName = "libsacloud_test_SSHKey"
const testPublicKey = `sh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFgFvUj3DrQyktz434X76N9IuOqYmWp3ffxcEb7Jzyg1GvfbzpcQDV9H0rIfGMXUhKkTYygLWeDOTGk1fd935lBdMUMv1lhtX9gPMZcyu945c313rpgnD/PrLVSoBGlpRVx29tA6t1x4b+LaVek4mQL2AojeRQdz8W3gF4dKdGi+Ci2ogV/dZkVsQuZRjLy09iixGB+vjF1tgnZQJqIz8CvFx8ULvcCUAzhRF8osALdSPPEBsAfaD3y5xXWHYnb+OFL3EZ1jb4rM6KB/LfaARFFrBk6rhqEjUYZgmAecMu79cY9Gc+6MhjONbdxT0gOhmZQK7kg/kwBU8prpJGLFGp ubuntu@sakura-dev`

func TestSSHKeyCRUD(t *testing.T) {
	api := client.SSHKey

	//CREATE
	newItem := api.New()
	newItem.Name = testSSHKeyName
	newItem.Description = "before"
	newItem.PublicKey = testPublicKey

	item, err := api.Create(newItem)

	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	id := item.ID

	//READ
	item, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, item)

	assert.NotEmpty(t, item.Fingerprint)
	assert.Equal(t, testPublicKey, item.PublicKey)

	//UPDATE
	item.Description = "after"
	item, err = api.Update(id, item)

	assert.NoError(t, err)
	assert.NotEqual(t, item.Description, "before")

	//Delete
	_, err = api.Delete(id)
	assert.NoError(t, err)
}

func TestSSHKeyGenerate(t *testing.T) {
	api := client.SSHKey

	// generate
	item, err := api.Generate(testSSHKeyName, "", "description")
	assert.NoError(t, err)
	assert.NotNil(t, item)

	// should have SSHKey properties + PrivateKey
	assert.NotEmpty(t, item.Name)
	assert.NotEmpty(t, item.PublicKey)
	assert.NotEmpty(t, item.Description)
	assert.NotEmpty(t, item.PrivateKey)

	//Delete
	_, err = api.Delete(item.ID)
	assert.NoError(t, err)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupSSHKey)
	testTearDownHandlers = append(testTearDownHandlers, cleanupSSHKey)
}

func cleanupSSHKey() {
	items, _ := client.SSHKey.Reset().WithNameLike(testSSHKeyName).Find()
	for _, item := range items.SSHKeys {
		client.SSHKey.Delete(item.ID)
	}
}
