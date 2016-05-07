package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testResourceMonitorJSON = `
    {
        "2016-05-07T12:15:00+09:00": {
            "CPU-TIME": 0.5
        },
        "2016-05-07T12:20:00+09:00": {
            "CPU-TIME": 0.1
        },
        "2016-05-07T12:25:00+09:00": {
            "CPU-TIME": 0
        },
        "2016-05-07T12:30:00+09:00": {
            "CPU-TIME": null
        }
    }
`

var testResourceMonitorResponseJSON = `
{
    "Data": ` + testResourceMonitorJSON + `,
    "is_ok" : true
}
`

func TestMarshalResourceMonitorJSON(t *testing.T) {
	var m MonitorValues
	err := json.Unmarshal([]byte(testResourceMonitorJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
}

func TestMarshalCPUResourceMonitorJSON(t *testing.T) {
	var m ResourceMonitorResponse
	err := json.Unmarshal([]byte(testResourceMonitorResponseJSON), &m)

	assert.NoError(t, err)
	assert.NotEmpty(t, m)
	assert.NotEmpty(t, m.Data)

}

func TestResourceMonitorCalc(t *testing.T) {
	var monitor MonitorValues
	json.Unmarshal([]byte(testResourceMonitorJSON), &monitor)

	var sum = 0.6
	var count float64 = 3
	var max = 0.5
	var min float64

	calcResult := monitor.Calc()
	assert.NotNil(t, calcResult)
	assert.NotNil(t, calcResult.CPU)
	assert.Equal(t, calcResult.CPU.Avg, sum/count)
	assert.Equal(t, calcResult.CPU.Count, count)
	assert.Equal(t, calcResult.CPU.Max, max)
	assert.Equal(t, calcResult.CPU.Min, min)

}
