package usecase

import (
	"github.com/google/uuid"
	"github.com/paulojr83/crm-backend/internal/entity"
)

type CustomerInputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type CustomerOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type CreateCustomerUseCase struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewCreateCustomerUseCase(
	CustomerRepository entity.CustomerRepositoryInterface,
) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		CustomerRepository: CustomerRepository,
	}
}

func (c *CreateCustomerUseCase) Execute(input CustomerInputDTO) (CustomerOutputDTO, error) {
	if input.ID == "" {
		input.ID = uuid.New().String()
	}
	customer, err := entity.NewCustomer(input.ID, input.Name, input.Role, input.Email, input.Phone, input.Contacted)
	if err != nil {
		return CustomerOutputDTO{}, err
	}

	if err := c.CustomerRepository.Create(customer); err != nil {
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
