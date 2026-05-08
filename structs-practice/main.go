package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
)

func main() {
	theNote, err := getNote()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Note Title : %s Note Content %s\n", theNote.GetNoteTitle(), theNote.GetNoteContent())

	err = theNote.Save() // Implementation save the struct json to the file is written inside the encapsulated Save() method
	if err != nil {
		fmt.Println("Saving to file fails due to errors", err)
	}
}

func getNote() (*note.Note, error) {
	noteTitle, noteTitleError := getUserInput("Note Title : ")
	noteContent, noteContentError := getUserInput("Note Content : ")
	if noteTitleError != nil || noteContentError != nil {
		fmt.Println("Either noteTitle or noteContent is nil")
		panic(noteContentError)
	}
	return note.New(noteTitle, noteContent)
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	inputValue, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	inputValue = strings.TrimSpace(inputValue)
	if inputValue == "" {
		return "", errors.New("Invalid Input !! ")
	}
	return inputValue, nil
}
