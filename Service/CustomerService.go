package Service

import "restAPI/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	//GetCustomer(id int) (domain.Customer, error)
}

type DefaultCustomerService struct {
	customerRepo domain.CustomerRepository
}

func (c DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return c.customerRepo.GetAll()
}

func NewCustomerService(customerRepo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{customerRepo}
}
