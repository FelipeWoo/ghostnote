package storage

import (
	"fmt"
	"os"
	"path/filepath"

	d "ghostnote/internal/domain"
)

type NoteRepository struct {
	basePath string
}

func NewNoteRepository(basePath string) *NoteRepository {
	return &NoteRepository{basePath: basePath}
}

func (r *NoteRepository) Save(note *d.Note) error {
	filename := filepath.Join(r.basePath, note.ID+".md")

	content := buildMarkdown(note)

	return os.WriteFile(filename, []byte(content), 0644)
}

func buildMarkdown(n *d.Note) string {
	return fmt.Sprintf(`---
id: %s
title: %s
tags: %s
links: %s
created_at: %d
updated_at: %d
---

%s
`,
		n.ID,
		n.Title,
		n.Tags,
		n.Links,
		n.CreatedAt,
		n.UpdatedAt,
		n.Payload,
	)
}
