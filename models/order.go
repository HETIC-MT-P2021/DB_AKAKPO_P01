package models

import "akakpo/db/types"

// ReadOrder reads all the orders
func ReadOrder(id string) ([]types.OrderResult, error) {
	sql := `
		SELECT
			(SELECT P.productName FROM products P WHERE P.productCode = OD.productCode) AS productName,
			OD.quantityOrdered,
			OD.priceEach,
			OD.orderLineNumber
		FROM orderdetails OD INNER JOIN orders O ON OD.orderNumber = O.orderNumber
		WHERE O.orderNumber = 10115
		GROUP BY OD.productCode
	`

	var response []types.OrderResult
	var err error

	rows, _ := Database.Query(sql)

	for rows.Next() {
		item := types.OrderResult{}
		err = rows.Scan(&item.ProductName, &item.QuantityOrdered, &item.PriceEach, &item.OrderLineNumber)

		response = append(response, item)
	}

	return response, err
}
