package hasura

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraPropertyIncidentRepository struct {
}

func (HasuraPropertyIncidentRepository) CreateIncident(incident models.Incident) (models.Incident, error) {
	return models.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) UpdateIncident(incident models.Incident) (models.Incident, error) {
	return models.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) GetIncident(propertyId string, incidentId string) (models.Incident, error) {
	return models.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) DeleteIncident(propertyId string, incidentId string) error {
	return nil
}
func (HasuraPropertyIncidentRepository) GetIncidents(propertyId string) ([]models.Incident, error) {
	return []models.Incident{}, nil
}
