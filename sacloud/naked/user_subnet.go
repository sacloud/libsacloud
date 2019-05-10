package naked

// UserSubnet ユーザーサブネット
type UserSubnet struct {
	DefaultRoute   string `json:",omitempty" yaml:"default_route,omitempty" structs:",omitempty" `
	NetworkMaskLen int    `json:",omitempty" yaml:"network_mask_len,omitempty" structs:",omitempty"`
}
