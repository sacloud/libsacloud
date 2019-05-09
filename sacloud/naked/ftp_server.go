package naked

// FTPServer FTPサーバ
type FTPServer struct {
	HostName  string `json:",omitemmpty" yaml:"host_name,ommitempty" structs:",omitempty"`
	IPAddress string `json:",omitemmpty" yaml:"ip_address,ommitempty" structs:",omitempty"`
}
