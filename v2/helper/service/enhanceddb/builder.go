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

package enhanceddb

import (
	"context"
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder エンハンスドデータベースのビルダー
type Builder struct {
	ID types.ID

	Name         string
	Description  string
	Tags         types.Tags
	IconID       types.ID
	DatabaseName string
	Password     string
	SettingsHash string
	Client       sacloud.EnhancedDBAPI
}

func (b *Builder) Build(ctx context.Context) (*sacloud.EnhancedDB, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*sacloud.EnhancedDB, error) {
	created, err := b.Client.Create(ctx, &sacloud.EnhancedDBCreateRequest{
		Name:         b.Name,
		Description:  b.Description,
		Tags:         b.Tags,
		IconID:       b.IconID,
		DatabaseName: b.DatabaseName,
	})
	if err != nil {
		return nil, err
	}

	return created, b.Client.SetPassword(ctx, created.ID, &sacloud.EnhancedDBSetPasswordRequest{
		Password: b.Password,
	})
}

func (b *Builder) update(ctx context.Context) (*sacloud.EnhancedDB, error) {
	current, err := b.Client.Read(ctx, b.ID)
	if err != nil {
		return nil, err
	}
	if current.DatabaseName != b.DatabaseName {
		return nil, errors.New("DatabaseName cannot be changed")
	}

	updated, err := b.Client.Update(ctx, b.ID, &sacloud.EnhancedDBUpdateRequest{
		Name:         b.Name,
		Description:  b.Description,
		Tags:         b.Tags,
		IconID:       b.IconID,
		SettingsHash: b.SettingsHash,
	})
	if err != nil {
		return nil, err
	}

	if b.Password != "" {
		err := b.Client.SetPassword(ctx, updated.ID, &sacloud.EnhancedDBSetPasswordRequest{
			Password: b.Password,
		})
		if err != nil {
			return nil, err
		}
	}

	return updated, nil
}
