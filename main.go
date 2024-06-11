package main

import (
	"cmp"
	"log"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{Network: "tcp"})

	// logger := zerolog.New(os.Stdout).Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
	// 	e.Str("msg", message)
	// }))

	// app.Use(fiberzerolog.New(fiberzerolog.Config{
	// 	Logger: &logger,
	// }))

	app.Use(fiberzerolog.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := cmp.Or(os.Getenv("PORT"), "3000")

	log.Fatal(app.Listen(":" + port))
}
