package models

import "akakpo/db/types"

// ReadCustomer TODO
func ReadCustomer(id string) (types.CustomerResult, error) {
	sql := `
		SELECT * FROM customers WHERE customerNumber = ?
	`

	var response types.CustomerResult
	var err error

	row := Database.QueryRow(sql, id)
	err = row.Scan(
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

	return response, err
}

// ReadCustomerOrders TODO
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
	var err error

	rows, _ := Database.Query(sql, customerNumber)

	for rows.Next() {
		item := types.CustomerOrderResult{}
		err = rows.Scan(
			&item.OrderNumber,
			&item.NbOrderedProducts,
			&item.TotalPrice,
		)

		response = append(response, item)
	}

	return response, err
}
