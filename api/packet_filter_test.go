package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPacketFilterName            = "libsacloud_test_packet_filter_name"
	testPacketFilterDesciprionAfter = "lisacloud_packetfilter_description_after"
)

func TestCRUDByPacketFilterAPI(t *testing.T) {
	defer initPacketFilter()()

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

func initPacketFilter() func() {
	cleanupPacketFilter()
	return cleanupPacketFilter
}

func cleanupPacketFilter() {
	api := client.PacketFilter
	res, _ := api.withNameLike(testPacketFilterName).Find()
	if res.Count > 0 {
		api.Delete(res.PacketFilters[0].ID)
	}
}
