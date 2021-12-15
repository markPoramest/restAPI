package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restAPI/Service"
	"strconv"
)

type CustomerHandler struct {
	customerService Service.CustomerService
}

func (cus *CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := cus.customerService.GetAllCustomers()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}
func (cus *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idNumber, _ := strconv.ParseInt(id, 10, 64)
	customer, err := cus.customerService.GetCustomer(idNumber)
	if err != nil {
		writeResponse(w, err.Code, err.Error())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
