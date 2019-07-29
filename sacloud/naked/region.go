package naked

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Region リージョン
type Region struct {
	ID          types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string   `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	NameServers []string `json:",omitempty" yaml:"name_servers,omitempty" structs:",omitempty"`
}
