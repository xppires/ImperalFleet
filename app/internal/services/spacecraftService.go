package services


import (
	"app/internal/repository"

)



type SpacecraftService struct {
	repo  *repository.SpacecraftRepository 
}

func NewSpacecraftService(spacecraftkRepo *repository.SpacecraftRepository ) *SpacecraftService {
 
	return &SpacecraftService{
		repo:  spacecraftkRepo , 
	}
}