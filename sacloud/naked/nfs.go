package naked

import (
	"time"

	"github.com/sacloud/libsacloud/sacloud/types"
)

// NFS NFS
type NFS struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string            `json:"" yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	Instance     *Instance           `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Interfaces   []*Interface        `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	Plan         *AppliancePlan      `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Remark       *ApplianceRemark    `json:",omitempty" yaml:"remark,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Switch       *Switch             `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
}
