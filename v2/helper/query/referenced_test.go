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
	"sync"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/testutil"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/require"
)

func TestCheckReferenced_withTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	caller := testutil.SingletonAPICaller()
	var f referenceFindFunc = func(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (bool, error) {
		time.Sleep(10 * time.Millisecond)
		return false, ctx.Err()
	}

	_, err := checkReferenced(ctx, caller, []string{"tk1v"}, types.ID(0), []referenceFindFunc{f, f})
	require.Error(t, err)
	require.EqualValues(t, "context deadline exceeded", err.Error())
}

func TestCheckReferenced_withMultipleZone(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	var mu sync.Mutex
	called := 0

	f := func(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (bool, error) {
		mu.Lock()
		defer mu.Unlock()
		called++
		return false, nil
	}

	zones := []string{"1", "2", "3"}
	funcs := []referenceFindFunc{f, f, f}
	result, err := checkReferenced(ctx, caller, zones, types.ID(0), funcs)
	require.Equal(t, len(zones)*len(funcs), called)
	require.Equal(t, result, false)
	require.NoError(t, err)
}

func TestCheckReferenced_shortCircuit(t *testing.T) {
	ctx := context.Background()
	caller := testutil.SingletonAPICaller()
	var mu sync.Mutex
	called := 0

	f := func(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (bool, error) {
		mu.Lock()
		defer mu.Unlock()
		called++
		if called == 2 {
			return true, nil
		}
		return false, nil
	}

	zones := []string{"1", "2", "3"}
	funcs := []referenceFindFunc{f, f, f}
	result, err := checkReferenced(ctx, caller, zones, types.ID(0), funcs)

	require.Equal(t, 2, called)
	require.Equal(t, result, true)
	require.NoError(t, err)
}
