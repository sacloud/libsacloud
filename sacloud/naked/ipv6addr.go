package naked

// IPv6Addr IPアドレス(IPv6)
type IPv6Addr struct {
	HostName  string     `yaml:"host_name"`                                         // ホスト名
	IPv6Addr  string     `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // IPv6アドレス
	Interface *Interface `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // インターフェース
	IPv6Net   *IPv6Net   `json:",omitempty" yaml:",omitempty" structs:",omitempty"` // IPv6サブネット
}
