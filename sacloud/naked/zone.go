package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// Zone ゾーン情報
type Zone struct {
	ID           types.ID   `json:",omitempty" yaml:"id,ommitempty" structs:",omitempty"`
	DisplayOrder int        `json:",omitempty" yaml:"display_order,ommitempty" structs:",omitempty"`
	Name         string     `json:",omitempty" yaml:"name,ommitempty" structs:",omitempty"`
	Description  string     `json:",omitempty" yaml:"description,ommitempty" structs:",omitempty"`
	IsDummy      bool       `json:",omitempty" yaml:"is_dummy,ommitempty" structs:",omitempty"`
	VNCProxy     *VNCProxy  `json:",omitempty" yaml:"vnc_proxy,ommitempty" structs:",omitempty"`
	FTPServer    *FTPServer `json:",omitempty" yaml:"ftp_server,ommitempty" structs:",omitempty"`
	Region       *Region    `json:",omitempty" yaml:"region,ommitempty" structs:",omitempty"`
}
