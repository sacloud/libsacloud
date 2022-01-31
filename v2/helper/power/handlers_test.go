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

package power

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func Test_serverHandler_boot(t *testing.T) {
	type fields struct {
		variables []string
	}
	tests := []struct {
		name        string
		fields      fields
		wantErr     bool
		checkCalled func(*dummyServerAPI) bool
	}{
		{
			name: "no cloudConfig",
			fields: fields{
				variables: nil,
			},
			wantErr: false,
			checkCalled: func(d *dummyServerAPI) bool {
				return d.bootIsCalled
			},
		},
		{
			name: "cloudConfig is an empty string",
			fields: fields{
				variables: []string{""},
			},
			wantErr: false,
			checkCalled: func(d *dummyServerAPI) bool {
				return d.bootIsCalled
			},
		},
		{
			name: "cloudConfig is a non-empty string",
			fields: fields{
				variables: []string{"string"},
			},
			wantErr: false,
			checkCalled: func(d *dummyServerAPI) bool {
				return d.bootWithVariablesIsCalled
			},
		},
		{
			name: "cloudConfig is a slice of empty string",
			fields: fields{
				variables: []string{"", ""},
			},
			wantErr: false,
			checkCalled: func(d *dummyServerAPI) bool {
				return d.bootIsCalled
			},
		},
		{
			name: "cloudConfig is a slice of string",
			fields: fields{
				variables: []string{"string1", "string2"},
			},
			wantErr: false,
			checkCalled: func(d *dummyServerAPI) bool {
				return d.bootWithVariablesIsCalled
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &dummyServerAPI{}
			h := &serverHandler{
				ctx:       context.Background(),
				client:    client,
				zone:      "is1a",
				id:        types.ID(1),
				variables: tt.fields.variables,
			}
			if err := h.boot(); (err != nil) != tt.wantErr {
				t.Errorf("boot() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.checkCalled(client) {
				t.Errorf("unexpected function was called")
			}
		})
	}
}

type dummyServerAPI struct {
	bootIsCalled              bool
	bootWithVariablesIsCalled bool
}

func (d *dummyServerAPI) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Server, error) {
	return nil, nil
}

func (d *dummyServerAPI) Boot(ctx context.Context, zone string, id types.ID) error {
	d.bootIsCalled = true
	return nil
}

func (d *dummyServerAPI) BootWithVariables(ctx context.Context, zone string, id types.ID, param *sacloud.ServerBootVariables) error {
	d.bootWithVariablesIsCalled = true
	return nil
}

func (d *dummyServerAPI) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	return nil
}
