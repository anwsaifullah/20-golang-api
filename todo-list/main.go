package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", handler)

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":

	}
}
