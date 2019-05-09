package naked

// Zone ゾーン情報
type Zone struct {
	ID           int64      `json:",omitemmpty" yaml:"id,ommitempty" structs:",omitempty"`
	DisplayOrder int        `json:",omitemmpty" yaml:"display_order,ommitempty" structs:",omitempty"`
	Name         string     `json:",omitemmpty" yaml:"name,ommitempty" structs:",omitempty"`
	Description  string     `json:",omitemmpty" yaml:"description,ommitempty" structs:",omitempty"`
	IsDummy      bool       `json:",omitemmpty" yaml:"is_dummy,ommitempty" structs:",omitempty"`
	VNCProxy     *VNCProxy  `json:",omitemmpty" yaml:"vnc_proxy,ommitempty" structs:",omitempty"`
	FTPServer    *FTPServer `json:",omitemmpty" yaml:"ftp_server,ommitempty" structs:",omitempty"`
	Region       *Region    `json:",omitemmpty" yaml:"region,ommitempty" structs:",omitempty"`
}
