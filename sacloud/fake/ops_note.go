package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *NoteOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) ([]*sacloud.Note, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Note
	for _, res := range results {
		dest := &sacloud.Note{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return values, nil
}

// Create is fake implementation
func (o *NoteOp) Create(ctx context.Context, zone string, param *sacloud.NoteCreateRequest) (*sacloud.Note, error) {
	result := &sacloud.Note{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	s.setNote(sacloud.APIDefaultZone, result)
	return result, nil
}

// Read is fake implementation
func (o *NoteOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.Note, error) {
	value := s.getNoteByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Note{}
	copySameNameField(value, dest)
	return dest, nil
}

// Update is fake implementation
func (o *NoteOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.NoteUpdateRequest) (*sacloud.Note, error) {
	value, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return value, nil
}

// Delete is fake implementation
func (o *NoteOp) Delete(ctx context.Context, zone string, id types.ID) error {
	_, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return err
	}
	s.delete(o.key, sacloud.APIDefaultZone, id)
	return nil
}
