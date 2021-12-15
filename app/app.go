package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"restAPI/Service"
	"restAPI/domain"
)

func Start() {
	ch := CustomerHandler{
		Service.NewCustomerService(domain.NewCustomerRepositoryDb()),
	}
	router := mux.NewRouter()
	router.HandleFunc("/customer", ch.getAllCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
