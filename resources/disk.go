package resources

import "time"

// Disk type of disk
type Disk struct {
	*Resource
	Name            string
	Connection      string `json:",omitempty"`
	ConnectionOrder int    `json:",omitempty"`
	ReinstallCount  int    `json:",omitempty"`
	*EAvailability
	SizeMB     int
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
	CreatedAt time.Time `json:",omitempty"`
	Icon      *Icon     `json:",omitempty"`
}

// DiskEditValue type of disk edit request value
type DiskEditValue struct {
	Password      string
	SSHKey        SSHKey
	DisablePWAuth bool
	Notes         []Resource
}
