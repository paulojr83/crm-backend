package database

import (
	"context"
	"database/sql"
	"github.com/paulojr83/crm-backend/internal/entity"
)

type CustomerRepository struct {
	Db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{Db: db}
}

func (r *CustomerRepository) Create(customer *entity.Customer) error {
	stmt, err := r.Db.Prepare("INSERT INTO customers (id, name, role, email, phone, contacted) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Role, customer.Email, customer.Phone, customer.Contacted)
	if err != nil {
		return err
	}
	return nil
}

func (q *CustomerRepository) FindAll() ([]entity.Customer, error) {
	rows, err := q.Db.Query("select id, name, role, email, phone, contacted from customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []entity.Customer
	for rows.Next() {
		var i entity.Customer
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Role,
			&i.Email,
			&i.Phone,
			&i.Contacted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomer = `-- name: getCustomer :one
select id, name, role, email, phone, contacted from customers where id = ?
`

func (q *CustomerRepository) GetCustomer(id string) (entity.Customer, error) {
	ctx := context.Background()
	row := q.Db.QueryRowContext(ctx, getCustomer, id)
	var customer entity.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Role, &customer.Email, &customer.Contacted)
	return customer, err
}

const updateCustomer = `-- name: UpdateCustomer :exec
update customers set name =?, role =?, email =?, phone =?, contacted =? where id = ?
`

func (q *CustomerRepository) UpdateCustomer(customer *entity.Customer) error {
	ctx := context.Background()
	_, err := q.Db.ExecContext(ctx, updateCustomer,
		customer.Name,
		customer.Role,
		customer.Phone,
		customer.Email,
		customer.Contacted,
		customer.ID)
	return err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
delete from customers where id = ?
`

func (q *CustomerRepository) DeleteCustomer(id string) error {
	ctx := context.Background()
	_, err := q.Db.ExecContext(ctx, deleteCustomer, id)
	return err
}
