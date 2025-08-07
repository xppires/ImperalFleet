package repository

import ( 
	"app/internal/models"
	"app/internal/interfaces"
	"context"
	"fmt" 
	
)


type SpacecraftRepositoryMysql struct {
	conn  interfaces.DBStore
}

func NewSpacecraftRepositoryMysql(conn interfaces.DBStore) *SpacecraftRepositoryMysql {
	return &SpacecraftRepositoryMysql{
		conn: conn,
	}
}

// Create an entry for a new spaceship.
func (r *SpacecraftRepositoryMysql) Create( ctx context.Context,craft *models.SpacecraftRequest) (int64, error) {
	qSpaceship := `INSERT INTO spaceships (name, class, status, image, crew, value) VALUES (?, ?, ?, ?, ?, ?)`
	qArmaments := `INSERT INTO armaments (spaceship_id,title, qty) VALUES (?, ?, ?)`

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	insertArmament, err := tx.Prepare(qArmaments)
	if err != nil {
		return 0, fmt.Errorf("spacecraft_repo: prepare stmt: %w", err)
	}

	result, err := tx.ExecContext(ctx, qSpaceship, craft.Name, craft.Class, craft.Status, craft.Image, craft.Crew, craft.Value)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("spacecraft_repo: insert spacecraft: %w", err)
	}
	lastID,_ := result.LastInsertId()   

	for _, armament := range craft.Armament {
		_, err := insertArmament.ExecContext(ctx, lastID, armament.Title, armament.Quantity)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("spacecraft_repo: insert armament: %w", err)
		}
	}
	tx.Commit()
	return lastID,nil
}

func (r *SpacecraftRepositoryMysql) Update( id int, craft *models.SpacecraftRequest) error {
  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryMysql) Delete( id int) error {
	  return fmt.Errorf("not implemented")
}
func (r *SpacecraftRepositoryMysql) GetByID( id int,filter *string) (models.Spacecraft, error) {
	var spacecraft models.Spacecraft
	  return spacecraft, fmt.Errorf("not implemented")
}

func (r *SpacecraftRepositoryMysql) Get( ctx context.Context,filters *map[string][]string) ([]models.Spacecraft, error) {
	// list := []models.Spacecraft{}
	// return list, fmt.Errorf("not implemented")

	// q := `
	// SELECT id, name, class, status , image, crew, value FROM spaceships
	// WHERE (LOWER(name) = LOWER($1)) OR $1 = ''
	// AND (LOWER(class) = LOWER($2)) OR $2 = ''
	// AND (LOWER(status) = LOWER($3)) OR $3 = ''
	// ORDER BY id
	// `
	q := `
	SELECT id, name, class, status , image, crew, value FROM spaceships 
	ORDER BY id
	`
	// var name, class ,status string = "", "", ""

	// name := filters.Get("name")
	// class := filters.Get("class")
	// status := filters.Get("status")

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("spacecraft_repo: begin tx: %w", err)
	}

	selectArmaments, err := tx.Prepare("SELECT id, spaceship_id, title, qty FROM armaments WHERE spaceship_id = ?")
	if err != nil {
		return nil, fmt.Errorf("spacecraft_repo: preparing armaments stmt: %w", err)
	}
	defer selectArmaments.Close()
	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("spacecraft_repo: retrieve spacecrafts spacecrafts: %w", err)
	}
	defer rows.Close()

	spaceshipsMap := make(map[int]*models.Spacecraft,0)
	for rows.Next() {
		var spacecraft models.Spacecraft
		if err := rows.Scan(&spacecraft.ID, &spacecraft.Name, &spacecraft.Class, &spacecraft.Status, &spacecraft.Image, &spacecraft.Crew, &spacecraft.Value); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("spacecraft_repo: retrieve spacecrafts: %w", err)
		}
		spaceshipsMap[len(spaceshipsMap)] = &spacecraft
		
	} 
	for craftIdx, spacecraft := range spaceshipsMap {
		armaments := make([]models.Armament, 0) 
		armRows, err := selectArmaments.Query(spacecraft.ID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err)
		}
		defer armRows.Close()

		for armRows.Next() {
			var armament models.Armament
			if err := armRows.Scan(&armament.ID, &armament.CraftID, &armament.Title, &armament.Quantity); err != nil {
				return nil, fmt.Errorf("spacecraft_repo: retrieve armaments: %w", err)
			}
			armaments = append(armaments, armament)
			spaceshipsMap[craftIdx].Armament = append(spaceshipsMap[craftIdx].Armament, armament)
		} 
		clear(armaments) 
		
	} 
	returnedSpacecrafts := make([]models.Spacecraft, 0, len(spaceshipsMap))
	for _, spacecraft := range spaceshipsMap {
		returnedSpacecrafts = append(returnedSpacecrafts, *spacecraft)
	}	
	return returnedSpacecrafts, nil 
}