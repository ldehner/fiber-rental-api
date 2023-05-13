package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraPropertyRepository struct {
}

func (HasuraPropertyRepository) CreateProperty(property storemodels.Property) (storemodels.Property, error) {
	return storemodels.Property{}, nil
}
func (HasuraPropertyRepository) UpdateProperty(property storemodels.Property) (storemodels.Property, error) {
	return storemodels.Property{}, nil
}
func (HasuraPropertyRepository) GetProperty(id string) (storemodels.Property, error) {
	return storemodels.Property{}, nil
}
func (HasuraPropertyRepository) DeleteProperty(id string) error {
	return nil
}
func (HasuraPropertyRepository) GetProperties() ([]storemodels.Property, error) {
	return []storemodels.Property{}, nil
}
func (HasuraPropertyRepository) UpdatePropertyTenant(id string, tenantId string) (storemodels.Property, error) {
	return storemodels.Property{}, nil
}
func (HasuraPropertyRepository) UpdatePropertyLandlord(id string, landlordId string) (storemodels.Property, error) {
	return storemodels.Property{}, nil
}
func (HasuraPropertyRepository) CreatePropertyInvite(invite storemodels.Invite) (storemodels.Invite, error) {
	return storemodels.Invite{}, nil
}
func (HasuraPropertyRepository) GetPropertyInvites(id string) (storemodels.Invite, error) {
	return storemodels.Invite{}, nil
}

func (HasuraPropertyRepository) DeletePropertyInvite(id string) error {
	return nil
}
