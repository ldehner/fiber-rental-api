package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraUserSearchProfileRepository struct {
}

func (HasuraUserSearchProfileRepository) CreateSearchProfile(searchprofile storemodels.SearchProfile) (storemodels.SearchProfile, error) {
	return storemodels.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) UpdateSearchProfile(searchprofile storemodels.SearchProfile) (storemodels.SearchProfile, error) {
	return storemodels.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) GetSearchProfile(id string) (storemodels.SearchProfile, error) {
	return storemodels.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) DeleteSearchProfile(id string) error {
	return nil
}
