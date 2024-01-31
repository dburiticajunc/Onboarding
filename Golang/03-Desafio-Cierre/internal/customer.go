package internal

import "errors"

var (
	ErrRepositoryCustomerDump  = errors.New("repository: error customer dump")
	ErrServiceCustomerDump     = errors.New("service: error customer dump")
	ErrServiceUnexpectError    = errors.New("service: unexpect error")
	ErrRepositoryUnexpectError = errors.New("repository: unexpect error")
)

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}
