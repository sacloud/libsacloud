package sacloud

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoteUpdateRequest_ToNaked(t *testing.T) {

	expects := []struct {
		model *NoteUpdateRequest
		naked *NakedNote
	}{
		{
			model: &NoteUpdateRequest{
				Name:    "test",
				Tags:    []string{"tag1", "tag2"},
				IconID:  2,
				Class:   "shell",
				Content: "content",
			},
			naked: &NakedNote{
				Name: "test",
				Tags: []string{"tag1", "tag2"},
				Icon: &NakedIcon{
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
		naked *NakedNote
	}{
		{
			model: &NoteUpdateRequest{
				Name:    "test",
				Tags:    []string{"tag1", "tag2"},
				IconID:  2,
				Class:   "shell",
				Content: "content",
			},
			naked: &NakedNote{
				Name: "test",
				Tags: []string{"tag1", "tag2"},
				Icon: &NakedIcon{
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
