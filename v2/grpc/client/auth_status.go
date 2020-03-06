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

package client

import (
	"context"
	"log"

	grpcSacloud "github.com/sacloud/libsacloud/v2/grpc/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"google.golang.org/grpc"
)

type AuthStatusOp struct {
	Addr string
}

func (s *AuthStatusOp) client() (grpcSacloud.AuthStatusAPIClient, func()) {
	conn, err := grpc.Dial(s.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}

	return grpcSacloud.NewAuthStatusAPIClient(conn), func() { conn.Close() }
}

func (s *AuthStatusOp) Read(ctx context.Context) (*sacloud.AuthStatus, error) {
	client, cleanup := s.client()
	defer cleanup()
	res, err := client.Read(ctx, &grpcSacloud.AuthStatusReadRequest{})
	if err != nil {
		return nil, err
	}
	auth := res.AuthStatus
	return &sacloud.AuthStatus{
		AccountID:    types.ID(auth.AccountID),
		AccountName:  auth.AccountName,
		AccountCode:  auth.AccountCode,
		AccountClass: auth.AccountClass,
		MemberCode:   auth.MemberCode,
		MemberClass:  auth.MemberClass,
		AuthClass:          types.EAuthClass(auth.AuthClass.String()),
		AuthMethod:         types.EAuthMethod(auth.AuthMethod.String()),
		IsAPIKey: auth.IsAPIKey,
		ExternalPermission: types.ExternalPermission(auth.ExternalPermission),
		OperationPenalty:   types.EOperationPenalty(auth.OperationPenalty.String()),
		Permission:         types.EPermission(auth.Permission.String()),
	}, nil

}
