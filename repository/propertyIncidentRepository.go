package repository

import "github.com/ldehner/fiber-rental-api/models"

type PropertyIncidentRepository interface {
	CreateIncident(incident models.Incident) (models.Incident, error)
	UpdateIncident(incident models.Incident) (models.Incident, error)
	GetIncident(propertyId string, incidentId string) (models.Incident, error)
	DeleteIncident(id string) error
	GetIncidents(propertyId string) ([]models.Incident, error)
}
