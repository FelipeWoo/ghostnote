package cli

import (
	"ghostnote/internal/logger"
	s "ghostnote/internal/services"
)

type NoteCommand struct {
	noteService *s.NoteService
}

func NewNoteCommand(noteService *s.NoteService) *NoteCommand {
	return &NoteCommand{
		noteService: noteService,
	}
}

func (c *NoteCommand) Create(title, payload string, tags, links []string) error {
	note, err := c.noteService.Create(title, payload, tags, links)
	if err != nil {
		return err
	}

	logger.Info("note created: %s\n", note.ID)
	return nil
}
