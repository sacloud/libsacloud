package sacloud

// PacketFilter type of PacketFilter
type PacketFilter struct {
	*Resource
	Index       int `json:",omitempty"`
	Name        string
	Description string `json:",omitempty"`

	//HACK API呼び出しルートにより数字/文字列が混在する
	// PackerFilterのCREATE時は文字列、以外は数値となる。現状利用しないためコメントとしておく
	// RequiredHostVersion int    `json:",omitempty"`

	Notice     string `json:",omitempty"`
	Expression []string
}
