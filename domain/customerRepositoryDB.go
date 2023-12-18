package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (cr CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllQuery := "select customer_id, name, city,zipcode, date_of_birth, status from customers"

	rows, err := cr.client.Query(findAllQuery)
	if err != nil {
		log.Println("Error querying the database for customers table" + err.Error())
	}

	customers := make([]Customer, 0)
	var c Customer
	for rows.Next() {
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error scanning the results to customers" + err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
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
