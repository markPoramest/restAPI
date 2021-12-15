package app

import (
	"encoding/json"
	"net/http"
	"restAPI/Service"
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
