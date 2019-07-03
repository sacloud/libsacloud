package naked

// Host 仮想マシンの起動しているホスト情報
type Host struct {
	InfoURL string `json:",omitempty" yaml:"info_url,omitempty" structs:",omitempty"`
	Name    string `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
}
