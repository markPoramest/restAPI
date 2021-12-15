package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		return
	}
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{{
		Name:    "John",
		City:    "New York",
		Zipcode: "10001",
	}, {
		Name:    "Mary",
		City:    "New York",
		Zipcode: "10001",
	}, {
		Name:    "Mike",
		City:    "New York",
		Zipcode: "10001",
	},
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}

//create function getCustomer by id
func getCustomer(w http.ResponseWriter, r *http.Request) {

}

func createCustomer(w http.ResponseWriter, r *http.Request) {

}
