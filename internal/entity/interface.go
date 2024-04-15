package entity

type CustomerRepositoryInterface interface {
	FindAll() ([]Customer, error)
	Create(customer *Customer) error
	GetCustomer(id string) (Customer, error)
	UpdateCustomer(customer *Customer) error
	DeleteCustomer(id string) error
}
