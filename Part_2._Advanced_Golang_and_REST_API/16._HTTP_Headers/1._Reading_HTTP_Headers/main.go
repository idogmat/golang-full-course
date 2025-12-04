package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("Название хеддера:", k, "-- значение, лежащее в хеддере:", v)
	}

	fmt.Println("payHandler закончил своё выполнение.")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// prepare message
	message := make(map[string]string)
	message["hello"] = "Hello, JSON!"

	// marshal the map to JSON
	// data, err := json.Marshal(message)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// w.Write(data)
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/json", jsonHandler)
	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
