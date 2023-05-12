package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func CreateResponseUSer(userModel models.User) User {
	return User{
		ID:        userModel.Id,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
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
	return c.Status(fiber.StatusCreated).JSON(CreateResponseUSer(dbuser))
}
