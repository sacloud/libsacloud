package naked

// DiskPlan ディスクプラン
type DiskPlan struct {
	ID           int    `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	StorageClass string `json:",omitempty" yaml:"storage_class,omitempty" structs:",omitempty"`
}
