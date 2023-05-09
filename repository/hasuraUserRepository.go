package repository

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/ldehner/fiber-rental-api/models"
)

type HasuraUserRepository struct {
}

func (rep HasuraUserRepository) CreateUser(user models.User) models.User {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	json, err := json.Marshal(user)
	fmt.Println(string(json))
	url := goDotEnvVariable("HASURA_URL") + "/user"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println(err)
		return models.User{}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-hasura-admin-secret", goDotEnvVariable("HASURA_ADMIN_SECRET"))
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return user
}
func (rep HasuraUserRepository) GetUser(id int) models.User {
	return models.User{}
}
func (rep HasuraUserRepository) GetUsers() []models.User {
	return []models.User{}
}
func (rep HasuraUserRepository) UpdateUser(user models.User) models.User {
	return user
}
func (rep HasuraUserRepository) DeleteUser(id int) error {
	return nil
}
func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
