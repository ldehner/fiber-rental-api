package repository

import "github.com/ldehner/fiber-rental-api/models"

type UserSearchprofileRepository interface {
	CreateSearchProfile(searchprofile models.SearchProfile) (models.SearchProfile, error)
	UpdateSearchProfile(searchprofile models.SearchProfile) (models.SearchProfile, error)
	GetSearchProfile(id string) (models.SearchProfile, error)
	DeleteSearchProfile(id string) error
}
