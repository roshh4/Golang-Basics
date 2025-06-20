package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Spidy 🕷️")
		fmt.Fprintln(w, "Hello, Alpha 🐺")
		fmt.Fprintln(w, "Hello, Chair 🪑")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
