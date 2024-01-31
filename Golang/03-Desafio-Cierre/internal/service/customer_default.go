package service

import (
	"app/internal"
	"errors"
)

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

func (s *CustomersDefault) SaveAll(c []internal.Customer) (err error) {
	// Call method SaveAll
	err = s.rp.SaveAll(c)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryCustomerDump) {
			err = internal.ErrServiceCustomerDump
			return
		}
		err = internal.ErrServiceUnexpectError
		return
	}
	return
}
