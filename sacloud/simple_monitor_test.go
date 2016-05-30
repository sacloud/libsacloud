package sacloud

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testSimpleMonitorJSONTemplate = `
	{
		"ID": "123456789012",
		"Name": "133.242.224.255",
		"Description": "sakura-dev",
		"Settings": {
		    "SimpleMonitor": {
			"DelayLoop": 3600,
			"HealthCheck": %s,
			"Enabled": "True",
			"NotifyEmail": {
			    "Enabled": "True"
			},
			"NotifySlack": {
			    "Enabled": "True",
			    "IncomingWebhooksURL": "https:\/\/hooks.slack.com\/services\/xxxxxxx"
			}
		    }
		},
		"Status": {
		    "Target": "133.242.224.255"
		},
		"ServiceClass": "cloud\/simplemon\/free",
		"CreatedAt": "2016-05-02T14:14:14+09:00",
		"ModifiedAt": "2016-05-02T14:14:14+09:00",
		"Provider": {
		    "ID": 3000001,
		    "Class": "simplemon",
		    "Name": "simplemon01",
		    "ServiceClass": "cloud\/simplemon"
		},
		"Icon": null,
		"Tags": [
		]
	    }
	`

	testPingMonitoringJSON = `
	{
		"Protocol": "ping"
	}`

	testTCPMonitoringJSON = `
	{
		"Protocol": "tcp",
		"Port": "22"
	}`

	testHTTPMonitoringJSON = `
	{
		"Protocol": "http",
		"Path": "\/index.html",
		"Status": "200",
		"Port": "80"
	}`

	testDNSMonitoringJSON = `
	{
		"Protocol": "dns",
		"QName": "www.example.com",
		"ExpectedData": "1.2.3.4"
	}`

	testSSHMonitoringJSON = `
	{
		"Protocol": "ssh",
		"Port": "22"
	}`
	testSMTPMonitoringJSON = `
	{
		"Protocol": "smtp",
		"Port": "25"
	}`
	testPOP3MonitoringJSON = `
	{
		"Protocol": "pop3",
		"Port": "110"
	}`

	testSNMPMonitoringJSON = `
	{
		"Protocol":"snmp",
		"Community":"gggg",
		"SNMPVersion":"2c",
		"OID":".1.3.6.1.2.1.1.5.0",
		"ExpectedData":"12"
	}
	`
)

func TestMarshalSimpleMonitorJSON(t *testing.T) {
	// ping
	var simpleMonitor SimpleMonitor
	err := json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testPingMonitoringJSON)), &simpleMonitor)

	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor)
	assert.NotEmpty(t, simpleMonitor.ID)
	assert.NotEmpty(t, simpleMonitor.Status.Target)
	assert.NotEmpty(t, simpleMonitor.Provider.Class)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)

	//tcp
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testTCPMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Port)

	//http
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testHTTPMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Path)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Port)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Status)

	//dns
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testDNSMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.QName)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.ExpectedData)

	//ssh
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testSSHMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Port)

	//smtp
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testSMTPMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Port)

	//pop3
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testPOP3MonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Port)

	//snmp
	err = json.Unmarshal([]byte(fmt.Sprintf(testSimpleMonitorJSONTemplate, testSNMPMonitoringJSON)), &simpleMonitor)
	assert.NoError(t, err)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Protocol)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.Community)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.SNMPVersion)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.OID)
	assert.NotEmpty(t, simpleMonitor.Settings.SimpleMonitor.HealthCheck.ExpectedData)

}
