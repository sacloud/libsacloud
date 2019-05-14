package naked

// FTPServer FTPサーバ
//
// ZoneAPIの戻り値などに含まれる
type FTPServer struct {
	HostName  string `json:",omitempty" yaml:"host_name,ommitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,ommitempty" structs:",omitempty"`
}

// OpeningFTPServer 接続可能な状態のFTPサーバ
//
// ISOイメージやアーカイブのOpenなどで返される
type OpeningFTPServer struct {
	HostName  string `json:",omitempty" yaml:"host_name,ommitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,ommitempty" structs:",omitempty"`
	User      string `json:",omitempty" yaml:"user,omitempty" structs:",omitempty"`
	Password  string `json:",omitempty" yaml:"password,ommitempty" structs:",omitempty"`
}
