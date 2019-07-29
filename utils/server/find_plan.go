package server

import (
	"context"
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// PlanFinder .
type PlanFinder interface {
	Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error)
}

// FindPlanRequest サーバプラン検索パラメータ
type FindPlanRequest struct {
	CPU        int
	MemoryGB   int
	Commitment types.ECommitment
	Generation types.EPlanGeneration
}

func (f *FindPlanRequest) findCondition() *sacloud.FindCondition {
	cond := &sacloud.FindCondition{
		Sort: search.SortKeys{
			{Key: "Generation", Order: search.SortDesc},
		},
		Filter: search.Filter{
			search.Key("Commitment"): types.Commitments.Standard,
		},
		Count: 1000,
	}
	if f.CPU > 0 {
		cond.Filter[search.Key("CPU")] = f.CPU
	}
	if f.MemoryGB > 0 {
		cond.Filter[search.Key("MemoryMB")] = f.MemoryGB * 1024
	}
	if f.Generation != types.PlanGenerations.Default {
		cond.Filter[search.Key("Generation")] = f.Generation
	}
	if f.Commitment != types.Commitments.Unknown && f.Commitment != types.Commitments.Standard {
		cond.Filter[search.Key("Commitment")] = f.Commitment
	}
	return cond
}

// FindPlan サーバプラン検索
func FindPlan(ctx context.Context, finder PlanFinder, zone string, param *FindPlanRequest) (*sacloud.ServerPlan, error) {
	var cond *sacloud.FindCondition
	if param != nil {
		cond = param.findCondition()
	}

	searched, err := finder.Find(ctx, zone, cond)
	if err != nil {
		return nil, err
	}
	if searched.Count == 0 || len(searched.ServerPlans) == 0 {
		return nil, errors.New("server plan not found")
	}
	return searched.ServerPlans[0], nil
}
