package propertyroutes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

// CreateInvite godoc
// @Summary Create an invite
// @Description Create an invite
// @Tags Property
// @Accept  json
// @Produce  json
// @Param invite body requestmodels.Invite true "Invite"
// @Success 200 {string} string "InviteId"
// @Router /property/invite/{propertyId} [get]
func CreateInvite(c *fiber.Ctx) error {
	var invite requestmodels.Invite
	if err := c.BodyParser(&invite); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	invite.Property = c.Params("propertyId")
	dbinvite, err := conf.Conf{}.GetPropertyRepository().CreatePropertyInvite(CreateStoreInvite(invite, uuid.New().String(), time.Now().AddDate(0, 0, 7)))
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
// @Success 200 {object} responsemodels.Property
// @Router /property/invite/{id}/{userId} [post]
func AcceptInvite(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Params("userId")
	var invite storemodels.Invite
	var prop responsemodels.Property
	invite, err := conf.Conf{}.GetPropertyRepository().GetPropertyInvites(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	if invite.ValidDue.Before(time.Now()) {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	if len(invite.Tenant) > 0 {
		property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyLandlord(invite.Property, userId)
		if err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}
		prop = CreateResponseProperty(property)
	} else {
		property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyTenant(invite.Property, userId)
		if err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}
		prop = CreateResponseProperty(property)
	}
	derr := conf.Conf{}.GetPropertyRepository().DeletePropertyInvite(id)
	if derr != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(prop)
}
