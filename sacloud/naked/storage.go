package naked

// Storage ストレージ
type Storage struct {
	ID          int64     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name        string    `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description string    `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Class       string    `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	DiskPlan    *DiskPlan `json:",omitempty" yaml:"disk_plan,omitempty" structs:",omitempty"`
	Zone        *Zone     `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
}
