package fake

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *CDROMOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.CDROMFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.CDROM
	for _, res := range results {
		dest := &sacloud.CDROM{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.CDROMFindResult{
		Total:  len(results),
		Count:  len(results),
		From:   0,
		CDROMs: values,
	}, nil
}

// Create is fake implementation
func (o *CDROMOp) Create(ctx context.Context, zone string, param *sacloud.CDROMCreateRequest) (*sacloud.CDROMCreateResult, error) {
	result := &sacloud.CDROM{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Availability = types.Availabilities.Uploading

	s.setCDROM(zone, result)
	return &sacloud.CDROMCreateResult{
		IsOk:  true,
		CDROM: result,
		FTPServer: &sacloud.FTPServer{
			HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
			IPAddress: "192.0.2.1",
			User:      fmt.Sprintf("cdrom%d", result.ID),
			Password:  "password-is-not-a-password",
		},
	}, nil
}

// Read is fake implementation
func (o *CDROMOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.CDROMReadResult, error) {
	value := s.getCDROMByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.CDROM{}
	copySameNameField(value, dest)
	return &sacloud.CDROMReadResult{
		IsOk:  true,
		CDROM: value,
	}, nil
}

// Update is fake implementation
func (o *CDROMOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.CDROMUpdateRequest) (*sacloud.CDROMUpdateResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.CDROM
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.CDROMUpdateResult{
		IsOk:  true,
		CDROM: value,
	}, nil
}

// Delete is fake implementation
func (o *CDROMOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, zone, id)
	return nil
}

// OpenFTP is fake implementation
func (o *CDROMOp) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *sacloud.OpenFTPRequest) (*sacloud.CDROMOpenFTPResult, error) {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.CDROM

	value.SetAvailability(types.Availabilities.Uploading)
	s.setCDROM(zone, value)

	return &sacloud.CDROMOpenFTPResult{
		IsOk: true,
		FTPServer: &sacloud.FTPServer{
			HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
			IPAddress: "192.0.2.1",
			User:      fmt.Sprintf("cdrom%d", id),
			Password:  "password-is-not-a-password",
		},
	}, nil
}

// CloseFTP is fake implementation
func (o *CDROMOp) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	readResult, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	value := readResult.CDROM
	if !value.Availability.IsUploading() {
		value.SetAvailability(types.Availabilities.Available)
	}
	s.setCDROM(zone, value)
	return nil
}
