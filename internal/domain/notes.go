package domain

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Note struct {
	ID        string   `json:"id" yaml:"id"`
	Title     string   `json:"title" yaml:"title"`
	Tags      []string `json:"tags" yaml:"tags"`
	Links     []string `json:"links" yaml:"links"`
	Payload   string   `json:"payload" yaml:"payload"`
	CreatedAt int64    `json:"created_at" yaml:"created_at"`
	UpdatedAt int64    `json:"updated_at" yaml:"updated_at"`
}

// Constructor
func NewNote(title, payload string, tags, links []string) *Note {
	now := now()

	return &Note{
		ID:        newID(),
		Title:     title,
		Tags:      tags,
		Links:     links,
		Payload:   payload,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Method (mutates itself)
func (n *Note) Touch() {
	n.UpdatedAt = now()
}

// -------- helpers --------

func now() int64 {
	return time.Now().UTC().UnixMilli()
}

func newID() string {
	id, _ := gonanoid.New()
	return id
}
