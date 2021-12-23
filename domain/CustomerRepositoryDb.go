package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"restAPI/errs"
	"restAPI/logger"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (c CustomerRepositoryDb) GetAll(status string) ([]Customer, *errs.AppError) {
	var findAllSQL string
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSQL = "SELECT customer_id , name , city , zipcode ,date_of_birth , status FROM customers"
		err = c.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL = "select * from customers where status = ?"
		err = c.client.Select(&customers, findAllSQL, status)
	}
	if err != nil {
		logger.Error("Error: " + err.Error())
		return nil, errs.NewUnExpectedError("Error while fetching data from database")
	}
	return customers, nil
}

func (c CustomerRepositoryDb) GetById(id int64) (*Customer, *errs.AppError) {
	findByIdSQL := "SELECT customer_id , name , city , zipcode ,date_of_birth , status FROM customers WHERE customer_id = ?"
	var customer Customer
	err := c.client.Get(&customer, findByIdSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error: " + err.Error())
			return nil, errs.NewUnExpectedError("Error while getting customer")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sqlx.Open("mysql", "root:1@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db}
}
