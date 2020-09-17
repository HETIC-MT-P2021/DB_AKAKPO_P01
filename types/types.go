package types

import (
	"database/sql"
	"reflect"
)

// Order Represents an order
type Order struct {
	OrderNumber    int
	OrderDate      string
	RequiredDate   string
	ShippedDate    string
	Status         string
	Comments       string
	CustomerNumber int
}

// Office Represents an office
type Office struct {
	OfficeCode   string
	City         string
	Phone        string
	AddressLine1 string
	AddressLine2 string
	State        string
	Country      string
	PostalCode   string
	Territory    string
}

// Employee Represents an employee
type Employee struct {
	EmployeeNumber int
	LastName       string
	FirstName      string
	Extension      string
	Email          string
	OfficeCode     string
	ReportsTo      int
	JobTitle       string
}

// Customer Represents a customer of an office
type Customer struct {
	CustomerNumber         int
	CustomerName           string
	ContactLastName        string
	ContactFirstName       string
	Phone                  string
	AddressLine1           string
	AddressLine2           string
	City                   string
	State                  string
	PostalCode             string
	Country                string
	SalesRepEmployeeNumber int
	CreditLimit            float32
}

// NullString TODO
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// OfficeResult TODO
type OfficeResult struct {
	EmployeeNumber int
	LastName       string
	FirstName      string
	Extension      string
	Email          string
	JobTitle       string
	ReportsTo      NullString
}

// EmployeeResult TODO
type EmployeeResult struct {
	EmployeeNumber     int
	LastName           string
	FirstName          string
	Extension          string
	Email              string
	JobTitle           string
	ReportsTo          NullString
	OfficeCity         string
	OfficePhone        string
	OfficeAddressLine1 string
	OfficeAddressLine2 string
}

// OrderResult TODO
type OrderResult struct {
	ProductName     string
	QuantityOrdered int
	PriceEach       float32
	OrderLineNumber int
}

// CustomerResult TODO
type CustomerResult struct {
	CustomerNumber         string
	CustomerName           string
	ContactLastName        string
	ContactFirstName       string
	Phone                  string
	AddressLine1           string
	AddressLine2           string
	City                   string
	State                  string
	PostalCode             string
	Country                string
	SalesRepEmployeeNumber int
	CreditLimit            float32
	SalesRespEmp           string
	OrderNumber            int
	NbOrderedProducts      int
	TotalPrice             float32
}
