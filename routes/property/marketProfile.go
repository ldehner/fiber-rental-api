package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type MarketProfile struct {
	Rent             float32 `json:"Rent"`
	PowerPrice       float32 `json:"PowerPrice"`
	Period           uint8   `json:"Period"`
	MinimumPeriod    uint8   `json:"MinimumPeriod"`
	MinimumIncome    float32 `json:"MinimumIncome"`
	Id               string  `json:"Id"`
	HeatPrice        float32 `json:"HeatPrice"`
	Description      string  `json:"Description"`
	AvailabilityDate string  `json:"AvailabilityDate"`
	Deposit          float32 `json:"Deposit"`
}

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
	var marketProfile models.MarketProfile
	if err := c.BodyParser(&marketProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbmarketProfile, err := conf.Conf{}.GetPropertyMarketProfileRepository().CreateMarketProfile(marketProfile)
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
	var marketProfile models.MarketProfile
	if err := c.BodyParser(&marketProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	marketProfile.Id = id
	dbmarketProfile, err := conf.Conf{}.GetPropertyMarketProfileRepository().UpdateMarketProfile(marketProfile)
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
	var responseProperties []MarketProfile
	for _, property := range marketProfiles {
		responseProperties = append(responseProperties, CreateResponseMarketProfile(property))
	}
	return c.Status(fiber.StatusOK).JSON(responseProperties)
}
