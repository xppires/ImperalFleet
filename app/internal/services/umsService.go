package  services

import (
	"app/internal/repository"
)

type UmsService struct {
	authRepo repository.AuthRepository
}
func NewUmsService(authRepo repository.AuthRepository) *UmsService {
	return &UmsService{authRepo: authRepo}
}
func (s *UmsService) Authenticate(username, password string) (bool, string, string, error) {
	return s.authRepo.Authenticate(username, password)
}
