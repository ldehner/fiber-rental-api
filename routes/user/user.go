package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body requestmodels.User true "User"
// @Success 201 {object} responsemodels.User
// @Router /user/create [post]
func CreateUser(c *fiber.Ctx) error {
	var user requestmodels.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().CreateUser(CreateStoreUser(user))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser.Id = user.Id
	return c.Status(fiber.StatusCreated).JSON(CreateResponseUser(dbuser))
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 302 {object} responsemodels.User
// @Router /user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := conf.Conf{}.GetUserRepository().GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	user.Id = id
	return c.Status(fiber.StatusFound).JSON(CreateResponseUser(user))
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 302 {object} []responsemodels.User
// @Router /user [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := conf.Conf{}.GetUserRepository().GetUsers()
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	var responseUsers []responsemodels.User
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}
	return c.Status(fiber.StatusFound).JSON(responseUsers)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept json
// @Produce json
// @Param user body requestmodels.User true "User"
// @Success 200 {object} responsemodels.User
// @Router /user/update [put]
func UpdateUser(c *fiber.Ctx) error {
	var user requestmodels.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().UpdateUser(CreateStoreUser(user))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseUser(dbuser))
}

// UpdateContactInfo godoc
// @Summary Update a user's contact info
// @Description Update a user's contact info
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param contactInfo body requestmodels.ContactInfo true "Contact Info"
// @Success 200 {object} responsemodels.User
// @Router /user/contactinfo/{id} [put]
func UpdateContactInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	var contactInfo requestmodels.ContactInfo
	if err := c.BodyParser(&contactInfo); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().UpdateContactInfo(CreateStoreContactInfo(contactInfo, id))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseUser(dbuser))
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetUserRepository().DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("User deleted")
}
