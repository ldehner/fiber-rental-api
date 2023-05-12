package hasura

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

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
func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
