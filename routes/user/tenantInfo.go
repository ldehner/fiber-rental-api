package userroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ldehner/fiber-rental-api/conf"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
)

// GetTenantInfo godoc
// @Summary Get a user's tenant info
// @Description Get a user's tenant info
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 302 {object} responsemodels.TenantInfo
// @Router /user/tenantinfo/{id} [get]
func GetTenantInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	tenantInfo, err := conf.Conf{}.GetUserTenantInfoRepository().GetTenantinfo(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseTenantInfo(tenantInfo))
}

// UpdateTenantInfo godoc
// @Summary Update a user's tenant info
// @Description Update a user's tenant info
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param tenantInfo body requestmodels.TenantInfo true "Tenant Info"
// @Success 200 {object} responsemodels.TenantInfo
// @Router /user/tenantinfo/{id} [put]
func UpdateTenantInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	var tenantInfo requestmodels.TenantInfo
	if err := c.BodyParser(&tenantInfo); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbtenantInfo, err := conf.Conf{}.GetUserTenantInfoRepository().UpdateTenantinfo(CreateStoreTenantInfo(tenantInfo, id))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseTenantInfo(dbtenantInfo))
}

// CreateTenantInfo godoc
// @Summary Create a user's tenant info
// @Description Create a user's tenant info
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param tenantInfo body requestmodels.TenantInfo true "Tenant Info"
// @Success 201 {object} responsemodels.TenantInfo
// @Router /user/tenantinfo/{id} [post]
func CreateTenantInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	var tenantInfo requestmodels.TenantInfo
	if err := c.BodyParser(&tenantInfo); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbtenantInfo, err := conf.Conf{}.GetUserTenantInfoRepository().CreateTenantinfo(CreateStoreTenantInfo(tenantInfo, id))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(CreateResponseTenantInfo(dbtenantInfo))
}

// DeleteTenantInfo godoc
// @Summary Delete a user's tenant info
// @Description Delete a user's tenant info
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /user/tenantinfo/{id} [delete]
func DeleteTenantInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := conf.Conf{}.GetUserTenantInfoRepository().DeleteTenantinfo(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Tenant info deleted")
}
