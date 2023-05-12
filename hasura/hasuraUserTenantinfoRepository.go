package hasura

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraUserTenantinfoRepository struct {
}

func (HasuraUserTenantinfoRepository) CreateTenantinfo(info models.TenantInfo) (models.TenantInfo, error) {
	return models.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) UpdateTenantinfo(info models.TenantInfo) (models.TenantInfo, error) {
	return models.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) GetTenantinfo(id string) (models.TenantInfo, error) {
	return models.TenantInfo{}, nil
}
func (HasuraUserTenantinfoRepository) DeleteTenantinfo(id string) error {
	return nil
}
