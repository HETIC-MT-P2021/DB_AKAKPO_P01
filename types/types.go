package types

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

// OfficeEmployeeResult Represents the return shape of an office employee
type OfficeEmployeeResult struct {
	EmployeeNumber int
	LastName       string
	FirstName      string
	Extension      string
	Email          string
	JobTitle       string
	ReportsTo      NullString
}

// OfficeResult Represents the return shape of an office
type OfficeResult struct {
	OfficeCode   string
	City         string
	Phone        string
	AddressLine1 string
	AddressLine2 NullString
	State        NullString
	Country      string
	PostalCode   string
	Territory    string
	Employees    []OfficeEmployeeResult
}

// EmployeeResult Represents the return shape of an employee
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
	OfficeAddressLine2 NullString
}

// OrderResult Represents the return shape of an order
type OrderResult struct {
	ProductName     string
	QuantityOrdered int
	PriceEach       float32
	OrderLineNumber int
}

// CustomerOrderResult Represents the return shape of a customer order
type CustomerOrderResult struct {
	OrderNumber       int
	NbOrderedProducts int
	TotalPrice        float32
}

// CustomerResult Represents the return shape of a customer
type CustomerResult struct {
	CustomerNumber         string
	CustomerName           string
	ContactLastName        string
	ContactFirstName       string
	Phone                  string
	AddressLine1           string
	AddressLine2           NullString
	City                   string
	State                  NullString
	PostalCode             NullString
	Country                string
	SalesRepEmployeeNumber int
	CreditLimit            float32
	Orders                 []CustomerOrderResult
}
