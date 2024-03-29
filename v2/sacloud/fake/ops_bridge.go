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

package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *BridgeOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.BridgeFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Bridge
	for _, res := range results {
		dest := &sacloud.Bridge{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.BridgeFindResult{
		Total:   len(results),
		Count:   len(results),
		From:    0,
		Bridges: values,
	}, nil
}

// Create is fake implementation
func (o *BridgeOp) Create(ctx context.Context, zone string, param *sacloud.BridgeCreateRequest) (*sacloud.Bridge, error) {
	result := &sacloud.Bridge{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	putBridge(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *BridgeOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Bridge, error) {
	value := getBridgeByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Bridge{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *BridgeOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.BridgeUpdateRequest) (*sacloud.Bridge, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	putBridge(zone, value)

	return value, nil
}

// Delete is fake implementation
func (o *BridgeOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, zone, id)
	return nil
}
