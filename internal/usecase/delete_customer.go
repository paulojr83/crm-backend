package usecase

import (
	"github.com/paulojr83/crm-backend/internal/entity"
)

type DeleteCustomerInputDTO struct {
	ID string `json:"id"`
}

type DeleteCustomerUseCase struct {
	CustomerRepository entity.CustomerRepositoryInterface
}

func NewDeleteCustomerUseCase(
	CustomerRepository entity.CustomerRepositoryInterface,
) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{
		CustomerRepository: CustomerRepository,
	}
}

func (c *DeleteCustomerUseCase) Execute(input DeleteCustomerInputDTO) error {

	err := c.CustomerRepository.DeleteCustomer(input.ID)
	if err != nil {
		return err
	}
	return nil
}
