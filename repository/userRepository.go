package repository

import "github.com/ldehner/fiber-rental-api/models"

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
	UpdateContactInfo(id string, contactInfo models.ContactInfo) (models.User, error)
}
