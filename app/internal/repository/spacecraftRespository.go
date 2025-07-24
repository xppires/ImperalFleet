package repository

import (
	"context" 
	"errors" 
	"app/internal/models"
	"app/internal/interfaces"
	"fmt"
	
)

var (
	errNotFound = errors.New("entry not found")
)

type SpacecraftRepository struct {
	conn  interfaces.DBStore
}

func NewSpacecraftRepository(conn interfaces.DBStore) *SpacecraftRepository {
	return &SpacecraftRepository{
		conn: conn,
	}
}

// Create an entry for a new spaceship.
func (r *SpacecraftRepository) Create(ctx context.Context, craft models.SpacecraftRequest) error {
	 return fmt.Errorf("not implemented")
}

func (r *SpacecraftRepository) Update(ctx context.Context, id int, craft models.SpacecraftRequest) error {
  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepository) Delete(ctx context.Context, id int) error {
	  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepository) GetByID(ctx context.Context, id int) (models.Spacecraft, error) {
	var spacecraft models.Spacecraft
	  return spacecraft, fmt.Errorf("not implemented")
}

// func (r *SpacecraftRepository) Get(ctx context.Context, filters url.Values) ([]models.Spacecraft, error) {
// 	 return [], fmt.Errorf("not implemented")
	 
// }