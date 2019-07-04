package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var zones = []string{"tk1a", "is1a", "is1b", "tk1v"}

var zoneIDs = map[string]types.ID{
	"tk1a": types.ID(21001),
	"is1a": types.ID(31001),
	"is1b": types.ID(31002),
	"tk1v": types.ID(29001),
}

var sharedSegmentSwitch = &sacloud.Switch{
	ID:             pool.generateID(),
	Name:           "スイッチ",
	Scope:          types.Scopes.Shared,
	Description:    "共有セグメント用スイッチ",
	NetworkMaskLen: pool.sharedNetMaskLen,
	DefaultRoute:   pool.sharedDefaultGateway.String(),
}

func init() {
	initArchives()
	initNotes()
	initSwitch()
	initZones()
}

func initArchives() {
	archives := []*sacloud.Archive{
		{
			ID:                   pool.generateID(),
			Name:                 "CentOS 7.6 (1810) 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-centos", "distro-ver-7.6", "os-linux"},
			DisplayOrder:         1,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.ID(2),
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
		{
			ID:                   pool.generateID(),
			Name:                 "Ubuntu Server 18.04.2 LTS 64bit",
			Tags:                 []string{"@size-extendable", "arch-64bit", "current-stable", "distro-ubuntu", "distro-ver-18.04.2", "os-linux"},
			DisplayOrder:         2,
			Scope:                types.Scopes.Shared,
			Availability:         types.Availabilities.Available,
			SizeMB:               20480,
			DiskPlanID:           types.ID(2),
			DiskPlanName:         "標準プラン",
			DiskPlanStorageClass: "iscsi9999",
		},
	}
	for _, zone := range zones {
		for _, archive := range archives {
			s.setArchive(zone, archive)
		}
	}
}

func initNotes() {
	notes := []*sacloud.Note{
		{
			ID:           1,
			Name:         "fake",
			Availability: types.Availabilities.Available,
			Scope:        types.Scopes.Shared,
			Class:        "shell",
		},
	}
	for _, note := range notes {
		s.setNote(sacloud.APIDefaultZone, note)
	}
}

func initSwitch() {
	for _, zone := range zones {
		s.setSwitch(zone, sharedSegmentSwitch)
	}
}

func initZones() {
	// zones
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           21001,
		Name:         "tk1a",
		Description:  "東京第1ゾーン",
		DisplayOrder: 1,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           31001,
		Name:         "is1a",
		Description:  "石狩第1ゾーン",
		DisplayOrder: 2,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           31002,
		Name:         "is1b",
		Description:  "石狩第2ゾーン",
		DisplayOrder: 3,
	})
	s.setZone(sacloud.APIDefaultZone, &sacloud.Zone{
		ID:           29001,
		Name:         "tk1v",
		Description:  "Sandbox",
		DisplayOrder: 4,
		IsDummy:      true,
	})
}
