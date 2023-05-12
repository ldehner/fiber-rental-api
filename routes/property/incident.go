package propertyroutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ldehner/fiber-rental-api/conf"
	"github.com/ldehner/fiber-rental-api/models"
)

type Incident struct {
	Description string `json:"Description"`
	Id          string `json:"Id"`
	Type        uint8  `json:"Type"`
	TenantId    string `json:"Tenant"`
	Status      int8   `json:"Status"`
	PropertyId  string `json:"Property"`
	LandlordId  string `json:"Landlord"`
}

// CreateIncident godoc
// @Summary Create a new incident
// @Description Create a new incident
// @Tags incidents
// @Accept  json
// @Produce  json
// @Param incident body Incident true "Incident"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /incidents/{propertyId} [post]
func CreateIncident(c *fiber.Ctx) error {
	var incident models.Incident
	if err := c.BodyParser(&incident); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	incident.PropertyId = c.Params("propertyId")
	incident.Id = uuid.New().String()
	dbincident, err := conf.Conf{}.GetPropertyIncidentRepository().CreateIncident(incident)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseIncident(dbincident))
}

// GetIncident godoc
// @Summary Get an incident
// @Description Get an incident
// @Tags incidents
// @Accept  json
// @Produce  json
// @Param id path string true "Incident ID"
// @Success 302 {object} Incident
// @Router /incident/{propertyId}/{incidentId} [get]
func GetIncident(c *fiber.Ctx) error {
	propertyId := c.Params("propertyId")
	incidentId := c.Params("incidentId")
	incident, err := conf.Conf{}.GetPropertyIncidentRepository().GetIncident(propertyId, incidentId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusFound).JSON(CreateResponseIncident(incident))
}

// GetIncidents godoc
// @Summary Get all incidents
// @Description Get all incidents
// @Tags incidents
// @Accept  json
// @Produce  json
// @Param propertyId path string true "Property ID"
// @Success 302 {array} Incident
// @Router /incident/{propertyId} [get]
func GetIncidents(c *fiber.Ctx) error {
	propertyId := c.Params("propertyId")
	incidents, err := conf.Conf{}.GetPropertyIncidentRepository().GetIncidents(propertyId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var responseIncidents []Incident
	for _, incident := range incidents {
		responseIncidents = append(responseIncidents, CreateResponseIncident(incident))
	}
	return c.Status(fiber.StatusFound).JSON(responseIncidents)
}

// UpdateIncident godoc
// @Summary Update an incident
// @Description Update an incident
// @Tags incidents
// @Accept  json
// @Produce  json
// @Param id path string true "Incident ID"
// @Param incident body Incident true "Incident"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /incident/{propertyId}/{incidentId} [patch]
func UpdateIncident(c *fiber.Ctx) error {
	propertyId := c.Params("propertyId")
	incidentId := c.Params("incidentId")
	var incident models.Incident
	if err := c.BodyParser(&incident); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	incident.Id = incidentId
	incident.PropertyId = propertyId
	dbincident, err := conf.Conf{}.GetPropertyIncidentRepository().UpdateIncident(incident)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseIncident(dbincident))
}

// DeleteIncident godoc
// @Summary Delete an incident
// @Description Delete an incident
// @Tags incidents
// @Accept  json
// @Produce  json
// @Param id path string true "Incident ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /incident/{propertyId}/{incidentId} [delete]
func DeleteIncident(c *fiber.Ctx) error {
	propertyId := c.Params("propertyId")
	incidentId := c.Params("incidentId")
	err := conf.Conf{}.GetPropertyIncidentRepository().DeleteIncident(propertyId, incidentId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("OK")
}
