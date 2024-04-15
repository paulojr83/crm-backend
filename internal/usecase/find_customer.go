package usecase

import (
	"github.com/paulojr83/crm-backend/internal/entity"
)

type FindCustomerInputDTO struct {
	ID string `json:"id"`
}

type FindCustomerOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type FindCustomerUseCase struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewFindCustomerUseCase(
	CustomerRepository entity.CustomerRepositoryInterface,
) *FindCustomerUseCase {
	return &FindCustomerUseCase{
		CustomerRepository: CustomerRepository,
	}
}

func (c *FindCustomerUseCase) Execute(input FindCustomerInputDTO) (CustomerOutputDTO, error) {

	customer, err := c.CustomerRepository.GetCustomer(input.ID)
	if err != nil {
		return CustomerOutputDTO{}, err
	}
	dto := CustomerOutputDTO{
		ID:        customer.ID,
		Name:      customer.Name,
		Role:      customer.Role,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Contacted: customer.Contacted,
	}
	return dto, nil
}
