package naked

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	testMonitorValuesCPUTimeJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "CPU-TIME": 0.5
        },
        "1970-01-01T00:00:02Z": {
            "CPU-TIME": 0.1
        },
        "1970-01-01T00:00:03Z": {
            "CPU-TIME": 0
        },
        "1970-01-01T00:00:04Z": {
            "CPU-TIME": null
        }
    }`

	testMonitorValuesDiskJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "Read": 10.0, 
            "Write": 20.0 
        },
        "1970-01-01T00:00:02Z": {
            "Read": 11.0, 
            "Write": 21.0 
        },
        "1970-01-01T00:00:03Z": {
            "Read": null, 
            "Write": null 
        },
        "1970-01-01T00:00:04Z": {
            "Read": 12.0, 
            "Write": 22.0 
        }
    }`

	testMonitorValuesInterfaceJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "Send": 10.0, 
            "Receive": 20.0 
        },
        "1970-01-01T00:00:02Z": {
            "Send": 11.0, 
            "Receive": 21.0 
        },
        "1970-01-01T00:00:03Z": {
            "Send": null, 
            "Receive": null 
        },
        "1970-01-01T00:00:04Z": {
            "Send": 12.0, 
            "Receive": 22.0 
        }
    }`

	testMonitorValuesRouterJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "In": 10.0, 
            "Out": 20.0 
        },
        "1970-01-01T00:00:02Z": {
            "In": 11.0, 
            "Out": 21.0 
        },
        "1970-01-01T00:00:03Z": {
            "In": null, 
            "Out": null 
        },
        "1970-01-01T00:00:04Z": {
            "In": 12.0, 
            "Out": 22.0 
        }
    }`

	testMonitorValuesDatabaseJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "Total-Memory-Size"  : 10.0, 
            "Used-Memory-Size"   : 20.0,
            "Total-Disk1-Size"   : 30.0,
            "Used-Disk1-Size"    : 40.0,
            "Total-Disk2-Size"   : 50.0,
            "Used-Disk2-Size"    : 60.0,
            "binlogUsedSizeKiB": 70.0,
            "delayTimeSec"     : 80.0
        },
        "1970-01-01T00:00:02Z": {
            "Total-Memory-Size"  : 11.0, 
            "Used-Memory-Size"   : 21.0,
            "Total-Disk1-Size"   : 31.0,
            "Used-Disk1-Size"    : 41.0,
            "Total-Disk2-Size"   : 51.0,
            "Used-Disk2-Size"    : 61.0,
            "binlogUsedSizeKiB": 71.0,
            "delayTimeSec"     : 81.0
        },
        "1970-01-01T00:00:03Z": {
            "Total-Memory-Size"  : null, 
            "Used-Memory-Size"   : null,
            "Total-Disk1-Size"   : null,
            "Used-Disk1-Size"    : null,
            "Total-Disk2-Size"   : null,
            "Used-Disk2-Size"    : null,
            "binlogUsedSizeKiB": null,
            "delayTimeSec"     : null 
        },
        "1970-01-01T00:00:04Z": {
            "Total-Memory-Size"  : 12.0, 
            "Used-Memory-Size"   : 22.0,
            "Total-Disk1-Size"   : 32.0,
            "Used-Disk1-Size"    : 42.0,
            "Total-Disk2-Size"   : 52.0,
            "Used-Disk2-Size"    : 62.0,
            "binlogUsedSizeKiB": 72.0,
            "delayTimeSec"     : 82.0
        }
    }`

	testMonitorValuesFreeDiskSizeJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "Free-Disk-Size": 0.5
        },
        "1970-01-01T00:00:02Z": {
            "Free-Disk-Size": 0.1
        },
        "1970-01-01T00:00:03Z": {
            "Free-Disk-Size": 0
        },
        "1970-01-01T00:00:04Z": {
            "Free-Disk-Size": null
        }
    }`

	testMonitorValuesResponseTimeSecJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "responsetimesec": 0.5
        },
        "1970-01-01T00:00:02Z": {
            "responsetimesec": 0.1
        },
        "1970-01-01T00:00:03Z": {
            "responsetimesec": 0
        },
        "1970-01-01T00:00:04Z": {
            "responsetimesec": null
        }
    }`

	testMonitorValuesLinkJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "UplinkBps": 10.0, 
            "DownlinkBps": 20.0 
        },
        "1970-01-01T00:00:02Z": {
            "UplinkBps": 11.0, 
            "DownlinkBps": 21.0 
        },
        "1970-01-01T00:00:03Z": {
            "UplinkBps": null, 
            "DownlinkBps": null 
        },
        "1970-01-01T00:00:04Z": {
            "UplinkBps": 12.0, 
            "DownlinkBps": 22.0 
        }
    }`

	testMonitorValuesConnexctionJSON = `
    {
        "1970-01-01T00:00:01Z": {
            "activeConnections": 10.0, 
            "connectionsPerSec": 20.0 
        },
        "1970-01-01T00:00:02Z": {
            "activeConnections": 11.0, 
            "connectionsPerSec": 21.0 
        },
        "1970-01-01T00:00:03Z": {
            "activeConnections": null, 
            "connectionsPerSec": null 
        },
        "1970-01-01T00:00:04Z": {
            "activeConnections": 12.0, 
            "connectionsPerSec": 22.0 
        }
    }`
)

func TestMonitorValues_UnmarshalJSON(t *testing.T) {

	expects := []struct {
		input  string
		expect MonitorValues
	}{
		{
			input: testMonitorValuesCPUTimeJSON,
			expect: MonitorValues{
				CPU: MonitorCPUTimeValues{
					{
						Time:    time.Unix(1, 0).UTC(),
						CPUTime: float64(0.5),
					},
					{
						Time:    time.Unix(2, 0).UTC(),
						CPUTime: float64(0.1),
					},
					{
						Time:    time.Unix(3, 0).UTC(),
						CPUTime: float64(0),
					},
				},
			},
		},
		{
			input: testMonitorValuesDiskJSON,
			expect: MonitorValues{
				Disk: MonitorDiskValues{
					{
						Time:  time.Unix(1, 0).UTC(),
						Read:  float64(10.0),
						Write: float64(20.0),
					},
					{
						Time:  time.Unix(2, 0).UTC(),
						Read:  float64(11.0),
						Write: float64(21.0),
					},
					{
						Time:  time.Unix(4, 0).UTC(),
						Read:  float64(12.0),
						Write: float64(22.0),
					},
				},
			},
		},
		{
			input: testMonitorValuesInterfaceJSON,
			expect: MonitorValues{
				Interface: MonitorInterfaceValues{
					{
						Time:    time.Unix(1, 0).UTC(),
						Send:    float64(10.0),
						Receive: float64(20.0),
					},
					{
						Time:    time.Unix(2, 0).UTC(),
						Send:    float64(11.0),
						Receive: float64(21.0),
					},
					{
						Time:    time.Unix(4, 0).UTC(),
						Send:    float64(12.0),
						Receive: float64(22.0),
					},
				},
			},
		},
		{
			input: testMonitorValuesRouterJSON,
			expect: MonitorValues{
				Router: MonitorRouterValues{
					{
						Time: time.Unix(1, 0).UTC(),
						In:   float64(10.0),
						Out:  float64(20.0),
					},
					{
						Time: time.Unix(2, 0).UTC(),
						In:   float64(11.0),
						Out:  float64(21.0),
					},
					{
						Time: time.Unix(4, 0).UTC(),
						In:   float64(12.0),
						Out:  float64(22.0),
					},
				},
			},
		},
		{
			input: testMonitorValuesDatabaseJSON,
			expect: MonitorValues{
				Database: MonitorDatabaseValues{
					{
						Time:              time.Unix(1, 0).UTC(),
						TotalMemorySize:   float64(10.0),
						UsedMemorySize:    float64(20.0),
						TotalDisk1Size:    float64(30.0),
						UsedDisk1Size:     float64(40.0),
						TotalDisk2Size:    float64(50.0),
						UsedDisk2Size:     float64(60.0),
						BinlogUsedSizeKiB: float64(70.0),
						DelayTimeSec:      float64(80.0),
					},
					{
						Time:              time.Unix(2, 0).UTC(),
						TotalMemorySize:   float64(11.0),
						UsedMemorySize:    float64(21.0),
						TotalDisk1Size:    float64(31.0),
						UsedDisk1Size:     float64(41.0),
						TotalDisk2Size:    float64(51.0),
						UsedDisk2Size:     float64(61.0),
						BinlogUsedSizeKiB: float64(71.0),
						DelayTimeSec:      float64(81.0),
					},
					{
						Time:              time.Unix(4, 0).UTC(),
						TotalMemorySize:   float64(12.0),
						UsedMemorySize:    float64(22.0),
						TotalDisk1Size:    float64(32.0),
						UsedDisk1Size:     float64(42.0),
						TotalDisk2Size:    float64(52.0),
						UsedDisk2Size:     float64(62.0),
						BinlogUsedSizeKiB: float64(72.0),
						DelayTimeSec:      float64(82.0),
					},
				},
			},
		},
		{
			input: testMonitorValuesFreeDiskSizeJSON,
			expect: MonitorValues{
				FreeDiskSize: MonitorFreeDiskSizeValues{
					{
						Time:         time.Unix(1, 0).UTC(),
						FreeDiskSize: float64(0.5),
					},
					{
						Time:         time.Unix(2, 0).UTC(),
						FreeDiskSize: float64(0.1),
					},
					{
						Time:         time.Unix(3, 0).UTC(),
						FreeDiskSize: float64(0),
					},
				},
			},
		},
		{
			input: testMonitorValuesResponseTimeSecJSON,
			expect: MonitorValues{
				ResponseTimeSec: MonitorResponseTimeSecValues{
					{
						Time:            time.Unix(1, 0).UTC(),
						ResponseTimeSec: float64(0.5),
					},
					{
						Time:            time.Unix(2, 0).UTC(),
						ResponseTimeSec: float64(0.1),
					},
					{
						Time:            time.Unix(3, 0).UTC(),
						ResponseTimeSec: float64(0),
					},
				},
			},
		},
		{
			input: testMonitorValuesLinkJSON,
			expect: MonitorValues{
				Link: MonitorLinkValues{
					{
						Time:        time.Unix(1, 0).UTC(),
						UplinkBPS:   float64(10.0),
						DownlinkBPS: float64(20.0),
					},
					{
						Time:        time.Unix(2, 0).UTC(),
						UplinkBPS:   float64(11.0),
						DownlinkBPS: float64(21.0),
					},
					{
						Time:        time.Unix(4, 0).UTC(),
						UplinkBPS:   float64(12.0),
						DownlinkBPS: float64(22.0),
					},
				},
			},
		},
		{
			input: testMonitorValuesConnexctionJSON,
			expect: MonitorValues{
				Connection: MonitorConnectionValues{
					{
						Time:              time.Unix(1, 0).UTC(),
						ActiveConnections: float64(10.0),
						ConnectionsPerSec: float64(20.0),
					},
					{
						Time:              time.Unix(2, 0).UTC(),
						ActiveConnections: float64(11.0),
						ConnectionsPerSec: float64(21.0),
					},
					{
						Time:              time.Unix(4, 0).UTC(),
						ActiveConnections: float64(12.0),
						ConnectionsPerSec: float64(22.0),
					},
				},
			},
		},
	}

	for _, tc := range expects {
		dest := MonitorValues{}
		err := json.Unmarshal([]byte(tc.input), &dest)
		require.NoError(t, err)
		require.Equal(t, tc.expect, dest)
	}

}
