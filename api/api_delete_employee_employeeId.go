package api

import (
	"PostgreSQL-Project/dao"
	"PostgreSQL-Project/utils"

	"github.com/gofiber/fiber/v2"
)

func Delete_EmployeeByEmployeeId_Api(c *fiber.Ctx) error {
	// Extract employee_id from the request parameters
	employeeID := c.Query("employeeId")
	if employeeID == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Employee ID is required")
	}

	// Call the DAO function
	err := dao.DB_DeleteEmployeeById(employeeID)
	if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Return the employee data as JSON
	return utils.SendSuccessResponse(c)
}
