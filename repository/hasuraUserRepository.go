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

func (rep HasuraUserRepository) CreateUser(user models.User) (models.User, error) {
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
func (rep HasuraUserRepository) GetUser(id int) (models.User, error) {
	return models.User{}, nil
}
func (rep HasuraUserRepository) GetUsers() ([]models.User, error) {
	return []models.User{}, nil
}
func (rep HasuraUserRepository) UpdateUser(user models.User) (models.User, error) {
	return user, nil
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

func PostRequest(body []byte, route string) (string, error) {
	start := time.Now()
	req, err := http.NewRequest("POST", goDotEnvVariable("HASURA_URL")+route, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-hasura-admin-secret", goDotEnvVariable("HASURA_ADMIN_SECRET"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println(time.Since(start))
	return resp.Status, nil
}

func GetRequest(route string) (string, error) {
	req, err := http.NewRequest("GET", goDotEnvVariable("HASURA_URL")+route, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-hasura-admin-secret", goDotEnvVariable("HASURA_ADMIN_SECRET"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Status, nil
}
