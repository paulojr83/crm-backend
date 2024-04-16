package main

import (
	"database/sql"
	"fmt"
	"github.com/paulojr83/crm-backend/configs"
	"github.com/paulojr83/crm-backend/internal/infra/web/webserver"
	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	webserver := webserver.NewWebServer(config.WebServerPort)
	webCustomerHandler := NewWebCustomerHandler(db)

	webserver.AddHandler("POST", "/customers", webCustomerHandler.AddCustomer)
	webserver.AddHandler("GET", "/customers", webCustomerHandler.GetCustomers)
	webserver.AddHandler("GET", "/customers/{id}", webCustomerHandler.GetCustomer)
	webserver.AddHandler("PUT", "/customers/{id}", webCustomerHandler.UpdateCustomer)
	webserver.AddHandler("DELETE", "/customers/{id}", webCustomerHandler.DeleteCustomer)

	fmt.Println("Starting web server on port", config.WebServerPort)
	webserver.Start()
}
