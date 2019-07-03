package fake

import (
	"fmt"
	"sync"

	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var s = &store{
	data: make(map[string]map[types.ID]interface{}),
}

type store struct {
	data map[string]map[types.ID]interface{}
	mu   sync.Mutex
}

func (s *store) key(resourceKey, zone string) string {
	return fmt.Sprintf("%s/%s", resourceKey, zone)
}

func (s *store) values(resourceKey, zone string) map[types.ID]interface{} {
	return s.data[s.key(resourceKey, zone)]
}

func (s *store) get(resourceKey, zone string) []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	var ret []interface{}
	for _, v := range values {
		ret = append(ret, v)
	}
	return ret
}

func (s *store) getByID(resourceKey, zone string, id types.ID) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		return nil
	}
	return values[id]
}

func (s *store) set(resourceKey, zone string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values == nil {
		values = map[types.ID]interface{}{}
	}
	if v, ok := value.(accessor.ID); ok {
		values[v.GetID()] = value
	} else {
		panic("value has no ID")
	}

	s.data[s.key(resourceKey, zone)] = values
}

func (s *store) delete(resourceKey, zone string, id types.ID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	values := s.values(resourceKey, zone)
	if values != nil {
		delete(values, id)
	}
}
