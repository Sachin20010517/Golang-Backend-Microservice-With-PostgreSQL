package dto


import "time"

// Employee struct matches your PostgreSQL table
type Employee struct {
    EmployeeID string    `json:"employee_id"`                      // Maps to varchar
    CreatedAt  time.Time `json:"created_at"`                       // Maps to timestamp
    FirstName  string    `json:"first_name" validate:"required"`   // Maps to varchar
    LastName   string    `json:"last_name" validate:"required"`    // Maps to varchar
    Country    string    `json:"country" validate:"required"`      // Maps to varchar (default 'Sri Lanka')
    MobileNo   string    `json:"mobile_no"`                        // Maps to varchar
    Deleted    bool      `json:"deleted" default:"false"`          // Maps to bool
    DOB        string    `json:"dob"`                              // Maps to date
    Branch     string    `json:"branch" validate:"required"`       // Maps to varchar
    Email      string    `json:"email" validate:"required,email"`  // Maps to varchar
}
