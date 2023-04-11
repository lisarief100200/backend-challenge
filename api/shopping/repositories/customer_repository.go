package api

import (
	"backend-challenge/shop"

	"database/sql"
)

type CustomerRepository interface {
	Create(customer shop.Customer) (shop.Customer, error) // Create
	FindAll() ([]shop.Customer, error)                    // Read
	Update(customer shop.Customer) (shop.Customer, error) // Update
	Delete(customer shop.Customer) (shop.Customer, error) // Delete
	FindByID(ID int) (shop.Customer, error)               // Find BY ID
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) FindAll() ([]shop.Customer, error) {
	var customers []shop.Customer

	rows, err := cr.db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c shop.Customer
		if err := rows.Scan(&c.ID, &c.CustomerName, &c.CustomerContNo, &c.CustomerAddress, &c.TotalBuy, &c.CreatorId, &c.CreatedAt); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (cr *customerRepository) Create(customer shop.Customer) (shop.Customer, error) {
	stmt, err := cr.db.Prepare("INSERT INTO customers (customer_name, customer_cont_no, customer_address, total_buy, creator_id, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return shop.Customer{}, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(customer.CustomerName, customer.CustomerContNo, customer.CustomerAddress, customer.TotalBuy, customer.CreatorId, customer.CreatedAt).Scan(&id)
	if err != nil {
		return shop.Customer{}, err
	}

	customer.ID = int(id)
	return customer, nil
}

func (cr *customerRepository) Update(customer shop.Customer) (shop.Customer, error) {
	query := "UPDATE customers SET customer_name = $2, customer_cont_no = $3, customer_address = $4, total_buy = $5, creator_id = $6, created_at = $7 WHERE id = $1"
	_, err := cr.db.Exec(query, customer.ID, customer.CustomerName, customer.CustomerContNo, customer.CustomerAddress, customer.TotalBuy, customer.CreatorId, customer.CreatedAt)

	return customer, err
}

func (cr *customerRepository) Delete(customer shop.Customer) (shop.Customer, error) {
	query := "DELETE FROM customers WHERE id = $1"
	_, err := cr.db.Exec(query, customer.ID)
	return customer, err
}

func (cr *customerRepository) FindByID(ID int) (shop.Customer, error) {
	var customer shop.Customer
	query := "SELECT * FROM customers WHERE id = $1"
	stmt, err := cr.db.Prepare(query)
	if err != nil {
		return customer, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(ID)
	err = row.Scan(&customer.ID, &customer.CustomerName, &customer.CustomerContNo, &customer.CustomerAddress, &customer.TotalBuy, &customer.CreatorId, &customer.CreatedAt)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
