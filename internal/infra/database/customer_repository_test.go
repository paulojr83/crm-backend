package database

import (
	"database/sql"
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
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *CustomerRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepositoryTestSuite))
}

func (suite *CustomerRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {

}

func (suite *CustomerRepositoryTestSuite) Test_List_Orders() {

}