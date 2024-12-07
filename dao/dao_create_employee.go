package dao

import (
	"database/sql"
	"log"
	"time"

	"PostgreSQL-Project/dto"
	"PostgreSQL-Project/dbConfig"
)

var db *sql.DB

// Initialize the database connection
func init() {
	db = dbConfig.ConnectDB()
}

// CreateEmployee inserts a new employee record into the database
func DB_CreateEmployee(employee *dto.Employee) error {
	query := `
		INSERT INTO employees 
		(employee_id, created_at, first_name, last_name, country, mobile_no, deleted, dob, branch, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	// Default to current timestamp if CreatedAt is not provided
	if employee.CreatedAt.IsZero() {
		employee.CreatedAt = time.Now()
	}

	// Execute the query
	_, err := db.Exec(
		query,
		employee.EmployeeID,
		employee.CreatedAt,
		employee.FirstName,
		employee.LastName,
		employee.Country,
		employee.MobileNo,
		employee.Deleted,
		employee.DOB,
		employee.Branch,
		employee.Email,
	)

	if err != nil {
		log.Printf("Error inserting employee record: %v", err)
		return err
	}

	log.Println("Employee record inserted successfully")
	return nil
}
