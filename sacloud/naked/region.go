package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// Region リージョン
type Region struct {
	ID          types.ID `json:",omitempty" yaml:"id,ommitempty" structs:",omitempty"`
	Name        string   `json:",omitempty" yaml:"name,ommitempty" structs:",omitempty"`
	Description string   `json:",omitempty" yaml:"description,ommitempty" structs:",omitempty"`
	NameServers []string `json:",omitempty" yaml:"name_servers,ommitempty" structs:",omitempty"`
}
