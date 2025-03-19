package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	isProd := os.Getenv("ENV") == "production"
	app := fiber.New()

	log.SetLevel(log.LevelInfo)

	app.Use(logger.New(logger.Config{
		Format: func() string {
			if isProd {
				return `{"time":"${time}","status":${status},"latency":"${latency}","method":"${method}","path":"${path}"}` + "\n"
			}
			return "${time} | ${status} | ${latency} | ${method} ${path}\n"
		}(),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Traitement de la requête sur la route principale")
		return c.SendString("Hello, World!")
	})

	// Démarrage du serveur
	log.Info("Démarrage du serveur sur le port 3000")
	app.Listen(":3000")
}
