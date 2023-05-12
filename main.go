package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/ldehner/fiber-rental-api/docs"
	propertyroutes "github.com/ldehner/fiber-rental-api/routes/property"
	userroutes "github.com/ldehner/fiber-rental-api/routes/user"
)

func setupPublicRoutes(app *fiber.App) {
	// user endpoints

	// property endpoints
}

func setupPrivateRoutes(app *fiber.App) {
	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	// user endpoints
	app.Post("/user/create", userroutes.CreateUser)
	app.Get("/user/:id", userroutes.GetUser)
	app.Get("/user", userroutes.GetUsers)
	app.Patch("/user/update", userroutes.UpdateUser)
	app.Patch("/user/contactinfo/:id", userroutes.UpdateContactInfo)
	app.Delete("/user/:id", userroutes.DeleteUser)
	app.Post("/user/tenantinfo/:id", userroutes.CreateTenantInfo)
	app.Patch("/user/tenantinfo/:id", userroutes.UpdateTenantInfo)
	app.Get("/user/tenantinfo/:id", userroutes.GetTenantInfo)
	app.Delete("/user/tenantinfo/:id", userroutes.DeleteTenantInfo)
	app.Post("/user/searchprofile/:id", userroutes.CreateSearchProfile)
	app.Patch("/user/searchprofile/:id", userroutes.UpdateSearchProfile)
	app.Get("/user/searchprofile/:id", userroutes.GetSearchProfile)
	app.Delete("/user/searchprofile/:id", userroutes.DeleteSearchProfile)
	// property endpoints
	app.Post("/property/create", propertyroutes.CreateProperty)
	app.Get("/property/:id", propertyroutes.GetProperty)
	app.Get("/property", propertyroutes.GetProperties)
	app.Patch("/property/update", propertyroutes.UpdateProperty)
	app.Delete("/property/:id", propertyroutes.DeleteProperty)
	app.Patch("/property/tenant/:id/:tenantId", propertyroutes.UpdatePropertyTenant)
	app.Patch("/property/landlord/:id/:landlordId", propertyroutes.UpdatePropertyLandlord)
	app.Post("/property/marketprofile/:id", propertyroutes.CreateMarketProfile)
	app.Patch("/property/marketprofile/:id", propertyroutes.UpdateMarketProfile)
	app.Get("/property/marketprofile/:id", propertyroutes.GetMarketProfile)
	app.Get("/property/marketprofile", propertyroutes.GetMarketProfiles)
	app.Delete("/property/marketprofile/:id", propertyroutes.DeleteMarketProfile)
	app.Post("/property/rentprofile/:id", propertyroutes.CreateRentProfile)
	app.Patch("/property/rentprofile/:id", propertyroutes.UpdateRentProfile)
	app.Get("/property/rentprofile/:id", propertyroutes.GetRentProfile)
	app.Delete("/property/rentprofile/:id", propertyroutes.DeleteRentProfile)
	app.Post("/property/invite/:id/:userId", propertyroutes.AcceptInvite)
	app.Get("/property/invite/:propertyId", propertyroutes.CreateInvite)
	app.Post("/property/incident:propertyId", propertyroutes.CreateIncident)
	app.Patch("/property/incident/:propertyId/:incidentId", propertyroutes.UpdateIncident)
	app.Get("/property/incident/:propertyId/:incidentId", propertyroutes.GetIncident)
	app.Get("/property/incident/:propertyId", propertyroutes.GetIncidents)
	app.Delete("/property/incident/:propertyId/:incidentId", propertyroutes.DeleteIncident)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	setupPublicRoutes(app)
	setupPrivateRoutes(app)

	app.Listen(":3000")
}
