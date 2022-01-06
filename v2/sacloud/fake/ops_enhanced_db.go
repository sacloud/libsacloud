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
func (o *EnhancedDBOp) Find(ctx context.Context, conditions *sacloud.FindCondition) (*sacloud.EnhancedDBFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.EnhancedDB
	for _, res := range results {
		dest := &sacloud.EnhancedDB{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.EnhancedDBFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		EnhancedDBs: values,
	}, nil
}

// Create is fake implementation
func (o *EnhancedDBOp) Create(ctx context.Context, param *sacloud.EnhancedDBCreateRequest) (*sacloud.EnhancedDB, error) {
	result := &sacloud.EnhancedDB{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.DatabaseType = "tidb"
	result.Region = "is1"
	result.Port = 3306
	result.HostName = result.DatabaseName + ".tidb-is1.db.sakurausercontent.com"
	result.MaxConnections = 50
	result.Availability = types.Availabilities.Available

	putEnhancedDB(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *EnhancedDBOp) Read(ctx context.Context, id types.ID) (*sacloud.EnhancedDB, error) {
	value := getEnhancedDBByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.EnhancedDB{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *EnhancedDBOp) Update(ctx context.Context, id types.ID, param *sacloud.EnhancedDBUpdateRequest) (*sacloud.EnhancedDB, error) {
	value, err := o.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	value.MaxConnections = 50
	fill(value, fillModifiedAt)

	putEnhancedDB(sacloud.APIDefaultZone, value)
	return value, nil
}

// Delete is fake implementation
func (o *EnhancedDBOp) Delete(ctx context.Context, id types.ID) error {
	_, err := o.Read(ctx, id)
	if err != nil {
		return err
	}

	ds().Delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}

// SetPassword is fake implementation
func (o *EnhancedDBOp) SetPassword(ctx context.Context, id types.ID, param *sacloud.EnhancedDBSetPasswordRequest) error {
	_, err := o.Read(ctx, id)
	return err
}
