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

package query

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/plans"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func findByPreviousIDCondition(id types.ID) *sacloud.FindCondition {
	return &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Tags.Name"): search.TagsAndEqual(fmt.Sprintf("%s=%s", plans.PreviousIDTagName, id)),
		},
	}
}

// ReadServer 指定のIDでサーバを検索、IDで見つからなかった場合は@previous-idタグで検索し見つかったサーバリソースを返す
//
// 対象が見つからなかった場合はsacloud.NoResultsErrorを返す
func ReadServer(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (*sacloud.Server, error) {
	serverOp := sacloud.NewServerOp(caller)

	server, err := serverOp.Read(ctx, zone, id)
	if err != nil {
		if !sacloud.IsNotFoundError(err) {
			return nil, err
		}

		found, err := serverOp.Find(ctx, zone, findByPreviousIDCondition(id))
		if err != nil {
			return nil, err
		}
		if len(found.Servers) == 0 {
			return nil, sacloud.NewNoResultsError()
		}

		// 複数ヒットした場合でも先頭だけ返す
		server = found.Servers[0]
	}

	return server, nil
}

// ReadRouter 指定のIDでルータを検索、IDで見つからなかった場合は@previous-idタグで検索し見つかったリソースを返す
//
// 対象が見つからなかった場合はsacloud.NoResultsErrorを返す
func ReadRouter(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (*sacloud.Internet, error) {
	routerOp := sacloud.NewInternetOp(caller)

	router, err := routerOp.Read(ctx, zone, id)
	if err != nil {
		if !sacloud.IsNotFoundError(err) {
			return nil, err
		}

		found, err := routerOp.Find(ctx, zone, findByPreviousIDCondition(id))
		if err != nil {
			return nil, err
		}
		if len(found.Internet) == 0 {
			return nil, sacloud.NewNoResultsError()
		}

		// 複数ヒットした場合でも先頭だけ返す
		router = found.Internet[0]
	}

	return router, nil
}

// ReadProxyLB 指定のIDでELBを検索、IDで見つからなかった場合は@previous-idタグで検索し見つかったリソースを返す
//
// 対象が見つからなかった場合はsacloud.NoResultsErrorを返す
func ReadProxyLB(ctx context.Context, caller sacloud.APICaller, id types.ID) (*sacloud.ProxyLB, error) {
	elbOp := sacloud.NewProxyLBOp(caller)

	elb, err := elbOp.Read(ctx, id)
	if err != nil {
		if !sacloud.IsNotFoundError(err) {
			return nil, err
		}

		found, err := elbOp.Find(ctx, findByPreviousIDCondition(id))
		if err != nil {
			return nil, err
		}
		if len(found.ProxyLBs) == 0 {
			return nil, sacloud.NewNoResultsError()
		}

		// 複数ヒットした場合でも先頭だけ返す
		elb = found.ProxyLBs[0]
	}

	return elb, nil
}
