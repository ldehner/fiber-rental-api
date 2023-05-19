package hasura

import (
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

type HasuraUserRepository struct {
}

func (HasuraUserRepository) CreateUser(user storemodels.User) (storemodels.User, error) {
	/*
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
	*/
	return user, nil
}
func (HasuraUserRepository) GetUser(id string) (storemodels.User, error) {
	return storemodels.User{}, nil
}
func (HasuraUserRepository) GetUsers() ([]storemodels.User, error) {
	return []storemodels.User{}, nil
}
func (HasuraUserRepository) UpdateUser(user storemodels.User) (storemodels.User, error) {
	return user, nil
}
func (HasuraUserRepository) UpdateContactInfo(contactInfo storemodels.ContactInfo) (storemodels.User, error) {
	return storemodels.User{}, nil
}
func (HasuraUserRepository) DeleteUser(id string) error {
	return nil
}
