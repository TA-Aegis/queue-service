package rest

import (
	"antrein/bc-dashboard/model/config"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func StartServer(cfg *config.Config, app *fiber.App) error {
	return app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
