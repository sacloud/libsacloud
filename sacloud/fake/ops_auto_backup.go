package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *AutoBackupOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.AutoBackupFindResult, error) {
	results, _ := find(o.key, zone, conditions)
	var values []*sacloud.AutoBackup
	for _, res := range results {
		dest := &sacloud.AutoBackup{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.AutoBackupFindResult{
		Total:       len(results),
		Count:       len(results),
		From:        0,
		AutoBackups: values,
	}, nil
}

// Create is fake implementation
func (o *AutoBackupOp) Create(ctx context.Context, zone string, param *sacloud.AutoBackupCreateRequest) (*sacloud.AutoBackup, error) {
	result := &sacloud.AutoBackup{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt)

	result.Availability = types.Availabilities.Available
	result.SettingsHash = "settingshash"
	result.AccountID = accountID
	result.ZoneID = zoneIDs[zone]
	result.ZoneName = zone

	s.setAutoBackup(zone, result)
	return result, nil
}

// Read is fake implementation
func (o *AutoBackupOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.AutoBackup, error) {
	value := s.getAutoBackupByID(zone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}
	dest := &sacloud.AutoBackup{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *AutoBackupOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.AutoBackupUpdateRequest) (*sacloud.AutoBackup, error) {
	value, err := o.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)

	return value, nil
}

// Delete is fake implementation
func (o *AutoBackupOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, zone, id)
	if err != nil {
		return err
	}

	s.delete(o.key, zone, id)
	return nil
}
