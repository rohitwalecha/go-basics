package note

import "errors"

type Note struct {
	noteTile    string
	noteContent string
}

func (note *Note) setNoteTitle(noteTitle string) {
	note.noteTile = noteTitle
}

func (note *Note) getNoteTitle() string {
	return note.noteTile
}

func (note *Note) setNoteContent(noteContent string) {
	note.noteContent = noteContent
}

func (note *Note) getNoteContent() string {
	return note.noteTile
}

func New(noteTile string, noteContent string) (*Note, error) {
	if noteTile == "" || noteContent == "" {
		return nil, errors.New("noteTitle and noteContent both are mandatory !!")
	}
	return &Note{
		noteTile, noteContent,
	}, nil
}
