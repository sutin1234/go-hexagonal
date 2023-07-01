package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/sutin1234/go-hexagonal/repository"
)

func main() {
	db, err := sqlx.Open("mysql", "root:password@/go_db")
	if err != nil {
		panic(err)
	}

	customerRepo := repository.NewCustomerRepositoryDB(db)
	_ = customerRepo
}
