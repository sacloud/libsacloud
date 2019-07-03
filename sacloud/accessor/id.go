package accessor

import "github.com/sacloud/libsacloud/sacloud/types"

/************************************************
 ID - StringID
************************************************/

// ID is accessor interface of ID field
type ID interface {
	GetID() types.ID
	SetID(id types.ID)
}

// GetStringID returns string id
func GetStringID(target ID) string {
	return target.GetID().String()
}

// SetStringID sets id from string
func SetStringID(target ID, id string) {
	target.SetID(types.StringID(id))
}

// GetInt64ID returns int64 id
func GetInt64ID(target ID) int64 {
	return target.GetID().Int64()
}

// SetInt64ID sets id from int64
func SetInt64ID(target ID, id int64) {
	target.SetID(types.ID(id))
}
