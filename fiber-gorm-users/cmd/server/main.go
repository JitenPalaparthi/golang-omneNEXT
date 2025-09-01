package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"fiber-gorm-users/internal/config"
	"fiber-gorm-users/internal/db"
	"fiber-gorm-users/internal/handlers"
	"fiber-gorm-users/internal/logger"
	"fiber-gorm-users/internal/models"
)

func main() {
	_ = godotenv.Load()

	log := logger.New()

	// Connect DB
	gormDB, err := db.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	// Auto-migrate
	if err := gormDB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal().Err(err).Msg("failed to automigrate")
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          10 * time.Second,
	})

	// CORS
	origin := os.Getenv("CORS_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: origin,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Routes
	usersHandler := handlers.NewUsersHandler(gormDB, log)
	app.Post("/users", usersHandler.Create)
	app.Get("/users", usersHandler.List)
	app.Get("/users/:id", usersHandler.GetByID)
	app.Put("/users/:id", usersHandler.Update)
	app.Delete("/users/:id", usersHandler.Delete)

	port := config.Port()
	log.Info().Msgf("server listening on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal().Err(err).Msg("fiber failed")
	}
	fmt.Printf("")
}
