package api

import (
	"prtask1/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(taskService service.Service) *fiber.App {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api/v1")

	api.Post("/tasks", taskService.CreateTask)       // POST /api/v1/tasks
	api.Get("/tasks", taskService.GetAllTasks)       // GET /api/v1/tasks
	api.Get("/tasks/:id", taskService.GetTaskByID)   // GET /api/v1/tasks/:id
	api.Put("/tasks/:id", taskService.UpdateTask)    // PUT /api/v1/tasks/:id
	api.Delete("/tasks/:id", taskService.DeleteTask) // DELETE /api/v1/tasks/:id

	return app
}
