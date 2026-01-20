package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type Post struct {
	ID   int
	Post string
}

var posts = []Post{
	// {ID: 1, Post: "Once upon a time"},
	// {ID: 2, Post: "Once upon a time, there lived a tiger."},
}

var nextId int = 0

func main() {
	http.HandleFunc("/posts", postHandler)
	http.HandleFunc("/posts/", postByIDHandler)

	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(posts)
	case "POST":
		var newPost Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			return
		}

		newPost.ID = nextId
		posts = append(posts, newPost)
		nextId++

		json.NewEncoder(w).Encode(map[string]string{"message": "Post added successfully"})
	}
}

func postByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	postId := strings.TrimPrefix(r.URL.Path, "/posts/")
	id, err := strconv.Atoi(postId)
	if err != nil {
		return
	}

	switch r.Method {
	case "GET":
		for _, post := range posts {
			if post.ID == id {
				json.NewEncoder(w).Encode(post)
				return
			}
		}
	case "PUT":
		var updatedPost Post
		json.NewDecoder(r.Body).Decode(&updatedPost)

		for i, post := range posts {
			if post.ID == id {
				posts[i].Post = updatedPost.Post
				json.NewEncoder(w).Encode(map[string]string{"message": "Post updated successfully."})
				return
			}
		}
	case "DELETE":
		for i, post := range posts {
			if post.ID == id {
				posts = slices.Delete(posts, i, i+1)
				json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted successfully."})
				return
			}
		}
	}
}
