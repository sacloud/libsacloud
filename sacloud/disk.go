package sacloud

import (
	"time"
)

// Disk type of disk
type Disk struct {
	*Resource
	Index           int             `json:",omitempty"`
	Name            string          `json:",omitempty"`
	Description     string          `json:",omitempty"`
	Connection      EDiskConnection `json:",omitempty"`
	ConnectionOrder int             `json:",omitempty"`
	ReinstallCount  int             `json:",omitempty"`
	*EAvailability
	SizeMB     int             `json:",omitempty"`
	MigratedMB int             `json:",omitempty"`
	Plan       *NumberResource `json:",omitempty"`
	Storage    struct {
		*Resource
		MountIndex string `json:",omitempty"`
		Class      string `json:",omitempty"`
	}
	SourceArchive *Archive `json:",omitempty"`
	SourceDisk    *Disk    `json:",omitempty"`
	//BundleInfo
	CreatedAt *time.Time `json:",omitempty"`
	Icon      *Icon      `json:",omitempty"`
	Tags      []string   `json:",omitempty"`
}

var (
	DiskPlanHDD                          = &NumberResource{ID: "2"}
	DiskPlanSSD                          = &NumberResource{ID: "4"}
	DiskConnectionVirtio EDiskConnection = "virtio"
	DiskConnectionIDE    EDiskConnection = "ide"
)

func (d *Disk) SetSourceArchive(sourceID string) {
	d.SourceArchive = &Archive{
		Resource: &Resource{ID: sourceID},
	}
}

func (d *Disk) SetSourceDisk(sourceID string) {
	d.SourceDisk = &Disk{
		Resource: &Resource{ID: sourceID},
	}
}

// DiskEditValue type of disk edit request value
type DiskEditValue struct {
	Password      string   `json:",omitempty"`
	SSHKey        SSHKey   `json:",omitempty"`
	SSHKeys       []SSHKey `json:",omitempty"`
	DisablePWAuth bool     `json:",omitempty"`
	HostName      string   `json:",omitempty"`
	UserIPAddress string   `json:",omitempty"`
	UserSubnet    struct {
		DefaultRoute   string `json:",omitempty"`
		NetworkMaskLen string `json:",omitempty"`
	} `json:",omitempty"`
	Notes []Resource `json:",omitempty"`
}

func (d *DiskEditValue) AddNote(noteID string) {
	if d.Notes == nil {
		d.Notes = []Resource{}
	}
	d.Notes = append(d.Notes, Resource{ID: noteID})
}
