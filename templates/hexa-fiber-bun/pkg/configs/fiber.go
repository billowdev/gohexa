package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberHttpServiceParams struct {
	Port    string
	Address string
}

func NewFiberHttpServiceParams() *FiberHttpServiceParams {

	fmt.Println("Port = ", SERVER_HTTP_PORT)

	return &FiberHttpServiceParams{
		Port:    SERVER_HTTP_PORT,
		Address: "",
	}
}

func NewFiberHTTPService(params *FiberHttpServiceParams) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:   false,
		BodyLimit: 200 * 1024 * 1024,
	})
	app.Use(logger.New())
	// storageservice.NewS3Client(configs.S3REGION)
	// app.Use(requestid.New())
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
		AllowHeaders: "*",
	}))

	log.Printf("Starting Fiber HTTP listener at Port [%s]...", params.Port)

	return app
}
