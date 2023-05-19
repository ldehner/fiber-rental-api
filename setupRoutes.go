package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/middleware"
	propertyroutes "github.com/ldehner/fiber-rental-api/routes/property"
	userroutes "github.com/ldehner/fiber-rental-api/routes/user"
)

func setupPublicRoutes(app *fiber.App) {
	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Post("/user", userroutes.CreateUser)
}

func setupPrivateRoutes(app *fiber.App) {
	app.Use(middleware.JwtMiddleware)

	// user endpoints
	app.Get("/user/:id", userroutes.GetUser)
	app.Get("/user", userroutes.GetUsers)
	app.Put("/user/:id", userroutes.UpdateUser)
	app.Put("/user/contactinfo/:id", userroutes.UpdateContactInfo)
	app.Delete("/user/:id", userroutes.DeleteUser)
	app.Post("/user/tenantinfo/:id", userroutes.CreateTenantInfo)
	app.Put("/user/tenantinfo/:id", userroutes.UpdateTenantInfo)
	app.Get("/user/tenantinfo/:id", userroutes.GetTenantInfo)
	app.Delete("/user/tenantinfo/:id", userroutes.DeleteTenantInfo)
	app.Post("/user/searchprofile/:id", userroutes.CreateSearchProfile)
	app.Put("/user/searchprofile/:id", userroutes.UpdateSearchProfile)
	app.Get("/user/searchprofile/:id", userroutes.GetSearchProfile)
	app.Delete("/user/searchprofile/:id", userroutes.DeleteSearchProfile)
	// property endpoints
	app.Post("/property", propertyroutes.CreateProperty)
	app.Get("/property/:id", propertyroutes.GetProperty)
	app.Get("/property", propertyroutes.GetProperties)
	app.Put("/property/:id", propertyroutes.UpdateProperty)
	app.Delete("/property/:id", propertyroutes.DeleteProperty)
	app.Put("/property/tenant/:id/:tenantId", propertyroutes.UpdatePropertyTenant)
	app.Put("/property/landlord/:id/:landlordId", propertyroutes.UpdatePropertyLandlord)
	app.Post("/property/marketprofile/:id", propertyroutes.CreateMarketProfile)
	app.Put("/property/marketprofile/:id", propertyroutes.UpdateMarketProfile)
	app.Get("/property/marketprofile/:id", propertyroutes.GetMarketProfile)
	app.Get("/property/marketprofile", propertyroutes.GetMarketProfiles)
	app.Delete("/property/marketprofile/:id", propertyroutes.DeleteMarketProfile)
	app.Post("/property/rentprofile/:id", propertyroutes.CreateRentProfile)
	app.Put("/property/rentprofile/:id", propertyroutes.UpdateRentProfile)
	app.Get("/property/rentprofile/:id", propertyroutes.GetRentProfile)
	app.Delete("/property/rentprofile/:id", propertyroutes.DeleteRentProfile)
	app.Post("/property/invite/:id/:userId", propertyroutes.AcceptInvite)
	app.Put("/property/invite/:propertyId", propertyroutes.CreateInvite)
	app.Post("/property/incident:propertyId", propertyroutes.CreateIncident)
	app.Put("/property/incident/:propertyId/:incidentId", propertyroutes.UpdateIncident)
	app.Get("/property/incident/:propertyId/:incidentId", propertyroutes.GetIncident)
	app.Get("/property/incident/:propertyId", propertyroutes.GetIncidents)
	app.Delete("/property/incident/:propertyId/:incidentId", propertyroutes.DeleteIncident)
}
