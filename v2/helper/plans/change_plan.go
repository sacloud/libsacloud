// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plans

import (
	"context"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var (
	PreviousIDTagName = "@previous-id"
	maxTags           = 10 // タグ上限数
)

// ChangeServerPlan 現在のIDをタグとして保持しつつプランを変更する
//
// もしすでにタグが上限(10)まで付与されている場合はプラン変更だけ行う
func ChangeServerPlan(
	ctx context.Context,
	caller sacloud.APICaller,
	zone string,
	id types.ID,
	cpu int,
	memoryGB int,
	commitment types.ECommitment,
	generation types.EPlanGeneration,
) (*sacloud.Server, error) {
	serverOp := sacloud.NewServerOp(caller)
	server, err := serverOp.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	if len(server.Tags) < maxTags {
		server.Tags = AppendPreviousIDTagIfAbsent(server.Tags, server.ID)

		updated, err := serverOp.Update(ctx, zone, server.ID, &sacloud.ServerUpdateRequest{
			Name:            server.Name,
			Description:     server.Description,
			Tags:            server.Tags,
			IconID:          server.IconID,
			PrivateHostID:   server.PrivateHostID,
			InterfaceDriver: server.InterfaceDriver,
		})
		if err != nil {
			return nil, err
		}
		server = updated
	}

	return serverOp.ChangePlan(ctx, zone, server.ID, &sacloud.ServerChangePlanRequest{
		CPU:                  cpu,
		MemoryMB:             memoryGB * size.GiB,
		ServerPlanGeneration: generation,
		ServerPlanCommitment: commitment,
	})
}

// ChangeRouterPlan 現在のIDをタグとして保持しつつプランを変更する
//
// もしすでにタグが上限(10)まで付与されている場合はプラン変更だけ行う
func ChangeRouterPlan(
	ctx context.Context,
	caller sacloud.APICaller,
	zone string,
	id types.ID,
	bandWidth int,
) (*sacloud.Internet, error) {
	internetOp := sacloud.NewInternetOp(caller)
	router, err := internetOp.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	if len(router.Tags) < maxTags {
		router.Tags = AppendPreviousIDTagIfAbsent(router.Tags, router.ID)

		updated, err := internetOp.Update(ctx, zone, router.ID, &sacloud.InternetUpdateRequest{
			Name:        router.Name,
			Description: router.Description,
			Tags:        router.Tags,
			IconID:      router.IconID,
		})
		if err != nil {
			return nil, err
		}
		router = updated
	}

	return internetOp.UpdateBandWidth(ctx, zone, router.ID, &sacloud.InternetUpdateBandWidthRequest{
		BandWidthMbps: bandWidth,
	})
}

// ChangeProxyLBPlan 現在のIDをタグとして保持しつつプランを変更する
//
// もしすでにタグが上限(10)まで付与されている場合はプラン変更だけ行う
func ChangeProxyLBPlan(
	ctx context.Context,
	caller sacloud.APICaller,
	id types.ID,
	cps int,
) (*sacloud.ProxyLB, error) {
	elbOp := sacloud.NewProxyLBOp(caller)
	elb, err := elbOp.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(elb.Tags) < maxTags {
		elb.Tags = AppendPreviousIDTagIfAbsent(elb.Tags, elb.ID)

		updated, err := elbOp.Update(ctx, elb.ID, &sacloud.ProxyLBUpdateRequest{
			HealthCheck:   elb.HealthCheck,
			SorryServer:   elb.SorryServer,
			BindPorts:     elb.BindPorts,
			Servers:       elb.Servers,
			Rules:         elb.Rules,
			LetsEncrypt:   elb.LetsEncrypt,
			StickySession: elb.StickySession,
			Timeout:       elb.Timeout,
			Gzip:          elb.Gzip,
			Syslog:        elb.Syslog,
			SettingsHash:  elb.SettingsHash,
			Name:          elb.Name,
			Description:   elb.Description,
			Tags:          elb.Tags,
			IconID:        elb.IconID,
		})
		if err != nil {
			return nil, err
		}
		elb = updated
	}

	return elbOp.ChangePlan(ctx, elb.ID, &sacloud.ProxyLBChangePlanRequest{
		ServiceClass: types.ProxyLBServiceClass(types.EProxyLBPlan(cps), elb.Region),
	})
}

func AppendPreviousIDTagIfAbsent(tags types.Tags, currentID types.ID) types.Tags {
	if len(tags) > maxTags {
		return tags
	}
	// すでに付けられたPreviousIDタグを消す
	updated := types.Tags{}
	for _, t := range tags {
		if !strings.HasPrefix(t, PreviousIDTagName) {
			updated = append(updated, t)
		}
	}
	updated = append(updated, fmt.Sprintf("%s=%s", PreviousIDTagName, currentID))
	updated.Sort()
	return updated
}
