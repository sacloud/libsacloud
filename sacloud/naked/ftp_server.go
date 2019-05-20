package naked

// FTPServer FTPサーバ
//
// ZoneAPIの戻り値などに含まれる
type FTPServer struct {
	HostName  string `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}

// OpeningFTPServer 接続可能な状態のFTPサーバ
//
// ISOイメージやアーカイブのOpenなどで返される
type OpeningFTPServer struct {
	HostName  string `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	User      string `json:",omitempty" yaml:"user,omitempty" structs:",omitempty"`
	Password  string `json:",omitempty" yaml:"password,omitempty" structs:",omitempty"`
}
