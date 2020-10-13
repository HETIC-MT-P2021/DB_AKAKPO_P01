package models

import (
	"akakpo/db/types"
)

// ReadEmployee Retreive from database an employee
func ReadEmployee(id string) (types.EmployeeResult, error) {
	sql := `
		SELECT
			E.employeeNumber,
			E.lastName,
			E.firstName,
			E.extension,
			E.email,
			E.jobTitle,
			(SELECT CONCAT(firstName, ' ', lastName) AS fullName FROM employees WHERE employeeNumber = E.reportsTo) AS reportsTo,
			O.city AS officeCity,
			O.phone AS officePhone,
			O.addressLine1 AS officeAddressLine1,
			O.addressLine2 AS officeAddressLine2
		FROM employees E
		INNER JOIN offices O ON E.officeCode = O.officeCode
		WHERE E.employeeNumber = ?
	`

	var response types.EmployeeResult

	row := Database.QueryRow(sql, id)
	err := row.Scan(&response.EmployeeNumber, &response.LastName, &response.FirstName, &response.Extension, &response.Email, &response.JobTitle, &response.ReportsTo, &response.OfficeCity, &response.OfficePhone, &response.OfficeAddressLine1, &response.OfficeAddressLine2)

	return response, err
}
