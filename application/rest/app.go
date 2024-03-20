package rest

import (
	"antrein/bc-dashboard/model/config"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func ApplicationDelegate(cfg *config.Config) (*fiber.App, error) {
	// ctx := context.Background()
	app := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s %s", cfg.Server.Rest.Name, cfg.Stage),
	})

	// setup gzip
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// setup cors
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// _, err := application.NewCommonResource(cfg, ctx)
	// if err != nil {
	// 	return nil, err
	// }

	dashboard := app.Group("bc/dashboard")

	dashboard.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Makan nasi pagi-pagi, ngapain kamu disini?")
	})

	dashboard.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong!")
	})

	return app, nil
}
