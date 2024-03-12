package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid todo data")
	}

	return Todo{
		Text: text,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf("Your todo: %v\n", t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
