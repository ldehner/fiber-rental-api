package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type UserSearchprofileRepository interface {
	CreateSearchProfile(searchprofile storemodels.SearchProfile) (storemodels.SearchProfile, error)
	UpdateSearchProfile(searchprofile storemodels.SearchProfile) (storemodels.SearchProfile, error)
	GetSearchProfile(id string) (storemodels.SearchProfile, error)
	DeleteSearchProfile(id string) error
}
