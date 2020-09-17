package types

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
