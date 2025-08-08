package repository
import (
	"app/internal/models"
	"context" 
	"errors"
)		

var (
	errNotFound = errors.New("entry not found")
)

type SpacecraftRepository  interface {
	Create(ctx context.Context, craft *models.SpacecraftRequest) (int64, error)
	Update(ctx context.Context,id string, craft *models.SpacecraftRequest) error
	Delete(id int) error
	GetByID(id int, filter *string) (models.Spacecraft, error)
	Get(ctx context.Context, filters *map[string][]string) ([]models.Spacecraft, error)
}
 