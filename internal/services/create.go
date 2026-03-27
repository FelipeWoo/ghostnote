package services

import (
	d "ghostnote/internal/domain"
)

type NoteRepository interface {
	Save(note *d.Note) error
	//more services for later
	// List() ([]d.Note, error)
	// Get(id string) (*d.Note, error)
}

type NoteService struct {
	repo NoteRepository
}

func NewNoteService(r NoteRepository) *NoteService {
	return &NoteService{repo: r}
}

func (s *NoteService) Create(title, payload string, tags, links []string) (*d.Note, error) {
	note := d.NewNote(title, payload, tags, links)

	err := s.repo.Save(note)
	if err != nil {
		return nil, err
	}
	return note, nil
}
