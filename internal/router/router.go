package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	apiV1Routes(api.Group("/v1"))
}

func apiV1Routes(router fiber.Router) {
	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
}
