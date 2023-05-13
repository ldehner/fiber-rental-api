package repository

import storemodels "github.com/ldehner/fiber-rental-api/models/store"

type UserRepository interface {
	CreateUser(user storemodels.User) (storemodels.User, error)
	GetUser(id string) (storemodels.User, error)
	GetUsers() ([]storemodels.User, error)
	UpdateUser(user storemodels.User) (storemodels.User, error)
	DeleteUser(id string) error
	UpdateContactInfo(contactInfo storemodels.ContactInfo) (storemodels.User, error)
}
