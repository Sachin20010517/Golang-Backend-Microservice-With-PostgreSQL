package api

import (
	//"PostgreSQL-Project/dao"
	"PostgreSQL-Project/dto"
	"PostgreSQL-Project/utils"
	"PostgreSQL-Project/dao"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateMemberApi(c *fiber.Ctx) error {
	inputObj := dto.Employee{}

	// Parse the incoming JSON data
	if err := c.BodyParser(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// Validate the input
	validate := validator.New()
	if validationErr := validate.Struct(&inputObj); validationErr != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
	}

	// Generate a unique Member ID
	inputObj.EmployeeID = utils.GenerateUniqueID("Emp")

	// Set defaults for other fields
	inputObj.Deleted = false

	// Insert the member into the database
	if err := dao.DB_CreateEmployee(&inputObj); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccessResponse(c)
}