package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraPropertyIncidentRepository struct {
}

func (HasuraPropertyIncidentRepository) CreateIncident(incident storemodels.Incident) (storemodels.Incident, error) {
	return storemodels.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) UpdateIncident(incident storemodels.Incident) (storemodels.Incident, error) {
	return storemodels.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) GetIncident(propertyId string, incidentId string) (storemodels.Incident, error) {
	return storemodels.Incident{}, nil
}
func (HasuraPropertyIncidentRepository) DeleteIncident(propertyId string, incidentId string) error {
	return nil
}
func (HasuraPropertyIncidentRepository) GetIncidents(propertyId string) ([]storemodels.Incident, error) {
	return []storemodels.Incident{}, nil
}
