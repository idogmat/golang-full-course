package http

import (
	"encoding/json"
	"errors"
)

type TaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("Title is required")
	}

	if t.Description == "" {
		return errors.New("Description is required")
	}
	return nil
}

type ErrorDto struct {
	Message string `json:"message"`
	Time    string `json:"string"`
}

func (e ErrorDto) ToString() string {
	b, err := json.MarshalIndent(e, "", "")
	if err != nil {
		panic(err)
	}
	return string(b)
}
