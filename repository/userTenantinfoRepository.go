package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type UserTenantinfoRepository interface {
	CreateTenantinfo(info storemodels.TenantInfo) (storemodels.TenantInfo, error)
	UpdateTenantinfo(info storemodels.TenantInfo) (storemodels.TenantInfo, error)
	GetTenantinfo(id string) (storemodels.TenantInfo, error)
	DeleteTenantinfo(id string) error
}
