package sacloud

import "time"

type License struct {
	*Resource
	Index       int       `json:",omitempty"`
	Name        string    `json:",omitempty"`
	CreatedAt   time.Time `json:",omitempty"`
	ModifiedAt  time.Time `json:",omitempty"`
	LicenseInfo struct {
		*NumberResource
		Name         string `json:",omitempty"`
		ServiceClass string `json:",omitempty"`
	} `json:",omitempty"`
}
