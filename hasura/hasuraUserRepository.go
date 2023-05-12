package hasura

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraUserRepository struct {
}

func (HasuraUserRepository) CreateUser(user models.User) (models.User, error) {
	json, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	status, err := PostRequest(json, "/user")
	if err != nil {
		return user, err
	}
	if status != "200 OK" {
		return user, fmt.Errorf("Error creating user: %s", status)
	}
	fmt.Println(status)
	return user, nil
}
func (HasuraUserRepository) GetUser(id string) (models.User, error) {
	return models.User{}, nil
}
func (HasuraUserRepository) GetUsers() ([]models.User, error) {
	return []models.User{}, nil
}
func (HasuraUserRepository) UpdateUser(user models.User) (models.User, error) {
	return user, nil
}
func (HasuraUserRepository) UpdateContactInfo(id string, contactInfo models.ContactInfo) (models.User, error) {
	return models.User{}, nil
}
func (HasuraUserRepository) DeleteUser(id string) error {
	return nil
}
