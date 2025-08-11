package repository

import (
	"app/internal/interfaces"
	"app/internal/models"
	"context"
	"fmt"
)

type SpacecraftRepositoryPosrtgresql struct {
	conn interfaces.DBStore
}

func NewSpacecraftRepositoryPosrtgresql(conn interfaces.DBStore) *SpacecraftRepositoryPosrtgresql {
	return &SpacecraftRepositoryPosrtgresql{
		conn: conn,
	}
}

// Create an entry for a new spaceship.
func (r *SpacecraftRepositoryPosrtgresql) Create(ctx context.Context, craft *models.SpacecraftRequest) (int64, error) {

	return 0, fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryPosrtgresql) Update(ctx context.Context, id string, craft *models.SpacecraftRequest) error {
	return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryPosrtgresql) Delete(ctx context.Context, id int) error {
	return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryPosrtgresql) GetByID(ctx context.Context, id int, filter *string) (models.Spacecraft, error) {
	var spacecraft models.Spacecraft
	return spacecraft, fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryPosrtgresql) Get(ctx context.Context, filters *map[string][]string) ([]models.Spacecraft, error) {
	var list []models.Spacecraft
	return list, fmt.Errorf("not implemented")
}
