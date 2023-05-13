package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
)

// CreateMarketProfile godoc
// @Summary Create a new market profile
// @Description Create a new market profile
// @Tags MarketProfile
// @Accept  json
// @Produce  json
// @Param marketProfile body MarketProfile true "Market Profile"
// @Success 200 {object} MarketProfile
// @Router /property/marketprofile/{id} [post]
func CreateMarketProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var marketProfile requestmodels.MarketProfile
	if err := c.BodyParser(&marketProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbmarketProfile, err := conf.Conf{}.GetPropertyMarketProfileRepository().CreateMarketProfile(CreateStoreMarketProfile(marketProfile, id))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseMarketProfile(dbmarketProfile))
}

// GetMarketProfile godoc
// @Summary Get a market profile
// @Description Get a market profile
// @Tags MarketProfile
// @Accept  json
// @Produce  json
// @Param id path string true "Market Profile ID"
// @Success 302 {object} MarketProfile
// @Router /property/marketprofile/{id} [get]
func GetMarketProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	marketProfile, err := conf.Conf{}.GetPropertyMarketProfileRepository().GetMarketProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseMarketProfile(marketProfile))
}

// UpdateMarketProfile godoc
// @Summary Update a market profile
// @Description Update a market profile
// @Tags MarketProfile
// @Accept  json
// @Produce  json
// @Param id path string true "Market Profile ID"
// @Param marketProfile body MarketProfile true "Market Profile"
// @Success 200 {object} MarketProfile
// @Router /property/marketprofile/{id} [put]
func UpdateMarketProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var marketProfile requestmodels.MarketProfile
	if err := c.BodyParser(&marketProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbmarketProfile, err := conf.Conf{}.GetPropertyMarketProfileRepository().UpdateMarketProfile(CreateStoreMarketProfile(marketProfile, id))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseMarketProfile(dbmarketProfile))
}

// DeleteMarketProfile godoc
// @Summary Delete a market profile
// @Description Delete a market profile
// @Tags MarketProfile
// @Accept  json
// @Produce  json
// @Param id path string true "Market Profile ID"
// @Success 200
// @Router /property/marketprofile/{id} [delete]
func DeleteMarketProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetPropertyMarketProfileRepository().DeleteMarketProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("MarketProfile deleted")
}

// GetMarketProfiles godoc
// @Summary Get all market profiles
// @Description Get all market profiles
// @Tags MarketProfile
// @Accept  json
// @Produce  json
// @Success 302 {array} MarketProfile
// @Router /property/marketprofiles [get]
func GetMarketProfiles(c *fiber.Ctx) error {
	marketProfiles, err := conf.Conf{}.GetPropertyMarketProfileRepository().GetMarketProfiles()
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	var responseProperties []responsemodels.MarketProfile
	for _, property := range marketProfiles {
		responseProperties = append(responseProperties, CreateResponseMarketProfile(property))
	}
	return c.Status(fiber.StatusOK).JSON(responseProperties)
}
