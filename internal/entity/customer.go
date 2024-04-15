package entity

import "errors"

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

func NewCustomer(id string, name, role, email, phone string, contacted bool) (*Customer, error) {

	order := &Customer{
		ID:        id,
		Name:      name,
		Role:      role,
		Email:     email,
		Phone:     phone,
		Contacted: contacted,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Customer) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Name == "" {
		return errors.New("invalid name")
	}
	if o.Role == "" {
		return errors.New("invalid role")
	}
	if o.Email == "" {
		return errors.New("invalid email")
	}
	if o.Phone == "" {
		return errors.New("invalid phone")
	}
	return nil
}
