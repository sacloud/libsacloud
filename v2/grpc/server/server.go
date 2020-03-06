// Copyright 2016-2020 The Libsacloud Authors
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

package server

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/sacloud/libsacloud/v2/grpc/proto"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	context "golang.org/x/net/context"
)

type ServerOpService struct {
}

func (s *ServerOpService) Boot(ctx context.Context, in *proto.ServerBootRequest) (*empty.Empty, error) {
	caller, err := sacloud.NewClientFromEnv()
	if err != nil {
		return &empty.Empty{}, err
	}

	if err := sacloud.NewServerOp(caller).Boot(ctx, in.Zone, types.ID(in.Id)); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *ServerOpService) Shutdown(ctx context.Context, in *proto.ServerShutdownRequest) (*empty.Empty, error) {
	caller, err := sacloud.NewClientFromEnv()
	if err != nil {
		return &empty.Empty{}, err
	}

	if err := sacloud.NewServerOp(caller).Shutdown(ctx, in.Zone, types.ID(in.Id), &sacloud.ShutdownOption{Force: in.Option.Force}); err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
