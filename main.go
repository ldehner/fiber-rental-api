package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/routes"
)

func setupPublicRoutes(app *fiber.App) {
	// user endpoints

	// property endpoints
}

func setupPrivateRoutes(app *fiber.App) {
	// user endpoints
	app.Post("/user", routes.CreateUser)
	// property endpoints
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	setupPublicRoutes(app)
	setupPrivateRoutes(app)

	app.Listen(":3000")
}
