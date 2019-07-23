package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Server サーバ
type Server struct {
	ID                types.ID               `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name              string                 `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description       string                 `yaml:"description"`
	Tags              types.Tags             `yaml:"tags"`
	Icon              *Icon                  `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt         *time.Time             `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt        *time.Time             `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability      types.EAvailability    `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	HostName          string                 `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	ServiceClass      string                 `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	InterfaceDriver   types.EInterfaceDriver `json:",omitempty" yaml:"interface_driver,omitempty" structs:",omitempty"`
	ServerPlan        *ServerPlan            `json:",omitempty" yaml:"server_plan,omitempty" structs:",omitempty"`
	Zone              *Zone                  `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Instance          *Instance              `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Disks             []*Disk                `json:",omitempty" yaml:"disks,omitempty" structs:",omitempty"`
	Interfaces        []*Interface           `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	PrivateHost       *PrivateHost           `json:",omitempty" yaml:"private_host,omitempty" structs:",omitempty"`
	WaitDiskMigration bool                   `json:",omitempty" yaml:"wait_disk_migration,omitempty" structs:",omitempty"`
	ConnectedSwitches []*ConnectedSwitch     `json:",omitempty" yaml:"connected_switches,omitempty" structs:",omitempty"`
}

// ConnectedSwitch サーバ作成時に指定する接続先スイッチ
type ConnectedSwitch struct {
	ID    types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
}
