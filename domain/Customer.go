package domain

import "restAPI/errs"

type Customer struct {
	Id          int64  `db:"customer_id" json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	GetAll(status string) ([]Customer, *errs.AppError)
	GetById(id int64) (*Customer, *errs.AppError)
}
