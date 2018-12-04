package api

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
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

	// NetworkOperator
	_, err = api.SetNetworkOperator(sim.ID, &sacloud.SIMNetworkOperatorConfig{
		Name:  sacloud.SIMOperatorsKDDI,
		Allow: true,
	})
	assert.NoError(t, err)

	operators, err := api.GetNetworkOperator(sim.ID)
	assert.NoError(t, err)
	assert.True(t, len(operators.NetworkOperatorConfigs) > 1) // response should have all career's setting
	for _, career := range operators.NetworkOperatorConfigs {
		expect := career.Name == sacloud.SIMOperatorsKDDI
		assert.Equal(t, expect, career.Allow, "career:%s has unexpected value: %v", career.Name, career.Allow)
	}

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
