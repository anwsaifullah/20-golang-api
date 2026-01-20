package main

import (
	"net/http"
	// "encoding/json"
)

func main() {
	http.HandleFunc("/weather", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// query := r.URL.Query().Get("city")

}
