package tests

import "time"

// bookProgressData defines the structure of an API book progress retrieval. Copy for test
type bookProgressData struct {
	UserID       string    `json:"user_id,omitempty"`
	BookID       string    `json:"book_id,omitempty"`
	LibraryID    string    `json:"library_id,omitempty"`
	Progress     string    `json:"progress"`
	LastReadDate time.Time `json:"last_read_date,omitempty"`
}
