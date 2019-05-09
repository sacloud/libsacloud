package naked

// Region リージョン
type Region struct {
	ID          int64    `json:",omitemmpty" yaml:"id,ommitempty" structs:",omitempty"`
	Name        string   `json:",omitemmpty" yaml:"name,ommitempty" structs:",omitempty"`
	Description string   `json:",omitemmpty" yaml:"description,ommitempty" structs:",omitempty"`
	NameServers []string `json:",omitemmpty" yaml:"name_servers,ommitempty" structs:",omitempty"`
}
