package Service

import (
	"restAPI/domain"
	"restAPI/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id int64) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	customerRepo domain.CustomerRepository
}

func (c DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return c.customerRepo.GetAll()
}

func (c DefaultCustomerService) GetCustomer(id int64) (*domain.Customer, *errs.AppError) {
	return c.customerRepo.GetById(id)
}

func NewCustomerService(customerRepo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{customerRepo}
}
