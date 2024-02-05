package main

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", Hangman)
	http.HandleFunc("/settings", Settings)

	fmt.Println("(http://localhost:8080) - Server started on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
