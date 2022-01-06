// Copyright 2016-2022 The Libsacloud Authors
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
	"testing"

	internetBuilder "github.com/sacloud/libsacloud/v2/helper/builder/internet"
	"github.com/sacloud/libsacloud/v2/helper/plans"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestReadServer(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	zone := testutil.TestZone()

	serverOp := sacloud.NewServerOp(caller)

	server, err := serverOp.Create(ctx, zone, &sacloud.ServerCreateRequest{
		Name:     testutil.ResourceName("query-server-with-previous-id"),
		CPU:      1,
		MemoryMB: 1 * size.GiB,
	})
	if err != nil {
		t.Fatal(err)
	}

	// ID指定でreadできる
	read, err := ReadServer(ctx, caller, zone, server.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.EqualValues(t, server.ID, read.ID)

	// プラン変更(@previous-idタグが付与される)
	_, err = plans.ChangeServerPlan(ctx, caller, zone, server.ID, 1, 2, types.Commitments.Standard, types.PlanGenerations.Default)
	if err != nil {
		t.Fatal(err)
	}

	// 旧IDでReadできる
	read2, err := ReadServer(ctx, caller, zone, server.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 2, read2.GetMemoryGB())

	// 存在しないIDだとNoResultErr
	_, err = ReadServer(ctx, caller, zone, types.ID(9999999999))
	require.True(t, sacloud.IsNoResultsError(err))

	// cleanup
	if err := serverOp.Delete(ctx, zone, read2.ID); err != nil {
		t.Fatal(err)
	}
}

func TestReadRouter(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	zone := testutil.TestZone()

	routerOp := sacloud.NewInternetOp(caller)
	routerBuilder := &internetBuilder.Builder{
		Name:           testutil.ResourceName("query-router-with-previous-id"),
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
		Client:         internetBuilder.NewAPIClient(caller),
	}

	router, err := routerBuilder.Build(ctx, zone)
	if err != nil {
		t.Fatal(err)
	}

	// ID指定でreadできる
	read, err := ReadRouter(ctx, caller, zone, router.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.EqualValues(t, router.ID, read.ID)

	// プラン変更(@previous-idタグが付与される)
	_, err = plans.ChangeRouterPlan(ctx, caller, zone, router.ID, 250)
	if err != nil {
		t.Fatal(err)
	}

	// 旧IDでReadできる
	read2, err := ReadRouter(ctx, caller, zone, router.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 250, read2.BandWidthMbps)

	// cleanup
	if err := routerOp.Delete(ctx, zone, read2.ID); err != nil {
		t.Fatal(err)
	}
}

func TestReadProxyLB(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()

	elbOp := sacloud.NewProxyLBOp(caller)
	elb, err := elbOp.Create(ctx, &sacloud.ProxyLBCreateRequest{
		Plan:   types.ProxyLBPlans.CPS100,
		Region: types.ProxyLBRegions.IS1,
		Name:   testutil.ResourceName("query-elb-with-previous-id"),
		HealthCheck: &sacloud.ProxyLBHealthCheck{
			Protocol:  types.ProxyLBProtocols.TCP,
			DelayLoop: 10,
		},
		Timeout: &sacloud.ProxyLBTimeout{
			InactiveSec: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// ID指定でreadできる
	read, err := ReadProxyLB(ctx, caller, elb.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.EqualValues(t, elb.ID, read.ID)

	// プラン変更(@previous-idタグが付与される)
	_, err = plans.ChangeProxyLBPlan(ctx, caller, elb.ID, types.ProxyLBPlans.CPS500.Int())
	if err != nil {
		t.Fatal(err)
	}

	// 旧IDでReadできる
	read2, err := ReadProxyLB(ctx, caller, elb.ID)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, types.ProxyLBPlans.CPS500, read2.Plan)

	// cleanup
	if err := elbOp.Delete(ctx, read2.ID); err != nil {
		t.Fatal(err)
	}
}
