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
	SizeMB      int       `json:",omitempty"`
	MigratedMB  int       `json:",omitempty"`
	Plan        *Resource `json:",omitempty"`
	DistantFrom []int64   `json:",omitempty"`
	Storage     struct {
		*Resource
		MountIndex int64  `json:",omitempty"`
		Class      string `json:",omitempty"`
	}
	SourceArchive *Archive            `json:",omitempty"`
	SourceDisk    *Disk               `json:",omitempty"`
	JobStatus     *MigrationJobStatus `json:",omitempty"`
	BundleInfo    interface{}         `json:",omitempty"`
	Server        *Server             `json:",omitempty"`
	CreatedAt     *time.Time          `json:",omitempty"`
	Icon          *Icon               `json:",omitempty"`
	*TagsType
}

type DiskPlanID int64

var (
	DiskPlanHDDID                        = DiskPlanID(2)
	DiskPlanSSDID                        = DiskPlanID(4)
	DiskPlanHDD                          = &Resource{ID: int64(DiskPlanHDDID)}
	DiskPlanSSD                          = &Resource{ID: int64(DiskPlanSSDID)}
	DiskConnectionVirtio EDiskConnection = "virtio"
	DiskConnectionIDE    EDiskConnection = "ide"
)

func (d DiskPlanID) ToResource() *Resource {
	return &Resource{ID: int64(d)}
}

func CreateNewDisk() *Disk {
	return &Disk{
		Plan:       DiskPlanSSD,
		Connection: DiskConnectionVirtio,
		SizeMB:     20480,
		TagsType:   &TagsType{},
	}
}

func (d *Disk) SetDiskPlanToHDD() {
	d.Plan = DiskPlanHDD
}
func (d *Disk) SetDiskPlanToSSD() {
	d.Plan = DiskPlanSSD
}

func (d *Disk) SetSourceArchive(sourceID int64) {
	d.SourceArchive = &Archive{
		Resource: &Resource{ID: sourceID},
	}
}

func (d *Disk) SetSourceDisk(sourceID int64) {
	d.SourceDisk = &Disk{
		Resource: &Resource{ID: sourceID},
	}
}

// DiskEditValue type of disk edit request value
type DiskEditValue struct {
	Password      *string   `json:",omitempty"`
	SSHKey        *SSHKey   `json:",omitempty"`
	SSHKeys       []*SSHKey `json:",omitempty"`
	DisablePWAuth *bool     `json:",omitempty"`
	HostName      *string   `json:",omitempty"`
	UserIPAddress *string   `json:",omitempty"`
	UserSubnet    *struct {
		DefaultRoute   string `json:",omitempty"`
		NetworkMaskLen string `json:",omitempty"`
	} `json:",omitempty"`
	Notes []*Resource `json:",omitempty"`
}

func (d *DiskEditValue) SetHostName(value string) {
	d.HostName = &value
}
func (d *DiskEditValue) SetPassword(value string) {
	d.Password = &value
}
func (d *DiskEditValue) SetSSHKeys(keyIDs []string) {
	d.SSHKeys = []*SSHKey{}
	for _, keyID := range keyIDs {
		d.SSHKeys = append(d.SSHKeys, &SSHKey{Resource: NewResourceByStringID(keyID)})
	}
}
func (d *DiskEditValue) SetDisablePWAuth(disable bool) {
	d.DisablePWAuth = &disable
}
func (d *DiskEditValue) SetNotes(noteIDs []string) {
	d.Notes = []*Resource{}
	for _, noteID := range noteIDs {
		d.Notes = append(d.Notes, NewResourceByStringID(noteID))
	}

}

func (d *DiskEditValue) AddNote(noteID string) {
	if d.Notes == nil {
		d.Notes = []*Resource{}
	}
	d.Notes = append(d.Notes, NewResourceByStringID(noteID))
}

func (d *DiskEditValue) SetUserIPAddress(ip string) {
	d.UserIPAddress = &ip
}
func (d *DiskEditValue) SetDefaultRoute(route string) {
	if d.UserSubnet == nil {
		d.UserSubnet = &struct {
			DefaultRoute   string `json:",omitempty"`
			NetworkMaskLen string `json:",omitempty"`
		}{}
	}
	d.UserSubnet.DefaultRoute = route
}

func (d *DiskEditValue) SetNetworkMaskLen(length string) {
	if d.UserSubnet == nil {
		d.UserSubnet = &struct {
			DefaultRoute   string `json:",omitempty"`
			NetworkMaskLen string `json:",omitempty"`
		}{}
	}
	d.UserSubnet.NetworkMaskLen = length
}
