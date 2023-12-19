package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (cr CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	findQuery := ""
	if status != "" {
		findQuery = "select customer_id, name, city,zipcode, date_of_birth, status from customers where status = " + status
	} else {
		findQuery = "select customer_id, name, city,zipcode, date_of_birth, status from customers"
	}

	rows, err := cr.client.Query(findQuery)
	if err != nil {
		log.Println("Error querying the database for customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	var c Customer
	for rows.Next() {
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("No Customers found")
			} else {
				log.Println("Error scanning the results to customers" + err.Error())
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (cr CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findCustomerQuery := "Select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	row := cr.client.QueryRow(findCustomerQuery, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error scanning the result to customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
