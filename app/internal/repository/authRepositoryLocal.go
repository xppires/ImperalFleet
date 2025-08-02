package repository

import (
	"app/internal/models"
)

var Users = []models.User{
	// Sample users
	// In a real application, you would fetch these from a database
	// For simplicity, we are using hardcoded values here
	{UID: "TK1", Username: "neo", Password: "keanu", Email: "", Role: "Technician"},
	{UID: "MG1", Username: "morpheus", Password: "lawrence", Email: "", Role: "Manager"},
}

type authRepositoryLocal struct {
	users []models.User
}

func NewAuthRepositoryLocal() *authRepositoryLocal {
	return &authRepositoryLocal{users: Users}
}
func (a *authRepositoryLocal) Authenticate(username, password string) (bool, string, string, error) {
	for _, user := range a.users {
		if user.Username == username && user.Password == password {
			return true, user.UID, user.Role, nil
		}
	}
	return false, "", "", nil
}