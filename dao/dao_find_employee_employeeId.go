package dao

import (
	"PostgreSQL-Project/dbConfig"
	"PostgreSQL-Project/dto"
	"database/sql"
	"log"
)

func DB_FindEmployeeByID(employeeID string) (*dto.Employee, error) {
	var employee dto.Employee

	query := `
		SELECT 
			employee_id, created_at, first_name, last_name, country, mobile_no, deleted, dob, branch, email
		FROM employees
		WHERE deleted = false AND employee_id = $1
	`

	err := dbConfig.ConnectDB().QueryRow(query, employeeID).Scan(
		&employee.EmployeeID,
		&employee.CreatedAt,
		&employee.FirstName,
		&employee.LastName,
		&employee.Country,
		&employee.MobileNo,
		&employee.Deleted,
		&employee.DOB,
		&employee.Branch,
		&employee.Email,
	)
	if err == sql.ErrNoRows {
		// No employee found for the given ID
		return nil, nil
	} else if err != nil {
		// General query execution error
		log.Printf("Error fetching employee by ID (%s): %v", employeeID, err)
		return nil, err
	}

	return &employee, nil
}
