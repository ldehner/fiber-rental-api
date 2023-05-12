package repository

import "github.com/ldehner/fiber-rental-api/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id string) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
	UpdateContactInfo(id string, contactInfo models.ContactInfo) (models.User, error)
}
