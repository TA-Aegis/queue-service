package rest

import (
	"antrein/bc-dashboard/application/common/resource"
	"antrein/bc-dashboard/application/common/usecase"
	"antrein/bc-dashboard/internal/handler/grpc/analytic"
	"antrein/bc-dashboard/internal/handler/rest/auth"
	"antrein/bc-dashboard/internal/handler/rest/project"
	"antrein/bc-dashboard/model/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ApplicationDelegate(cfg *config.Config, uc *usecase.CommonUsecase, rsc *resource.CommonResource) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName: "BC Dashboard",
	})

	// setup gzip
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// setup cors
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	dashboard := app.Group("bc/dashboard")

	dashboard.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Makan nasi pagi-pagi, ngapain kamu disini?")
	})

	dashboard.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong!")
	})

	// routes

	// auth
	authRoute := auth.New(cfg, uc.AuthUsecase, rsc.Vld)
	authRoute.RegisterRoute(app)

	// project
	projectRoute := project.New(cfg, uc.ProjectUsecase, uc.ConfigUsecase, rsc.Vld)
	projectRoute.RegisterRoute(app)

	// analytic
	analyticRouter := analytic.New(cfg, rsc.GRPC)
	analyticRouter.RegisterRoute(app)

	return app, nil
}
