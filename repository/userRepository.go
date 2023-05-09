package repository

import "github.com/ldehner/fiber-rental-api/models"

type UserRepository interface {
	CreateUser(user models.User) models.User
	GetUser(id int) models.User
	GetUsers() []models.User
	UpdateUser(user models.User) models.User
	DeleteUser(id int) error
}
