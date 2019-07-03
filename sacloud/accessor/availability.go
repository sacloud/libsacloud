package accessor

import "github.com/sacloud/libsacloud/sacloud/types"

// Availability Availabilityを持つリソース向けのインターフェース
type Availability interface {
	GetAvailability() types.EAvailability
	SetAvailability(types.EAvailability)
}
