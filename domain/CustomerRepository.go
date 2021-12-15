package domain

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	//GetById(id int64) (Customer, error)
	//Create(customer Customer) (Customer, error)
	//Update(customer Customer) (Customer, error)
	//Delete(id int64) error
}

type CustomerRepositoryImpl struct {
	customers []Customer
}

func (repo CustomerRepositoryImpl) GetAll() ([]Customer, error) {
	return repo.customers, nil
}
