package dao

import (
	"PostgreSQL-Project/dbConfig"
	"PostgreSQL-Project/dto"
	"fmt"
	"log"
)

func DB_UpdateEmployee(employee *dto.Employee) error {
	query := `
		UPDATE employees
		SET 
			first_name = $1,
			last_name = $2,
			country = $3,
			mobile_no = $4,
			dob = $5,
			branch = $6,
			email = $7,
			created_at = now()
		WHERE employee_id = $8 AND deleted = false
	`

	result, err := dbConfig.ConnectDB().Exec(
		query,
		employee.FirstName,
		employee.LastName,
		employee.Country,
		employee.MobileNo,
		employee.DOB,
		employee.Branch,
		employee.Email,
		employee.EmployeeID,
	)

	if err != nil {
		log.Printf("Error updating employee with ID (%s): %v", employee.EmployeeID, err)
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected for update: %v", err)
		return err
	}

	if rowsAffected == 0 {
		// No rows were updated (employee not found or already deleted)
		return fmt.Errorf("employee with ID %s not found or already deleted", employee.EmployeeID)
	}

	return nil
}
