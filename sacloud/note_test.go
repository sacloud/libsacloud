package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud-v2/sacloud/naked"
	"github.com/stretchr/testify/require"
)

func TestNoteUpdateRequest_ToNaked(t *testing.T) {

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
		naked, err := expect.model.ToNaked()
		require.NoError(t, err)
		require.Equal(t, expect.naked, naked)
	}

}

func TestNoteUpdateRequest_ParseNaked(t *testing.T) {

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
		err := model.ParseNaked(expect.naked)
		require.NoError(t, err)
		require.Equal(t, expect.model, model)
	}

}
