package resources

import "time"

// Instance type of instance
type Instance struct {
	Server            Resource  `json:",omitempty"`
	Status            string    `json:",omitempty"`
	BeforeStatus      string    `json:",omitempty"`
	StatusChangedAt   time.Time `json:",omitempty"`
	MigrationProgress string    `json:",omitempty"`
	MigrationSchedule string    `json:",omitempty"`
	IsMigrating       bool      `json:",omitempty"`
	MigrationAllowed  string    `json:",omitempty"`
	ModifiedAt        time.Time `json:",omitempty"`
	Host              struct {
		Name          string `json:",omitempty"`
		InfoURL       string `json:",omitempty"`
		Class         string `json:",omitempty"`
		Version       int    `json:",omitempty"`
		SystemVersion string `json:",omitempty"`
	} `json:",omitempty"`
	CDROM struct {
		*Resource
		DisplayOrder string `json:",omitempty"`
		StorageClass string `json:",omitempty"`
		Name         string `json:",omitempty"`
		Description  string `json:",omitempty"`
		SizeMB       int    `json:",omitempty"`
		Scope        string `json:",omitempty"`
		*EAvailability
		ServiceClass string    `json:",omitempty"`
		CreatedAt    time.Time `json:",omitempty"`
		Icon         string    `json:",omitempty"`
		Storage      *Storage  `json:",omitempty"`
	} `json:",omitempty"`
	CDROMStorage *Storage `json:",omitempty"`
}

// Storage type of Storage
type Storage struct {
	*Resource
	Class       string `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Zone        *Zone  `json:",omitempty"`
	DiskPlan    struct {
		*NumberResource
		StorageClass string `json:",omitempty"`
		Name         string `json:",omitempty"`
	} `json:",omitempty"`
	//Capacity []string `json:",omitempty"`
}
