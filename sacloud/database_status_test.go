// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sacloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var databaseStatusJSON = `
    {
      "Status": "up",
      "DBConf": {
        "version": {
          "lastmodified": "2017-06-28 14:02:34 +0900",
          "commithash": "ea15c88152d0c37df9ffe9ffda45500024f1f1a0",
          "status": "latest",
          "tag": "1.0",
          "expire": "2017-08-01 15:00:00 +0900"
        },
        "MariaDB": {
          "status": "running"
        },
        "log": [
           {
            "name": "systemctl",
	    "data": "log1\nlog1-2\nlog1-3"
	   },
           {
             "name": "mariadb.log",
	    "data": "log2"
           }
	],
	"backup": {
          "history": [
            {
              "createdat": "2017-07-25T11:50:58+09:00",
              "availability": "discontinued",
              "recoveredat": "0000-00-00T00:00:00+09:00",
              "size": "3661539"
            },
            {
              "createdat": "2017-10-10T10:10:10+09:00",
              "availability": "discontinued",
              "recoveredat": "2017-11-11T11:11:11+09:00",
              "size": "3661539"
            }
          ]
        }
      },
      "is_ok": true
    }
`

func TestMarshalDatabaseStatusJSON(t *testing.T) {
	// ping
	var status DatabaseStatus
	err := json.Unmarshal([]byte(databaseStatusJSON), &status)

	assert.NoError(t, err)
	assert.EqualValues(t, "up", status.Status)
	assert.NotNil(t, status.DBConf)

	// version
	assert.NotNil(t, status.DBConf.Version)
	versionInfo := status.DBConf.Version
	assert.NotEmpty(t, versionInfo.LastModified)
	assert.NotEmpty(t, versionInfo.CommitHash)
	assert.NotEmpty(t, versionInfo.Status)
	assert.NotEmpty(t, versionInfo.Tag)
	assert.NotEmpty(t, versionInfo.Expire)

	// log
	assert.Len(t, status.DBConf.Log, 2)

	// for systemctl log
	syslog := status.DBConf.Log[0]
	assert.NotEmpty(t, syslog.Name)
	assert.NotEmpty(t, syslog.Data)
	assert.True(t, syslog.IsSystemdLog())
	assert.Len(t, syslog.Logs(), 3)

	// for normal log
	normalLog := status.DBConf.Log[1]
	assert.NotEmpty(t, normalLog.Name)
	assert.NotEmpty(t, normalLog.Data)
	assert.False(t, normalLog.IsSystemdLog())
	assert.Len(t, normalLog.Logs(), 1)

	// backup history
	assert.NotNil(t, status.DBConf.Backup)
	assert.Len(t, status.DBConf.Backup.History, 2)

	// have RecoveredAt
	// format times
	layout := "20060102150405" // yyyyMMddHHmmss

	h0 := status.DBConf.Backup.History[0]
	assert.NotEmpty(t, h0.CreatedAt)
	assert.NotEmpty(t, h0.Availability)
	assert.NotEmpty(t, h0.Size)
	assert.Nil(t, h0.RecoveredAt)
	assert.EqualValues(t, "", h0.FormatRecoveredAt(layout))

	// have not RecoveredAt
	h1 := status.DBConf.Backup.History[1]
	assert.NotNil(t, h1.RecoveredAt)
	assert.EqualValues(t, "20171010101010", h1.FormatCreatedAt(layout))
	assert.EqualValues(t, "20171111111111", h1.FormatRecoveredAt(layout))
}
