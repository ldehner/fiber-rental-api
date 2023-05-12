package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type Invite struct {
	Id       string `json:"Id"`
	Property string `json:"Property"`
	Tenant   string `json:"Tenant"`
	Landlord string `json:"Landlord"`
	ValidDue string `json:"ValidDue"`
}

// CreateInvite godoc
// @Summary Create an invite
// @Description Create an invite
// @Tags Property
// @Accept  json
// @Produce  json
// @Param invite body Invite true "Invite"
// @Success 200 {object} Invite
// @Router /property/invite/{propertyId} [get]
func CreateInvite(c *fiber.Ctx) error {
	var invite models.Invite
	if err := c.BodyParser(&invite); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	invite.Id = uuid.New().String()
	invite.Property = c.Params("propertyId")
	dbinvite, err := conf.Conf{}.GetPropertyRepository().CreatePropertyInvite(invite)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString(dbinvite.Id)
}

// AcceptInvite godoc
// @Summary Accept an invite
// @Description Accept an invite
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Invite ID"
// @Param userId path string true "User ID"
// @Success 200 {object} Invite
// @Router /property/invite/{id}/{userId} [post]
func AcceptInvite(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Params("userId")

	var prop models.Property
	invite, err := conf.Conf{}.GetPropertyRepository().GetPropertyInvites(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	if len(invite.Tenant) > 0 {
		property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyLandlord(invite.Property, userId)
		if err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}
		prop = property
	} else {
		property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyTenant(invite.Property, userId)
		if err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}
		prop = property
	}
	derr := conf.Conf{}.GetPropertyRepository().DeletePropertyInvite(id)
	if derr != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(prop))
}
