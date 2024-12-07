package apiHandlers

import (
	"PostgreSQL-Project/api"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/PostgreSQL-Project/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {

	cg.Post("/CreateEmployee", api.CreateMemberApi)
	cg.Get("/FindAllEmployees", api.FindAllEmployeesApi)
	cg.Get("/FindEmployee/EmployeeId",api.FindEmployeeByEmployeeIDApi)
	cg.Delete("/DeleteEmployee/EmployeeId",api.Delete_EmployeeByEmployeeId_Api)
	cg.Put("UpdateEmployee", api.Update_EmployeeApi)
	

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "PostgreSQL-Project service is up and running", "version": "1.0"})
	})
}
