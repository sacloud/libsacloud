package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *ArchiveOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ArchiveFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.Archive
	for _, res := range results {
		dest := &sacloud.Archive{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.ArchiveFindResult{
		Total:    len(results),
		Count:    len(results),
		From:     0,
		Archives: values,
	}, nil
}

// Create is fake implementation
func (o *ArchiveOp) Create(ctx context.Context, zone string, param *sacloud.ArchiveCreateRequest) (*sacloud.ArchiveCreateResult, error) {
	result := &sacloud.Archive{}

	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	if !param.SourceArchiveID.IsEmpty() {
		source, err := o.Read(ctx, zone, param.SourceArchiveID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceArchive is not found")
		}
		result.SourceArchiveAvailability = source.Archive.Availability
	}
	if !param.SourceDiskID.IsEmpty() {
		diskOp := NewDiskOp()
		source, err := diskOp.Read(ctx, zone, param.SourceDiskID)
		if err != nil {
			return nil, newErrorBadRequest(o.key, types.ID(0), "SourceDisk is not found")
		}
		result.SourceDiskAvailability = source.Disk.Availability
	}

	result.DisplayOrder = random(100)
	result.Availability = types.Availabilities.Migrating
	result.DiskPlanID = types.ID(2)
	result.DiskPlanName = "標準プラン"
	result.DiskPlanStorageClass = "iscsi9999"

	s.setArchive(zone, result)

	id := result.ID
	startDiskCopy(o.key, zone, func() (interface{}, error) {
		res, err := o.Read(context.Background(), zone, id)
		if err != nil {
			return nil, err
		}
		return res.Archive, nil
	})

	return &sacloud.ArchiveCreateResult{
		IsOk:    true,
		Archive: result,
	}, nil
}

// CreateBlank is fake implementation
func (o *ArchiveOp) CreateBlank(ctx context.Context, zone string, param *sacloud.ArchiveCreateBlankRequest) (*sacloud.ArchiveCreateBlankResult, error) {
	result := &sacloud.Archive{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillScope)

	result.Availability = types.Availabilities.Uploading

	s.setArchive(zone, result)

	return &sacloud.ArchiveCreateBlankResult{
		IsOk:    true,
		Archive: result,
		FTPServer: &sacloud.FTPServer{
			HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
			IPAddress: "192.0.2.1",
			User:      fmt.Sprintf("archive%d", result.ID),
			Password:  "password-is-not-a-password",
		},
	}, nil
}

// Read is fake implementation
func (o *ArchiveOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.ArchiveReadResult, error) {
	value := s.getArchiveByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.Archive{}
	copySameNameField(value, dest)
	return &sacloud.ArchiveReadResult{
		IsOk:    true,
		Archive: dest,
	}, nil
}

// Update is fake implementation
func (o *ArchiveOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.ArchiveUpdateRequest) (*sacloud.ArchiveUpdateResult, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value.Archive)
	fill(value.Archive, fillModifiedAt)
	return &sacloud.ArchiveUpdateResult{
		IsOk:    true,
		Archive: value.Archive,
	}, nil
}

// Delete is fake implementation
func (o *ArchiveOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// OpenFTP is fake implementation
func (o *ArchiveOp) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *sacloud.OpenFTPRequest) (*sacloud.ArchiveOpenFTPResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Archive

	value.SetAvailability(types.Availabilities.Uploading)
	s.setArchive(zone, value)

	return &sacloud.ArchiveOpenFTPResult{
		IsOk: true,
		FTPServer: &sacloud.FTPServer{
			HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
			IPAddress: "192.0.2.1",
			User:      fmt.Sprintf("archive%d", id),
			Password:  "password-is-not-a-password",
		},
	}, nil
}

// CloseFTP is fake implementation
func (o *ArchiveOp) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.Archive

	if !value.Availability.IsUploading() {
		value.SetAvailability(types.Availabilities.Available)
	}
	s.setArchive(zone, value)
	return nil
}
