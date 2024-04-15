package database

import (
	"database/sql"
	"github.com/paulojr83/crm-backend/internal/entity"
	"testing"

	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type CustomerRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *CustomerRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec("create table customers(  id varchar(255) not null primary key,  name  varchar(255) not null,  role varchar(16) not null,  email varchar(32)  null,  phone varchar(32),  contacted bool)")
	if err != nil {
		return
	}
	suite.Db = db
}

func (suite *CustomerRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}

func (suite *CustomerRepositoryTestSuite) Test_SaveCustomer() {
	customer, err := entity.NewCustomer("123", "Test", "Test", "test@test.com", "Test", false)
	suite.NoError(err)
	repo := NewCustomerRepository(suite.Db)
	err = repo.Create(customer)
	suite.NoError(err)

	var customerResult entity.Customer
	err = suite.Db.QueryRow("select id, name, role, email, phone, contacted from customers where id = ?", customer.ID).
		Scan(&customerResult.ID, &customerResult.Name, &customerResult.Role, &customerResult.Email, &customerResult.Phone, &customerResult.Contacted)

	suite.NoError(err)
	suite.Equal(customer.ID, customerResult.ID)
	suite.Equal(customer.Email, customerResult.Email)
	suite.Equal(customer.Role, customerResult.Role)
	suite.Equal(customer.Name, customerResult.Name)
	suite.Equal(customer.Contacted, customerResult.Contacted)
}
