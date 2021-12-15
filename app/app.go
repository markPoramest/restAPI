package app

import (
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customer", getAllCustomer)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		return
	}
}
