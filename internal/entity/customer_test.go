package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewCustomer_ThenShouldReceiveAnError(t *testing.T) {
	customer := Customer{}
	assert.Error(t, customer.IsValid(), "invalid id")
}

func TestGivenAnEmptyName_WhenCreateANewCustomer_ThenShouldReceiveAnError(t *testing.T) {
	customer := Customer{ID: "123"}
	assert.Error(t, customer.IsValid(), "invalid price")
}
