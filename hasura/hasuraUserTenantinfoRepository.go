package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraUserTenantinfoRepository struct {
}

func (HasuraUserTenantinfoRepository) CreateTenantinfo(info storemodels.TenantInfo) (storemodels.TenantInfo, error) {
	return storemodels.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) UpdateTenantinfo(info storemodels.TenantInfo) (storemodels.TenantInfo, error) {
	return storemodels.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) GetTenantinfo(id string) (storemodels.TenantInfo, error) {
	return storemodels.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) DeleteTenantinfo(id string) error {
	return nil
}
