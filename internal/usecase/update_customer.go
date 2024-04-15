package usecase

import (
	"github.com/paulojr83/crm-backend/internal/entity"
)

type UpdateCustomerInputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type UpdateCustomerOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type UpdateCustomerUseCase struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewUpdateCustomerUseCase(
	CustomerRepository entity.CustomerRepositoryInterface,
) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		CustomerRepository: CustomerRepository,
	}
}

func (c *UpdateCustomerUseCase) Execute(input UpdateCustomerInputDTO) (UpdateCustomerOutputDTO, error) {

	customer, err := entity.NewCustomer(input.ID, input.Name, input.Role, input.Email, input.Phone, input.Contacted)
	if err != nil {
		return UpdateCustomerOutputDTO{}, err
	}

	err = c.CustomerRepository.UpdateCustomer(customer)
	if err != nil {
		return UpdateCustomerOutputDTO{}, err
	}
	dto := UpdateCustomerOutputDTO{
		ID:        input.ID,
		Name:      input.Name,
		Role:      input.Role,
		Email:     input.Email,
		Phone:     input.Phone,
		Contacted: input.Contacted,
	}
	return dto, nil
}
