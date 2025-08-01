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
	Create(craft *models.SpacecraftRequest) (int, error)
	Update(id int, craft *models.SpacecraftRequest) error
	Delete(id int) error
	GetByID(id int, filter *string) (models.Spacecraft, error)
	Get(ctx context.Context, filters *map[string][]string) ([]models.Spacecraft, error)
}
 