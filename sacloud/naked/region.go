package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// Region リージョン
type Region struct {
	ID          types.ID `json:",omitemmpty" yaml:"id,ommitempty" structs:",omitempty"`
	Name        string   `json:",omitemmpty" yaml:"name,ommitempty" structs:",omitempty"`
	Description string   `json:",omitemmpty" yaml:"description,ommitempty" structs:",omitempty"`
	NameServers []string `json:",omitemmpty" yaml:"name_servers,ommitempty" structs:",omitempty"`
}
