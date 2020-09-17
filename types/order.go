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
