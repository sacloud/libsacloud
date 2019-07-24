package nfs

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// NoteFinder スタートアップスクリプト(Note)検索インターフェース
type NoteFinder interface {
	Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.NoteFindResult, error)
}

type nfsPlansEnvelope struct {
	Plans nfsPlans `json:"plans"`
}

type nfsPlans struct {
	HDD []nfsPlanValue
	SSD []nfsPlanValue
}

func (p nfsPlans) findPlanID(diskPlanID types.ID, size types.ENFSSize) types.ID {
	var plans []nfsPlanValue
	switch diskPlanID {
	case types.NFSPlans.HDD:
		plans = p.HDD
	case types.NFSPlans.SSD:
		plans = p.SSD
	default:
		return types.ID(0)
	}

	for _, plan := range plans {
		if plan.Availability.IsAvailable() && plan.Size == int(size) {
			return plan.PlanID
		}
	}

	return types.ID(0)
}

type nfsPlanValue struct {
	Size         int                 `json:"size"`
	Availability types.EAvailability `json:"availability"`
	PlanID       types.ID            `json:"planId"`
}

// FindNFSPlanID ディスクプランとサイズからNFSのプランIDを取得
func FindNFSPlanID(ctx context.Context, finder NoteFinder, diskPlanID types.ID, size types.ENFSSize) (types.ID, error) {

	// find note
	searched, err := finder.Find(ctx, &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key(keys.Name): "sys-nfs",
			search.Key("Class"):   "json",
		},
	})
	if err != nil {
		return types.ID(0), err
	}
	if searched.Count == 0 || len(searched.Notes) == 0 {
		return types.ID(0), errors.New("note[sys-nfs] not found")
	}
	note := searched.Notes[0]

	// parse note's content
	var pe nfsPlansEnvelope
	if err := json.Unmarshal([]byte(note.Content), &pe); err != nil {
		return types.ID(0), err
	}

	return pe.Plans.findPlanID(diskPlanID, size), nil
}
