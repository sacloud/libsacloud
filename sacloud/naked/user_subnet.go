package naked

import "encoding/json"

// UserSubnet ユーザーサブネット
type UserSubnet struct {
	DefaultRoute   string `yaml:"default_route"`
	NetworkMaskLen int    `yaml:"network_mask_len"`
}

// UnmarshalJSON DefaultRouteがからの場合に"0.0.0.0"となることへの対応
func (s *UserSubnet) UnmarshalJSON(data []byte) error {
	type alias UserSubnet
	var tmp alias
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp.DefaultRoute == "0.0.0.0" {
		tmp.DefaultRoute = ""
	}
	*s = UserSubnet(tmp)
	return nil
}
