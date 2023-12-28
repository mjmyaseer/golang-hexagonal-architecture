package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (cr CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error

	if status != "" {
		findQuery := "select customer_id, name, city,zipcode, date_of_birth, status from customers where status = ?"
		err = cr.client.Select(&customers, findQuery, status)
	} else {
		findQuery := "select customer_id, name, city,zipcode, date_of_birth, status from customers"
		err = cr.client.Select(&customers, findQuery)
	}

	if err != nil {
		logger.Error("Error querying the database for customers table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return customers, nil
}

func (cr CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findCustomerQuery := "Select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	var customer Customer
	err := cr.client.Get(&customer, findCustomerQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error scanning the result to customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
