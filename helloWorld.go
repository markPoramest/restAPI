package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)

	http.ListenAndServe("localhost:8080", nil)
}
func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World")
}
