package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}
