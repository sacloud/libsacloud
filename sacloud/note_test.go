package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/stretchr/testify/require"
)

func TestNoteUpdateRequest_Validate(t *testing.T) {
	expects := []struct {
		model    *NoteUpdateRequest
		hasError bool
	}{
		{
			model:    &NoteUpdateRequest{},
			hasError: true,
		},
		{
			model: &NoteUpdateRequest{
				Name: "foo",
			},
			hasError: false,
		},
	}

	for _, tc := range expects {
		err := tc.model.Validate()
		if tc.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestNoteUpdateRequest_convertTo(t *testing.T) {

	expects := []struct {
		model *NoteUpdateRequest
		naked *naked.Note
	}{
		{
			model: &NoteUpdateRequest{
				Name:    "test",
				Tags:    []string{"tag1", "tag2"},
				IconID:  2,
				Class:   "shell",
				Content: "content",
			},
			naked: &naked.Note{
				Name: "test",
				Tags: []string{"tag1", "tag2"},
				Icon: &naked.Icon{
					ID: 2,
				},
				Class:   "shell",
				Content: "content",
			},
		},
	}

	for _, expect := range expects {
		naked, err := expect.model.convertTo()
		require.NoError(t, err)
		require.Equal(t, expect.naked, naked)
	}

}

func TestNoteUpdateRequest_convertFrom(t *testing.T) {

	expects := []struct {
		model *NoteUpdateRequest
		naked *naked.Note
	}{
		{
			model: &NoteUpdateRequest{
				Name:    "test",
				Tags:    []string{"tag1", "tag2"},
				IconID:  2,
				Class:   "shell",
				Content: "content",
			},
			naked: &naked.Note{
				Name: "test",
				Tags: []string{"tag1", "tag2"},
				Icon: &naked.Icon{
					ID: 2,
				},
				Class:   "shell",
				Content: "content",
			},
		},
	}

	for _, expect := range expects {
		model := &NoteUpdateRequest{}
		err := model.convertFrom(expect.naked)
		require.NoError(t, err)
		require.Equal(t, expect.model, model)
	}

}
