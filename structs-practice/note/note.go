package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Note struct {
	NoteTitle   string `json:"note_title"`
	NoteContent string `json:"note_content"`
}

func (note *Note) SetNoteTitle(noteTitle string) {
	note.NoteTitle = noteTitle
}

func (note Note) GetNoteTitle() string {
	return note.NoteTitle
}

func (note *Note) SetNoteContent(noteContent string) {
	note.NoteContent = noteContent
}

func (note Note) GetNoteContent() string {
	return note.NoteContent
}

func New(noteTitle string, noteContent string) (*Note, error) {
	if noteTitle == "" || noteContent == "" {
		return nil, errors.New("noteTitle and noteContent both are mandatory !!")
	}
	return &Note{
		NoteTitle:   noteTitle,
		NoteContent: noteContent,
	}, nil
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.NoteTitle, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note) // Marshal function returns an Error object as well so check and return it in case the parsing/marshal fails
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // Same with this method this also returns an error object
}
