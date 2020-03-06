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
	grpcSacloud "github.com/sacloud/libsacloud/v2/grpc/sacloud"
	"github.com/sacloud/libsacloud/v2/grpc/sacloud/types"
	"github.com/sacloud/libsacloud/v2/sacloud"
	context "golang.org/x/net/context"
)

type AuthStatusService struct {
}

func (s *AuthStatusService) Read(ctx context.Context, in *grpcSacloud.AuthStatusReadRequest) (*grpcSacloud.AuthStatusReadResult, error) {
	caller, err := sacloud.NewClientFromEnv()
	if err != nil {
		return &grpcSacloud.AuthStatusReadResult{}, err
	}

	auth, err := sacloud.NewAuthStatusOp(caller).Read(ctx)
	if err != nil {
		return &grpcSacloud.AuthStatusReadResult{}, err
	}

	return &grpcSacloud.AuthStatusReadResult{
		AuthStatus: &grpcSacloud.AuthStatus{
			AccountID:          auth.AccountID.Int64(),
			AccountName:        auth.AccountName,
			AccountCode:        auth.AccountCode,
			AccountClass:       auth.AccountClass,
			MemberCode:         auth.MemberCode,
			MemberClass:        auth.MemberClass,
			AuthClass:          types.AuthClass(types.AuthClass_value[string(auth.AuthClass)]),
			AuthMethod:         types.AuthMethod(types.AuthMethod_value[string(auth.AuthMethod)]),
			IsAPIKey:           auth.IsAPIKey,
			ExternalPermission: string(auth.ExternalPermission),
			OperationPenalty:   types.OperationPenalty(types.OperationPenalty_value[string(auth.OperationPenalty)]),
			Permission:         types.APIKeyPermission(types.APIKeyPermission_value[string(auth.Permission)]),
		},
	}, nil
}
