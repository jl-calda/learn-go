package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/learngo/note"
	"example.com/learngo/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func noteTaking() {
	title, content := getNoteData()
	todoText := getTodoData()

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	outputData(userNote)
	outputData(userTodo)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return err
	}

	fmt.Println("Note and todo saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getNoteInput("Enter the title of the note:")

	content := getNoteInput("Enter the content of the note:")

	return title, content
}

func getTodoData() string {
	text := getNoteInput("Enter the todo text:")
	return text
}

func getNoteInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
