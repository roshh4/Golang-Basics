package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleGeneral)
	http.HandleFunc("/rosh", handleRosh)
	http.HandleFunc("/srihari", handleSrihari)

	fmt.Println("Server started at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func handleRosh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "rosh is peak")
}
func handleSrihari(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "srihari is gay")
}
func handleGeneral(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
