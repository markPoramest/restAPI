package domain

type CustomerRepositoryImpl struct {
	customers []Customer
}

func (repo CustomerRepositoryImpl) GetAll() ([]Customer, error) {
	return repo.customers, nil
}
