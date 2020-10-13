package models

import (
	"akakpo/db/types"
)

// ReadCustomer Retreive from database a customer
func ReadCustomer(id string) (types.CustomerResult, error) {
	sql := `
		SELECT * FROM customers WHERE customerNumber = ?
	`

	var response types.CustomerResult
	var errors []error

	rows, queryErr := Database.Query(sql, id)

	if queryErr != nil {
		errors = append(errors, queryErr)
	}

	rows.Next()
	scanError := rows.Scan(
		&response.CustomerNumber,
		&response.CustomerName,
		&response.ContactLastName,
		&response.ContactFirstName,
		&response.Phone,
		&response.AddressLine1,
		&response.AddressLine2,
		&response.City,
		&response.State,
		&response.PostalCode,
		&response.Country,
		&response.SalesRepEmployeeNumber,
		&response.CreditLimit,
	)

	if scanError != nil {
		errors = append(errors, scanError)
	}

	if len(errors) == 0 {
		return response, nil
	}
	return response, errors[0]
}

// ReadCustomerOrders Retreive from database the orders of a customer
func ReadCustomerOrders(customerNumber string) ([]types.CustomerOrderResult, error) {
	sql := `
		SELECT
			O.orderNumber,
			COUNT(OD.productCode) AS nbOrderedProducts,
			SUM(OD.quantityOrdered * OD.priceEach) AS totalPrice
		FROM orders O
		INNER JOIN orderdetails OD ON O.orderNumber = OD.orderNumber
		WHERE O.customerNumber = ?
		GROUP BY O.orderNumber
	`

	var response []types.CustomerOrderResult
	var errors []error

	rows, queryErr := Database.Query(sql, customerNumber)

	if queryErr != nil {
		errors = append(errors, queryErr)
	}

	for rows.Next() {
		item := types.CustomerOrderResult{}
		scanErr := rows.Scan(
			&item.OrderNumber,
			&item.NbOrderedProducts,
			&item.TotalPrice,
		)

		if scanErr != nil {
			errors = append(errors, scanErr)
		}

		response = append(response, item)
	}

	if len(errors) == 0 {
		return response, nil
	}
	return response, errors[0]
}
