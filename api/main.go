package main

import (
	"boomin_game_api/src/database"
	"boomin_game_api/src/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCecheChannel()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":1129")
}
