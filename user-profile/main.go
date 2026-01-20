package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}

var id = 1

func main() {
	http.HandleFunc("/users/", userByIDHandler)
	http.HandleFunc("/users", addUserHandler)

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(path)
	if r.Method == "GET" {
		for _, user := range users {
			if user.ID == id {
				json.NewEncoder(w).Encode(user)
			}
		}
	}
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Invalid JSON",
			})

			return
		}

		user.ID = id
		users = append(users, user)
		id++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
