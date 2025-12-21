package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HttpHandlers *HTTPHandlers
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HttpHandlers: handlers,
	}
}

func (s *HTTPServer) Start() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.HttpHandlers.CreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.HttpHandlers.GetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.HttpHandlers.GetCompletedTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.HttpHandlers.GetUnompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.HttpHandlers.GetTasks)
	router.Path("/tasks/{title}").Methods("PUTCH").HandlerFunc(s.HttpHandlers.CompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.HttpHandlers.DeleteTask)
	fmt.Println("Started server")
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
