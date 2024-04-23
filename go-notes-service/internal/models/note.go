package models

import "time"

type Note struct {
    ID           int       `json:"id"`
    CreatedAt    time.Time `json:"created_at"`
    AuthorID     int       `json:"author_id"`
    Text         string    `json:"text"`
    IsPublic     bool      `json:"is_public"`
}

func NewNote(id, authorID int, text string, isPublic bool) *Note {
    return &Note{
        ID:           id,
        CreatedAt:    time.Now(),
        AuthorID:     authorID,
        Text:         text,
        IsPublic:     isPublic,
    }
}
