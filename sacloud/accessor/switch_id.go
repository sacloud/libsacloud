package accessor

import "github.com/sacloud/libsacloud/sacloud/types"

/************************************************
 switchID
************************************************/

// SwitchID is accessor interface of SwitchID field
type SwitchID interface {
	GetSwitchID() types.ID
	SetSwitchID(id types.ID)
}
