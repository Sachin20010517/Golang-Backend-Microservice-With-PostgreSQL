package api

import (
	"PostgreSQL-Project/dao"
	"PostgreSQL-Project/utils"

	"github.com/gofiber/fiber/v2"
)

func FindAllEmployeesApi(c *fiber.Ctx) error {

	employees, err := dao.FindAllEmployees()
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&employees)
}
