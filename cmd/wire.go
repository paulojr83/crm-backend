//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/paulojr83/crm-backend/internal/entity"
	"github.com/paulojr83/crm-backend/internal/infra/database"
	"github.com/paulojr83/crm-backend/internal/infra/web"
)

var setCustomerRepositoryDependency = wire.NewSet(
	database.NewCustomerRepository,
	wire.Bind(new(entity.CustomerRepositoryInterface), new(*database.CustomerRepository)),
)

func NewWebCustomerHandler(db *sql.DB) *web.WebCustomerHandler {
	wire.Build(
		setCustomerRepositoryDependency,
		web.NewWebCustomerHandler,
	)
	return &web.WebCustomerHandler{}
}
