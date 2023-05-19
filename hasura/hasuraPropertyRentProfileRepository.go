package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraPropertyRentProfileRepository struct {
}

func (HasuraPropertyRentProfileRepository) CreateRentProfile(rentprofile storemodels.RentProfile) (storemodels.RentProfile, error) {
	return storemodels.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) UpdateRentProfile(rentprofile storemodels.RentProfile) (storemodels.RentProfile, error) {
	return storemodels.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) GetRentProfile(id string) (storemodels.RentProfile, error) {
	return storemodels.RentProfile{}, nil
}
func (HasuraPropertyRentProfileRepository) DeleteRentProfile(id string) error {
	return nil
}
