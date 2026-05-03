package main

import (
	"errors"
	"fmt"

	"example.com/structs-practice/note"
)

func main() {
	theNote, err := getNote()
	if err != nil {
		panic(err)
	}
	fmt.Println(theNote)
}

func getNote() (*note.Note, error) {
	noteTitle, err := getUserInput("Note Title : ")
	noteContent, err := getUserInput("Note Content : ")
	if err != nil {
		panic(err)
	}
	return note.New(noteTitle, noteContent)
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var inputValue string
	fmt.Scanln(&inputValue)
	if inputValue == "" {
		return "", errors.New("Invalid Input !! ")
	}
	return inputValue, nil
}
