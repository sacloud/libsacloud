package fake

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *DiskOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.DiskFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Disk
	for _, res := range results {
		dest := &sacloud.Disk{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.DiskFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Disks: values,
	}, nil
}

// Create is fake implementation
func (o *DiskOp) Create(ctx context.Context, zone string, param *sacloud.DiskCreateRequest) (*sacloud.DiskCreateResult, error) {
	result := &sacloud.Disk{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillDiskPlan)

	if result.Connection == types.EDiskConnection("") {
		result.Connection = types.DiskConnections.VirtIO
	}
	if !param.SourceArchiveID.IsEmpty() {
		archiveOp := NewArchiveOp()
		source, err := archiveOp.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Archive.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Disk.Availability
	}

	s.setDisk(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		readResult, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return readResult.Disk, nil
	})

	return &sacloud.DiskCreateResult{
		IsOk: true,
		Disk: result,
	}, nil
}

// CreateDistantly is fake implementation
func (o *DiskOp) CreateDistantly(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, distantFrom []types.ID) (*sacloud.DiskCreateDistantlyResult, error) {
	res, err := o.Create(ctx, zone, createParam)
	if err != nil {
		return nil, err
	}
	return &sacloud.DiskCreateDistantlyResult{
		IsOk: res.IsOk,
		Disk: res.Disk,
	}, nil
}

// Config is fake implementation
func (o *DiskOp) Config(ctx context.Context, zone string, id types.ID, edit *sacloud.DiskEditRequest) error {
	// TODO ディスクに接続されたサーバのIDを拾ってInterfaces[0].UserSubnet.DefaultRoute/UserIPAddressなども書き換えた方がいいかも?
	return nil
}

// CreateWithConfig is fake implementation
func (o *DiskOp) CreateWithConfig(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, editParam *sacloud.DiskEditRequest, bootAtAvailable bool) (*sacloud.DiskCreateWithConfigResult, error) {
	// check
	if !createParam.ServerID.IsEmpty() {
		serverOp := NewServerOp()
		_, err := serverOp.Read(ctx, zone, createParam.ServerID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), fmt.Sprintf("Server %s is not found", createParam.ServerID))
		}
	}

	createResult, err := o.Create(ctx, zone, createParam)
	if err != nil {
		return nil, err
	}
	result := createResult.Disk

	if !createParam.ServerID.IsEmpty() && bootAtAvailable {
		waiter := sacloud.WaiterForReady(func() (interface{}, error) {
			res, err := o.Read(ctx, zone, result.ID)
			if err != nil {
				return nil, err
			}
			return res.Disk, nil
		})
		res, err := waiter.WaitForState(ctx)
		if err != nil {
			return nil, err
		}
		result = res.(*sacloud.Disk)

		// boot server
		serverOp := NewServerOp()
		if err := serverOp.Boot(ctx, zone, createParam.ServerID); err != nil {
			return nil, err
		}
	}
	return &sacloud.DiskCreateWithConfigResult{
		IsOk: true,
		Disk: result,
	}, nil
}

// CreateWithConfigDistantly is fake implementation
func (o *DiskOp) CreateWithConfigDistantly(ctx context.Context, zone string, createParam *sacloud.DiskCreateRequest, editParam *sacloud.DiskEditRequest, bootAtAvailable bool, distantFrom []types.ID) (*sacloud.DiskCreateWithConfigDistantlyResult, error) {
	res, err := o.CreateWithConfig(ctx, zone, createParam, editParam, bootAtAvailable)
	if err != nil {
		return nil, err
	}
	return &sacloud.DiskCreateWithConfigDistantlyResult{
		IsOk: res.IsOk,
		Disk: res.Disk,
	}, nil
}

// ToBlank is fake implementation
func (o *DiskOp) ToBlank(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Disk

	value.SourceArchiveID = types.ID(0)
	value.SourceArchiveAvailability = types.Availabilities.Unknown
	value.SourceDiskID = types.ID(0)
	value.SourceDiskAvailability = types.Availabilities.Unknown

	s.setDisk(zone, value)
	return nil
}

// ResizePartition is fake implementation
func (o *DiskOp) ResizePartition(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	return nil
}

// ConnectToServer is fake implementation
func (o *DiskOp) ConnectToServer(ctx context.Context, zone string, id types.ID, serverID types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Disk

	serverOp := NewServerOp()
	serverReadResult, err := serverOp.Read(ctx, zone, serverID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", serverID))
	}
	server := serverReadResult.Server

	for _, connected := range server.Disks {
		if connected.ID == value.ID {
			return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is already connected to Server[%d]", id, serverID))
		}
	}

	// TODO とりあえず同時実行制御は考慮しない。更新対象リソースが増えるようであれば実装方法を考える

	server.Disks = append(server.Disks, value)
	s.setServer(zone, server)
	value.ServerID = serverID
	s.setDisk(zone, value)

	return nil
}

// DisconnectFromServer is fake implementation
func (o *DiskOp) DisconnectFromServer(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Disk

	if value.ServerID.IsEmpty() {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Disk[%d] is not connected to Server", id))
	}

	serverOp := NewServerOp()
	serverReadResult, err := serverOp.Read(ctx, zone, value.ServerID)
	if err != nil {
		return newErrorBadRequest(o.key, id, fmt.Sprintf("Server[%d] is not exists", value.ServerID))
	}
	server := serverReadResult.Server

	var disks []*sacloud.Disk
	for _, connected := range server.Disks {
		if connected.ID != value.ID {
			disks = append(disks, connected)
		}
	}
	if len(disks) == len(server.Disks) {
		return newInternalServerError(o.key, id, fmt.Sprintf("Disk[%d] is not found on server's connected disks", id))
	}

	server.Disks = disks
	s.setServer(zone, server)
	value.ServerID = types.ID(0)
	s.setDisk(zone, value)

	return nil
}

// InstallDistantFrom is fake implementation
func (o *DiskOp) InstallDistantFrom(ctx context.Context, zone string, id types.ID, installParam *sacloud.DiskInstallRequest, distantFrom []types.ID) (*sacloud.DiskInstallDistantFromResult, error) {
	res, err := o.Install(ctx, zone, id, installParam)
	if err != nil {
		return nil, err
	}
	return &sacloud.DiskInstallDistantFromResult{
		IsOk: res.IsOk,
		Disk: res.Disk,
	}, nil
}

// Install is fake implementation
func (o *DiskOp) Install(ctx context.Context, zone string, id types.ID, installParam *sacloud.DiskInstallRequest) (*sacloud.DiskInstallResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Disk

	fill(value, fillDiskPlan)
	s.setDisk(zone, value)
	return &sacloud.DiskInstallResult{
		IsOk: true,
		Disk: value,
	}, nil
}

// Read is fake implementation
func (o *DiskOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.DiskReadResult, error) {
	value := s.getDiskByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Disk{}
	copySameNameField(value, dest)
	return &sacloud.DiskReadResult{
		IsOk: true,
		Disk: dest,
	}, nil
}

// Update is fake implementation
func (o *DiskOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.DiskUpdateRequest) (*sacloud.DiskUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Disk
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.DiskUpdateResult{
		IsOk: true,
		Disk: value,
	}, nil
}

// Delete is fake implementation
func (o *DiskOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// Monitor is fake implementation
func (o *DiskOp) Monitor(ctx context.Context, zone string, id types.ID, condition *sacloud.MonitorCondition) (*sacloud.DiskMonitorResult, error) {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	now := time.Now().Truncate(time.Second)
	m := now.Minute() % 5
	if m != 0 {
		now.Add(time.Duration(m) * time.Minute)
	}

	res := &sacloud.DiskActivity{}
	for i := 0; i < 5; i++ {
		res.Values = append(res.Values, &sacloud.MonitorDiskValue{
			Time:  now.Add(time.Duration(i*-5) * time.Minute),
			Read:  float64(random(1000)),
			Write: float64(random(1000)),
		})
	}

	return &sacloud.DiskMonitorResult{
		IsOk: true,
		Data: res,
	}, nil
}
