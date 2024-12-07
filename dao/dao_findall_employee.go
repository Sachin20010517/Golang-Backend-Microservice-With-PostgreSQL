package dao

import (
	// "database/sql"
	"log"

	"PostgreSQL-Project/dto"
	"PostgreSQL-Project/dbConfig"
)

func FindAllEmployees() ([]dto.Employee, error) {
	var employees []dto.Employee

	query := `
		SELECT 
			employee_id, created_at, first_name, last_name, country, mobile_no, deleted, dob, branch, email
		FROM employees
		WHERE deleted = false
	`

	rows, err := dbConfig.ConnectDB().Query(query)
	if err != nil {
		log.Printf("Error executing query to find all employees: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee dto.Employee
		err = rows.Scan(
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
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		employees = append(employees, employee)
	}

	if rows.Err() != nil {
		log.Printf("Error in rows iteration: %v", rows.Err())
		return nil, rows.Err()
	}

	return employees, nil
}
