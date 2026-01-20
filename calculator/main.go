package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	http.ListenAndServe(":8080", nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	a, b := q.Get("a"), q.Get("b")

	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)

	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "Invalid input")
		return
	}

	result := strconv.Itoa(num1 + num2)
	fmt.Fprintf(w, result)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	a, b := q.Get("a"), q.Get("b")

	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)

	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "Invalid input")
		return
	}

	result := strconv.Itoa(num1 * num2)
	fmt.Fprintf(w, result)
}
