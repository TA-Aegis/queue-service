package main

import (
	"antrein/bc-dashboard/model/config"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func startServer(cfg *config.Config, app *fiber.App) error {
	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
