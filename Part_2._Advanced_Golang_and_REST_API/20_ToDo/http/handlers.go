package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todolist *todo.List
}

func NewHTTPHandlers(todolist *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todolist: todolist,
	}
}

func (h *HTTPHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var dto TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := dto.ValidateForCreate(); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	task := todo.NewTask(dto.Title, dto.Description)
	if err := h.todolist.AddTask(task); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}
	b, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed ti write http response", err)
	}
}

func (h *HTTPHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	mux := mux.Vars(r)
	title := mux["title"]
	task, err := h.todolist.GetTask(title)
	if err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}
	b, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed ti write http response", err)
	}
}

func (h *HTTPHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todolist.GetTasks()
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed ti write http response", err)
	}
}

func (h *HTTPHandlers) GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todolist.GetTasks()
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed ti write http response", err)
	}
}

func (h *HTTPHandlers) GetUnompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todolist.GetTasks()
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed ti write http response", err)
	}
}

func (h *HTTPHandlers) CompleteTask(w http.ResponseWriter, r *http.Request) {
	mux := mux.Vars(r)
	title := mux["title"]
	_, err := h.todolist.GetTask(title)
	if err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}
	if err := h.todolist.CompleteTask(title); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	mux := mux.Vars(r)
	title := mux["title"]
	_, err := h.todolist.GetTask(title)
	if err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
			return
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			return
		}
	}
	if err := h.todolist.DeleteTask(title); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now().String(),
		}
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
