package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sutin1234/go-hexagonal/service"
)

type customerHandle struct {
	cusService service.CustomerService
}

func NewCustomerHandle(cusService service.CustomerService) customerHandle {
	return customerHandle{cusService: cusService}
}

func (h customerHandle) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.cusService.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func (h customerHandle) GetCustomer(w http.ResponseWriter, r *http.Request) {
	cusId, _ := strconv.Atoi(mux.Vars(r)["cusId"])
	customer, err := h.cusService.GetCustomer(cusId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
