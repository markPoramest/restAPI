package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customer", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer}", createCustomer).Methods(http.MethodPost)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
