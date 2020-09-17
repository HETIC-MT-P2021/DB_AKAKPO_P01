package models

import "akakpo/db/types"

// ReadCustomer reads all the customers
func ReadCustomer(id string) ([]types.CustomerResult, error) {
	sql := `
		SELECT
			(SELECT CONCAT(E.lastName, ' ', E.firstName) FROM employees E WHERE E.employeeNumber = C.salesRepEmployeeNumber) AS salesRespEmp,
			O.orderNumber,
			COUNT(OD.productCode) AS nbOrderedProducts,
			SUM(OD.quantityOrdered * OD.priceEach) AS totalPrice
		FROM orders O
			INNER JOIN orderdetails OD ON O.orderNumber = OD.orderNumber
			INNER JOIN customers C ON O.customerNumber = C.customerNumber
		WHERE O.customerNumber = 112
		GROUP BY O.orderNumber
	`

	var response []types.CustomerResult
	var err error

	rows, _ := Database.Query(sql)
	for rows.Next() {
		item := types.CustomerResult{}
		err = rows.Scan(
			&item.CustomerNumber,
			&item.CustomerName,
			&item.ContactLastName,
			&item.ContactFirstName,
			&item.Phone,
			&item.AddressLine1,
			&item.AddressLine2,
			&item.City,
			&item.State,
			&item.PostalCode,
			&item.Country,
			&item.SalesRepEmployeeNumber,
			&item.CreditLimit,
			&item.SalesRespEmp,
			&item.OrderNumber,
			&item.NbOrderedProducts,
			&item.TotalPrice,
		)

		response = append(response, item)
	}

	return response, err
}
