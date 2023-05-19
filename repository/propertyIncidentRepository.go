package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type PropertyIncidentRepository interface {
	CreateIncident(incident storemodels.Incident) (storemodels.Incident, error)
	UpdateIncident(incident storemodels.Incident) (storemodels.Incident, error)
	GetIncident(propertyId string, incidentId string) (storemodels.Incident, error)
	DeleteIncident(propertyId string, incidentId string) error
	GetIncidents(propertyId string) ([]storemodels.Incident, error)
}
