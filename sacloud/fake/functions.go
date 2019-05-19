package fake

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/accessor"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func random(max int) int {
	return rand.Intn(max)
}

func newErrorNotFound(resourceKey string, id types.ID) error {
	return sacloud.NewAPIError("", nil, "", http.StatusNotFound, &sacloud.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "404 NotFound",
		ErrorCode:    fmt.Sprintf("%d", http.StatusNotFound),
		ErrorMessage: fmt.Sprintf("%s[ID:%s] is not found", resourceKey, id),
	})
}

func newErrorBadRequest(resourceKey string, id types.ID, msgs ...string) error {
	return sacloud.NewAPIError("", nil, "", http.StatusBadRequest, &sacloud.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "400 BadRequest",
		ErrorCode:    fmt.Sprintf("%d", http.StatusBadRequest),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is bad: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func newErrorConflict(resourceKey string, id types.ID, msgs ...string) error {
	return sacloud.NewAPIError("", nil, "", http.StatusConflict, &sacloud.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "409 Conflict",
		ErrorCode:    fmt.Sprintf("%d", http.StatusConflict),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is conflicted: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func newInternalServerError(resourceKey string, id types.ID, msgs ...string) error {
	return sacloud.NewAPIError("", nil, "", http.StatusInternalServerError, &sacloud.APIErrorResponse{
		IsFatal:      true,
		Serial:       "",
		Status:       "500 Internal Server Error",
		ErrorCode:    fmt.Sprintf("%d", http.StatusInternalServerError),
		ErrorMessage: fmt.Sprintf("request to %s[ID:%s] is failed: %s", resourceKey, id, strings.Join(msgs, " ")),
	})
}

func find(resourceKey, zone string, conditions *sacloud.FindCondition) ([]interface{}, error) {
	var results []interface{}
	if conditions == nil {
		conditions = &sacloud.FindCondition{}
	}

	targets := s.get(resourceKey, zone)
	for i, target := range targets {
		// count
		if conditions.Count != 0 && len(results) >= conditions.Count {
			break
		}

		// from
		if i < conditions.From {
			continue
		}

		results = append(results, target)
	}

	// TODO sort/filter/include/exclude is not implemented
	return results, nil
}

func copySameNameField(source interface{}, dest interface{}) {
	data, _ := json.Marshal(source)
	json.Unmarshal(data, dest)
}

func fill(target interface{}, fillFuncs ...func(interface{})) {
	for _, f := range fillFuncs {
		f(target)
	}
}

func fillID(target interface{}) {
	if v, ok := target.(accessor.ID); ok {
		id := v.GetID()
		if id.IsEmpty() {
			v.SetID(pool.generateID())
		}
	}
}

func fillAvailability(target interface{}) {
	if v, ok := target.(accessor.Availability); ok {
		value := v.GetAvailability()
		if value == types.Availabilities.Unknown {
			v.SetAvailability(types.Availabilities.Available)
		}
	}
}

func fillScope(target interface{}) {
	if v, ok := target.(accessor.Scope); ok {
		value := v.GetScope()
		if value == types.EScope("") {
			v.SetScope(types.Scopes.User)
		}
	}
}

func fillDiskPlan(target interface{}) {
	if v, ok := target.(accessor.DiskPlan); ok {
		id := v.GetDiskPlanID()
		switch id {
		case types.ID(2):
			v.SetDiskPlanName("標準プラン")
		case types.ID(4):
			v.SetDiskPlanName("SSDプラン")
		}
		v.SetDiskPlanStorageClass("iscsi9999")
	}
}

func fillCreatedAt(target interface{}) {
	if v, ok := target.(accessor.CreatedAt); ok {
		value := v.GetCreatedAt()
		if value == nil {
			now := time.Now()
			v.SetCreatedAt(&now)
		}
	}
}

func fillModifiedAt(target interface{}) {
	if v, ok := target.(accessor.ModifiedAt); ok {
		value := v.GetModifiedAt()
		if value == nil {
			now := time.Now()
			v.SetModifiedAt(&now)
		}
	}
}
