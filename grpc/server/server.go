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
