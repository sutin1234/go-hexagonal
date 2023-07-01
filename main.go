package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sutin1234/go-hexagonal/repository"
)

func main() {

	db, err := sqlx.Open("mysql", "root:password@/go_db")
	if err != nil {
		panic(err)
	}

	customerRepo := repository.NewCustomerRepositoryDB(db)
	customers, err := customerRepo.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("customers %v", customers)

	customer, err := customerRepo.GetById(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("customer %v", customer)
}
