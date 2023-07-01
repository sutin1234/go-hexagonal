package service

import (
	"database/sql"

	"github.com/sutin1234/go-hexagonal/errs"
	"github.com/sutin1234/go-hexagonal/logs"
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
		logs.Error(err)
		return nil, err
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
			return nil, errs.NewNotFoundError("customer not found 777")
		}
		return nil, errs.NewUnExpectedError()
	}

	cusResp := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &cusResp, nil
}
