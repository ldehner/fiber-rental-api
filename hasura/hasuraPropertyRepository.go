package hasura

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraPropertyRepository struct {
}

func (HasuraPropertyRepository) CreateProperty(property models.Property) (models.Property, error) {
	return models.Property{}, nil
}
func (HasuraPropertyRepository) UpdateProperty(property models.Property) (models.Property, error) {
	return models.Property{}, nil
}
func (HasuraPropertyRepository) GetProperty(id string) (models.Property, error) {
	return models.Property{}, nil
}
func (HasuraPropertyRepository) DeleteProperty(id string) error {
	return nil
}
func (HasuraPropertyRepository) GetProperties() ([]models.Property, error) {
	return []models.Property{}, nil
}
func (HasuraPropertyRepository) UpdatePropertyTenant(id string, tenantId string) (models.Property, error) {
	return models.Property{}, nil
}
func (HasuraPropertyRepository) UpdatePropertyLandlord(id string, landlordId string) (models.Property, error) {
	return models.Property{}, nil
}
func (HasuraPropertyRepository) CreatePropertyInvite(invite models.Invite) (models.Invite, error) {
	return models.Invite{}, nil
}
func (HasuraPropertyRepository) GetPropertyInvites(id string) (models.Invite, error) {
	return models.Invite{}, nil
}

func (HasuraPropertyRepository) DeletePropertyInvite(id string) error {
	return nil
}
