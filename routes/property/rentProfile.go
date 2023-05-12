package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type RentProfile struct {
	ContractDue string  `json:"ContractDue"`
	HeatCosts   float32 `json:"HeatCosts"`
	Deposit     float32 `json:"Deposit"`
	Id          string  `json:"Id"`
	Minimum     uint8   `json:"Minimum"`
	Rent        float32 `json:"Rent"`
	RentalStart string  `json:"RentalStart"`
}

// CreateRentProfile godoc
// @Summary Create a rent profile
// @Description Create a rent profile
// @ID create-rent-profile
// @Accept  json
// @Produce  json
// @Param rentProfile body RentProfile true "Rent Profile"
// @Success 200 {object} RentProfile
// @Router /property/rentprofile/{id} [post]
func CreateRentProfile(c *fiber.Ctx) error {
	var rentProfile models.RentProfile
	if err := c.BodyParser(&rentProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbrentProfile, err := conf.Conf{}.GetPropertyRentProfileRepository().CreateRentProfile(rentProfile)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseRentProfile(dbrentProfile))
}

// GetRentProfile godoc
// @Summary Get a rent profile
// @Description Get a rent profile
// @ID get-rent-profile
// @Accept  json
// @Produce  json
// @Param id path string true "Rent Profile ID"
// @Success 302 {object} RentProfile
// @Router /property/rentprofile/{id} [get]
func GetRentProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	rentProfile, err := conf.Conf{}.GetPropertyRentProfileRepository().GetRentProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseRentProfile(rentProfile))
}

// UpdateRentProfile godoc
// @Summary Update a rent profile
// @Description Update a rent profile
// @ID update-rent-profile
// @Accept  json
// @Produce  json
// @Param id path string true "Rent Profile ID"
// @Param rentProfile body RentProfile true "Rent Profile"
// @Success 200 {object} RentProfile
// @Router /property/rentprofile/{id} [patch]
func UpdateRentProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var rentProfile models.RentProfile
	if err := c.BodyParser(&rentProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	rentProfile.Id = id
	dbrentProfile, err := conf.Conf{}.GetPropertyRentProfileRepository().UpdateRentProfile(rentProfile)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseRentProfile(dbrentProfile))
}

// DeleteRentProfile godoc
// @Summary Delete a rent profile
// @Description Delete a rent profile
// @ID delete-rent-profile
// @Accept  json
// @Produce  json
// @Param id path string true "Rent Profile ID"
// @Success 200
// @Router /property/rentprofile/{id} [delete]
func DeleteRentProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetPropertyRentProfileRepository().DeleteRentProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Rent Profile deleted")
}