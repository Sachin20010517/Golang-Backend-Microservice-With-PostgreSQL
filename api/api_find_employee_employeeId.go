package api

import (
	"PostgreSQL-Project/dao"
	"PostgreSQL-Project/utils"

	"github.com/gofiber/fiber/v2"
)

func FindEmployeeByEmployeeIDApi(c *fiber.Ctx) error {
	// Extract employee_id from the request parameters
	employeeID := c.Query("employeeId")
	if employeeID == "" {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Employee ID is required")
	}

	// Call the DAO function
	employee, err := dao.DB_FindEmployeeByID(employeeID)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	// Check if the employee was not found
	if employee == nil {
		return utils.SendErrorResponse(c, fiber.StatusNotFound, "Employee not found")
	}

	// Return the employee data as JSON
	return c.Status(fiber.StatusOK).JSON(employee)
}
