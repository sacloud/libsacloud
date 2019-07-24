package naked

// VNCProxy VNCプロキシ
type VNCProxy struct {
	HostName  string `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	IPAddress string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
}

// VNCProxyInfo サーバに対するVNCProxy
type VNCProxyInfo struct {
	Status       string `json:",omitempty"` // ステータス
	Host         string `json:",omitempty"` // プロキシホスト
	IOServerHost string `json:",omitempty"` // 新プロキシホスト(Hostがlocalhostの場合にこちらを利用する)
	Port         string `json:",omitempty"` // ポート番号
	Password     string `json:",omitempty"` // VNCパスワード
	VNCFile      string `json:",omitempty"` // VNC接続情報ファイル(VNCビューア用)
}
