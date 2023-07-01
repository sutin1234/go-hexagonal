package service

import (
	"database/sql"
	"errors"
	"log"

	"github.com/sutin1234/go-hexagonal/repository"
)

type customerService struct {
	cusRepo repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) customerService {
	return customerService{cusRepo: cusRepo}
}

func (c customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := c.cusRepo.GetAll()
	if err != nil {
		log.Println(err)
	}
	customerResposes := []CustomerResponse{}
	for _, customer := range customers {
		cusResp := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResposes = append(customerResposes, cusResp)
	}
	return customerResposes, nil
}
func (c customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := c.cusRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		log.Println(err)
		return nil, err
	}
	cusResp := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &cusResp, nil
}
