package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testPacketFilterName            = "libsacloud_test_packet_filter_name"
	testPacketFilterDesciprionAfter = "lisacloud_packetfilter_description_after"
)

func TestCRUDByPacketFilterAPI(t *testing.T) {
	api := client.PacketFilter

	//CREATE
	packetFilter := api.New()
	packetFilter.Name = testPacketFilterName
	packetFilter.Description = "aaaaaaa"
	packetFilter.AddTCPRule("", "", "", "", false)
	packetFilter.AddUDPRule("", "", "", "", false)

	res, err := api.Create(packetFilter)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	//for READ
	var id = res.ID

	//READ
	res, err = api.Read(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Description)

	assert.Equal(t, len(res.Expression), 2)

	//UPDATE
	packetFilter.Description = testPacketFilterDesciprionAfter

	res, err = api.Update(id, packetFilter)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res.Description)
	assert.Equal(t, res.Description, testPacketFilterDesciprionAfter)

	//DELETE
	res, err = api.Delete(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func init() {
	testSetupHandlers = append(testSetupHandlers, cleanupTestPacketFilter)

	testTearDownHandlers = append(testTearDownHandlers, cleanupTestPacketFilter)
}

func cleanupTestPacketFilter() {
	api := client.PacketFilter
	res, _ := api.withNameLike(testPacketFilterName).Find()
	if res.Count > 0 {
		api.Delete(res.PacketFilters[0].ID)
	}
}
