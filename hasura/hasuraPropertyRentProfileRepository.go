package hasura

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraPropertyRentProfileRepository struct {
}

func (HasuraPropertyRentProfileRepository) CreateRentProfile(rentprofile models.RentProfile) (models.RentProfile, error) {
	return models.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) UpdateRentProfile(rentprofile models.RentProfile) (models.RentProfile, error) {
	return models.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) GetRentProfile(id string) (models.RentProfile, error) {
	return models.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) DeleteRentProfile(id string) error {
	return nil
}
