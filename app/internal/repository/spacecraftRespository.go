package repository

import ( 
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
func (r *SpacecraftRepository) Create( craft *models.SpacecraftRequest) (int, error) {

	 return 0, fmt.Errorf("not implemented")
}

func (r *SpacecraftRepository) Update( id int, craft *models.SpacecraftRequest) error {
  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepository) Delete( id int) error {
	  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepository) GetByID( id int,filter *string) (models.Spacecraft, error) {
	var spacecraft models.Spacecraft
	  return spacecraft, fmt.Errorf("not implemented")
}

func (r *SpacecraftRepository) Get( filters *string) ([]models.Spacecraft, error) {
	list := []models.Spacecraft{}
	return list, fmt.Errorf("not implemented")
	 
}