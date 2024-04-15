package usecase

import (
	"github.com/paulojr83/crm-backend/internal/entity"
)

type ListCustomerOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type ListCustomerUseCase struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewListCustomerUseCase(
	CustomerRepository entity.CustomerRepositoryInterface,
) *ListCustomerUseCase {
	return &ListCustomerUseCase{
		CustomerRepository: CustomerRepository,
	}
}

func (c *ListCustomerUseCase) Execute() ([]ListCustomerOutputDTO, error) {
	customers, err := c.CustomerRepository.FindAll()
	if err != nil {
		return []ListCustomerOutputDTO{}, err
	}
	var ordersOutPut []ListCustomerOutputDTO
	for _, customer := range customers {
		dto := ListCustomerOutputDTO{
			ID:        customer.ID,
			Name:      customer.Name,
			Role:      customer.Role,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Contacted: customer.Contacted,
		}
		ordersOutPut = append(ordersOutPut, dto)
	}

	return ordersOutPut, nil
}
