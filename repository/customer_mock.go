package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 001, Name: "Sutin", DateOfBirth: "17/01/2352", City: "Surinn", ZipCode: "32120", Status: 1},
		{CustomerID: 002, Name: "Sutin", DateOfBirth: "17/01/2352", City: "Surinn", ZipCode: "32120", Status: 1},
		{CustomerID: 003, Name: "Sutin", DateOfBirth: "17/01/2352", City: "Surinn", ZipCode: "32120", Status: 0},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("not found")
}
