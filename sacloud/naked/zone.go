package naked

import "github.com/sacloud/libsacloud/sacloud/types"

// Zone ゾーン情報
type Zone struct {
	ID           types.ID   `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	DisplayOrder int        `json:",omitempty" yaml:"display_order,omitempty" structs:",omitempty"`
	Name         string     `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string     `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	IsDummy      bool       `json:",omitempty" yaml:"is_dummy,omitempty" structs:",omitempty"`
	VNCProxy     *VNCProxy  `json:",omitempty" yaml:"vnc_proxy,omitempty" structs:",omitempty"`
	FTPServer    *FTPServer `json:",omitempty" yaml:"ftp_server,omitempty" structs:",omitempty"`
	Region       *Region    `json:",omitempty" yaml:"region,omitempty" structs:",omitempty"`
}
