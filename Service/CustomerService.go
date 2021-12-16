package Service

import (
	"restAPI/domain"
	"restAPI/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id int64) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	customerRepo domain.CustomerRepository
}

func (c DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "Active" {
		status = "1"
	} else if status == "Inactive" {
		status = "0"
	} else {
		status = ""
	}
	return c.customerRepo.GetAll(status)
}

func (c DefaultCustomerService) GetCustomer(id int64) (*domain.Customer, *errs.AppError) {
	return c.customerRepo.GetById(id)
}

func NewCustomerService(customerRepo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{customerRepo}
}
