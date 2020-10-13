package models

import "akakpo/db/types"

// ReadOffice TODO
func ReadOffice(id string) (types.OfficeResult, error) {
	sql := `
		SELECT
			officeCode,
			city,
			phone,
			addressLine1,
			addressLine2,
			state,
			country,
			postalCode,
			territory
		FROM offices
		WHERE officeCode = ?
	`

	var response types.OfficeResult
	var errors []error

	rows, queryErr := Database.Query(sql, id)

	if queryErr != nil {
		errors = append(errors, queryErr)
	}

	rows.Next()
	scanError := rows.Scan(&response.OfficeCode, &response.City, &response.Phone, &response.AddressLine1, &response.AddressLine2, &response.State, &response.Country, &response.PostalCode, &response.Territory)

	if scanError != nil {
		errors = append(errors, scanError)
	}

	if len(errors) == 0 {
		return response, nil
	}
	return response, errors[0]
}

// ReadOfficeEmployees TODO
func ReadOfficeEmployees(id string) ([]types.OfficeEmployeeResult, error) {
	sql := `
		SELECT
			E.employeeNumber,
			E.lastName,
			E.firstName,
			E.extension,
			E.email,
			E.jobTitle,
			(SELECT CONCAT(firstName, ' ', lastName) AS fullName FROM employees WHERE employeeNumber = E.reportsTo) AS reportsTo
		FROM employees E
		WHERE E.officeCode = ?
	`

	var response []types.OfficeEmployeeResult
	var err error

	rows, _ := Database.Query(sql, id)
	defer rows.Close()

	for rows.Next() {
		item := types.OfficeEmployeeResult{}
		err = rows.Scan(&item.EmployeeNumber, &item.LastName, &item.FirstName, &item.Extension, &item.Email, &item.JobTitle, &item.ReportsTo)

		response = append(response, item)
	}

	return response, err
}
