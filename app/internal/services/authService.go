package  services

import (
	"app/internal/repository"
)

type AuthService struct {
	authRepo repository.AuthRepository
}
func NewAuthService(authRepo repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}
func (s *AuthService) Authenticate(username, password string) (bool, string, string, error) {
	return s.authRepo.Authenticate(username, password)
}
