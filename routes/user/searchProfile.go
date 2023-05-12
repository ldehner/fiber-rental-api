package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type SearchProfile struct {
	Budget  float32 `json:"Budget"`
	City    string  `json:"City"`
	Country string  `json:"Country"`
	Radius  float32 `json:"Radius"`
	Rooms   uint8   `json:"Rooms"`
	Size    uint16  `json:"Size"`
	Street  string  `json:"Street"`
	Type    uint8   `json:"Type"`
	Zipcode string  `json:"Zipcode"`
}

// CreateSearchProfile godoc
// @Summary Create a user's search profile
// @Description Create a user's search profile
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param searchProfile body SearchProfile true "Search Profile"
// @Success 201 {object} SearchProfile
// @Router /user/searchprofile/{id} [post]
func CreateSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var searchProfile models.SearchProfile
	if err := c.BodyParser(&searchProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	searchProfile.UserId = id
	dbsearchProfile, err := conf.Conf{}.GetUserSearchProfileRepository().CreateSearchProfile(searchProfile)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(CreateResponseSearchProfile(dbsearchProfile))
}

// GetSearchProfile godoc
// @Summary Get a user's search profile
// @Description Get a user's search profile
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 302 {object} SearchProfile
// @Router /user/searchprofile/{id} [get]
func GetSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	searchProfile, err := conf.Conf{}.GetUserSearchProfileRepository().GetSearchProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseSearchProfile(searchProfile))
}

// UpdateSearchProfile godoc
// @Summary Update a user's search profile
// @Description Update a user's search profile
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param searchProfile body SearchProfile true "Search Profile"
// @Success 200 {object} SearchProfile
// @Router /user/searchprofile/{id} [patch]
func UpdateSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var searchProfile models.SearchProfile
	if err := c.BodyParser(&searchProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	searchProfile.UserId = id
	dbsearchProfile, err := conf.Conf{}.GetUserSearchProfileRepository().UpdateSearchProfile(searchProfile)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseSearchProfile(dbsearchProfile))
}

// DeleteSearchProfile godoc
// @Summary Delete a user's search profile
// @Description Delete a user's search profile
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /user/searchprofile/{id} [delete]
func DeleteSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetUserSearchProfileRepository().DeleteSearchProfile(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Search profile deleted")
}
