package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Docker!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
