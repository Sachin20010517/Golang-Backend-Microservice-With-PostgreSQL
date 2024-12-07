package dao

import (
	"PostgreSQL-Project/dbConfig"
	// "database/sql"
	"fmt"
	"log"
)

func DB_DeleteEmployeeById(employeeID string) (error) {
	

	query := `
		UPDATE employees
		set deleted = true
		WHERE employee_id = $1 AND deleted = false
	`

	result, err := dbConfig.ConnectDB().Exec(query,employeeID)
	if err != nil {
		log.Printf("Error deleting employee with ID (%s): %v", employeeID, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected for deletion: %v", err)
		return err
	}

	if rowsAffected == 0 {
		// No rows were updated (employee not found or already deleted)
		return fmt.Errorf("employee with ID %s not found or already deleted", employeeID)
	}
	
	return nil
}
