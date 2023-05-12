package repository

import "github.com/ldehner/fiber-rental-api/models"

type PropertyRentProfileRepository interface {
	CreateRentProfile(rentprofile models.RentProfile) (models.RentProfile, error)
	UpdateRentProfile(rentprofile models.RentProfile) (models.RentProfile, error)
	GetRentProfile(id string) (models.RentProfile, error)
	DeleteRentProfile(id string) error
}
