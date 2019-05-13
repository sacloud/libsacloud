package naked

import "github.com/sacloud/libsacloud-v2/sacloud/types"

// Provider CommonServiceItemなどで利用されるProvider
type Provider struct {
	ID           types.ID `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string   `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Class        string   `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	ServiceClass string   `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}
