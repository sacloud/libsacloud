package api

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	testSIMName        = "libsacloud-test-sim"
	testSIMEnvICCID    = "LIBSACLOUD_TEST_ICCID"
	testSIMEnvPASSCODE = "LIBSACLOUD_TEST_PASSCODE"
)

func TestSIMBasicCRUD(t *testing.T) {

	iccID := os.Getenv(testSIMEnvICCID)
	passcode := os.Getenv(testSIMEnvPASSCODE)

	if iccID == "" || passcode == "" {
		t.Skipf("%s and %s is required. skip", testSIMEnvICCID, testSIMEnvPASSCODE)
	}

	defer initSIM()()

	// create
	api := client.SIM
	req := api.New(testSIMName, iccID, passcode)

	sim, err := api.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, sim)

	// update
	sim.Description = "description_updated"
	updSIM, err := api.Update(sim.ID, sim)
	assert.NoError(t, err)
	assert.NotNil(t, updSIM)

	// read
	readSIM, err := api.Read(sim.ID)
	assert.NoError(t, err)
	assert.NotNil(t, readSIM)
	assert.Equal(t, "description_updated", readSIM.Description)

	// delete
	_, err = api.Delete(sim.ID)
	assert.NoError(t, err)

}

func initSIM() func() {
	cleanupSIM()
	return cleanupSIM
}

func cleanupSIM() {
	items, _ := client.SIM.Reset().WithNameLike(testSIMName).Find()

	for _, item := range items.CommonServiceSIMItems {
		client.DNS.Delete(item.ID)
	}
}
