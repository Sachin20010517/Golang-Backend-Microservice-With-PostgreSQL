package api

import (
	"PostgreSQL-Project/dao"
	"PostgreSQL-Project/dto"
	"PostgreSQL-Project/utils"


	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Update_EmployeeApi(c *fiber.Ctx) error {
	// Parse request body into dto.Employee
	//var employee dto.Employee

	inputObj := dto.Employee{}

	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }

	// Call the DAO function
	err := dao.DB_UpdateEmployee(&inputObj)
	if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }

	// Return success response
	return utils.SendSuccessResponse(c)
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": "Employee updated successfully",
	// })
}
