package main

import (
	"demo/handlers"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	PORT string
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "3000", "--port=3000 or -port=3000 or --port 3000 or -port 3000")
		flag.Parse()
	}

	app := fiber.New()
	log.Println("fiber started and running on port", PORT)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	UserHandler := handlers.NewUserHandler("users.data")
	app.Post("/users", UserHandler.Create)

	app.Listen(":" + PORT)
}
