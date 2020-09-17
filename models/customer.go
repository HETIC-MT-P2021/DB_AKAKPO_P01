package models

import "akakpo/db/types"

// ReadCustomer TODO
func ReadCustomer(id string) (types.CustomerResult, error) {
	sql := `
		SELECT * FROM customers WHERE customerNumber = 112
	`

	var response types.CustomerResult
	var err error

	row := Database.QueryRow(sql)
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
func ReadCustomerOrders(customberNumber string) ([]types.CustomerOrderResult, error) {
	sql := `
		SELECT
			O.orderNumber,
			COUNT(OD.productCode) AS nbOrderedProducts,
			SUM(OD.quantityOrdered * OD.priceEach) AS totalPrice
		FROM orders O
		INNER JOIN orderdetails OD ON O.orderNumber = OD.orderNumber
		WHERE O.customerNumber = 112
		GROUP BY O.orderNumber
	`

	var response []types.CustomerOrderResult
	var err error

	rows, _ := Database.Query(sql)

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
