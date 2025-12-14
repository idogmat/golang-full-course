package main

import (
	"restapi/http"
	"restapi/todo"
)

func main() {

	todoList := todo.NewList()
	handlers := http.NewHTTPHandlers(todoList)
	server := http.NewHTTPServer(handlers)
	server.Start()
}
