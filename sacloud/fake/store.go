package fake

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Store fakeドライバーでのバックエンド(永続化)を担当するドライバーインターフェース
type Store interface {
	Init() error
	NeedInitData() bool
	Put(resourceKey, zone string, id types.ID, value interface{})
	Get(resourceKey, zone string, id types.ID) interface{}
	List(resourceKey, zone string) []interface{}
	Delete(resourceKey, zone string, id types.ID)
}
