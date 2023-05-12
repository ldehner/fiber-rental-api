package repository

import (
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraUserSearchProfileRepository struct {
}

func (HasuraUserSearchProfileRepository) CreateSearchProfile(searchprofile models.SearchProfile) (models.SearchProfile, error) {
	return models.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) UpdateSearchProfile(searchprofile models.SearchProfile) (models.SearchProfile, error) {
	return models.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) GetSearchProfile(id string) (models.SearchProfile, error) {
	return models.SearchProfile{}, nil
}
func (HasuraUserSearchProfileRepository) DeleteSearchProfile(id string) error {
	return nil
}
