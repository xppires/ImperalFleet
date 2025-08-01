package services


import (
	"app/internal/repository"
	"app/internal/models"
	"context"

)



type SpacecraftService struct {
	SpacecraftRepo  repository.SpacecraftRepository 
}

func NewSpacecraftService(spacecraftkRepo repository.SpacecraftRepository ) *SpacecraftService {
 
	return &SpacecraftService{
		SpacecraftRepo:  spacecraftkRepo , 
	}
}


func (s *SpacecraftService) Create(Spacecraft *models.SpacecraftRequest) (int, error) {
	SpacecraftID, err := s.SpacecraftRepo.Create(Spacecraft)
	return  SpacecraftID, err
}

func (s *SpacecraftService) Get(ctx context.Context, filter *map[string][]string) ( [] models.Spacecraft, error) {
	return s.SpacecraftRepo.Get(ctx,filter)
}

func (s *SpacecraftService) Delete(SpacecraftID int) (  error) {
	return s.SpacecraftRepo.Delete(SpacecraftID)
}

func (s *SpacecraftService) GetByID(SpacecraftID int ,filter *string) ( models.Spacecraft, error) {
	return s.SpacecraftRepo.GetByID(SpacecraftID, filter)
}

func (s *SpacecraftService) Update(SpacecraftID int, craft *models.SpacecraftRequest) (  error) {
	return  s.SpacecraftRepo.Update(SpacecraftID, craft) 
}
