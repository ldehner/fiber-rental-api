package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type ContactInfo struct {
	Phone string `json:"Phone"`
	Mail  string `json:"Mail"`
}

type User struct {
	Id          string `json:"Id"`
	FirstName   string `json:"First"`
	LastName    string `json:"Last"`
	Phone       string `json:"Phone"`
	Mail        string `json:"Mail"`
	Country     string `json:"Country"`
	City        string `json:"City"`
	Street      string `json:"Street"`
	Housenumber string `json:"Housenumber"`
	Apartment   string `json:"Apartment"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "User"
// @Success 201 {object} User
// @Router /user/create [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(CreateResponseUser(dbuser))
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 302 {object} User
// @Router /user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := conf.Conf{}.GetUserRepository().GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseUser(user))
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 302 {object} []User
// @Router /user [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := conf.Conf{}.GetUserRepository().GetUsers()
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	var responseUsers []User
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}
	return c.Status(fiber.StatusFound).JSON(responseUsers)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body User true "User"
// @Success 200 {object} User
// @Router /user/update [patch]
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseUser(dbuser))
}

// UpdateContactInfo godoc
// @Summary Update a user's contact info
// @Description Update a user's contact info
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param contactInfo body ContactInfo true "Contact Info"
// @Success 200 {object} User
// @Router /user/contactinfo/{id} [patch]
func UpdateContactInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	var contactInfo models.ContactInfo
	if err := c.BodyParser(&contactInfo); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().UpdateContactInfo(id, contactInfo)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseUser(dbuser))
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
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
