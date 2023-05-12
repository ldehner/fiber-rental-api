package repository

import "github.com/ldehner/fiber-rental-api/models"

type UserTenantinfoRepository interface {
	CreateTenantinfo(info models.TenantInfo) (models.TenantInfo, error)
	UpdateTenantinfo(info models.TenantInfo) (models.TenantInfo, error)
	GetTenantinfo(id string) (models.TenantInfo, error)
	DeleteTenantinfo(id string) error
}
