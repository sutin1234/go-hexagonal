package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sutin1234/go-hexagonal/handler"
	"github.com/sutin1234/go-hexagonal/repository"
	"github.com/sutin1234/go-hexagonal/service"
)

func main() {

	db, err := sqlx.Open("mysql", "root:password@/go_db")
	if err != nil {
		panic(err)
	}

	customerRepo := repository.NewCustomerRepositoryDB(db)
	_ = customerRepo
	// Mock
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepositoryMock)
	// customerService := service.NewCustomerService(customerRepo)
	customerHandle := handler.NewCustomerHandle(customerService)

	// customers, err := customerRepo.GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("customers %v", customers)

	// customer, err := customerRepo.GetById(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("customer %v", customer)

	// servcies

	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("customers Resp %v", customers)

	customer, err := customerService.GetCustomer(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("customer Resp %v", customer)

	// Router

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandle.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{cusId:[0-9]+}", customerHandle.GetCustomer).Methods(http.MethodGet)
	http.ListenAndServe(":8081", router)
}
