package models

import "akakpo/db/types"

// ReadOffice TODO
func ReadOffice(id string) ([]types.OfficeResult, error) {
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
		WHERE E.officeCode = 1
	`

	var response []types.OfficeResult
	var err error

	rows, _ := Database.Query(sql)
	defer rows.Close()

	for rows.Next() {
		item := types.OfficeResult{}
		err = rows.Scan(&item.EmployeeNumber, &item.LastName, &item.FirstName, &item.Extension, &item.Email, &item.JobTitle, &item.ReportsTo)

		response = append(response, item)
	}

	return response, err
}
