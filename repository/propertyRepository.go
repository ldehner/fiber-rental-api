package repository

import "github.com/ldehner/fiber-rental-api/models"

type PropertyRepository interface {
	CreateProperty(property models.Property) (models.Property, error)
	UpdateProperty(property models.Property) (models.Property, error)
	GetProperty(id string) (models.Property, error)
	DeleteProperty(id string) error
	GetProperties() ([]models.Property, error)
	UpdatePropertyTenant(id string, tenantId string) (models.Property, error)
	UpdatePropertyLandlord(id string, landlordId string) (models.Property, error)
	CreatePropertyInvite(invite models.Invite) (models.Invite, error)
	GetPropertyInvites(id string) (models.Invite, error)
	DeletePropertyInvite(id string) error
}
