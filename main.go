package main

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/sutin1234/go-hexagonal/handler"
	"github.com/sutin1234/go-hexagonal/logs"
	"github.com/sutin1234/go-hexagonal/repository"
	"github.com/sutin1234/go-hexagonal/service"
)

func main() {

	initConfig()
	dsn := fmt.Sprintf("%v:%v@/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	customerRepo := repository.NewCustomerRepositoryDB(db)
	// _ = customerRepo
	// Mock
	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	// customerService := service.NewCustomerService(customerRepositoryMock)
	customerService := service.NewCustomerService(customerRepo)
	customerHandle := handler.NewCustomerHandle(customerService)

	accountRepositoryDB := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

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

	// services

	// customers, err := customerService.GetCustomers()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("customers Resp %v", customers)

	// customer, err := customerService.GetCustomer(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("customer Resp %v", customer)

	// Router

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandle.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", customerHandle.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info("Application StartAndServe at " + viper.GetString("app.host") + ":" + viper.GetString("app.port"))
	err = http.ListenAndServe(":"+viper.GetString("app.port"), router)
	if err != nil {
		panic(err)
	}

}

func initConfig() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
