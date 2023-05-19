package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type PropertyRepository interface {
	CreateProperty(property storemodels.Property) (storemodels.Property, error)
	UpdateProperty(property storemodels.Property) (storemodels.Property, error)
	GetProperty(id string) (storemodels.Property, error)
	DeleteProperty(id string) error
	GetProperties() ([]storemodels.Property, error)
	UpdatePropertyTenant(id string, tenantId string) (storemodels.Property, error)
	UpdatePropertyLandlord(id string, landlordId string) (storemodels.Property, error)
	CreatePropertyInvite(invite storemodels.Invite) (storemodels.Invite, error)
	GetPropertyInvites(id string) (storemodels.Invite, error)
	DeletePropertyInvite(id string) error
}
