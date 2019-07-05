package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Find is fake implementation
func (o *NoteOp) Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.NoteFindResult, error) {
	results, _ := find(o.key, sacloud.APIDefaultZone, conditions)
	var values []*sacloud.Note
	for _, res := range results {
		dest := &sacloud.Note{}
		copySameNameField(res, dest)
		values = append(values, dest)
	}
	return &sacloud.NoteFindResult{
		Total: len(results),
		Count: len(results),
		From:  0,
		Notes: values,
	}, nil
}

// Create is fake implementation
func (o *NoteOp) Create(ctx context.Context, zone string, param *sacloud.NoteCreateRequest) (*sacloud.NoteCreateResult, error) {
	result := &sacloud.Note{}
	copySameNameField(param, result)
	fill(result, fillID, fillCreatedAt, fillAvailability, fillScope)
	s.setNote(sacloud.APIDefaultZone, result)
	return &sacloud.NoteCreateResult{
		IsOk: true,
		Note: result,
	}, nil
}

// Read is fake implementation
func (o *NoteOp) Read(ctx context.Context, zone string, id types.ID) (*sacloud.NoteReadResult, error) {
	value := s.getNoteByID(sacloud.APIDefaultZone, id)
	if value == nil {
		return nil, newErrorNotFound(o.key, id)
	}

	dest := &sacloud.Note{}
	copySameNameField(value, dest)
	return &sacloud.NoteReadResult{
		IsOk: true,
		Note: dest,
	}, nil
}

// Update is fake implementation
func (o *NoteOp) Update(ctx context.Context, zone string, id types.ID, param *sacloud.NoteUpdateRequest) (*sacloud.NoteUpdateResult, error) {
	readResult, err := o.Read(ctx, sacloud.APIDefaultZone, id)
	if err != nil {
		return nil, err
	}
	value := readResult.Note

	copySameNameField(param, value)
	fill(value, fillModifiedAt)
	return &sacloud.NoteUpdateResult{
		IsOk: true,
		Note: value,
	}, nil
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
