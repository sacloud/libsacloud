package client

import (
	"context"
	"log"

	"github.com/sacloud/libsacloud/v2/grpc/proto"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"google.golang.org/grpc"
)

type ServerOp struct {
	Addr string
}

func (s *ServerOp) client() (proto.ServerOpClient, func()) {
	conn, err := grpc.Dial(s.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}

	return proto.NewServerOpClient(conn), func() { conn.Close() }
}

func (s *ServerOp) Boot(ctx context.Context, zone string, id types.ID) error {
	client, cleanup := s.client()
	defer cleanup()
	_, err := client.Boot(ctx, &proto.ServerBootRequest{Zone: zone, Id: id.Int64()})
	return err
}

func (s *ServerOp) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *sacloud.ShutdownOption) error {
	client, cleanup := s.client()
	defer cleanup()
	_, err := client.Shutdown(ctx, &proto.ServerShutdownRequest{
		Zone:   zone,
		Id:     id.Int64(),
		Option: &proto.ShutdownOption{Force: shutdownOption.Force},
	})
	return err
}

func (s *ServerOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerFindResult, error) {
	return nil, nil
}
func (s *ServerOp) Create(ctx context.Context, zone string, param *sacloud.ServerCreateRequest) (*sacloud.Server, error) {
	return nil, nil
}
func (s *ServerOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Server, error) {
	return nil, nil
}
func (s *ServerOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ServerUpdateRequest) (*sacloud.Server, error) {
	return nil, nil
}
func (s *ServerOp) Delete(ctx context.Context, zone string, id types.ID) error { return nil }
func (s *ServerOp) DeleteWithDisks(ctx context.Context, zone string, id types.ID, disks *sacloud.ServerDeleteWithDisksRequest) error {
	return nil
}
func (s *ServerOp) ChangePlan(ctx context.Context, zone string, id types.ID, plan *sacloud.ServerChangePlanRequest) (*sacloud.Server, error) {
	return nil, nil
}
func (s *ServerOp) InsertCDROM(ctx context.Context, zone string, id types.ID, insertParam *sacloud.InsertCDROMRequest) error {
	return nil
}
func (s *ServerOp) EjectCDROM(ctx context.Context, zone string, id types.ID, ejectParam *sacloud.EjectCDROMRequest) error {
	return nil
}

func (s *ServerOp) Reset(ctx context.Context, zone string, id types.ID) error {
	return nil
}
func (s *ServerOp) SendKey(ctx context.Context, zone string, id types.ID, keyboardParam *sacloud.SendKeyRequest) error {
	return nil
}
func (s *ServerOp) GetVNCProxy(ctx context.Context, zone string, id types.ID) (*sacloud.VNCProxyInfo, error) {
	return nil, nil
}
func (s *ServerOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.CPUTimeActivity, error) {
	return nil, nil
}
