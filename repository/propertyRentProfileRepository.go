package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type PropertyRentProfileRepository interface {
	CreateRentProfile(rentprofile storemodels.RentProfile) (storemodels.RentProfile, error)
	UpdateRentProfile(rentprofile storemodels.RentProfile) (storemodels.RentProfile, error)
	GetRentProfile(id string) (storemodels.RentProfile, error)
	DeleteRentProfile(id string) error
}
