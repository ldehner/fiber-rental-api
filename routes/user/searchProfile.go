package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
)

// CreateSearchProfile godoc
// @Summary Create a user's search profile
// @Description Create a user's search profile
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param searchProfile body requestmodels.SearchProfile true "Search Profile"
// @Success 201 {object} responsemodels.SearchProfile
// @Router /user/searchprofile/{id} [post]
func CreateSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var searchProfile requestmodels.SearchProfile
	if err := c.BodyParser(&searchProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbsearchProfile, err := conf.Conf{}.GetUserSearchProfileRepository().CreateSearchProfile(CreateStoreSearchProfile(searchProfile, id))
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
// @Success 302 {object} responsemodels.SearchProfile
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
// @Param searchProfile body requestmodels.SearchProfile true "Search Profile"
// @Success 200 {object} responsemodels.SearchProfile
// @Router /user/searchprofile/{id} [put]
func UpdateSearchProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var searchProfile requestmodels.SearchProfile
	if err := c.BodyParser(&searchProfile); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbsearchProfile, err := conf.Conf{}.GetUserSearchProfileRepository().UpdateSearchProfile(CreateStoreSearchProfile(searchProfile, id))
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
