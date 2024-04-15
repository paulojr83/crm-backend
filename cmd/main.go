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
	webOrderHandler := NewWebCustomerHandler(db)

	webserver.AddHandler("POST", "/customers", webOrderHandler.AddCustomer)
	webserver.AddHandler("GET", "/customers", webOrderHandler.GetCustomers)
	webserver.AddHandler("GET", "/customers/{id}", webOrderHandler.GetCustomer)
	webserver.AddHandler("PUT", "/customers/{id}", webOrderHandler.UpdateCustomer)
	webserver.AddHandler("DELETE", "/customers/{id}", webOrderHandler.DeleteCustomer)

	fmt.Println("Starting web server on port", config.WebServerPort)
	webserver.Start()
}
