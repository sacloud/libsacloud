package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNoteJSON = `
{
	"ID": "112800309940",
	"Name": "lb-dsr",
	"Class": "shell",
	"Scope": "shared",
	"Content": "#!\/bin\/bash\n# @sacloud-once\n# @sacloud-desc-begin\n# \u30ed\u30fc\u30c9\u30d0\u30e9\u30f3\u30b9\u5bfe\u8c61\u306e\u30b5\u30fc\u30d0\u306e\u8a2d\u5b9a\u306e\u305f\u3081\u306e\u30b9\u30af\u30ea\u30d7\u30c8\u3067\u3059\u3002\n# \u3053\u306e\u30b9\u30af\u30ea\u30d7\u30c8\u306f\u3001CentOS6.X\u3082\u3057\u304f\u306fScientific Linux6.X\u3067\u306e\u307f\u52d5\u4f5c\u3057\u307e\u3059\u3002\n# @sacloud-desc-end\n# @sacloud-require-archive distro-centos distro-ver-6.*\n# @sacloud-require-archive distro-sl distro-ver-6.*\n# @sacloud-text required shellarg maxlen=20 para1 \u0022\u30ed\u30fc\u30c9\u30d0\u30e9\u30f3\u30b5\u30fc\u306eVIP\u0022\n\nPARA1=@@@para1@@@\nPARA2=\u0022net.ipv4.conf.all.arp_ignore = 1\u0022\nPARA3=\u0022net.ipv4.conf.all.arp_announce = 2\u0022\nPARA4=\u0022DEVICE=lo:0\u0022\nPARA5=\u0022IPADDR=\u0022$PARA1\nPARA6=\u0022NETMASK=255.255.255.255\u0022\n\ncp --backup \/etc\/sysctl.conf \/tmp\/ || exit 1\n\necho $PARA2 \u003E\u003E \/etc\/sysctl.conf\necho $PARA3 \u003E\u003E \/etc\/sysctl.conf\nsysctl -p 1\u003E\/dev\/null\n\ncp --backup \/etc\/sysconfig\/network-scripts\/ifcfg-lo:0 \/tmp\/ 2\u003E\/dev\/null\n\ntouch \/etc\/sysconfig\/network-scripts\/ifcfg-lo:0\necho $PARA4 \u003E \/etc\/sysconfig\/network-scripts\/ifcfg-lo:0\necho $PARA5 \u003E\u003E \/etc\/sysconfig\/network-scripts\/ifcfg-lo:0\necho $PARA6 \u003E\u003E \/etc\/sysconfig\/network-scripts\/ifcfg-lo:0\n\nifup lo:0 || exit 1\n\nexit 0\n",
	"Description": "\u30ed\u30fc\u30c9\u30d0\u30e9\u30f3\u30b9\u5bfe\u8c61\u306e\u30b5\u30fc\u30d0\u306e\u8a2d\u5b9a\u306e\u305f\u3081\u306e\u30b9\u30af\u30ea\u30d7\u30c8\u3067\u3059\u3002\n\u3053\u306e\u30b9\u30af\u30ea\u30d7\u30c8\u306f\u3001CentOS6.X\u3082\u3057\u304f\u306fScientific Linux6.X\u3067\u306e\u307f\u52d5\u4f5c\u3057\u307e\u3059\u3002",
	"Remark": {
		"Form": [
			{
				"type": "text",
				"name": "para1",
				"label": "\u30ed\u30fc\u30c9\u30d0\u30e9\u30f3\u30b5\u30fc\u306eVIP",
				"options": {
					"maxlen": "20",
					"required": true,
					"shellarg": true
				}
			}
		],
		"Require": {
		    "Archive": {
			"Tags": [
			    [ "distro-centos" , "distro-ver-6.*" ],
			    [ "distro-sl" , "distro-ver-6.*"]
			]
		    }
		}
	},
	"Availability": "available",
	"CreatedAt": "2016-03-23T16:49:29+09:00",
	"ModifiedAt": "2016-03-28T19:07:20+09:00",
	"Icon": null,
	"Tags": []
}

`

func TestMarshalNoteJSON(t *testing.T) {
	var note Note
	err := json.Unmarshal([]byte(testNoteJSON), &note)

	assert.NoError(t, err)
	assert.NotEmpty(t, note)

	assert.NotEmpty(t, note.ID)
}
