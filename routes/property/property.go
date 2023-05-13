package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	responsemodels "github.com/ldehner/fiber-rental-api/models/response"
)

// CreateProperty godoc
// @Summary Create a new property
// @Description Create a new property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param property body requestmodels.Property true "Property"
// @Success 200 {object} responsemodels.Property
// @Router /property [post]
func CreateProperty(c *fiber.Ctx) error {
	var property requestmodels.Property
	if err := c.BodyParser(&property); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	dbproperty, err := conf.Conf{}.GetPropertyRepository().CreateProperty(CreateStoreProperty(property, uuid.New().String()))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(dbproperty))
}

// GetProperty godoc
// @Summary Get a property
// @Description Get a property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Property ID"
// @Success 200 {object} responsemodels.Property
// @Router /property/{id} [get]
func GetProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	property, err := conf.Conf{}.GetPropertyRepository().GetProperty(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(property))
}

// GetProperties godoc
// @Summary Get all properties
// @Description Get all properties
// @Tags Property
// @Accept  json
// @Produce  json
// @Success 200 {array} responsemodels.Property
// @Router /property [get]
func GetProperties(c *fiber.Ctx) error {
	properties, err := conf.Conf{}.GetPropertyRepository().GetProperties()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	var responseProperties []responsemodels.Property
	for _, property := range properties {
		responseProperties = append(responseProperties, CreateResponseProperty(property))
	}
	return c.Status(fiber.StatusOK).JSON(responseProperties)
}

// UpdateProperty godoc
// @Summary Update a property
// @Description Update a property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Property ID"
// @Param property body requestmodels.Property true "Property"
// @Success 200 {object} responsemodels.Property
// @Router /property/{id} [put]
func UpdateProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var property requestmodels.Property
	if err := c.BodyParser(&property); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	dbproperty, err := conf.Conf{}.GetPropertyRepository().UpdateProperty(CreateStoreProperty(property, id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(dbproperty))
}

// DeleteProperty godoc
// @Summary Delete a property
// @Description Delete a property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Property ID"
// @Success 200
// @Router /property/{id} [delete]
func DeleteProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetPropertyRepository().DeleteProperty(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

// UpdatePropertyTenant godoc
// @Summary Update a property tenant
// @Description Update a property tenant
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Property ID"
// @Param tenantId path string true "Tenant ID"
// @Success 200
// @Router /property/tenant/{id}/{tenantId} [put]
func UpdatePropertyTenant(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantId := c.Params("tenantId")
	property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyTenant(id, tenantId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(property))
}

// UpdatePropertyLandlord godoc
// @Summary Update a property landlord
// @Description Update a property landlord
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path string true "Property ID"
// @Param landlordId path string true "Landlord ID"
// @Success 200
// @Router /property/landlord/{id}/{landlordId} [put]
func UpdatePropertyLandlord(c *fiber.Ctx) error {
	id := c.Params("id")
	landlordId := c.Params("landlordId")
	property, err := conf.Conf{}.GetPropertyRepository().UpdatePropertyLandlord(id, landlordId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseProperty(property))
}
