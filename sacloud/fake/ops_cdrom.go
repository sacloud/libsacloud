package fake

import (
	"context"
	"fmt"

	"github.com/imdario/mergo"
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
func (o *CDROMOp) Create(ctx context.Context, zone string, param *sacloud.CDROMCreateRequest) (*sacloud.CDROM, *sacloud.FTPServer, error) {
	result := &sacloud.CDROM{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	result.Availability = types.Availabilities.Uploading

	putCDROM(zone, result)
	return result, &sacloud.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("cdrom%d", result.ID),
		Password:  "password-is-not-a-password",
	}, nil
}

// Read is fake implementation
func (o *CDROMOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.CDROM, error) {
	value := getCDROMByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.CDROM{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *CDROMOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.CDROMUpdateRequest) (*sacloud.CDROM, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	putCDROM(zone, value)
	return value, nil
}

// Patch is fake implementation
func (o *CDROMOp) Patch(ctx context.Context, zone string, id types.ID, param *sacloud.CDROMPatchRequest) (*sacloud.CDROM, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	patchParam := make(map[string]interface{})
	if err := mergo.Map(&patchParam, value); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(&patchParam, param); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	if err := mergo.Map(param, &patchParam); err != nil {
		return nil, fmt.Errorf("patch is failed: %s", err)
	}
	copySameNameField(param, value)

	if param.PatchEmptyToDescription {
		value.Description = ""
	}
	if param.PatchEmptyToTags {
		value.Tags = nil
	}
	if param.PatchEmptyToIconID {
		value.IconID = types.ID(int64(0))
	}

	putCDROM(zone, value)
	return value, nil
}

// Delete is fake implementation
func (o *CDROMOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	ds().Delete(o.key, zone, id)
	return nil
}

// OpenFTP is fake implementation
func (o *CDROMOp) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *sacloud.OpenFTPRequest) (*sacloud.FTPServer, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	value.SetAvailability(types.Availabilities.Uploading)
	putCDROM(zone, value)

	return &sacloud.FTPServer{
		HostName:  fmt.Sprintf("sac-%s-ftp.example.jp", zone),
		IPAddress: "192.0.2.1",
		User:      fmt.Sprintf("cdrom%d", id),
		Password:  "password-is-not-a-password",
	}, nil
}

// CloseFTP is fake implementation
func (o *CDROMOp) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}
	if !value.Availability.IsUploading() {
		value.SetAvailability(types.Availabilities.Available)
	}
	putCDROM(zone, value)
	return nil
}
