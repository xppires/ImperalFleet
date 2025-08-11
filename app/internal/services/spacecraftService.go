package services

import (
	"app/internal/models"
	"app/internal/repository"
	"context"
)

type SpacecraftService struct {
	SpacecraftRepo repository.SpacecraftRepository
}

func NewSpacecraftService(spacecraftkRepo repository.SpacecraftRepository) *SpacecraftService {

	return &SpacecraftService{
		SpacecraftRepo: spacecraftkRepo,
	}
}

func (s *SpacecraftService) Create(ctx context.Context, Spacecraft *models.SpacecraftRequest) (int64, error) {
	SpacecraftID, err := s.SpacecraftRepo.Create(ctx, Spacecraft)
	return SpacecraftID, err
}

func (s *SpacecraftService) Get(ctx context.Context, filter *map[string][]string) ([]models.Spacecraft, error) {
	return s.SpacecraftRepo.Get(ctx, filter)
}

func (s *SpacecraftService) Delete(ctx context.Context, SpacecraftID int) error {
	return s.SpacecraftRepo.Delete(ctx, SpacecraftID)
}

func (s *SpacecraftService) GetByID(ctx context.Context, SpacecraftID int, filter *string) (models.Spacecraft, error) {
	return s.SpacecraftRepo.GetByID(ctx, SpacecraftID, filter)
}

func (s *SpacecraftService) Update(ctx context.Context, SpacecraftID string, craft *models.SpacecraftRequest) error {
	return s.SpacecraftRepo.Update(ctx, SpacecraftID, craft)
}
