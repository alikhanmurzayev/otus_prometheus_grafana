package main

import (
	"strconv"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port int, userController *userController) error {
	app := fiber.New()

	app.Use(Instrumenting)
	app.Use(Recover)
	app.Get("/health", userController.HealthCheck)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/panic", func(ctx *fiber.Ctx) error { panic("hand-made panic") })

	userGroup := app.Group("/user")
	userGroup.Post("/", userController.CreateUser)
	userGroup.Get("/:id", userController.GetUser)
	userGroup.Delete("/:id", userController.DeleteUser)
	userGroup.Put("/:id", userController.UpdateUser)

	return app.Listen(":" + strconv.Itoa(port))
}
